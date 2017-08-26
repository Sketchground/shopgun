package shopgun

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Session request
type Session struct {
	APIKey string `json:"api_key"`
}

// SessionResponse response
type SessionResponse struct {
	Token       string      `json:"token"`
	Reference   string      `json:"reference"`
	Expires     string      `json:"expires"`
	User        interface{} `json:"user"`
	Permissions struct {
		Guest []string `json:"guest"`
	} `json:"permissions"`
	Provider interface{} `json:"provider"`
	ClientID string      `json:"client_id"`
}

func (c *Client) session() (SessionResponse, error) {
	//TODO:  If session is still valid, then return session.

	path := fmt.Sprintf("%v/sessions", c.basePath)
	data, err := json.Marshal(&Session{c.apiKey})
	if err != nil {
		return SessionResponse{}, err
	}
	body := bytes.NewReader(data)
	req, err := http.NewRequest("POST", path, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.c.Do(req)
	if err != nil {
		return SessionResponse{}, err
	}
	if resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Print(string(respBody))
		return SessionResponse{}, errors.New("got error: %v" + resp.Status)
	}

	var sr SessionResponse
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&sr)
	if err != nil {
		return SessionResponse{}, err
	}

	c.sr = sr
	return sr, nil
}

func (c *Client) doGet(path string, out interface{}) error {
	path = fmt.Sprintf("%v%v", c.basePath, path)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	sr, err := c.session()
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Token", sr.Token)

	sign := sha256.New()
	sign.Write([]byte(c.secret + sr.Token))
	sum := sign.Sum(nil)
	req.Header.Set("X-Signature", fmt.Sprintf("%x", sum))

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	//TODO: Keep an eye on X-RateLimit
	bod, _ := ioutil.ReadAll(resp.Body)
	log.Print(string(bod))

	//dec := json.NewDecoder(resp.Body)
	//err = dec.Decode(&out)
	err = json.Unmarshal(bod, &out)
	if err != nil {
		return err
	}
	return nil
}
