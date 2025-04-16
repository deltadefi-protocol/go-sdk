package deltadefi

type AppClient struct {
	pathUrl string
	client  *Client
}

func newAppClient(client *Client) *AppClient {
	return &AppClient{
		pathUrl: "/app",
		client:  client,
	}
}

func (c *AppClient) GetTermsAndConditions() (interface{}, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/terms-and-conditions")
	if err != nil {
		return nil, err
	}

	return &bodyBytes, nil
}

func (c *AppClient) GetHydraCycle() (interface{}, error) {
	bodyBytes, err := c.client.get(c.pathUrl + "/hydra-cycle")
	if err != nil {
		return nil, err
	}

	return &bodyBytes, nil
}
