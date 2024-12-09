package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
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
	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
	id := c.Param("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
	}
	for _, a := range receipts {
		if a.Id == id_int {
			if a.Retailer != "" {
				for i := 0; i < len(a.Retailer); i++ {
					if alphanumeric.MatchString(string(a.Retailer[i])) {
						fmt.Println(string(a.Retailer[i]))
						rec_points++
					}
				}
			}
			if a.Total > 0.0 {
				if a.Total == float32(math.Trunc(float64(a.Total))) {
					rec_points += 50
				}
				if (math.Mod(float64(a.Total), 0.25)) == 0.0 {
					rec_points += 25
				}
			}
			if len(a.Items) > 0 {
				item_points := int(len(a.Items) / 2)
				rec_points += item_points * 5
			}
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
