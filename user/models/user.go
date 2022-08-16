package models

import (
	"log"
	"user/db"
)

type Person struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Desc string `json:"desc" form:"desc"`
}

//新增一条记录
func AddRecord(person *Person) int {
	rs, err := db.MysqlDB.Exec("insert into person(name, `desc`) value(?,?)", person.Name, person.Desc)
	if err != nil {
		log.Fatal(err)
	}
	id, err := rs.LastInsertId()
	if err != err {
		log.Fatal(err)
	}
	return int(id)
}

//查询一条记录
func QueryRecord(p *Person) (person Person, err error) {
	person = Person{}
	err = db.MysqlDB.QueryRow("select id,name,`desc` from person where id = ?", p.Id).Scan(&person.Id,
		&person.Name, &person.Desc)
	return person, err
}

//查询所有记录
func QueryAllRecord() (persons []Person, err error) {
	rows, err := db.MysqlDB.Query("select id,name,`desc` from person")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.Id, &person.Name, &person.Desc)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}
	rows.Close()
	return persons, err
}

//更新记录
func UpdateRecord(person *Person) int {
	rs, err := db.MysqlDB.Exec("update person set `desc` = ? where id = ?", person.Desc, person.Id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int(rows)
}

//删除一条记录
func DeleteRecord(id int) int64 {
	rs, err := db.MysqlDB.Exec("delete from person where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return int64(rows)
}
