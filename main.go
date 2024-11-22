package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Receipt struct {
	Id           int     `json:"id,omitempty"`
	Retailer     string  `json:"retailer,omitempty"`
	PurchaseDate string  `json:"purchaseDate,omitempty"`
	PurchaseTime string  `json:"purchaseTime,omitempty"`
	Total        float32 `json:"total,omitempty"`
	Items        []Item  `json:"items,omitempty"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription,omitempty"`
	Price            float32 `json:"price,omitempty"`
}

var receipts []Receipt

func createReceipt(c *gin.Context) {
	var newReceipt Receipt
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func getReceiptById(c *gin.Context) {
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	for _, a := range receipts {
		if a.Id == id_int {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func getReceiptPointsById(c *gin.Context) {
	rec_points := 0
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	for _, a := range receipts {
		if a.Id == id_int {
			//TODO Points Calculation
			c.IndentedJSON(http.StatusOK, rec_points)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}

func handleRequests() {
	router := gin.Default()
	router.GET("/receipts", getReceipts)
	router.GET("/receipts/:id", getReceiptById)
	router.GET("/receipts/:id/points", getReceiptPointsById)
	router.POST("/receipts/process", createReceipt)
	router.Run("localhost:8080")
}

func main() {
	handleRequests()
}
