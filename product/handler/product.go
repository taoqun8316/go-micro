package handler

import (
	"common"
	"context"
	"product/domain/model"
	"product/domain/service"
	pb "product/proto"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p Product) AddProduct(ctx context.Context, request *pb.ProductInfo, response *pb.ResponseProduct) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	id, err := p.ProductDataService.AddProduct(product)
	if err != nil {
		return err
	}
	response.ProductId = id
	return nil
}

func (p Product) FindProductByID(ctx context.Context, request *pb.RequestID, response *pb.ProductInfo) error {
	product, err := p.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err = common.SwapTo(product, response); err != nil {
		return err
	}
	return nil
}

func (p Product) UpdateProduct(ctx context.Context, request *pb.ProductInfo, response *pb.Response) error {
	product := &model.Product{}
	if err := common.SwapTo(request, product); err != nil {
		return err
	}
	err := p.ProductDataService.UpdateProduct(product)
	if err != nil {
		return err
	}
	response.Message = "修改成功"
	return nil
}

func (p Product) DeleteProductByID(ctx context.Context, request *pb.RequestID, response *pb.Response) error {
	err := p.ProductDataService.DeleteProduct(request.ProductId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

func (p Product) FindAllProduct(ctx context.Context, request *pb.RequestAll, response *pb.AllProduct) error {
	products, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}

	for _, v := range products {
		productInfo := &pb.ProductInfo{}
		if err = common.SwapTo(v, productInfo); err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
