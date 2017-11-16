package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/johnswanson/ttc"
	"net/http"
	"strings"
)

func Save(api ttc.API, p *ttc.Ping) {
	j, _ := json.Marshal(p)
	req, err := http.NewRequest("POST", api.URL+"/api/ping", strings.NewReader(string(j)))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", api.Token)
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("http Do error %v\n", err)
	}
	defer resp.Body.Close()
	fmt.Printf("[%v] %v\n", resp.Status, resp.Body)
	return
}

func GetConfig(api ttc.API, config *ttc.Config) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", api.URL+"/api/config", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", api.Token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to retrieve config from server!")
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(config)
	return nil

}
