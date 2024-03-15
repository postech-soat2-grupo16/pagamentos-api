package entities

import (
	"time"

	"golang.org/x/exp/slices"
)

type Pagamento struct {
	ID        uint32          `gorm:"primary_key;auto_increment" json:"id"`
	PedidoID  string          `gorm:"not null" json:"pedido_id"`
	Amount    float64         `gorm:"not null" json:"amount"`
	Status    PagamentoStatus `gorm:"not null" json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (p *Pagamento) IsStatusValid() bool {
	status := []PagamentoStatus{StatusPagamentoAprovado, StatusPagamentoRecusado}
	return slices.Contains(status, p.Status)
}
