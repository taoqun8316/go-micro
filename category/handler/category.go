package handler

import (
	"category/common"
	"category/domain/model"
	"category/domain/service"
	pb "category/proto"
	"context"
	"go-micro.dev/v4/util/log"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	id, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.CategoryId = id
	response.Message = "添加成功"
	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, request *pb.CategoryRequest, response *pb.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "修改成功"
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, request *pb.DeleteCategoryRequest, response *pb.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, request *pb.FindByNameRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByID(ctx context.Context, request *pb.FindByIDRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *pb.FindByLevelRequest, response *pb.FindAllResponse) error {
	categories, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categories, response)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *pb.FindByParentRequest, response *pb.FindAllResponse) error {
	categories, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categories, response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *pb.FindAllRequest, response *pb.FindAllResponse) error {
	categories, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categories, response)
	return nil
}

func categoryToResponse(categories []model.Category, response *pb.FindAllResponse) {
	for _, cg := range categories {
		cr := &pb.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Categories = append(response.Categories, cr)
	}
}
