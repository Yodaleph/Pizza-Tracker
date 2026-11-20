package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{
		"Order Placed", 
		"Preparing", 
		"Baking", 
		"Quality Check", 
		"Ready",
	}
	PizzaTypes = []string{
		"Margherita", 
		"Pepperoni",
		"Vegetarian",
		"BBQ Chicken",
		"Hawaiian",
		"Meat Lovers",
		"Supreme",
		"Buffalo Chicken",
		"Four Cheese",
		"Truffle Mushroom",
	}
	PizzaSizes = []string{
		"Small", 
		"Medium",
		"Large",
		"X-Large",
	}
)
type OrderModel struct {
	DB *gorm.DB
}
type Order struct {	
	ID 				string		`gorm:"primaryKey;size:14" json:"id"`
	Status 			string 		`gorm:"not null" json:"status"`
	CustomerName 	string 		`gorm:"not null" json:"customer_name"`
	Phone 			string 		`gorm:"not null" json:"phone"`
	Address 		string 		`gorm:"not null" json:"address"`
	Items 			[]OrderItem `gorm:"foreignKey:OrderID" json:"pizzas"`
	CreatedAt 		time.Time 	`json:"created_at"`
}

type OrderItem struct {
	ID 				string 		`gorm:"primaryKey;size:14" json:"id"`
	OrderID 		string 		`gorm:"not null" json:"order_id"`
	Size 			string 		`gorm:"not null" json:"size"`
	Pizza 			string 		`gorm:"not null" json:"pizza"`
	Instructions 	string 		`json:"instructions"`
}
func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderModel) CreateOrder(order *Order) error {
	return oi.DB.Create(order).Error
}
func(oi *OrderModel) GetOrderByID(id string) (*Order, error) {
	var order Order
	if err := oi.DB.Preload("Items").First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
func (o *OrderModel) GetAllOrders() ([]Order, error) {
	var orders []Order
	err := o.DB.Preload("Items").Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (o *OrderModel) UpdateOrderStatus(id, status string) error {
	return o.DB.Model(&Order{}).Where("id = ?", id).Update("status", status).Error
}

func (o *OrderModel) DeleteOrder(id string) error {
	return o.DB.Select("items").Delete(&Order{ID:id}).Error
}