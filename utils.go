package maddenutils

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"gopkg.in/guregu/null.v4"
)

// MaddenTransactionItems holds transcation items from Madden
type MaddenTransactionItems struct {
	ProductName       string  `json:"productName"`
	Sku               string  `json:"sku"`
	Ean               string  `json:"ean"`
	BrandName         string  `json:"brandName"`
	Supplier          string  `json:"supplier,omitempty"`
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
	ExternalID        string  `json:"externalID"`
}

type Seller struct {
	Name string `json:"name"`
}

type Payments struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// MaddenTransaction holds transaction from Madden
type MaddenTransaction struct {
	TransactionDate         time.Time                `json:"transactionDate"`
	TransactionID           int                      `json:"id,omitempty"`
	CompletedAt             null.Time                `json:"completed_at"`
	CancelledAt             null.Time                `json:"cancelled_at"`
	ExternalStoreID         string                   `json:"externalStoreId"`
	ExternalTransactionID   string                   `json:"externalTransactionId"`
	MoneyFinalNet           float64                  `json:"moneyFinalNet"`
	MoneyFinalVat           float64                  `json:"moneyFinalVat"`
	MoneyTotalGrossRoundOff float64                  `json:"MoneyTotalGrossRoundOff"`
	CurrencyCode            string                   `json:"currencyCode"`
	Market                  string                   `json:"market"`
	ExternalRowID           string                   `json:"externalRowId"`
	CartDiscount            float64                  `json:"cartDiscount"`
	ShippingName            null.String              `json:"shippingName"`
	ShippingCost            null.Float               `json:"shippingCost"`
	Seller                  Seller                   `json:"seller"`
	Payments                []Payments               `json:"payments"`
	Items                   []MaddenTransactionItems `json:"items"`
}

// MaddenStock holds stock items
type MaddenStock struct {
	Sku       string `json:"sku"`
	EAN       string `json:"ean"`
	BrandName string `json:"brandName"`
	Quantity  int    `json:"quantity"`
}

// MaddenPIM holds product data
type MaddenPIM struct {
	ProductName string `json:"productName"`
	VariantName string `json:"variantName"`
	Sku         string `json:"sku"`
	EAN         string `json:"ean"`
	BrandName   string `json:"brandName"`
	Category    string `json:"category"`
	Collection  string `json:"collection"`
	Season      string `json:"season"`
	Year        string `json:"year"`
	Size        string `json:"size"`
	Color       string `json:"color"`
	AgeGroup    string `json:"ageGroup"`
	Gender      string `json:"gender"`
	ModelNumber string `json:"ModelNumber"`
}

// MaddenBearer holds the madden token reponse
type MaddenBearer struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
}

// APIKey holds the Madden apiKey
type APIKey struct {
	AccountID string `json:"accountID"`
	APIKey    string `json:"apiKey"`
}

// GetAuth generates bearer token from Madden
func GetAuth(baseURL string, accountID string, key string) string {
	maddenKey := APIKey{
		AccountID: accountID,
		APIKey:    key,
	}

	maddenKeyJSON, _ := json.Marshal(maddenKey)

	// Get Madden Bearer Token
	_, maddenAuthResponse := Post(
		baseURL+"/auth",
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

// BatchPIM batches pim struct into given size
func BatchPIM(batchSize int, batchSlice []MaddenPIM) [][]MaddenPIM {

	batches := make([][]MaddenPIM, 0, (len(batchSlice)+batchSize-1)/batchSize)
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
