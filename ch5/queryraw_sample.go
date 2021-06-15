package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hello"
	"strconv"
)

// Mydata is json structure.
type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

// Str get string value.
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail +
		"," + strconv.Itoa(m.Age) + ">"
}

var qry = "select * from mydata where id = ?"

func main() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for true {
		s := hello.Input("Id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs := con.QueryRow(qry, n)
		var md Mydata
		er2 := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er2 != nil {
			panic(er2)
		}
		fmt.Println(md.Str())
	}
	fmt.Println("***end***")
}
