package controller

import (
	"encoding/json"
	"net/http"
)

var BlockChain = NewBlockchain()

type status map[string]interface{}

type Message struct {
	Data int `json:"data"`
}

func WriteBlock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(status{"error": "could not decode message"})
			return
		}

		go func() {
			BlockChain.AddBlock(m.Data)
		}()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("working: please check your terminal")
	}
}

func GetBlockchain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BlockChain.Blocks)
	}
}
