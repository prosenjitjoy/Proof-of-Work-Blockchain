package main

import (
	"encoding/json"
	"fmt"
	"main/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", controller.GetBlockchain())
	router.Post("/", controller.WriteBlock())

	go func() {
		for _, block := range controller.BlockChain.Blocks {
			fmt.Printf("Prev hash: %v\n", block.PrevHash)
			bytes, _ := json.MarshalIndent(block.Data, "", "  ")
			fmt.Printf("Data: %v\n", string(bytes))
			fmt.Printf("Hash: %v\n", block.Hash)
			fmt.Println()
		}
	}()

	http.ListenAndServe(":5000", router)
}
