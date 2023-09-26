package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/jinzhu/gorm"
	"go-micro.dev/v4/registry"
	"product/domain/repository"
	"product/domain/service"
	"product/handler"
	pb "product/proto"

	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	common "product/common"
)

var (
	serviceName = "product"
	version     = "latest"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	//创建数据库服务
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
	)

	//只执行一次
	rp := repository.NewProductRepository(db)
	rp.InitTable()

	productDataService := service.NewProductDataService(repository.NewProductRepository(db))

	// Register handler
	if err := pb.RegisterProductHandler(srv.Server(), &handler.Product{ProductDataService: productDataService}); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
