package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

// Адрес Menu Service
const menuServiceURL = "http://localhost:8081"

func main() {
    dsn := "host=localhost user=postgres password=4545 dbname=orders_microservices-http-practice port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Не удалось подключиться к БД: " + err.Error())
    }

    db.AutoMigrate(&Order{})

    router := gin.Default()

    router.POST("/orders", createOrder)
    router.GET("/orders/:id", getOrder)
    router.GET("/orders", listOrders)

    router.Run(":8082")
}