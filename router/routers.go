package router
import(
	"awesomeProject1/controllers"
	"log"
	"net/http"
)
func HandleRequest(){
	http.HandleFunc("/getBooksByTitle",controllers.GetBooksByTitle)
	http.HandleFunc("/getAllBooks",controllers.GetAllBooks)
	log.Fatal(http.ListenAndServe(":10000", nil))
}