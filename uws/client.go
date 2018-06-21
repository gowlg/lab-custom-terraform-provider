package uws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Instance struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Private bool   `json:"private"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateInstanceInput struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Private bool   `json:"private"`
}

type UpdateInstanceInput struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Private bool   `json:"private"`
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(baseUrl string) (*Client, error) {
	client := &Client{
		httpClient: http.DefaultClient,
		baseURL:    baseUrl,
	}

	return client, nil
}

func (c *Client) CreateInstance(newInstance CreateInstanceInput) (*Instance, error) {
	url := fmt.Sprintf("%s/instances", c.baseURL)

	requestBody, err := json.Marshal(newInstance)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := ioutil.ReadAll(resp.Body)

		return nil, fmt.Errorf("Expected status code 201 but got %v, %v", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var instance Instance

	err = json.Unmarshal(body, &instance)

	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (c *Client) ReadInstance(instanceID int) (*Instance, error) {
	url := fmt.Sprintf("%s/instances/%d", c.baseURL, instanceID)

	resp, err := c.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status code 200 but got %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var instance Instance

	err = json.Unmarshal(body, &instance)

	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (c *Client) UpdateInstance(instanceID int, updatedInstance UpdateInstanceInput) (*Instance, error) {
	url := fmt.Sprintf("%s/instances/%d", c.baseURL, instanceID)

	requestBody, err := json.Marshal(updatedInstance)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status code 200 but got %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var instance Instance

	err = json.Unmarshal(body, &instance)

	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (c *Client) DeleteInstance(instanceID int) error {
	url := fmt.Sprintf("%s/instances/%d", c.baseURL, instanceID)

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Expected status code 200 but got %v", resp.StatusCode)
	}

	return nil
}
