package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"user/domain/repository"
	"user/domain/service"
	"user/handler"
	pb "user/proto"

	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName = "go.micro.service.user"
	version     = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
	)

	//创建数据库服务
	db, err := gorm.Open("mysql", "root:root@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次
	/*rp := repository.NewUserRepository(db)
	rp.InitTable()*/
	userDataService := service.NewUserDataService(repository.NewUserRepository(db))

	// Register handler
	err = pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
