package external

import (
	"context"
	"encoding/json"

	"github.com/dacharat/circuit-breaker-go/internal/httpbreaker"
)

type ThirdApi interface {
	GetData(ctx context.Context) (getDataResponse, error)
}

type ThirdApiImpl struct {
	client httpbreaker.Client
}

func NewThirdApi(client httpbreaker.Client) ThirdApi {
	return &ThirdApiImpl{
		client: client,
	}
}

func (t *ThirdApiImpl) GetData(ctx context.Context) (getDataResponse, error) {
	url := "http://localhost:3000"

	var response getDataResponse
	body, err := t.client.Get(ctx, url, nil)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)

	return response, err
}

type getDataResponse struct {
	Message string `json:"message"`
}
