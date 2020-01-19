package resas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	pathCities = "cities"
)

type CitiesResponse struct {
	Message interface{} `json:"message"`
	Result  []struct {
		PrefCode    int    `json:"prefCode"`
		CityCode    string `json:"cityCode"`
		CityName    string `json:"cityName"`
		BigCityFlag string `json:"bigCityFlag"`
	} `json:"result"`
}

func (c *Client) GetCities(ctx context.Context) (*CitiesResponse, error) {
	var result CitiesResponse

	err := c.call(ctx, http.MethodGet, path.Join(APIV1, pathCities), nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetCitiesByPrefCode(ctx context.Context, prefCode int64) (*CitiesResponse, error) {
	var result CitiesResponse

	q, err := url.ParseQuery(fmt.Sprintf("prefCode=%d", prefCode))
	if err != nil {
		return nil, err
	}
	err = c.call(ctx, http.MethodGet, path.Join(APIV1, pathCities), q, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
