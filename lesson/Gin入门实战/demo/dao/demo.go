package dao

import (
	"go_gin_study/lesson/Gin入门实战/demo/public"
	"time"
)

type Area struct {
	Id       int       `json:"id" orm:"column(id);auto" description:"自增主键"`
	AreaName string    `json:"area_name" orm:"column(area_name);size(191)" description:"区域名称"`
	CityID   int       `json:"city_id" orm:"column(city_id)" description:"城市id"`
	UserID   int       `json:"user_id" orm:"column(user_id)" description:"操作人"`
	UpdateAt time.Time `json:"update_at" orm:"column(update_at);type(datetime)" description:"更新时间"`
	CreateAt time.Time `json:"create_at" orm:"column(create_at);type(datetime)" description:"创建时间"`
	DeleteAt time.Time `json:"delete_at" orm:"column(delete_at);type(datetime)" description:"删除时间"`
}

func (a *Area) TableName() string {
	return "area"
}

func (a *Area) Find(id string) ([]*Area, error) {
	var area []*Area
	err := public.GormPool.Where("id=?", id).Find(&area).Error
	if err != nil {
		return nil, err
	}
	return area, nil
}
