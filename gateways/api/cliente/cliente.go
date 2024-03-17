package cliente

import (
	"encoding/json"
	"fmt"
	"github.com/joaocampari/postech-soat2-grupo16/adapters/cliente"
	"net/http"
)

type ClientesAPIRepository struct {
	ApiURL string
}

func NewGateway(apiURL string) *ClientesAPIRepository {
	return &ClientesAPIRepository{
		apiURL,
	}
}

func (p *ClientesAPIRepository) GetByID(clienteID uint32) (*cliente.Cliente, error) {
	url := fmt.Sprintf("%s%s%s", p.ApiURL, "/clientes/", clienteID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("erro na resposta da API Clientes: %s", resp.Status)
	}

	var clienteResponse cliente.Cliente
	err = json.NewDecoder(resp.Body).Decode(&clienteResponse)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON Clientes: %v", err)
	}

	return &clienteResponse, nil
}
