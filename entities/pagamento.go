package entities

import (
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pedido"
	"time"

	"golang.org/x/exp/slices"
)

type Pagamento struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	PedidoID  uint32    `gorm:"not null" json:"-"`
	Amount    float64   `gorm:"not null" json:"amount"`
	Status    string    `gorm:"not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Pagamento) IsStatusValid() bool {
	status := []string{pedido.StatusPagamentoAprovado, pedido.StatusPagamentoNegado}
	return slices.Contains(status, p.Status)
}
