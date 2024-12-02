package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	name STRING,
	email STRING,
	password STRING,
	created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
	//エラーの確認
	// _, err = Db.Exec(cmdU)
	// if err != nil {
	// 	log.Println(err)
	// }

	//スペルミス注意
	//cmdTなのでDb.Exec(cmdT)
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	email STRING,
	user_id INTEGER,
	created_at DATETIME)` , tableNameSession)

	Db.Exec(cmdS)
}

// 返り値　uuidのオブジェクト→uuidobj
// UUIDを作成し、uuidobjに格納される
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// 引数　plaintext　返り値　cryptext
// sha1 160ビット（20バイト）のハッシュ値 を生成
// %xにはplaintext が []byte に変換された値が入る
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
