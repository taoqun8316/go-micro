package handler

import (
	"context"
	"product/domain/service"
	pb "product/proto"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p Product) AddProduct(ctx context.Context, info *pb.ProductInfo, product *pb.ResponseProduct) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) FindProductByID(ctx context.Context, id *pb.RequestID, info *pb.ProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) UpdateProduct(ctx context.Context, info *pb.ProductInfo, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) DeleteProductByID(ctx context.Context, id *pb.RequestID, response *pb.Response) error {
	//TODO implement me
	panic("implement me")
}

func (p Product) FindAllProduct(ctx context.Context, all *pb.RequestAll, product *pb.AllProduct) error {
	//TODO implement me
	panic("implement me")
}
