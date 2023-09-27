package handler

import (
	"cart/common"
	"cart/domain/model"
	"cart/domain/service"
	pb "cart/proto"
	"context"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c Cart) AddCart(ctx context.Context, request *pb.CartInfo, response *pb.AddResponse) error {
	category := &model.Cart{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	id, err := c.CartDataService.AddCart(category)
	if err != nil {
		return err
	}
	response.CartId = id
	response.Msg = "添加成功"
	return nil
}

func (c Cart) CleanCart(ctx context.Context, request *pb.Clean, response *pb.Response) error {
	err := c.CartDataService.CleanCart(request.UserId)
	if err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (c Cart) Incr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	err := c.CartDataService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (c Cart) Decr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	err := c.CartDataService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (c Cart) DeleteItemByID(ctx context.Context, request *pb.CartID, response *pb.Response) error {
	err := c.CartDataService.DeleteCart(request.Id)
	if err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (c Cart) GetAll(ctx context.Context, request *pb.CartFindAll, response *pb.CartAll) error {
	category, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}
