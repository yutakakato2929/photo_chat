package mysql

import (
	"database/sql"
	"fmt"
	"photo_chat/infra/conf"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      string `db:"ID"`
	Account string `db:"ACCOUNT"`
	Name    string `db:"NAME"`
	Passwd  string `db:"PASSWD"`
	Created string `db:"CREATED"`
}

type Message struct {
	Username   string `json:"username"`
	Targetname string `json:"targetname"`
	Message    string `json:"message"`
}

//Overall database function
func Openmysql() (*sql.DB, error) {
	confDB, err := conf.ReadConfDB()
	if err != nil {
		return nil, err
	}
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", confDB.User, confDB.Pass, confDB.Host, confDB.Port, confDB.DbName, confDB.Charset)
	db, err := sql.Open("mysql", str)
	if err != nil {
		return nil, err
	}
	return db, nil
}

//Only for user table
func InsertUser(acc, name, pw string, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO USER(ACCOUNT,NAME,PASSWD,CREATED) VALUES(?,?,?,now())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(acc, name, pw)
	if err != nil {
		return err
	}
	return nil
}

func GetUserAll(db *sql.DB) ([]User, error) {
	var users []User
	stmt, err := db.Prepare("SELECT * FROM USER")
	if err != nil {
		return users, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Account, &user.Name, &user.Passwd, &user.Created)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByFlag(value, flag string, db *sql.DB) (User, error) {
	var user User
	stmt, err := db.Prepare("SELECT ID,ACCOUNT,NAME,PASSWD,CREATED FROM USER WHERE ? = ?")
	if err != nil {
		return user, err
	}
	rows, err := stmt.Query(flag, value)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&user.Id, &user.Account, &user.Name, &user.Passwd, &user.Created)
	if err != nil {
		return user, err
	}
	return user, nil
}

//Only for message table
func InsertMessage(msg Message, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO MESSAGE(MESSAGE,SENDER,RECIEVER,CREATED) VALUES(?,?,?,now())")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(msg.Message, msg.Username, msg.Targetname)
	if err != nil {
		return err
	}
	return nil
}

func GetMessageByName(sender, reciever string, db *sql.DB) ([]Message, error) {
	var msgs []Message
	stmt, err := db.Prepare("SELECT MESSAGE,SENDER,RECIEVER FROM message WHERE (SENDER = ? AND RECIEVER = ?) OR (SENDER = ? AND RECIEVER = ?)")
	if err != nil {
		return msgs, err
	}
	rows, err := stmt.Query(sender, reciever, reciever, sender)
	if err != nil {
		return msgs, err
	}
	defer rows.Close()
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.Message, &msg.Username, &msg.Targetname)
		if err != nil {
			return msgs, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, nil
}
