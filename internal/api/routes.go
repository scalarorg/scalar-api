package api

import (
	"github.com/go-chi/chi"
	_ "github.com/scalarorg/xchains-api/docs"
	"github.com/scalarorg/xchains-api/internal/api/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *Server) SetupRoutes(r *chi.Mux) {
	handlers := a.handlers
	r.Get("/healthcheck", registerHandler(handlers.HealthCheck))
	r.Get("/v1/global-params", registerHandler(handlers.GetScalarGlobalParams))

	registerParamsHandler(r, handlers)
	registerDAppHandler(r, handlers)
	registerGmpHandler(r, handlers)
	registerTransferHandler(r, handlers)
	registerVaultHandler(r, handlers)
	registerCustodialHandler(r, handlers)
	registerValidatorHandler(r, handlers)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
}
func registerParamsHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Get("/v1/params/covenant", registerHandler(handlers.GetCovenantParams))
}
func registerDAppHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Get("/v1/dApp", registerHandler(handlers.GetDApp))
	r.Post("/v1/dApp", registerHandler(handlers.CreateDApp))
	r.Put("/v1/dApp", registerHandler(handlers.UpdateDApp))
	r.Patch("/v1/dApp", registerHandler(handlers.ToggleDApp))
	r.Delete("/v1/dApp", registerHandler(handlers.DeleteDApp))
}
func registerGmpHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/gmp/GMPStats", registerHandler(handlers.GMPStats))
	r.Post("/v1/gmp/GMPStatsAVGTimes", registerHandler(handlers.GMPStatsAVGTimes))
	r.Post("/v1/gmp/GMPChart", registerHandler(handlers.GMPChart))
	r.Post("/v1/gmp/GMPCumulativeVolume", registerHandler(handlers.GMPCumulativeVolume))
	r.Post("/v1/gmp/GMPTotalVolume", registerHandler(handlers.GMPTotalVolume))
	r.Post("/v1/gmp/GMPTotalFee", registerHandler(handlers.GMPTotalFee))
	r.Post("/v1/gmp/GMPTotalActiveUsers", registerHandler(handlers.GMPTotalActiveUsers))
	r.Post("/v1/gmp/GMPTopUsers", registerHandler(handlers.GMPTopUsers))
	r.Post("/v1/gmp/GMPTopITSAssets", registerHandler(handlers.GMPTopITSAssets))
	r.Post("/v1/gmp/searchGMP", registerHandler(handlers.GMPSearch))
	r.Post("/v1/gmp/getDataMapping", registerHandler(handlers.GetGMPDataMapping))
	r.Post("/v1/gmp/estimateTimeSpent", registerHandler(handlers.EstimateTimeSpent))

}
func registerTransferHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/transfer/searchTransfers", registerHandler(handlers.TransferSearch))
}

func registerVaultHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/vault/searchVault", registerHandler(handlers.SearchVault))
}

func registerCustodialHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/custodial", registerHandler(handlers.CreateCustodial))
	r.Get("/v1/custodials", registerHandler(handlers.GetCustodials))
	r.Get("/v1/custodial/{name}", registerHandler(handlers.GetCustodialByName))
	r.Post("/v1/custodial/group", registerHandler(handlers.CreateCustodialGroup))
	r.Get("/v1/custodial/groups", registerHandler(handlers.GetCustodialGroups))
	r.Get("/v1/custodial/groups/shorten", registerHandler(handlers.GetShortenCustodialGroups))
	r.Get("/v1/custodial/group/{name}", registerHandler(handlers.GetCustodialGroupByName))
}

func registerValidatorHandler(r *chi.Mux, handlers *handlers.Handler) {
	r.Post("/v1/validator/searchBlocks", registerHandler(handlers.SearchBlocks))
	r.Post("/v1/validator/searchBlock/{height}", registerHandler(handlers.SearchBlockByHeight))
	r.Post("/v1/validator/getTransactions", registerHandler(handlers.GetTransactions))
	r.Post("/v1/validator/searchTransactions", registerHandler(handlers.SearchTransactions))
	r.Post("/v1/validator/getTransaction/{hash}", registerHandler(handlers.GetTransactionByHash))
}
