package client

import (
	"context"
	"io"
	"net/http"
)

type Api struct{}

func (a *Api) ResolveAxiosData(ctx context.Context, response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
