package controllers

import (
	"awesomeProject1/config"
	"awesomeProject1/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println(err)

	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.FindAll()
		if err!=nil{
			fmt.Println(err)
		}else{
			context.JSON(200,books)

		}
	}
}

func GetBooksByTitle(context *gin.Context){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println(err)
	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.SearchBooksByTitle(context.Query("title"))
		if err!=nil{
			fmt.Println(err)
		}else{
			context.JSON(200,books)
			//json.NewEncoder(w).Encode(bookks)
		}
	}
}