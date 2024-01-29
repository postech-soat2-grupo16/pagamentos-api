package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/joaocampari/postech-soat2-grupo16/adapters/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"github.com/joaocampari/postech-soat2-grupo16/util"

	"github.com/go-chi/chi/v5"
)

type PagamentoController struct {
	useCase interfaces.PagamentoUseCase
}

func NewPagamentoController(useCase interfaces.PagamentoUseCase, r *chi.Mux) *PagamentoController {
	controller := PagamentoController{useCase: useCase}
	r.Route("/pagamentos", func(r chi.Router) {
		r.Get("/{idPedido}/qr-code", controller.GetQRCodeByPedidoID())
		r.Post("/mp-webhook", controller.PaymentWebhookCreate())
		r.Get("/{idPagamento}", controller.GetPaymentById())
		r.Get("/health", controller.Health())
	})
	return &controller
}

// @Summary	Get QR Code pagamento
//
// @Tags		Payments
//
// @ID			get-qr-code-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	pagamento.QRCode
// @Failure	404
// @Router		/pagamentos/{id}/qr-code [get]
func (c *PagamentoController) GetQRCodeByPedidoID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "idPedido")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("Criação de Pagamento para o Pedido %s\n", id)
		payment, err := c.useCase.CreatePayment(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Criação de QR Code para o Pedido %s\n", payment.PedidoID)
		qrCodeStr, err := c.useCase.CreateQRCode(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if qrCodeStr == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		qrCode := pagamento.QRCode{
			QRCode: *qrCodeStr,
		}
		log.Printf("QR CODE Criado para o Pedido %s\n", payment.PedidoID)

		_, err = c.useCase.UpdatePaymentStatusByPaymentID(payment.ID)

		json.NewEncoder(w).Encode(qrCode)
	}
}

// @Summary	Receive payment callback from MercadoPago
//
// @Tags		Payments
//
// @ID			receive-callback
// @Produce	json
// @Param		data	body		pagamento.PaymentCallback	true	"Order data"
// @Success	200		{object}	pagamento.Pagamento
// @Failure	400
// @Router		/pagamentos/mp-webhook [post]
func (c *PagamentoController) PaymentWebhookCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payment pagamento.PaymentCallback
		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updatedPayment, err := c.useCase.UpdatePaymentStatusByPaymentID(uint32(payment.Id))
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if updatedPayment == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedPayment)
	}

}

// @Summary	Get payment by ID
//
// @Tags		Payments
//
// @ID			get-payment-by-id
// @Produce	json
// @Param		id	path		string	true	"Payment ID"
// @Success	200	{object}	pagamento.Pagamento
// @Failure	404
// @Router		/pagamentos/{id} [get]
func (c *PagamentoController) GetPaymentById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "idPagamento")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payment, err := c.useCase.GetByID(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if payment == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(payment)
	}
}

// @Summary	Health check
//
// @Tags		Payments
//
// @ID			health-check
// @Produce	json
// @Success	200	{object}	string
// @Router		/pagamentos/health [get]
func (c *PagamentoController) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
	}
}
