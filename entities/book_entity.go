package entities


type Book struct{
	Id int64 `json:"Id""`
	BookName string `json:"BookName""`
	Author string `json:"Author""`
	Image string `json:Image`
	Description string `json:"Descritpion"`
}

