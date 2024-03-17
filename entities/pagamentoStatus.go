package entities

type PagamentoStatus string

const (
	// Status do pagamento
	StatusPagamentoIniciado     PagamentoStatus = "INICIADO"
	StatusPagamentoQRCodeCriado                 = "QR_CODE_CRIADO"
	StatusPagamentoQRCodeErro                   = "QR_CODE_ERRO"
	StatusPagamentoAprovado                     = "APROVADO"
	StatusPagamentoRecusado                     = "RECUSADO"
)
