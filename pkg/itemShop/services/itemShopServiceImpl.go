package services

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	_itemShopException "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/exception"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	"github.com/labstack/echo/v4"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
	logger             echo.Logger
}

func NewItemShopServiceImpl(itemShopRepository _itemShopRepository.ItemShopRepository, logger echo.Logger) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository, logger}
}

func (s *itemShopServiceImpl) Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		s.logger.Errorf("Failed to list item or counting: %s", err.Error())
		return nil, &_itemShopException.ItemListing{}
	}
	itemCounting := int64(len(itemEntityList))
	totalPage := s.totalPageCalculation(&itemCounting, itemFilter.Size)
	return s.toItemResultResponse(itemEntityList, itemFilter.Page, totalPage), err
}

func (s *itemShopServiceImpl) totalPageCalculation(totalItems *int64, size int64) int64 {
	totalPage := *totalItems / size
	if *totalItems%size != 0 {
		totalPage++
	}
	return totalPage
}

func (s *itemShopServiceImpl) toItemResultResponse(itemEntityList []*entities.Item, page, totalPage int64) *_itemShopModel.ItemResult {
	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemEntityList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return &_itemShopModel.ItemResult{
		Items: itemModelList,
		Paginate: _itemShopModel.PaginateResult{
			Page:      page,
			TotalPage: totalPage,
		},
	}
}