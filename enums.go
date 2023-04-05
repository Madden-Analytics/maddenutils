package maddenutils

import "strconv"

type TransChannel uint8

const (
	TransChannel_D2C TransChannel = iota + 1
	TransChannel_B2B
)

func (t TransChannel) String() string {
	return strconv.Itoa(int(t))
}

type TransType uint8

const (
	TransType_SALE TransType = iota + 1
	TransType_RETURN
)

func (t TransType) String() string {
	return strconv.Itoa(int(t))
}

type TransStatementType uint8

const (
	TransStatementType_FEE TransStatementType = iota + 1
	TransStatementType_VOUCHER
	TransStatementType_SHIPPING
	TransStatementType_SERVICE
	TransStatementType_CART_DISCOUNT
)

func (t TransStatementType) String() string {
	return strconv.Itoa(int(t))
}

type SKUStatus uint8

const (
	SKUStatus_Inactive SKUStatus = 1
	SKUStatus_Active   SKUStatus = 2
	SKUStatus_Outgoing SKUStatus = 3
	SKUStatus_Draft    SKUStatus = 4
)

func (t SKUStatus) String() string {
	return strconv.Itoa(int(t))
}

type DeliveryStatus uint8

const (
	DeliveryStatus_NotDelivered       DeliveryStatus = 0
	DeliveryStatus_PartiallyDelivered DeliveryStatus = 1
	DeliveryStatus_FullyDelivered     DeliveryStatus = 2
	DeliveryStatus_Cancelled          DeliveryStatus = 3
	DeliveryStatus_Draft              DeliveryStatus = 4
	DeliveryStatus_Open               DeliveryStatus = 5
)

func (t DeliveryStatus) String() string {
	return strconv.Itoa(int(t))
}
