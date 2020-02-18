package catbot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const catAPIURL = "https://api.thecatapi.com/v1"

type Cat struct {
	APIKey string
}

type catImage struct {
	ID         string        `json:"id"`
	URL        string        `json:"url"`
	Categories []catCategory `json:"categories"`
	Breeds     []catBreeds   `json:"breeds"`
}

type catCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type catBreeds struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Temperament    string `json:"temperament"`
	LifeSpan       string `json:"life_span"`
	AltNames       string `json:"alt_names"`
	WikipediaURL   string `json:"wikipedia_url"`
	Origin         string `json:"origin"`
	WeightImperial string `json:"weight_imperial"`
}

func (c Cat) Verify() error {
	_, err := c.GetRandomURL()
	return err
}

func (c Cat) GetRandomURL() (string, error) {
	client := http.Client{
		Timeout: time.Minute,
	}
	url := catAPIURL + "/images/search"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("x-api-key", c.APIKey)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", fmt.Errorf("Invalid status code %d", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var images []catImage
	err = json.Unmarshal(body, &images)
	if err != nil {
		return "", err
	}
	if len(images) <= 0 {
		return "", errors.New("Did not find any cats")
	}
	return images[0].URL, nil
}
