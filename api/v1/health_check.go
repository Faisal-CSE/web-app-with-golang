package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func SampleHealthCheckAPI(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Rest API v1.0 - Health Check")
	Articles := []Article{
		{Id: "000000001", Title: "Sample Article 1", Desc: "Article Description 1", Content: "Article Content 1"},
		{Id: "000000002", Title: "Sample Article 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}
	_ = json.NewEncoder(w).Encode(Articles)
}
