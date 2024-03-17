package mercadoPago

import (
	"time"
)

type Payment struct {
	ID                int     `json:"id"`
	Status            string  `json:"status"`
	StatusDetail      string  `json:"status_detail"`
	PaymentTypeID     string  `json:"payment_type_id"`
	PaymentMethodID   string  `json:"payment_method_id"`
	TransactionAmount float64 `json:"transaction_amount"`
	Installments      int     `json:"installments"`
	Description       string  `json:"description"`
	Capture           bool    `json:"capture"`
	ExternalReference string  `json:"external_reference"`
}

type WalletPayment struct {
	TransactionAmount float64 `json:"transaction_amount"`
	Description       string  `json:"description"`
	ExternalReference string  `json:"external_reference"`
}

type Disbursement struct {
	CollectorID string `json:"collector_id"`
}

type Payer struct {
	ID int `json:"id"`
}

type Data struct {
	ID              int            `json:"id"`
	Status          string         `json:"status"`
	WalletPayment   WalletPayment  `json:"wallet_payment"`
	Payments        []Payment      `json:"payments"`
	Disbursements   []Disbursement `json:"disbursements"`
	Payer           Payer          `json:"payer"`
	SiteID          string         `json:"site_id"`
	BinaryMode      bool           `json:"binary_mode"`
	DateCreated     time.Time      `json:"date_created"`
	DateLastUpdated time.Time      `json:"date_last_updated"`
}
