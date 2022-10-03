package maddenutils

import "strconv"

type TransactionChannel uint8

const (
	D2C TransactionChannel = iota + 1
	B2B
)

func (t TransactionChannel) String() string {
	return strconv.Itoa(int(t))
}

type TransactionType uint8

const (
	Sale TransactionType = iota + 1
	Return
)

func (t TransactionType) String() string {
	return strconv.Itoa(int(t))
}

type TransactionStatementType uint8

const (
	Fee TransactionStatementType = iota + 1
	Voucher
	Shipping
	Service
	CartDiscount
)

func (t TransactionStatementType) String() string {
	return strconv.Itoa(int(t))
}
