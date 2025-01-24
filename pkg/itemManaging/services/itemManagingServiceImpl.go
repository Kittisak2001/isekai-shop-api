package services

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/repositories"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopRepository "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/repositories"
	"github.com/labstack/echo/v4"
)

type (
	itemManagingService struct {
		itemManagingRepository _itemManagingRepository.ItemManagingRepository
		itemShopRepository     _itemShopRepository.ItemShopRepository
		logger                 echo.Logger
	}
)

func NewItemManagingService(itemManagingRepository _itemManagingRepository.ItemManagingRepository, itemShopRepository _itemShopRepository.ItemShopRepository, logger echo.Logger) ItemManagingService {
	return &itemManagingService{itemManagingRepository, itemShopRepository, logger}
}

func (s *itemManagingService) Creating(itemCreatingReq *model.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}
	if err := s.itemManagingRepository.Creating(itemEntity); err != nil {
		s.logger.Errorf("Failed to creating item : %s", err.Error())
		return nil, &exception.ItemCreating{}
	}
	return itemEntity.ToItemModel(), nil
}

func (s *itemManagingService) Editing(itemID *uint64, itemEditingModel *model.ItemEditingReq) (*_itemShopModel.Item, error) {
	if err := s.itemManagingRepository.Editing(itemID, itemEditingModel); err != nil {
		if err.Error() == "record not found"{
			return nil, &exception.ItemNotfound{ItemID: *itemID}
		}
		s.logger.Errorf("Failed to editing item : %s", err.Error())
		return nil, &exception.ItemEditing{}
	}
	item, err := s.itemShopRepository.FindById(itemID)
	if err != nil {
		s.logger.Errorf("Failed to find item by id : %s", err.Error())
		return nil, &exception.ItemNotfound{ItemID: *itemID}
	}
	return item.ToItemModel(), nil
}

func (s *itemManagingService) Archiving(itemID *uint64) error{
	if err := s.itemManagingRepository.Archiving(itemID); err != nil{
		s.logger.Errorf("Failed to archive item by id : %s", err.Error())
		return &exception.ItemArchiving{ItemID: *itemID}
	}
	return nil
}