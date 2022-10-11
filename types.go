package maddenutils

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

type Transaction struct {
	TransactionDate       time.Time              `json:"transactionDate"`
	TransactionID         int                    `json:"id,omitempty"`
	CompletedAt           null.Time              `json:"completed_at"`
	CancelledAt           null.Time              `json:"cancelled_at"`
	ExternalStoreID       string                 `json:"externalStoreID"`
	ExternalTransactionID string                 `json:"externalTransactionID"`
	ExternalParentID      *string                `json:"externalParentID"`
	CurrencyCode          string                 `json:"currencyCode"`
	ExternalRowID         string                 `json:"externalRowId"`
	Market                string                 `json:"market"`
	TransactionType       TransactionType        `json:"transactionType"`
	TransactionChannel    TransactionChannel     `json:"transactionChannel"`
	CustomerName          string                 `json:"customerName"`
	ExternalCustomerID    string                 `json:"externalCustomerID"`
	DeliveryDate          null.Time              `json:"deliveryDate"`
	Statements            []TransactionStatement `json:"statements"`
	Items                 []TransactionItem      `json:"items"`
}

type TransactionStatement struct {
	Name          string                   `json:"name"`
	StatementType TransactionStatementType `json:"type"`
	PriceNet      float64                  `json:"priceNet"`
	PriceVat      float64                  `json:"priceVat"`
}

type TransactionItem struct {
	ProductName        string                    `json:"productName"`
	VariantName        string                    `json:"variantName"`
	SKU                string                    `json:"sku"`
	EAN                string                    `json:"ean"`
	Key                string                    `json:"key"`
	VendorSKU          string                    `json:"vendorSKU"`
	BrandName          string                    `json:"brandName"`
	Supplier           string                    `json:"supplier,omitempty"`
	Category           string                    `json:"category"`
	Collection         string                    `json:"collection"`
	Season             string                    `json:"season"`
	Year               string                    `json:"year"`
	Size               string                    `json:"size"`
	Color              string                    `json:"color"`
	AgeGroup           string                    `json:"ageGroup"`
	Gender             string                    `json:"gender"`
	ProductGroupID     string                    `json:"productGroupID"`
	VariantGroupID     string                    `json:"variantGroupID"`
	Quantity           int                       `json:"quantity"`
	VatValue           int                       `json:"vatValue"`
	MoneyItemTotalNet  float64                   `json:"moneyItemTotalNet"`
	MoneyItemTotalVat  float64                   `json:"moneyItemTotalVat"`
	CostPrice          float64                   `json:"costPrice"`
	CostPriceCurrency  string                    `json:"costPriceCurrency"`
	MoneyDiscount      float64                   `json:"moneyDiscount"`
	ExternalID         string                    `json:"externalID"`
	ProductType        int                       `json:"productType,omitempty"`
	FulfillmentType    int                       `json:"fulfillmentType"`
	Shipments          []TransactionItemShipment `json:"shipments"`
	DeliveryWindowName *string                   `json:"deliveryWindow"`
}

type TransactionItemShipment struct {
	Quantity            int       `json:"quantity"`
	Date                time.Time `json:"shipmentDate"`
	ShipmentID          string    `json:"shipmentID"`
	ExternalWarehouseID string    `json:"externalWarehouseID"`
}

// DeliveryEvent holds each POI deliveries
type DeliveryEvent struct {
	Quantity   int       `json:"quantity"`
	Date       time.Time `json:"deliveryDate"`
	DeliveryID string    `json:"deliveryID"`
}

// PurchaseOrderItem holds each purchase order row
type PurchaseOrderItem struct {
	ProductName        string          `json:"productName"`
	Sku                string          `json:"sku"`
	EAN                string          `json:"ean"`
	Key                string          `json:"key"`
	BrandName          string          `json:"brandName"`
	Quantity           int             `json:"quantity"`
	TotalPriceNet      float64         `json:"totalPriceNet"`
	TotalPriceVat      float64         `json:"totalPriceVat"`
	TotalPriceCurrency string          `json:"totalPriceCurrency"`
	Supplier           null.String     `json:"supplier"`
	ExternalRowID      string          `json:"externalRowID"`
	Deliveries         []DeliveryEvent `json:"deliveries"`
}

// PurchaseOrder main holder of purchase order items
type PurchaseOrder struct {
	ID                  int                 `json:"id,omitempty"`
	OrderDate           time.Time           `json:"orderDate" gorm:"not null"`
	DeliveryDate        null.Time           `json:"deliveryDate"`
	DeliveryStatus      int                 `json:"deliveryStatus"` // 0 - Not Delivered, 10 - Partially Delivered, 20 - Fully Delivered
	WarehouseExternalID string              `json:"externalWarehouseId"`
	ExternalPurchaseNo  string              `json:"externalPurchaseNo"`
	ExternalComment     string              `json:"externalComment"`
	Items               []PurchaseOrderItem `json:"items"`
}

type Seller struct {
	Name string `json:"name"`
}

type Payment struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// MaddenStock holds stock items
type Stock struct {
	SKU               string `json:"sku"`
	EAN               string `json:"ean"`
	Key               string `json:"key"`
	AvailableQuantity *int   `json:"availableQuantity"`
	PhysicalQuantity  *int   `json:"physicalQuantity"`
	Supplier          string `json:"supplier,omitempty"`
}

// Products - Holds all SKU Data
type Product struct {
	ID                int            `json:"id"`
	ProductName       string         `json:"productName"`
	VariantName       string         `json:"variantName"`
	SKU               string         `json:"sku"`
	VendorSKU         string         `json:"vendorSKU"`
	EAN               string         `json:"ean"`
	BrandName         string         `json:"brandName" example:"Acme"`
	CategoryName      string         `json:"categoryName"`
	Collection        string         `json:"collection"`
	Season            string         `json:"season"`
	Year              string         `json:"year"`
	Size              string         `json:"size"`
	Color             string         `json:"color"`
	AgeGroup          string         `json:"ageGroup"`
	Gender            string         `json:"gender"`
	ProductGroupID    string         `json:"productGroupID"`
	VariantGroupID    string         `json:"variantGroupID"`
	ExternalID        string         `json:"externalID"`
	Key               string         `json:"key"`
	SkuSynonym        []string       `json:"skuSynonyms"`
	ProductType       int            `json:"productType"`
	ProductID         uint           `json:"productID"`
	VariantID         uint           `json:"variantID"`
	Info              datatypes.JSON `json:"info"`
	CostPrice         float64        `json:"costPrice"`
	CostPriceCurrency string         `json:"costPriceCurrency"`
	Price             float64        `json:"price"`
	PriceCurrency     string         `json:"priceCurrency"`
	Active            bool           `json:"active"`
	LeadTime          int            `json:"leadTime"`
	MOQ               int            `json:"moq"`
	Supplier          string         `json:"supplier"`
}

type Store struct {
	Name            string      `json:"name"`
	ExternalStoreID string      `json:"externalStoreId"`
	ChannelType     int         `json:"channelType"` //enum
	Adress          string      `json:"adress"`
	Adress2         string      `json:"adress2"`
	Zip             string      `json:"zip"`
	City            string      `json:"city"`
	State           string      `json:"state"`
	CountryCode     string      `json:"countryCode"`
	CurrencyCode    string      `json:"currencyCode"`
	Warehouses      []Warehouse `json:"warehouses"`
}

// Warehouse serves WarehouseItems too some Stores
type Warehouse struct {
	Name                string `json:"name"`
	WarehouseExternalID string `json:"externalId"`
}

type SKU struct {
	ID                int            `json:"id"`
	ProductName       string         `json:"productName"`
	ProductGroupID    string         `json:"productGroupID"`
	Brand             Brand          `json:"brand"`
	Category          Category       `json:"category"`
	ProductType       int            `json:"productType"`
	Info              datatypes.JSON `json:"info"`
	Sku               string         `json:"sku"`
	VendorSKU         string         `json:"vendorSKU"`
	Ean               string         `json:"ean"`
	Key               string         `json:"key"`
	SkuSynonyms       []SkuSynonym   `json:"skuSynonyms"`
	VariantName       string         `json:"variantName"`
	Size              string         `json:"size"`
	Color             string         `json:"color"`
	Collection        string         `json:"collection"`
	Season            string         `json:"season"`
	Year              string         `json:"year"`
	AgeGroup          string         `json:"ageGroup"`
	Gender            string         `json:"gender"`
	Active            bool           `json:"active"`
	VariantGroupID    string         `json:"variantGroupID"`
	ExternalID        string         `json:"externalID"`
	Source            int            `json:"source"`
	CostPrice         int            `json:"costPrice"`
	CostPriceCurrency string         `json:"costPriceCurrency"`
	Price             int            `json:"price"`
	PriceCurrency     string         `json:"priceCurrency"`
	Moq               int            `json:"moq"`
	LeadTime          int            `json:"leadTime"`
	Supplier          string         `json:"supplier"`
}

// Brand - The Brand and synonyms of the product
type Brand struct {
	ID       uint     `json:"id" example:"1"`
	Name     string   `json:"name" example:"Acme"`
	Synonyms []string `json:"synonyms"`
}

// Category - The category of the product
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Synonyms - Hold the sku_synonyms of a product
type SkuSynonym struct {
	ID            int    `json:"id"`
	FromKey       string `json:"FromKey"`
	FromAccountID string `json:"FromAccountID"`
	SkuID         int    `json:"skuId"`
	ToAccountID   string `json:"ToAccountID"`
}

// MaddenBearer holds the madden token reponse
type Bearer struct {
	TokenType   string `json:"tokenType"`
	AccessToken string `json:"accessToken"`
}

// APIKey holds the Madden apiKey
type APIKey struct {
	AccountID string `json:"accountID"`
	APIKey    string `json:"apiKey"`
}

type StoreTrafficEvent struct {
	ID                    int       `json:"id,omitempty"`
	EventTime             time.Time `json:"eventTime"`
	Name                  string    `json:"name,omitempty"`
	StoreID               string    `json:"storeID,omitempty"`
	StoreTrafficCounterID string    `json:"storeTrafficCounterID,omitempty"`
	IncomingTraffic       int       `json:"incomingTraffic"`
	OutgoingTraffic       int       `json:"outgoingTraffic"`
}
