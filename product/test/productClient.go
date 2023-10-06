package main

import (
	"common"
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	opentracing2 "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	product "product/proto"
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路跟踪
	t, io, err := common.NewTracer("go.micro.service.client", "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	srv := micro.NewService()
	srv.Init(
		micro.Name("go.micro.service.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8090"),
		micro.Registry(consulRegistry), //添加注册中心
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())), //绑定链路跟踪
	)

	productService := product.NewProductService("go.micro.service.product", srv.Client())
	productAdd := &product.ProductInfo{
		ProductName:        "goods1",
		ProductSku:         "cap",
		ProductPrice:       99.9,
		ProductDescription: "test-products",
		ProductImage: []*product.ProductImage{
			{
				ImageName: "pic1",
				ImageCode: "pic1",
				ImageUrl:  "www.baidu.com",
			},
			{
				ImageName: "pic2",
				ImageCode: "pic2",
				ImageUrl:  "www.google.com",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "size1",
				SizeCode: "size1",
			},
			{
				SizeName: "size2",
				SizeCode: "size2",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "seo_title",
			SeoKeywords:    "seo_title",
			SeoDescription: "seo_title",
			SeoCode:        "seo_title1",
		},
		ProductCategoryId: 1,
	}

	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
