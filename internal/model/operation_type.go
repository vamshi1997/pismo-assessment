package model

type OperationType int

const (
	NormalPurchase       OperationType = iota + 1 // 1
	PurchaseInstallments                          // 2
	Withdrawal                                    // 3
	CreditVoucher                                 // 4
)

// Map for converting enum to string
var transactionTypeToString = map[OperationType]string{
	NormalPurchase:       "Normal Purchase",
	PurchaseInstallments: "Purchase with Installments",
	Withdrawal:           "Withdrawal",
	CreditVoucher:        "Credit Voucher",
}

// String method for pretty printing
func (t OperationType) String() string {
	return transactionTypeToString[t]
}
