package models

import (
	"log"
	"reakgo/utility"
)

type FormAddView struct {
	Id      int32
	Name    string
	Address string
	Email string
	Password string
}

func (form FormAddView) Add(name string, address string,email string,password string) {
	utility.Db.MustExec("INSERT INTO simpleForm (name, address,email,password) VALUES (?, ?, ?, ?)", name, address, email,password)
}

func (form FormAddView) View() ([]FormAddView, error) {
	var resultSet []FormAddView

	rows, err := utility.Db.Query("SELECT * FROM simpleForm")
	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()

		for rows.Next() {
			var singleRow FormAddView
			err = rows.Scan(&singleRow.Id, &singleRow.Name, &singleRow.Address,&singleRow.Email,&singleRow.Password,)
			if err != nil {
				log.Println(err)
			} else {
				resultSet = append(resultSet, singleRow)
			}
		}
	}
	return resultSet, err
}
