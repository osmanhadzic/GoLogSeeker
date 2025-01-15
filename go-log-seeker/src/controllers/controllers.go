package controllers

import (
	"go-log-seeker/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var items = []models.Item{
	{Id: 1, Name: "Item 1", Description: "Description 1"},
	{Id: 2, Name: "Item 2", Description: "Description 2"},
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for _, item := range items {
		if item.Id == idInt {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem models.Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, item := range items {
		if item.Id == idInt {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	for i, item := range items {
		if item.Id == idInt {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
