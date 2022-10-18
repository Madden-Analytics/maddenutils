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