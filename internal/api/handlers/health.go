package handlers

import (
	"net/http"

	"github.com/scalarorg/xchains-api/internal/types"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Health check the service, including ping database connection
// @Produce json
// @Success 200 {string} PublicResponse[string] "Server is up and running"
// @Router /healthcheck [get]
func (h *Handler) HealthCheck(request *http.Request) (*Result, *types.Error) {
	// TODO: Add health check for db connection
	return NewResult("Server is up and running"), nil
}
