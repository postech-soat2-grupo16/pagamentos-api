package api

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joaocampari/postech-soat2-grupo16/gateways/message"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/controllers"
	"github.com/joaocampari/postech-soat2-grupo16/external"
	apimercadopago "github.com/joaocampari/postech-soat2-grupo16/gateways/api/mercadopago"
	apipedido "github.com/joaocampari/postech-soat2-grupo16/gateways/api/pedido"
	pagamentogateway "github.com/joaocampari/postech-soat2-grupo16/gateways/db/pagamento"
	"github.com/joaocampari/postech-soat2-grupo16/usecases/pagamento"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

const (
	// This is only a test token, not a real one and will be removed in the future replacing by a secret service.
	authToken     = "TEST-8788837371574102-082018-c29a1c5da797dbf70a8c99b842da2850-144255706"
	awsAccessKey  = "xxx"
	awsSecretKey  = "xxx"
	awsRegion     = "us-east-1"
	apiPedidosURL = "https://dk3sau3iu3.execute-api.us-east-1.amazonaws.com"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupQueue() *sqs.SQS {
	return external.GetSqsClient()
}

func SetupRouter(db *gorm.DB, queue *sqs.SQS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db, queue)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB, queue *sqs.SQS) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Gateways
	mercadoPagoGateway := apimercadopago.NewGateway(authToken)
	queueGateway := message.NewGateway(queue)
	pedidosApiGateway := apipedido.NewGateway(apiPedidosURL)
	pagamentoGateway := pagamentogateway.NewGateway(orm)
	// Use cases
	pagamentoUseCase := pagamento.NewUseCase(pagamentoGateway, mercadoPagoGateway, queueGateway, pedidosApiGateway)
	// Handlers
	_ = controllers.NewPagamentoController(pagamentoUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
