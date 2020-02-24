package main

import (
	"awesomeProject1/config"
	"awesomeProject1/models"
	"fmt"
)

func main() {
	demo1_callFindAll()
}
func demo1_callFindAll(){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println(err)
	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.FindAll()
		if err!=nil{
			fmt.Println(err)
		}else{
			for _,book:=range books{
				fmt.Println("Id:",book.Id)
				fmt.Println("Name :",book.BookName)
				fmt.Println("Author :",book.Author)
				fmt.Println("Image :",book.Image)
				fmt.Println("Description :",book.Description)
			}
		}
	}
}