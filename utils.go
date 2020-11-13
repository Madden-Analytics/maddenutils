package maddenutils

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// MaddenTransactionItems holds transcation items from Madden
type MaddenTransactionItems struct {
	ProductName       string  `json:"productName"`
	Sku               string  `json:"sku"`
	Ean               string  `json:"ean"`
	BrandName         string  `json:"brandName"`
	Category          string  `json:"category"`
	Collection        string  `json:"collection"`
	Season            string  `json:"season"`
	Year              string  `json:"year"`
	Size              string  `json:"size"`
	Color             string  `json:"color"`
	AgeGroup          string  `json:"ageGroup"`
	Gender            string  `json:"gender"`
	ModelNumber       string  `json:"modelNumber"`
	Quantity          int     `json:"quantity"`
	VatValue          int     `json:"vatValue"`
	MoneyItemTotalNet float64 `json:"moneyItemTotalNet"`
	MoneyItemTotalVat float64 `json:"moneyItemTotalVat"`
	CostPrice         float64 `json:"costPrice"`
	CostPriceCurency  string  `json:"costPriceCurrency"`
	MoneyDiscount     float64 `json:"moneyDiscount"`
}

// MaddenTransaction holds transaction from Madden
type MaddenTransaction struct {
	TransactionDate         time.Time                `json:"transactionDate"`
	ExternalStoreID         string                   `json:"externalStoreId"`
	ExternalTransactionID   string                   `json:"externalTransactionId"`
	MoneyFinalNet           float64                  `json:"moneyFinalNet"`
	MoneyFinalVat           float64                  `json:"moneyFinalVat"`
	MoneyTotalGrossRoundOff float64                  `json:"MoneyTotalGrossRoundOff"`
	CurrencyCode            string                   `json:"currencyCode"`
	ExternalRowID           string                   `json:"externalRowId"`
	Items                   []MaddenTransactionItems `json:"items"`
}

// MaddenStock holds stock items
type MaddenStock struct {
	Sku       string `json:"sku"`
	EAN       string `json:"ean"`
	BrandName string `json:"brandName"`
	Quantity  int    `json:"quantity"`
}

// MaddenBearer holds the madden token reponse
type MaddenBearer struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
}

// APIKey holds the Madden apiKey
type APIKey struct {
	APIKey string `json:"apiKey"`
}

// GetAuth generates bearer token from Madden
func GetAuth(baseURL string, accountID string, key string) string {
	maddenKey := APIKey{
		APIKey: key,
	}

	maddenKeyJSON, _ := json.Marshal(maddenKey)

	// Get Madden Bearer Token
	maddenAuthResponse := Post(
		baseURL+"/accounts/"+accountID+"/auth",
		"",
		maddenKeyJSON,
	)
	var maddenAuth MaddenBearer
	json.Unmarshal(maddenAuthResponse, &maddenAuth)

	return maddenAuth.AccessToken
}

// BatchTransactions batches transaction struct into given size
func BatchTransactions(batchSize int, batchSlice []MaddenTransaction) [][]MaddenTransaction {

	batches := make([][]MaddenTransaction, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// BatchStocks batches stock struct into given size
func BatchStocks(batchSize int, batchSlice []MaddenStock) [][]MaddenStock {

	batches := make([][]MaddenStock, 0, (len(batchSlice)+batchSize-1)/batchSize)
	for batchSize < len(batchSlice) {
		batchSlice, batches = batchSlice[batchSize:], append(batches, batchSlice[0:batchSize:batchSize])
	}
	batches = append(batches, batchSlice)

	return batches
}

// StrToInt converts decimal string to int
func StrToInt(str string) int {
	nonFractionalPart := strings.Split(str, ".")
	theInt, _ := strconv.Atoi(nonFractionalPart[0])
	return theInt
}
