package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// 可以不在本地安装postgresQL
	// 1.使用ElephantSQL创建一个SQL服务
	// 2.docker + 端口映射 + volume挂载
	dsn := os.Getenv("DB_URL")
	// 切记使用`=`而非`:=`, 否则会重新初始化DB，导致nil reference
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
