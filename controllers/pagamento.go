package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	pedido2 "github.com/joaocampari/postech-soat2-grupo16/adapters/pedido"
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
		r.Get("/{id}/qr-code", controller.GetQRCodeByPedidoID)
		r.Post("/mp-webhook", controller.PaymentWebhookCreate)
	})
	return &controller
}

// @Summary	Get QR Code pedido
//
// @Tags		Orders
//
// @ID			get-qr-code-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	pedido2.Pedido
// @Failure	404
// @Router		/pedidos/{id}/qr-code [get]
func (c *PagamentoController) GetQRCodeByPedidoID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	qrCodeStr, err := c.useCase.CreateQRCode(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if qrCodeStr == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	qrCode := pedido2.QRCode{
		QRCode: *qrCodeStr,
	}
	json.NewEncoder(w).Encode(qrCode)
}

// @Summary	Receive payment callback from MercadoPago
//
// @Tags		Orders
//
// @ID			receive-callback
// @Produce	json
// @Param		data	body		pedido2.PaymentCallback	true	"Order data"
// @Success	200		{object}	pedido2.Pedido
// @Failure	400
// @Router		/pedidos/mp-webhook [post]
func (c *PagamentoController) PaymentWebhookCreate(w http.ResponseWriter, r *http.Request) {
	var payment pedido2.PaymentCallback
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(payment.Data.ID, 10, 32)
	pagamento, err := c.useCase.UpdatePaymentStatusByPaymentID(uint32(id))
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pagamento == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pagamento)
}
