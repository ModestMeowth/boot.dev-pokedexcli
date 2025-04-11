package pokeapi

import (
    "io"
    "net/http"
    "time"

    "github.com/ModestMeowth/boot.dev-pokedexcli/internal/pokecache"
)

type Client struct {
    httpClient http.Client
    cache pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
        cache: pokecache.NewCache(interval),
    }
}

func (c *Client) Do(url string) ([]byte, error) {
    var data []byte
    var err error

    if cache, ok := c.cache.Get(url); ok {
        data = cache
    } else {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return []byte{}, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return []byte{}, err
        }
        defer resp.Body.Close()

        data, err = io.ReadAll(resp.Body)
        c.cache.Add(url, data)
    }

    return data, err
}
