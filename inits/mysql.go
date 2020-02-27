package inits

// 数据库连接初始化
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
)

// DB 数据库模型
var DB *sql.DB

const dsn = "root:ectripaes2013@tcp(202.96.155.121:3306)/golist"

// InitDb 初始化数据连接
func InitDb() {
	db, err := sql.Open("mysql", dsn)
	defer db.Close()
	DB = db
	if err != nil {
		fmt.Println("数据库连接失败==>", err)
	}
	fmt.Println("数据库已连接！")
}
