package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goTraining/Week02/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", GetArticle)
	http.ListenAndServe(":8080", nil)
}

func GetArticle(w http.ResponseWriter, req *http.Request) {
	title := "golang"
	// 略去参数检验
	article, err := service.GetArticleByTitle(title)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	if article == nil {
		msg := fmt.Sprintf("There is not article which title is %s", title)
		w.Write([]byte(msg))
		return
	}
	
	bytes, err := json.Marshal(article)
	if err != nil {
		log.Printf("handler Marshal err: %+v",err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(bytes)
}
