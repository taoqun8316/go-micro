package main

import (
	"category/domain/repository"
	"category/handler"
	pb "category/proto"
	"fmt"
	"github.com/jinzhu/gorm"

	"category/domain/service"
	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName = "category"
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
	rp := repository.NewCategoryRepository(db)
	rp.InitTable()
	categoryDataService := service.NewCategoryDataService(repository.NewCategoryRepository(db))

	// Register handler
	err = pb.RegisterCategoryHandler(srv.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
