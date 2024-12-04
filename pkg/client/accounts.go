package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/deltadefi-protocol/go-sdk/pkg/api/requests"
	"github.com/deltadefi-protocol/go-sdk/pkg/api/responses"
)

type Accounts struct {
	httpClient *http.Client
}

func NewAccounts(httpClient *http.Client) *Accounts {
	return &Accounts{
		httpClient: httpClient,
	}
}

func (a *Accounts) SignIn(ctx context.Context, data *requests.SignInRequest) (*responses.SignInResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/accounts/signin", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("auth_key", data.AuthKey)
	req.URL.Query().Set("wallet_address", data.WalletAddress)

	res, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp responses.SignInResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *Accounts) GetBalance(ctx context.Context) (*responses.GetBalanceResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "/accounts/balance", nil)
	if err != nil {
		return nil, err
	}

	res, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp responses.GetBalanceResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
