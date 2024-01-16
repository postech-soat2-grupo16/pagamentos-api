package pagamento

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"gorm.io/gorm"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (p *Repository) UpdatePaymentStatusByPaymentID(pagamentoID uint32, status string) (*entities.Pagamento, error) {
	pagamento := entities.Pagamento{
		ID:     pagamentoID,
		Status: status,
	}
	result := p.repository.Model(&pagamento).Where("pagamento_id = ?", pagamentoID).Update("status", status)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pagamento, nil
}
