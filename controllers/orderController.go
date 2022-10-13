package controllers

import (
	"assignment-2/models"
	"assignment-2/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreteOrder(ctx *gin.Context) {
	var newOrder *models.Order

	err := ctx.ShouldBindJSON(&newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}

	err = repository.CreateOrder(newOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, newOrder)
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, _ := strconv.Atoi(orderID)

	err := repository.DeleteOrder(orderIDInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order %d has been delete", orderIDInt),
	})
}

func GetAllOrders(ctx *gin.Context) {
	var orderList []models.Order

	orderList, err := repository.GetAllOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, orderList)
}

func GetOrderByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, _ := strconv.Atoi(orderID)

	var order models.Order

	order, err := repository.GetOrderByID(orderIDInt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderIDInt, _ := strconv.Atoi(orderID)

	var updatedOrder models.Order

	err := ctx.ShouldBindJSON(&updatedOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = repository.UpdateOrder(orderIDInt, &updatedOrder)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedOrder)
}
