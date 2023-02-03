package practice

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"giao/util"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	db = NewDB()
}

func NewDB() *sql.DB {
	fmt.Println("connecting db")
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3308)/vpea_erp_local")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		fmt.Println("connect fail")
		return nil
	}
	fmt.Println("connect success")
	return db
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUser(id string) user {
	m := QueryToMap("select * from users where id = ?", id)
	var users []user
	err := MapToStruct(m, &users)
	util.CheckErr(err)
	if len(users) > 0 {
		return users[0]
	}
	return user{}
}

func QueryToMap(queryStr string, args ...any) []map[string]interface{} {
	prepare, err := db.Prepare(queryStr)
	if err != nil {
		util.CheckErr(err)
		return nil
	}
	rows, err := prepare.Query(args...)
	if err != nil {
		util.CheckErr(err)
		return nil
	}
	defer rows.Close()
	columns, err := rows.Columns()
	util.CheckErr(err)
	columnsLen := len(columns)

	// 返回的map切片
	resMapData := make([]map[string]interface{}, 0)

	// 一条数据各列的值
	values := make([]interface{}, columnsLen)
	// 一条数据各列的值的地址
	columnsProp := make([]interface{}, columnsLen)

	for rows.Next() {
		for i := 0; i < columnsLen; i++ {
			columnsProp[i] = &values[i]
		}
		// Scan 将值映射到地址上，这里映射到columnsProp，变相为values赋值，因为columnsProp存储的是values各值的指针
		rows.Scan(columnsProp...)

		// 一条数据的map
		rowMap := make(map[string]interface{})

		for i, col := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		resMapData = append(resMapData, rowMap)
	}
	return resMapData
}

func MapToStruct(m interface{}, s interface{}) error {
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, s)
	if err != nil {
		return err
	}
	return nil
}
