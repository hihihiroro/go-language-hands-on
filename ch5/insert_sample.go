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

	nm := hello.Input("name")
	ml := hello.Input("mail")
	age := hello.Input("age")
	ag, _ := strconv.Atoi(age)

	qry := "insert into mydata (name,mail,age) values (?,?,?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

// print all record.
func showRecord(con *sql.DB) {
	qry = "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

// get Mydata from Rows.
func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

// get Mydata from Row.
func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
