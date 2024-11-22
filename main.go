package main

import (
	"log"
	"net/http"
)

const creditScoreMin = 500
const creditScoreMax = 900

type credit_rating struct {
	CreditRating int `json:"credit_rating"`
}

type Rceipt struct {
	ID           int     `json:"id"`
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchase-date"`
	PurchaseTime string  `json:"purchase-time"`
	Total        float64 `json:"total"`
	Items        []Item  `json:"items"`
}

type Item struct {
	ShortDesc string `json:"short-description"`
	Price     string `json:"price"`
}

func postReceipts(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	http.Handle("/receipts/process", http.HandlerFunc(postReceipts))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	handleRequests()
}
