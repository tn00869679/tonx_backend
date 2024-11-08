package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func createListRequest(departureTime string) *httptest.ResponseRecorder {
	queryParams := url.Values{}
	queryParams.Add("departureAirport", "台北桃園國際機場 (TPE)")
	queryParams.Add("arrivalAirport", "東京羽田機場 (HND)")
	queryParams.Add("departureTime", departureTime)
	queryParams.Add("page", "1")
	queryParams.Add("perPage", "5")

	req := httptest.NewRequest(http.MethodGet, "/flight/list?"+queryParams.Encode(), nil)

	w := httptest.NewRecorder()
	r := setupRouter()
	r.ServeHTTP(w, req)
	return w
}

func TestGetFlightList_ShouldGetFail(t *testing.T) {
	departureTime, _ := time.Parse(time.RFC3339, "2024-11-01T00:00:00Z")
	w := createListRequest(departureTime.String())
	assert.Equal(t, 400, w.Code)
}

func TestGetFlightList_ShouldGetSuccess(t *testing.T) {
	w := createListRequest("2024-11-01T00:00:00Z")
	assert.Equal(t, 200, w.Code)
}
