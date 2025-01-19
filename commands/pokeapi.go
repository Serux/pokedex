package commands

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Serux/pokedex/apiresponses"
)

func Get[V apiresponses.ExploreArea | apiresponses.Locationarea | apiresponses.Pokemon](url string, c *ConfigPoke) V {
	val, ok := c.Cache.Get(url)

	var body []byte
	if ok {
		body = val
	} else {
		resp, err := http.Get(url)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, _ = io.ReadAll(resp.Body)

		c.Cache.Add(url, body)
	}

	var la V
	json.Unmarshal(body, &la)

	return la
}
