package repository

import (
	"assignment-2/database"
	"assignment-2/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateOrder(order *models.Order) error {
	order.OrderedAt = time.Now()
	err := database.DB.Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(id int) error {
	order, err := GetOrderByID(id)
	if err != nil {
		return err
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&order).Association("Items").Clear()
		if err != nil {
			return err
		}

		err = tx.Delete(&order, id).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func GetAllOrders() ([]models.Order, error) {
	order := []models.Order{}

	err := database.DB.Model(&order).Preload("Items").Find(&order).Error
	if err != nil {
		fmt.Println(err)
		return order, err
	}

	results := make([]models.Order, len(order))
	for index := range order {
		results[index] = models.Order(order[index])
	}

	return results, nil
}

func GetOrderByID(id int) (models.Order, error) {
	order := models.Order{}
	err := database.DB.Model(&models.Order{}).Preload("Items").Take(&order, id).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func UpdateOrder(id int, updatedData *models.Order) error {
	order, err := GetOrderByID(id)
	if err != nil {
		return err
	}

	if updatedData.CustomerName != "" {
		order.CustomerName = updatedData.CustomerName
	}

	updatedData.OrderedAt = time.Now()

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		if updatedData.Items != nil {
			err := tx.Model(&order).Association("Items").Replace(updatedData.Items)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
