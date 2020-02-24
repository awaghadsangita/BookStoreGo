package controllers

import (
	"awesomeProject1/config"
	"awesomeProject1/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter,r *http.Request){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println(err)

	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.FindAll()
		if err!=nil{
			fmt.Println(err)
		}else{
			json.NewEncoder(w).Encode(books)
		}
	}
}

func GetBooksByTitle(w http.ResponseWriter,r *http.Request){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println(err)
	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.SearchBooksByTitle(r.URL.Query().Get("title"))
		if err!=nil{
			fmt.Println(err)
		}else{
			json.NewEncoder(w).Encode(books)
		}
	}
}