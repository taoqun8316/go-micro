package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	opentracing2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"product/domain/repository"
	"product/domain/service"
	"product/handler"
	pb "product/proto"

	"common"
	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	serviceName = "go.micro.service.product"
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

	//链路跟踪
	t, io, err := common.NewTracer(serviceName, "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	//创建数据库服务
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry), //添加注册中心
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())), //绑定链路跟踪
	)

	//只执行一次
	/*rp := repository.NewProductRepository(db)
	rp.InitTable()*/

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
