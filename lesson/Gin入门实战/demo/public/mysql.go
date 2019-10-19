package public

import (
	"github.com/e421083458/gorm"
	"go_gin_study/lesson/Gin入门实战/common/lib"
)

var (
	GormPool *gorm.DB
)

func InitMysql() error {
	dbPool, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	GormPool = dbPool
	return nil
}
