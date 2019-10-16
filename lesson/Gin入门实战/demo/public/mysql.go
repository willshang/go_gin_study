package public

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
