package services

import (
	"assignment2/pkg/helper"
	"assignment2/pkg/models"
	"assignment2/pkg/params"
	"assignment2/pkg/repositories"
	"net/http"
)

type ItemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

var createdItems []params.ItemResponse

func (i *ItemService) CreateItem(responseOrder params.Response, request params.CreateOder) *params.Response {
	createdItems = nil

	orderData := responseOrder.Payload
	order, err := orderData.(*models.Order)

	if !err {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
		}
	}

	items := request.Items

	for _, item := range items {
		itemModel := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     int(order.ID),
		}

		itemData, err := i.itemRepo.CreateItem(&itemModel)

		if err != nil {
			return &params.Response{
				Status:         400,
				Error:          "BAD REQUEST",
			}
		}

		createdItems = append(createdItems, params.ItemResponse{ItemID: int(itemData.ID),
			ItemCode:    itemData.ItemCode,
			Description: itemData.Description,
			Quantity:    itemData.Quantity,
			OrderID:     order.ID})
	}

	data  := params.AllResponseData{
		OrderID:      order.ID,
		OrderedAt:    *order.OrderedAt,
		CustomerName: order.CustomerName,
		Items:	createdItems,
	}

	return helper.SuccessCreateResponse(data, "succes")
}

func (i *ItemService) GetItemsByOrderID(orderId int) (*[]models.Item, *params.Response) {
	response, err := i.itemRepo.GetItemsByOrderID(orderId)
	if err != nil {
		return response, &params.Response{
			Status:         http.StatusNotFound,
			Error:          "Error - Item Not Found",
			AdditionalInfo: err.Error(),
		}
	}

	return response, &params.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Payload: response,
	}
}

func (i *ItemService) UpdateItemByID(itemModel *[]models.Item, request params.CreateOder) *params.Response {

	items := request.Items

	for _, v := range *itemModel {
		for _, itemRequest := range items {
			if uint(v.ID) == uint(itemRequest.ItemID) {
				updateItem := models.Item{
					ItemCode:    itemRequest.ItemCode,
					Description: itemRequest.Description,
					Quantity:    itemRequest.Quantity,
				}

				_, err := i.itemRepo.UpdateItemByID(itemRequest.ItemID, &updateItem)

				if err != nil {
					return &params.Response{
						Status:         400,
						Error:          "BAD REQUEST",
						AdditionalInfo: err,
					}
				}
			}
		}
	}

	return &params.Response{
		Status:  200,
		Message: "Success - Update Order & Items",
	}
}

func (i *ItemService) DeleteItems(orderId int) *params.Response {
	err := i.itemRepo.DeleteItem(orderId)

	if err != nil {
		return &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Success - Delete Items",
	}
}