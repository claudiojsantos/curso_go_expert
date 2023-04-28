package main

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Product{})

	// products := []Product{
	// 	{Name: "Product 1", Price: 19.99},
	// 	{Name: "Product 2", Price: 29.99},
	// 	{Name: "Product 3", Price: 39.99},
	// }

	// db.Create(&products)

	// Select One
	// var product Product

	// db.First(&product, 1)

	// fmt.Println(product)

	// db.First(&product, "name = ?", "Product 2")

	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// select limit e offset
	// db.Limit(2).Offset(2).Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 20).Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// like
	// var products []Product
	// db.Where("name LIKE ?", "%uct 3%").Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update, delecao
	// var prod Product
	// var prods []Product

	// db.First(&prod, 1)
	// prod.Name = "Product 1 Updated"
	// db.Save(&prod)

	// fmt.Println(prod)

	// db.Delete(&prod)

	// db.Find(&prods)

	// for _, product := range prods {
	// 	fmt.Println(product)
	// }
}
