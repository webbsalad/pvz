package integration

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateAndCloseReception(t *testing.T) {
	pvzID := uuid.New().String()

	{ // create pvz
		req := map[string]interface{}{
			"id":                pvzID,
			"registration_date": time.Now().UTC().Format(time.RFC3339Nano),
			"city":              "Москва",
		}
		var resp struct {
			Id string `json:"id"`
		}
		doRequest(t, "POST", baseURL+"/pvz", req, &resp, modToken)
		require.Equal(t, pvzID, resp.Id)
	}

	{ // open reception
		var resp struct {
			Id string `json:"id"`
		}
		doRequest(t, "POST", baseURL+"/receptions", map[string]interface{}{"pvzId": pvzID}, &resp, empToken)
		require.NotEmpty(t, resp.Id)
	}

	// add products
	for i := 0; i < 50; i++ {
		var resp struct {
			Id string `json:"id"`
		}
		req := map[string]interface{}{"type": "электроника", "pvzId": pvzID}
		doRequest(t, "POST", baseURL+"/products", req, &resp, empToken)
		require.NotEmpty(t, resp.Id, "product %d", i)
	}

	{ // close reception
		var resp struct {
			Id     string `json:"id"`
			Status string `json:"status"`
		}
		doRequest(t, http.MethodPost, baseURL+"/pvz/"+pvzID+"/close_last_reception", map[string]interface{}{}, &resp, empToken)

		require.Equal(t, "close", resp.Status)
	}
}
