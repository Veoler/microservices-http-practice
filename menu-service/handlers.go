package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// GET /drinks — список всех напитков
func listDrinks(c *gin.Context) {
    var drinks []Drink
    db.Find(&drinks)
    c.JSON(http.StatusOK, drinks)
}

// GET /drinks/:id — один напиток
func getDrink(c *gin.Context) {
    id := c.Param("id")

    var drink Drink
    if err := db.First(&drink, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Напиток не найден"})
        return
    }

    c.JSON(http.StatusOK, drink)
}

// POST /drinks — создать напиток
func createDrink(c *gin.Context) {
    var req CreateDrinkRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    drink := Drink{
        Name:    req.Name,
        Price:   req.Price,
        InStock: req.InStock,
    }

    db.Create(&drink)
    c.JSON(http.StatusCreated, drink)
}