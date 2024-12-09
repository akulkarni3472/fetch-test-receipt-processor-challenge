package main

import (
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Receipt struct {
	Id           int     `json:"id,omitempty"`
	Retailer     string  `json:"retailer,omitempty"`
	PurchaseDate string  `json:"purchaseDate,omitempty"`
	PurchaseTime string  `json:"purchaseTime,omitempty"`
	Total        float64 `json:"total,omitempty"`
	Items        []Item  `json:"items,omitempty"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription,omitempty"`
	Price            float64 `json:"price,omitempty"`
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
				//Alphanumeric Retailer Points
				for i := 0; i < len(a.Retailer); i++ {
					if alphanumeric.MatchString(string(a.Retailer[i])) {
						rec_points++
					}
				}
			}
			if a.Total > 0.0 {
				//Round Total Points
				if a.Total == float64(math.Trunc(float64(a.Total))) {
					rec_points += 50
				}
				//Divisble By 0.25 Points
				if (math.Mod(float64(a.Total), 0.25)) == 0.0 {
					rec_points += 25
				}
			}
			if len(a.Items) > 0 {
				//Item List Points
				item_points := int(len(a.Items) / 2)
				rec_points += item_points * 5
				//Trimmed Length Points
				for i := 0; i < len(a.Items); i++ {
					if len(strings.TrimSpace(a.Items[i].ShortDescription))%3 == 0 {
						rec_points += int(math.Round(a.Items[i].Price * 0.2))
					}
				}
			}
			//Odd day points
			if a.PurchaseDate != "" {
				var day = string(a.PurchaseDate[8:10])
				day_int, err := strconv.Atoi(day)
				if err != nil {
					c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
				}
				if day_int%2 != 0 {
					rec_points += 6
				}
			}
			//Time Points
			if a.PurchaseTime != "" {
				var timeHours = string(a.PurchaseTime[0:2])
				hours_int, err := strconv.Atoi(timeHours)
				if err != nil {
					c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id not found"})
				}
				if (hours_int >= 14) && (hours_int <= 16) {
					rec_points += 10
				}
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
