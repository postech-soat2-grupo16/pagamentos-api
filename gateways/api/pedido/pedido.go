package pedido

import (
	"encoding/json"
	"fmt"
	"github.com/joaocampari/postech-soat2-grupo16/adapters/pedido"
	"net/http"
)

type PedidosAPIRepository struct {
	ApiURL string
}

func NewGateway(apiURL string) *PedidosAPIRepository {
	return &PedidosAPIRepository{
		apiURL,
	}
}

func (p *PedidosAPIRepository) GetByID(pedidoID uint32) (*pedido.Pedido, error) {
	apiURL := fmt.Sprintf("%s%s%s", p.ApiURL, "/pedidos/", pedidoID)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("erro na resposta da API: %s", resp.Status)
	}

	var pedidoResponse pedido.Pedido
	err = json.NewDecoder(resp.Body).Decode(&pedidoResponse)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	return &pedidoResponse, nil
}
