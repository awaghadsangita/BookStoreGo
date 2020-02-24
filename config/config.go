package config
import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB()(db *sql.DB,err error){
	dbDriver:="mysql"
	dbUser:="root"
	dbPass:="Ravi@9011907937"
	dbName:="BookStore"
	db, err=sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	return
}
