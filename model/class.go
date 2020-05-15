package model

import "fmt"

type Class struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

//list
func ListClass() ([]Class, error) {
	mods := make([]Class, 0, 4)
	err := DB.Select(&mods, "SELECT * FROM class")
	return mods, err
}

//add
func AddClass(class Class) int64 {
	Tx, _ := DB.Begin()
	result, err := Tx.Exec("insert into class (`name`,`desc`) values (?,?)", class.Name, class.Desc)
	if err == nil {
		affects, _ := result.RowsAffected()
		id, _ := result.LastInsertId()
		if affects == 1 {
			Tx.Commit()
			return id
		}
		return 0
	}
	return 0
}

//edit
func EditClass(class Class) bool {
	Tx, _ := DB.Begin()
	result, err := Tx.Exec("update class set `name`=?,`desc` =? where `id`=?", class.Name, class.Desc, class.Id)
	if err == nil {
		affects, _ := result.RowsAffected()
		if affects == 1 {
			Tx.Commit()
			return true
		}
		return false
	}
	return false
}

//delete
func DelClass(id int) bool {
	Tx, err := DB.Begin()
	if err != nil {
		fmt.Println("事务创建失败")
	}
	result, errs := Tx.Exec("delete from class where id=?", id)
	if errs == nil {
		affects, _ := result.RowsAffected()
		if affects == 1 {
			Tx.Commit()
			return true
		}
		return false
	}
	Tx.Rollback()
	return false
}

//getById
func GetClassById(id int) (*Class, error) {
	mod := &Class{}
	err := DB.Get(mod, "select * from class where id=?", id)
	return mod, err
}
