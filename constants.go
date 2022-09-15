package maddenutils

import "strconv"

type TransactionType uint8

const (
	D2C TransactionType = iota + 1
	D2B
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
