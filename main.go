package main

import (
	"errors"
	"fmt"
	"mini-challenge-5/database"
	"mini-challenge-5/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createProduct("Samsung")
	// updateProductById(1, "Apple")
	// getProductById(1)
	// createVariant(1, "12", 10)
	// getProductWithVariant()
	// deleteVariantById(1)
	// deleteProductWithVariants(2)

}

func createProduct(name string) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil.")
		return
	}

	product := models.Product{
		Name: name,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error creating product data: ", err)
		return
	}

	fmt.Println("New Product Data", product)
}

func updateProductById(id int, name string) {
	db := database.GetDB()

	product := models.Product{}

	err := db.Model(&product).Where("id = ?", id).Updates(models.Product{Name: name}).Error

	if err != nil {
		fmt.Println("Error updating product data:", err)
		return
	}
	fmt.Printf("Update product's name: %+v \n", product.Name)
}

func getProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}

	err := db.First(&product, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Product data not found")
			return
		}
		print("Error finding product:", err)
	}

	fmt.Printf("Product data: %+v \n", product)
}

func createVariant(productId uint, variantName string, quantity int) {
	db := database.GetDB()

	Variant := models.Variant{
		ProductID:   productId,
		VariantName: variantName,
		Quantity:    quantity,
	}

	err := db.Create(&Variant).Error

	if err != nil {
		fmt.Println("Error creating Variant data:", err)
		return
	}

	fmt.Println("New Variant Data", Variant)
}

func getProductWithVariant() {
	db := database.GetDB()

	var products []models.Product
	err := db.Preload("Variants").Find(&products).Error

	if err != nil {
		fmt.Println("Error getting products data with variants:", err.Error())
		return
	}

	fmt.Println("Product Data With Variants")
	for _, product := range products {
		fmt.Printf("Product ID: %d, Name: %s\n", product.ID, product.Name)
		fmt.Println("Variants:")
		for _, variant := range product.Variants {
			fmt.Printf("  Variant ID: %d, Name: %s, Quantity: %d\n", variant.ID, variant.VariantName, variant.Quantity)
		}
		fmt.Println("---------------------------")
	}
}

func deleteVariantById(id uint) {
	db := database.GetDB()

	variant := models.Variant{}

	err := db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		fmt.Println("Error deleting variant:", err.Error())
		return
	}

	fmt.Printf("Variant with id %d has been successfully deleted", id)
}

func deleteProductWithVariants(id uint) {
	db := database.GetDB()
	variant := models.Variant{}
	product := models.Product{}

	result := db.Transaction(func(tx *gorm.DB) error {

		tx.Delete(&variant, "product_id = ?", id)

		tx.Delete(&product, id)

		return nil
	})

	if result != nil {
		fmt.Println("Error deleting product and variants:", result.Error())
		return
	}

	fmt.Println("Product and variants deleted successfully")
}
