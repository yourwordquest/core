package tigergraph

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/yourwordquest/core/utils/requests"
)

var username string
var password string
var server string
var graph string
var tokenExpiry int64
var savedToken string
var tokenMutex sync.Mutex

func init() {
	username = os.Getenv("TG_USERNAME")
	password = os.Getenv("TG_PASSWORD")
	graph = os.Getenv("TG_GRAPH")
	server = os.Getenv("TG_SERVER")
	tokenExpiry = 0
	savedToken = ""
}

func GetToken() (token string, err error) {
	// A token can only be acquired by one process at a time.
	// This ensures that we don't end up refreshing the token multiple times
	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	err = nil
	if time.Now().UnixMilli() < tokenExpiry {
		return savedToken, nil
	}

	// Get a new token
	resp := struct {
		Code       string `json:"code"`
		Expiration int64  `json:"expiration"`
		HasError   bool   `json:"error"`
		Message    string `json:"message"`
		Token      string `json:"token"`
	}{}
	_, err = requests.Request{
		Method: requests.POST,
		Headers: requests.Headers{
			"Authorization": fmt.Sprintf("Basic %s:%s", username, password),
		},
		JSON: requests.Payload{
			"graph":    graph,
			"lifetime": "1800", // Token valid for 30 min
		},
	}.Send(fmt.Sprintf("%s/requesttoken", server), &resp)
	if err != nil {
		return
	}
	if !resp.HasError {
		err = errors.New(resp.Message)
		return
	}
	savedToken = resp.Token
	token = savedToken
	// Set our expiration 5 min behind
	tokenExpiry = resp.Expiration - 1800
	return
}
