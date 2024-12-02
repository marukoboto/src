package models

import (
	"log"
	"time"
)

// userデータの箱のような物を作成
type User struct {
	ID int
	//UUIDデータ識別用の名前や番号
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos	  []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// User型を*ポインタ経由で書く→直接変更を行うため
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
	uuid,
	name,
	email,
	password,
	created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		//Encryptでパスワードをハッシュ値に
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err

}

// ユーザーの取得
// id intからユーザー取得　id intの部分をcmdで照らし合わせ合致するものを取得
func GetUser(id int) (user User, err error) {
	user = User{}
	//userから取得した情報が入る
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`
	//Dbを呼び出し、cmd,idが引数
	//cmdにて取得した値をuserに格納
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

// createで作成されたUser部分の更新
func (u *User) UpdateUser() (err error) {
	//テーブルの更新
	cmd := `update users set name = ?, email = ? where id = ?`
	//u.Name, u.Email, u.IDから値を取って、置き換えを行う
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// 消す
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
	uuid,
	email,
	user_id,
	created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
		if err != nil {
			log.Panicln(err)
		}

		cmd2 := `select id, uuid, email, user_id, created_at
		from sessions where user_id = ? and email = ?`

		err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
			&session.ID, 
			&session.UUID, 
			&session.Email, 
			&session.UserID,
			&session.CreatedAt)

		return session, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at
	from sessions where uuid =?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

//UUIDの削除
func (sess *Session) DeleteSessionByUUID() (err error) {
	//uuidと一致するもの
	cmd := `delete from sessions where uuid =?`
	_, err = Db. Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//sessionを元にユーザー情報を取得
func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd :=`select id, uuid, name, email, created_at From users
	where id = ?`
	// DBからユーザー情報を取得し、userのフィールドに格納する
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID, 
		&user.UUID, 
		&user.Name, 
		&user.Email, 
		&user.CreatedAt)

	return user, err
}