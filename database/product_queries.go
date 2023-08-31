package database

import "online_shop/models"

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	result := DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
func GetProductById(id int) (models.Product, error) {
	var product models.Product
	result := DB.First(&product, id)
	if result.Error != nil {
		return models.Product{}, result.Error
	}
	return product, nil
}
func DeleteProductById(id int) error {
	var product models.Product
	result := DB.Delete(&product, id)
	return result.Error
}
