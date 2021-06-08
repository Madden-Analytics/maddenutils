package maddenutils

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// MaddenTransactionItems holds transcation items from Madden
type MaddenTransactionItems struct {
	ProductName       string  `json:"productName"`
	VariantName       string  `json:"variantName"`
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
	StyleNumber       string  `json:"styleNumber"`
	Quantity          int     `json:"quantity"`
	VatValue          int     `json:"vatValue"`
	MoneyItemTotalNet float64 `json:"moneyItemTotalNet"`
	MoneyItemTotalVat float64 `json:"moneyItemTotalVat"`
	CostPrice         float64 `json:"costPrice"`
	CostPriceCurency  string  `json:"costPriceCurrency"`
	MoneyDiscount     float64 `json:"moneyDiscount"`
	ExternalID        string  `json:"externalID"`
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
	Seller                  seller                   `json:"seller"`
	Payments                []payments               `json:"payments"`
	Items                   []MaddenTransactionItems `json:"items"`
}

type seller struct {
	Name string `json:"name"`
}

type payments struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// MaddenStock holds stock items
type MaddenStock struct {
	Sku       string `json:"sku"`
	EAN       string `json:"ean"`
	BrandName string `json:"brandName"`
	Quantity  int    `json:"quantity"`
	Supplier  string `json:"supplier,omitempty"`
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
	ModelNumber string `json:"modelNumber"`
	StyleNumber string `json:"styleNumber"`
}

// MaddenVariant holds variant data
type MaddenVariant struct {
	ID        int    `json:"id"`
	Sku       string `json:"sku"`
	Ean       string `json:"ean"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	ProductID int    `json:"productID"`
	Product   struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		ModelNumber string      `json:"modelNumber"`
		ImageURL    string      `json:"imageUrl"`
		BrandID     int         `json:"brandID"`
		CategoryID  int         `json:"categoryID"`
		ActivityID  interface{} `json:"activityID"`
	} `json:"product"`
	Collection  string `json:"collection"`
	Season      string `json:"season"`
	Year        string `json:"year"`
	AgeGroup    string `json:"ageGroup"`
	Gender      string `json:"gender"`
	StyleNumber string `json:"styleNumber"`
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