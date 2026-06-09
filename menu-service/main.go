package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    dsn := "host=localhost user=postgres password=4545 dbname=menu_microservices-http-practice port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Не удалось подключиться к БД: " + err.Error())
    }

    db.AutoMigrate(&Drink{})

    router := gin.Default()

    router.GET("/drinks", listDrinks)
    router.GET("/drinks/:id", getDrink)
    router.POST("/drinks", createDrink)

    router.Run(":8081")
}