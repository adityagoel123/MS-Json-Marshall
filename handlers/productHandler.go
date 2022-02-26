package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adityagoel/product-api/data"
)

type Products struct {
	thisLogger *log.Logger
}

func NewProducts(thisLogger *log.Logger) *Products {
	return &Products{thisLogger}
}

func (h *Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	h.thisLogger.Println("Serving Products LIST...")

	listOfProducts := data.GetProducts()
	jsonRepOfListOfProd, errorWhileMarshalling := json.Marshal(listOfProducts)

	if errorWhileMarshalling != nil {
		http.Error(responseWriter, "Unable to marshall the JSON", http.StatusInternalServerError)
	}

	responseWriter.Write(jsonRepOfListOfProd)
}
