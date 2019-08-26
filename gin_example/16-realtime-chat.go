package main

import (
	"fmt"
	"github.com/dustin/go-broadcast"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"math/rand"
	"net/http"
)

func main() {
	router := gin.Default()
	router.SetHTMLTemplate(htmla)

	router.GET("/room/:roomid", roomGET01)
	router.GET("/room/:roomid", roomPOST01)
	router.DELETE("/room/:roomid", roomDELETE)
	router.GET("/stream/:roomid", stream)

	router.Run(":8080")
}

func stream(c *gin.Context) {
	roomid := c.Param("roomid")
	listener := openListener(roomid)
	defer closeListener(roomid, listener)

	clientGone := c.Writer.CloseNotify()
	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case messages := <-listener:
			c.SSEvent("message", message)
			return true
		}
	})
}

func roomGET01(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := fmt.Sprint(rand.Int31())
	c.HTML(http.StatusOK, "chat_room", gin.H{
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST01(c *gin.Context) {
	roomid := c.Param("roomid")
	userid := c.PostForm("user")
	message := c.PostForm("message")
	room(roomid).Submit(userid + ":" + message)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	deleteBroadcast(roomid)
}

var roomChannels = make(map[string]broadcast.Broadcaster)

func openListener(roomid string) chan interface{} {
	listener := make(chan interface{})
	room(roomid).Register(listener)
	return listener
}

func closeListener(roomid string, listener chan interface{}) {
	room(roomid).Unregister(listener)
	close(listener)
}

func deleteBroadcast(roomid string) {
	b, ok := roomChannels[roomid]
	if ok {
		b.Close()
		delete(roomChannels, roomid)
	}
}

func room(roomid string) broadcast.Broadcaster {
	b, ok := roomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		roomChannels[roomid] = b
	}
	return b
}

var htmla = template.Must(template.New("chat_room").Parse(`
<html>
<head>
    <title>{{.roomid}}</title>
    <link rel="stylesheet" type="text/css" href="http://meyerweb.com/eric/tools/css/reset/reset.css">
    <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7/jquery.js"></script>
    <script src="http://malsup.github.com/jquery.form.js"></script>
    <script>
        $('#message_form').focus();
        $(document).ready(function() {
            // bind 'myForm' and provide a simple callback function
            $('#myForm').ajaxForm(function() {
                $('#message_form').val('');
                $('#message_form').focus();
            });

            if (!!window.EventSource) {
                var source = new EventSource('/stream/{{.roomid}}');
                source.addEventListener('message', function(e) {
                    $('#messages').append(e.data + "</br>");
                    $('html, body').animate({scrollTop:$(document).height()}, 'slow');

                }, false);
            } else {
                alert("NOT SUPPORTED");
            }
        });
    </script>
    </head>
    <body>
    <h1>Welcome to {{.roomid}} room</h1>
    <div id="messages"></div>
    <form id="myForm" action="/room/{{.roomid}}" method="post">
    User: <input id="user_form" name="user" value="{{.userid}}">
    Message: <input id="message_form" name="message">
    <input type="submit" value="Submit">
    </form>
</body>
</html>
`))
