package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"example.com/m/v2/store"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func Test_getAllBooksShouldReturn200(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://localhost:8000/api/books")

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
}

type DelResponse struct {
	Result string `json:"result"`
}

func Test_postBookAndDeleteBook(t *testing.T) {
	//Create mock data
	mockBook := store.Book{Isbn: "122", Title: "Boxer", Price: 300.20}
	myResponse := store.Book{}
	client := resty.New()
	resp, _ := client.R().SetHeader("Content-Type", "application/json").SetBody(`{"isbn":"122","title":"Boxer","price":300.20}`).Post("http://localhost:8000/api/books")
	err := json.Unmarshal(resp.Body(), &myResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, mockBook.Isbn, myResponse.Isbn)
	assert.Equal(t, mockBook.Title, myResponse.Title)
	assert.Equal(t, mockBook.Price, myResponse.Price)
	//Delete mock data

	delResp := DelResponse{Result: "success"}
	acResp := DelResponse{}
	resp2, _ := client.R().Delete("http://localhost:8000/api/book/" + strconv.Itoa(myResponse.ID))
	err2 := json.Unmarshal(resp2.Body(), &acResp)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	assert.Equal(t, delResp.Result, acResp.Result)

}
