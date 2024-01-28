package pagamento

import (
	"github.com/joaocampari/postech-soat2-grupo16/entities"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Repository {
	return &Repository{repository: repository}
}

func (p *Repository) Update(pagamento entities.Pagamento) (*entities.Pagamento, error) {
	result := p.repository.Updates(&pagamento)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &pagamento, nil
}

func (p *Repository) CreatePayment(pagamento entities.Pagamento) (*entities.Pagamento, error) {
	result := p.repository.Create(&pagamento)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &pagamento, nil
}

func (p *Repository) GetByID(paymentID uint32) (*entities.Pagamento, error) {
	pagamento := entities.Pagamento{
		ID: paymentID,
	}
	result := p.repository.First(&pagamento)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pagamento, nil
}
