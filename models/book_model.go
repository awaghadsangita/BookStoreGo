package models

import (
	"awesomeProject1/entities"
	"database/sql"
)

type BookModel struct{
	Db *sql.DB
}
func (bookModel BookModel)FindAll() ([]entities.Book,error){
	rows,err:=bookModel.Db.Query("select  * from books")
	if err!=nil{
		return nil,err
	}else {
		books := []entities.Book{}
		for rows.Next() {
			var id int64
			var bookname string
			var author string
			var image string
			var description string
			err2 := rows.Scan(&id, &bookname, &author, &image, &description)
			if err2 != nil {
				return nil, err2
			} else {
				book := entities.Book{id, bookname, author, image, description}
				books = append(books, book)
			}
		}
		return books, nil
	}
}
func (bookModel BookModel)SearchBooksByTitle(title string) ([]entities.Book,error){
	rows,err:=bookModel.Db.Query("select  * from books where bookname REGEXP '"+title+"'")
	if err!=nil{
		return nil,err
	}else {
		books := []entities.Book{}
		for rows.Next() {
			var id int64
			var bookname string
			var author string
			var image string
			var description string
			err2 := rows.Scan(&id, &bookname, &author, &image, &description)
			if err2 != nil {
				return nil, err2
			} else {
				book := entities.Book{id, bookname, author, image, description}
				books = append(books, book)
			}
		}
		return books, nil
	}
}