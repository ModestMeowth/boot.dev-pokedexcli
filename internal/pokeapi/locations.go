package pokeapi

import (
    "encoding/json"
)

type LocationList struct {
    Count int `json:"count"`
    Next *string `json:"next"`
    Previous *string `json:"previous"`
    Results []struct {
        Name string `json:"name"`
        URL string `json:"url"`
    } `json:"results"`
}

type ExploreList struct {
    Encounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL string `json:"url"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}


func (c *Client) GetLocation(area string) (ExploreList, error) {
    var data []byte
    var err error

    url := baseURL + "/location-area/" + area

    data, err = c.Do(url)
    if err != nil {
        return ExploreList{}, err
    }

    exploreList := ExploreList{}
    err = json.Unmarshal(data, &exploreList)

    return exploreList, err
}

func (c *Client) ListLocations(pageURL *string) (LocationList, error) {
    var data []byte
    var err error

    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    data, err = c.Do(url)
    if err != nil {
        return LocationList{}, err
    }

    locationsResp := LocationList{}
    err = json.Unmarshal(data, &locationsResp)
    if err != nil {
        return LocationList{}, err
    }

    return locationsResp, nil
}
