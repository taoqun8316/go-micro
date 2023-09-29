package main

import (
	"cart-client/handler"
	cartApi "cart-client/proto"
	go_micro_service_cart "cart/proto"
	"common"
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/select/roundrobin"
	opentracing2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"net"
	"net/http"
)

var (
	serviceName = "go.micro.api.cartApi"
	version     = "latest"
)

func main() {
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

	//熔断器
	hytrixStreamHandler := hystrix.NewStreamHandler()
	hytrixStreamHandler.Start()

	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hytrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("0.0.0.0:8086"),
		micro.Registry(consulRegistry), //添加注册中心
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())), //绑定链路跟踪
		micro.WrapClient(NewClientHystrixWrapper()),                                 //添加熔断
		micro.WrapClient(roundrobin.NewClientWrapper()),                             //添加负载均衡
	)

	cartService := go_micro_service_cart.NewCartService("go.micro.service.cart", srv.Client())

	// Register handler
	err = cartApi.RegisterCartApiHandler(srv.Server(), &handler.CartApi{CartService: cartService})
	if err != nil {
		log.Error(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, resp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
