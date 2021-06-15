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
		// begin
		s := hello.Input("Id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs, er := con.Query(qry, n)
		// end
		if er != nil {
			panic(er)
		}
		for rs.Next() {
			var md Mydata
			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
			if er != nil {
				panic(er)
			}
			fmt.Println(md.Str())
		}
	}
	fmt.Println("***end***")
}
