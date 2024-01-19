package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
)

func TestGetPedidos(t *testing.T) {
	t.Run("given_existing_pedido_id_should_return_pagamento_details", func(t *testing.T) {
		orderID := 3

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d/pagamentos/status", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %s", res.Status)
		}

		var response entities.Pagamento
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response.Status != "APROVADO" {
			t.Fatalf("expected status APROVADO; got %s", response.Status)
		}
	})

	t.Run("given_nonexistent_pedido_id_should_return_404", func(t *testing.T) {
		orderID := 999

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status Not Found; got %s", res.Status)
		}
	})
}

func TestSavePedidos(t *testing.T) {
	t.Run("given_valid_pedido_should_create_new_pedido", func(t *testing.T) {
		newOrder := pagamento.Pedido{
			Items:     []pagamento.Item{{ItemID: 1, Quantity: 2}, {ItemID: 2, Quantity: 3}},
			Notes:     "Novo pagamento",
			ClienteID: 1,
		}

		jsonOrder, err := json.Marshal(newOrder)
		if err != nil {
			t.Fatalf("could not marshal pagamento: %v", err)
		}

		// Cria uma requisição POST com o JSON do novo pagamento no corpo
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/pedidos", baseURL), bytes.NewBuffer(jsonOrder))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected status Created; got %s", res.Status)
		}
	})

	t.Run("given_nonexisting_pedido_id_should_return_404_when_updating", func(t *testing.T) {
		newNote := "Pedido atualizado"
		orderID := 999
		orderUpdated := pagamento.Pedido{
			Items:     []pagamento.Item{{ItemID: 1, Quantity: 5}, {ItemID: 2, Quantity: 3}},
			ClienteID: 1,
			Notes:     newNote,
		}

		jsonOrder, err := json.Marshal(orderUpdated)
		if err != nil {
			t.Fatalf("could not marshal pagamento: %v", err)
		}

		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), bytes.NewBuffer(jsonOrder))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status not found; got %s", res.Status)
		}
	})

	t.Run("given_existing_pedido_id_should_delete_pedido", func(t *testing.T) {
		orderID := 1

		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("expected status No Content; got %s", res.Status)
		}

		req, err = http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status NOT FOUND; got %s", res.Status)
		}
	})
}
