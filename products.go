package maddenutils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func GetProductsMap(maBaseURL string, token string) (map[string]Product, error) {

	// Get all products from MA and make map
	maProductsMap := make(map[string]Product)
	page := 1
	for {
		// convert page to str
		pageString := strconv.Itoa(page)

		// Get all SKUS
		statusCode, response := Request("GET", maBaseURL+"/v1/products?pageSize=1000&page="+pageString, token, nil)
		if statusCode == http.StatusOK {
			var maProducts []Product
			json.Unmarshal(response, &maProducts)

			if len(maProducts) > 0 {
				for _, product := range maProducts {
					maProductsMap[product.Key] = product
				}
				page += 1

			} else {
				break
			}
		} else {
			err := errors.New("could not get products from MA")
			return nil, err
		}
	}

	return maProductsMap, nil

}

func GetSkusMap(maBaseURL string, token string) (map[string]SKU, error) {

	// Get all products from MA and make map
	maSkusMap := make(map[string]SKU)
	page := 1
	for {
		// convert page to str
		pageString := strconv.Itoa(page)

		// Get all SKUS
		statusCode, response := Request("GET", maBaseURL+"/v1/skus?pageSize=1000&page="+pageString, token, nil)
		if statusCode == http.StatusOK {
			var maSkus []SKU
			json.Unmarshal(response, &maSkus)

			if len(maSkus) > 0 {
				for _, sku := range maSkus {
					maSkusMap[sku.Key] = sku
				}
				page += 1

			} else {
				break
			}
		} else {
			err := errors.New("could not get skus from MA")
			return nil, err
		}
	}

	return maSkusMap, nil

}
