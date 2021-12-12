package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	perr "github.com/pkg/errors"
)
var name string

func main()  {
	id := 1
	name,err := GetName(id)
	if perr.Cause(err)!= sql.ErrNoRows && err != nil{
		fmt.Printf("get name error:\n%+v\n",err)
	}else if perr.Cause(err) == sql.ErrNoRows {
		fmt.Printf("Get the name with ID %v is empty",id)
	}
	fmt.Println(name)
}

func GetName(id int) (string,error) {
	db,_ := sql.Open("mysql","root:54e566yfrzxe9t@(127.0.0.1:3306)/golangdb")
	errPing := db.Ping()
	if errPing != nil {
		perr.Wrap(errPing,"sql Open error")
		return "", errPing
	}
	err := db.QueryRow("select name from user where id=?",id).Scan(&name)
	if err != nil && err !=  sql.ErrNoRows {
		perr.Wrap(err,"GetName sql error")
		return "", err
	}
	return name,err
}