package driver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/database"
	clienteHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	itemHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	pedidoHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/pedido"
	clienterepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/cliente"
	itemrepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/item"
	pedidorepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/item"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/pedido"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dialector := database.GetPostgresDialector()
	db := database.NewORM(dialector)

	return db
}

func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Handler
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Repositories
	pedidoRepository := pedidorepo.NewPedidoRepository(orm)
	clienteRepository := clienterepo.NewClienteRepository(orm)
	itemRepository := itemrepo.NewItemRepository(orm)
	// Use cases
	itemUseCase := item.NewItemUseCase(itemRepository)
	pedidoUseCase := pedido.NewUseCase(pedidoRepository)
	clienteUseCase := cliente.NewClienteUseCase(clienteRepository)
	// Handlers
	_ = itemHandler.NewHandler(itemUseCase, r)
	_ = pedidoHandler.NewHandler(pedidoUseCase, r)
	_ = clienteHandler.NewHandler(clienteUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
