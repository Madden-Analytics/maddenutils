package maddenutils

import (
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/datatypes"
)

// MaddenTransactionItems holds transcation items from Madden
type MaddenTransactionItems struct {
	ProductName       string  `json:"productName"`
	VariantName       string  `json:"variantName"`
	Sku               string  `json:"sku"`
	Ean               string  `json:"ean"`
	VendorSKU         string  `json:"vendorSKU"`
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
	ProductType       int     `json:"productType,omitempty"`
	FulfillmentType   int     `json:"fulfillmentType"`
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
	ProductID         int            `json:"productID,omitempty"`
	VariantID         int            `json:"variantID,omitempty"`
	ProductName       string         `json:"productName"`
	VariantName       string         `json:"variantName"`
	Sku               string         `json:"sku"`
	EAN               string         `json:"ean"`
	BrandName         string         `json:"brandName"`
	Category          string         `json:"category"`
	Collection        string         `json:"collection"`
	Season            string         `json:"season"`
	Year              string         `json:"year"`
	Size              string         `json:"size"`
	Color             string         `json:"color"`
	AgeGroup          string         `json:"ageGroup"`
	Gender            string         `json:"gender"`
	ModelNumber       string         `json:"modelNumber"`
	StyleNumber       string         `json:"styleNumber"`
	VariantExternalID string         `json:"variantExternalID"`
	VariantKey        string         `json:"variantKey"`
	ProductType       int            `json:"productType"`
	ProductInfo       datatypes.JSON `json:"info"`
	VendorSKU         string         `json:"vendorSKU"`
}

// PIM - New type to hold pims data
type PIM struct {
	ProductName       string         `json:"productName"`
	VariantName       string         `json:"variantName"`
	SKU               string         `json:"sku"`
	VendorSKU         string         `json:"vendorSKU"`
	EAN               string         `json:"ean"`
	BrandName         string         `json:"brandName"`
	Category          string         `json:"category"`
	Collection        string         `json:"collection"`
	Season            string         `json:"season"`
	Year              string         `json:"year"`
	Size              string         `json:"size"`
	Color             string         `json:"color"`
	AgeGroup          string         `json:"ageGroup"`
	Gender            string         `json:"gender"`
	ModelNumber       string         `json:"modelNumber"`
	StyleNumber       string         `json:"styleNumber"`
	VariantExternalID string         `json:"variantExternalID"`
	Key               string         `json:"variantKey"`
	ProductType       int            `json:"productType"`
	ProductID         int            `json:"productID"`
	VariantID         int            `json:"variantID"`
	Info              datatypes.JSON `json:"info"`
	CostPrice         float64        `json:"costPrice"`
	CostPriceCurrency string         `json:"costPriceCurrency"`
	Active            bool           `json:"active"`
	Prices            Price          `json:"prices"`
}

// MaddenProduct holds main product
type MaddenProduct struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Modelnumber string            `json:"modelNumber"`
	Imageurl    string            `json:"imageUrl"`
	Brandid     int               `json:"brandID"`
	Brand       *MaddenBrands     `json:"brand,omitempty"`
	Categoryid  int               `json:"categoryID"`
	Category    *MaddenCategories `json:"category"`
	Activityid  int               `json:"activityID"`
	ProductType int               `json:"productType"`
	ProductInfo datatypes.JSON    `json:"info"`
	Variants    []*MaddenVariant  `json:"variants,omitempty"`
}

// Product - New updated product type
type Product struct {
	ID          int            `json:"id" example:"1"`
	Name        string         `json:"name" example:"Petter"`
	ModelNumber string         `json:"modelNumber" example:"123456"`
	BrandID     uint           `json:"-"`
	Brand       Brand          `json:"brand"`
	CategoryID  uint           `json:"-"`
	Category    Category       `json:"category"`
	ProductInfo datatypes.JSON `json:"info"`
	Variants    []Variant      `json:"variants"`
}

// Variant - New updated variant type
type Variant struct {
	ID                uint    `json:"id"`
	ProductID         uint    `json:"productID"`
	SKU               string  `json:"sku"`
	VendorSKU         string  `json:"vendorSKU"`
	EAN               string  `json:"ean"`
	Key               string  `json:"key"`
	Name              string  `json:"name"`
	Size              string  `json:"size"`
	Color             string  `json:"color"`
	Collection        string  `json:"collection"`
	Season            string  `json:"season"`
	Year              string  `json:"year"`
	AgeGroup          string  `json:"ageGroup"`
	Gender            string  `json:"gender"`
	Active            bool    `json:"active"`
	StyleNumber       string  `json:"styleNumber"`
	ExternalID        string  `json:"externalID"`
	Source            int     `json:"source"`
	CostPrice         float64 `json:"costPrice"`
	CostPriceCurrency string  `json:"costPriceCurrency"`
	Price             *Price  `json:"prices,omitempty"`
	Meta              Meta    `json:"meta"`
}

// Price - Prices assosciated with the variant
type Price struct {
	SalePrice    float64 `json:"price,omitempty"`
	RRPPrice     float64 `json:"retailPrice,omitempty"`
	CostPrice    float64 `json:"costPrice,omitempty"`
	CurrencyCode string  `json:"currencyCode,omitempty"`
}

// Meta - Read-only field with additional variant data
type Meta struct {
	BrandName     string         `json:"brandName"`
	BrandSynonyms []string       `json:"brandSynonyms"`
	ProductName   string         `json:"productName"`
	ProductType   int            `json:"productType"`
	CategoryName  string         `json:"categoryName"`
	ModelNumber   string         `json:"modelNumber"`
	ProductInfo   datatypes.JSON `json:"info"`
}

// Category - The category of the product
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// MaddenVariant holds variant data
type MaddenVariant struct {
	ID           int            `json:"id"`
	Sku          string         `json:"sku"`
	VendorSKU    string         `json:"vendorSKU"`
	Ean          string         `json:"ean"`
	Name         string         `json:"name"`
	Key          string         `json:"key"`
	Size         string         `json:"size"`
	Color        string         `json:"color"`
	ProductID    int            `json:"productID"`
	Product      *MaddenProduct `json:"product,omitempty"`
	Collection   string         `json:"collection"`
	Season       string         `json:"season"`
	Year         string         `json:"year"`
	AgeGroup     string         `json:"ageGroup"`
	Gender       string         `json:"gender"`
	StyleNumber  string         `json:"styleNumber"`
	Active       bool           `json:"active"`
	ProjectionID string         `json:"projectionID"`
	ExternalID   string         `json:"externalID"`
	Source       int            `json:"source"`
}

type MaddenStores struct {
	Name            string            `json:"name"`
	ExternalStoreID string            `json:"externalStoreId"`
	ChannelType     int               `json:"channelType"` //enum
	Adress          string            `json:"adress"`
	Adress2         string            `json:"adress2"`
	Zip             string            `json:"zip"`
	City            string            `json:"city"`
	State           string            `json:"state"`
	CountryCode     string            `json:"countryCode"`
	CurrencyCode    string            `json:"currencyCode"`
	Warehouses      []MaddenWarehouse `json:"warehouses"`
}

// Warehouse serves WarehouseItems too some Stores
type MaddenWarehouse struct {
	Name                string `json:"name"`
	WarehouseExternalID string `json:"externalId"`
}

// MaddenBrands hold brand data
type MaddenBrands struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Logourl  string   `json:"logoUrl"`
	Synonyms []string `json:"synonyms"`
	File     string   `json:"file"`
}

type Brand struct {
	ID       uint     `json:"id" example:"1"`
	Name     string   `json:"name" example:"Acme"`
	Synonyms []string `json:"synonyms"`
}

// MaddenCategories holds category data
type MaddenCategories struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Synonyms []string `json:"synonyms"`
	ParentID int      `json:"parent_id"`
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

// Transaction main holder of transaction items
type WholesaleTransaction struct {
	TransactionID           int                        `json:"id,omitempty"`
	TransactionDate         time.Time                  `json:"transactionDate"`
	CompletedAt             null.Time                  `json:"completed_at"`
	CancelledAt             null.Time                  `json:"cancelled_at"`
	AskedDeliveryDate       null.Time                  `json:"askedDeliveryDate"`
	CustomerName            string                     `json:"customerName"`
	CustomerExternalID      string                     `json:"externalCustomerID"`
	ExternalStoreID         string                     `json:"externalStoreId"`
	ExternalTransactionID   string                     `json:"externalTransactionId"`
	MoneyFinalNet           float64                    `json:"moneyFinalNet"`
	MoneyFinalVat           float64                    `json:"moneyFinalVat"`
	MoneyTotalGrossRoundOff float64                    `json:"MoneyTotalGrossRoundOff"`
	CurrencyCode            string                     `json:"currencyCode"`
	ExternalRowID           string                     `json:"externalRowId"`
	Items                   []WholesaleTransactionItem `json:"items"`
	Market                  string                     `json:"market"`
	CartDiscount            float64                    `json:"cartDiscount"`
	ShippingName            null.String                `json:"shippingName"`
	ShippingCost            null.Float                 `json:"shippingCost"`
}

// TransactionItem or event log
type WholesaleTransactionItem struct {
	ProductName       string                   `json:"productName"`
	SKU               string                   `json:"sku"`
	VendorSKU         string                   `json:"vendorSKU"`
	EAN               string                   `json:"ean"`
	BrandName         string                   `json:"brandName"`
	Category          string                   `json:"category"`
	Collection        string                   `json:"collection"`
	Season            string                   `json:"season"`
	Year              string                   `json:"year"`
	Size              string                   `json:"size"`
	Color             string                   `json:"color"`
	AgeGroup          string                   `json:"ageGroup"`
	Gender            string                   `json:"gender"`
	ModelNumber       string                   `json:"ModelNumber"`
	Quantity          int                      `json:"quantity"`
	VatValue          int                      `json:"vatValue"`
	MoneyItemTotalNet float64                  `json:"moneyItemTotalNet"`
	MoneyItemTotalVat float64                  `json:"moneyItemTotalVat"`
	CostPrice         float64                  `json:"costPrice"`
	CostPriceCurrency string                   `json:"costPriceCurrency"`
	MoneyDiscount     float64                  `json:"moneyDiscount"`
	Supplier          string                   `json:"supplier"`
	ExternalID        string                   `json:"externalID"`
	StyleNumber       string                   `json:"styleNumber"`
	VariantName       string                   `json:"variantName"`
	ProductType       int                      `json:"productType"`
	FulfillmentType   int                      `json:"fulfillmentType"`
	ShipmentEvent     []WholesaleShipmentEvent `json:"shipments"`
}

// ShippingEvent holds each TIW shipment
type WholesaleShipmentEvent struct {
	Quantity            int       `json:"quantity"`
	Date                time.Time `json:"shipmentDate"`
	ShipmentID          string    `json:"shipmentID"`
	ExternalWarehouseID string    `json:"externalWarehouseID"`
}
