package pagamento

import (
	"time"
)

const (
	// Status do pagamento
	StatusPagamentoIniciado = "INICIADO"
	StatusPagamentoAprovado = "APROVADO"
	StatusPagamentoNegado   = "NEGADO"
)

type Pedido struct {
	ID        uint32 `json:"id"`
	Items     []Item `json:"items"`
	Notes     string `json:"notes"`
	ClienteID uint32 `json:"clienteId"`
	Status    string `json:"status"`
}

type Item struct {
	ItemID   uint32  `json:"itemId"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type QRCode struct {
	QRCode string `json:"qr_code"`
}

type Pagamento struct {
	ID        uint32    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentCallback struct {
	Id          int       `json:"id"`
	LiveMode    bool      `json:"live_mode"`
	Type        string    `json:"type"`
	DateCreated time.Time `json:"date_created"`
	UserId      int       `json:"user_id"`
	ApiVersion  string    `json:"api_version"`
	Action      string    `json:"action"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
}

func (p *Pedido) GetAmount() float64 {
	var amount float64
	for _, item := range p.Items {
		amount += float64(item.Price) * float64(item.Quantity)
	}
	return amount
}
