package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"umbrella.github.com/go-zero-example/db_demo/sql"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3307)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewSqlConn("mysql", dsn)

	// Insert user
	r, err := conn.ExecCtx(context.Background(), "insert into user (type, name) values (?, ?)", 1, "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())

	// Query user
	var u sql.User
	query := "select id, name, type, create_at, update_at from user where id=?"
	err = conn.QueryRowCtx(context.Background(), &u, query, 1)
	if err != nil {
		panic(err)
	}

	// Update user
	_, err = conn.ExecCtx(context.Background(), "update user set type = ? where name = ?", 2, "test")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Delete user
	_, err = conn.ExecCtx(context.Background(), "delete from user where `id` = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}
