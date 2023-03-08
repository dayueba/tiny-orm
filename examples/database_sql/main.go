package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int64
	Ctime time.Time
	Mtime time.Time // 更新时间
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@/db1?parseTime=true")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
}

func FindUsers(ctx context.Context) ([]*User, error) {
	rows, err := db.QueryContext(ctx, "SELECT `id`,`name`,`age`,`ctime`,`mtime` FROM user WHERE `age`<? ORDER BY `id` LIMIT 20 ", 20)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // 不能忘记关闭
	result := []*User{}
	for rows.Next() {
		a := &User{}
		if err := rows.Scan(&a.Id, &a.Name, &a.Age, &a.Ctime, &a.Mtime); err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	if rows.Err() != nil { // 需要处理错误
		return nil, rows.Err()
	}
	return result, nil
}

func main() {
	defer db.Close()

	users, err := FindUsers(context.Background())
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
