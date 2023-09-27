package handler

import (
	"cart/common"
	"cart/domain/model"
	"cart/domain/service"
	pb "cart/proto"
	"context"
	"go-micro.dev/v4/util/log"
)

type Cart struct {
	CartDataService service.ICartDataService
}

func (c Cart) AddCart(ctx context.Context, request *pb.CartInfo, response *pb.AddResponse) error {
	cart := &model.Cart{}
	err := common.SwapTo(request, cart)
	if err != nil {
		return err
	}
	id, err := c.CartDataService.AddCart(cart)
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
	response.Msg = "清空成功"
	return nil
}

func (c Cart) Incr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	err := c.CartDataService.IncrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "增加成功"
	return nil
}

func (c Cart) Decr(ctx context.Context, request *pb.Item, response *pb.Response) error {
	err := c.CartDataService.DecrNum(request.Id, request.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "扣除成功"
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
	categories, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	cartToResponse(categories, response)
	return nil
}

func cartToResponse(carts []model.Cart, response *pb.CartAll) {
	for _, cg := range carts {
		cr := &pb.CartInfo{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.CartInfo = append(response.CartInfo, cr)
	}
}
