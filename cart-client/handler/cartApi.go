package handler

import (
	pb "cart-client/proto"
	go_micro_service_cart "cart/proto"
	"context"
	"encoding/json"
	"errors"
	"go-micro.dev/v4/util/log"
	"strconv"
)

type CartApi struct {
	CartService go_micro_service_cart.CartService
}

func (c *CartApi) FindAll(ctx context.Context, request *pb.Request, response *pb.Response) error {
	log.Info("接收到请求")
	if _, ok := request.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}
	userIdString := request.Get["user_id"].Values[0]
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		return err
	}
	all, err := c.CartService.GetAll(context.TODO(), &go_micro_service_cart.CartFindAll{
		UserId: userId,
	})
	if err != nil {
		return err
	}
	b, err := json.Marshal(all)
	response.StatusCode = 200
	response.Body = string(b)
	return nil
}
