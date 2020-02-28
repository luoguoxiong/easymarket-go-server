package mysqlinit

// 数据库连接初始化
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jmoiron/sqlx"
)

// DB 数据库模型
var DB *sqlx.DB

const dsn = "root:ectripaes2013@tcp(202.96.155.121:3306)/nideShop"

// InitDb 初始化数据连接
func InitDb() (db *sqlx.DB, err error) {
	db, err = sqlx.Connect("mysql", dsn)
	DB = db
	if err != nil {
		fmt.Println("数据库连接失败==>", err)
	}
	fmt.Println("数据库已连接！")
	return
}
