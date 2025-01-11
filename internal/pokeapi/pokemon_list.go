package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(locationName string) (RespDeepLocation, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespDeepLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespDeepLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	locationsResp := RespDeepLocation{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	c.cache.Add(url, data)
	return locationsResp, nil
}
