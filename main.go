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

/*{
	{Id:1,Retailer:"Walmart",PurchaseDate:"2023-04-05",PurchaseTime:"12:32",Total:31.43,Items:[
        {ShortDescription:"Digital Clock",Price:30.00},
        {ShortDescription:"Candybar",Price:1.43}
        ]
    },
	{Id:2,Retailer:"Home Depot",PurchaseDate:"2022-07-12",PurchaseTime:"15:34",Total:15.35,Items:[
        {ShortDescription:"Hammer",Price:15.00},
        {ShortDescription:"Nails",Price:0.35}
        ]
    },
	{Id:3,Retailer:"Petsmart",PurchaseDate:"2024-11-25",PurchaseTime:"16:02",Total:30.00,Items:[
        {ShortDescription:"Chew Toy",Price:10.00},
        {ShortDescription:"Dog Treats",Price:20.00}
        ]
    },
}*/

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
