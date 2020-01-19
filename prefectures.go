package resas

import (
	"context"
	"net/http"
	"path"
)

const (
	pathPrefectures = "prefectures"
)

type PrefecturesResponse struct {
	Message interface{} `json:"message"`
	Result  []struct {
		PrefCode int    `json:"prefCode"`
		PrefName string `json:"prefName"`
	} `json:"result"`
}

func (c *Client) GetPrefectures(ctx context.Context) (*PrefecturesResponse, error) {
	var result PrefecturesResponse

	err := c.call(ctx, http.MethodGet, path.Join(APIV1, pathPrefectures), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
