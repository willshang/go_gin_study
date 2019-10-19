# Gin入门实战
- https://github.com/e421083458/gin_scaffold
- https://github.com/e421083458/golang_common
- https://github.com/taylorchen709/vue-admin

## API
- 1.GET localhost:8880/demo/index
- 2.GET localhost:8880/demo/bind
``` 
  -F name=zhangsanxx \
  -F age=10 \
  -F password=xx
```
- 3.GET localhost:8880/demo/redis
- 4.GET localhost:8880/demo/dao
- 5.POST localhost:8880/api/login
``` 
  -F username=admin \
  -F password=123456
```
- 6.GET localhost:8880/api/loginOut
- 7.GET localhost:8880/api/user/listPage 
``` 
  -F page=1
```
- 8.GET localhost:8880/api/user/add
``` 
  -F name=lisi \
  -F sex=1 \
  -F birth=1994-01-01 \
  -F age=15 \
  -F addr=xxx
```
- 9.GET localhost:8880/api/user/edit
``` 
  -F name=lisi \
  -F sex=1 \
  -F birth=1994-01-01 \
  -F age=16 \
  -F addr=xxx \
  -F id=2
```
- 10.GET localhost:8880/api/user/remove
``` 
  -F ids=3
```
- 11.GET localhost:8880/api/user/batchRemove
