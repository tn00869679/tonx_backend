package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	_ "tonx_backend/init"
	"tonx_backend/internal/database"
	"tonx_backend/internal/routes"
	"tonx_backend/internal/structure/request"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func createUpdateRequest(flight string) *httptest.ResponseRecorder {
	departureTime, _ := time.Parse(time.RFC3339, "2024-11-01T00:00:00Z")
	reqInfo := &request.BookTicket{
		Flight:        flight,
		DepartureTime: departureTime,
	}

	reqJson, _ := json.Marshal(reqInfo)
	req := httptest.NewRequest(http.MethodPatch, "/flight/bookTicket", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r := setupRouter()
	r.ServeHTTP(w, req)
	return w
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	pdb := database.Connect()
	routes.Setup(r, pdb)
	return r
}

func TestUpdateFlightInfo_ShouldSearchFail(t *testing.T) {
	w := createUpdateRequest("TEST123")
	assert.Equal(t, 503, w.Code)
}

func TestUpdateFlightInfo_ShouldBookTicketFail(t *testing.T) {
	w := createUpdateRequest("航班 ForTest")
	assert.Equal(t, 406, w.Code)
}

func TestUpdateFlightInfo_ShouldBookTicketSuccess(t *testing.T) {
	w := createUpdateRequest("航班 ForTestBookSuccess")
	assert.Equal(t, 200, w.Code)
}
