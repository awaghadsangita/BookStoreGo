package router
import(
	"awesomeProject1/controllers"
	"github.com/gin-gonic/gin"
)
func HandleRequest(){
	r:=gin.Default()
	r.GET("/getBookByTitle",controllers.GetBooksByTitle)
	r.GET("/getAllBooks",controllers.GetAllBooks)
	r.Run()
}