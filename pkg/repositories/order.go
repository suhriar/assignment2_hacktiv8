package repositories

import (
	"assignment2/pkg/models"

	"github.com/jinzhu/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetOrderByID(orderId int) (*models.Order, error)
	GetAllOrdersWithItems() (*[]models.Order, error)
	GetOrderByIDWithItems(orderId int) (*[]models.Order, error)
	UpdateOrderByID(orderId int, order *models.Order) (*models.Order, error)
	DeleteOrder(orderId int) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (r *orderRepo) CreateOrder(order *models.Order) (*models.Order, error) {
	return order, r.db.Create(order).Error
}

func (r *orderRepo) GetOrderByIDWithItems(orderId int) (*[]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Items").Where("id=?", orderId).Find(&order).Error
	return &order, err
}

func (r *orderRepo) GetAllOrdersWithItems() (*[]models.Order, error)  {
	var orders []models.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return &orders, err
}

func (r *orderRepo) GetOrderByID(orderId int) (*models.Order, error) {
	var order models.Order

	err := r.db.Preload("Items").First(&order, "id=?", orderId).Error
	return &order, err
}

func (r *orderRepo) UpdateOrderByID(orderId int, updateOrder *models.Order) (*models.Order, error) {
	var order models.Order

	err := r.db.Model(&order).Where("id=?", orderId).Updates(updateOrder).Error
	return &order, err
}

func (r *orderRepo) DeleteOrder(orderId int) error {
	var order models.Order

	err := r.db.Where("id=?", orderId).Delete(&order).Error
	return err
}