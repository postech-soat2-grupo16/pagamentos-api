package pedido

import (
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
)

type PedidosAPIRepository struct {
	ApiURL string
}

func NewGateway(apiURL string) *PedidosAPIRepository {
	return &PedidosAPIRepository{
		apiURL,
	}
}

func (p *PedidosAPIRepository) GetByID(pedidoID uint32) (*pagamento.Pedido, error) {
	/*url := fmt.Sprintf("/%s%s%s", p.ApiURL, "/pedidos/", strconv.Itoa(int(pedidoID)))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("erro na resposta da API: %s", resp.Status)
	}

	var pedidoResponse pagamento.Pedido
	err = json.NewDecoder(resp.Body).Decode(&pedidoResponse)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}*/
	// Objeto pagamento para teste
	pedidoResponse := pagamento.Pedido{
		Items: []pagamento.Item{{
			ItemID:   1,
			Name:     "teste",
			Price:    10,
			Quantity: 2,
		}},
		Notes:     "Novo pagamento",
		ClienteID: 1,
	}

	return &pedidoResponse, nil
}
