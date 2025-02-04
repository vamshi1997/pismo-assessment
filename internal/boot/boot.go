package boot

import (
	"fmt"
	"github.com/vamshi1997/pismo-assessment/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitApp() {
	InitConfig()
	InitDb()
}

func GetDB() *gorm.DB {
	return db
}

func InitDb() {
	cfg = GetConfig()

	fmt.Println(cfg.AppConfig.DB.Host)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.AppConfig.DB.UserName,
		cfg.AppConfig.DB.Password,
		cfg.AppConfig.DB.Host,
		cfg.AppConfig.DB.Port,
		cfg.AppConfig.DB.DBName,
		cfg.AppConfig.DB.Charset,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Not able connect to database")
		panic(err)
	}
	fmt.Println("Application connected to database successfully ...")

	err := db.AutoMigrate(&model.Account{}, &model.Transaction{})
	if err != nil {
		fmt.Println("Not able migrate account or transaction table")
		panic(err)
	}
	fmt.Println("migrated account table successfully ...")

}
