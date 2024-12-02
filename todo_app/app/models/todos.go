package models

import (
	"log"
	"time"
)

// todo用のストラクト
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// userのメソッド　引数　content 返り値 error
func (u *User) CreateTodo(content string) (err error) {
	//createdat
	cmd := `insert into todos (
	content, 
	user_id, 
	created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	
	return err
}

// userのメソッド　引数　content 返り値 error
func GetTodo(id int) (todo Todo, err error) {
	//GetTodoで取得したIDとの合致を確認
	//todostableからidに関するレコード取得
	cmd := `select id, content, user_id, created_at from todos 
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)
	//Todo型のtodoに値が返る　(todo Todo, err error) 部分
	return todo, err
}

//返り値　todos
func GetTodos() (todos []Todo, err error) {
	//全てのtodoを取得する
	cmd := `select id, content, user_id, created_at from todos`
	//Db.Query(cmd) でtodosテーブル全行取得
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	//結果セットから1行ずつデータを取得
	for rows.Next() {
		var todo Todo
		//
		err = rows.Scan(&todo.ID, 
			&todo.Content, 
			&todo.UserID, 
			&todo.CreatedAt)
		if err != nil{
			log.Fatalln(err)
		}
		//追加した部分に後から新しい要素が1件入る
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//特定のユーザーの情報を取得
func (u * User) GetTodosByUser() (todos []Todo, err error){
	cmd := `select id, content, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content, 
			&todo.UserID, 
			&todo.CreatedAt)
	
		if err != nil{
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos,err
}

//値の更新　Todoのメソッド
func (t *Todo) UpdateTodo() error{
	//t.ID部分とwhere idを合致させ、content,user_idを取ってくる
	cmd := `update todos set content = ?, user_id = ?
	where id = ?`
	//Db.Execで変更した内容を更新する
	_, err =Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}


//削除
func (t *Todo) DeleteTodo() error {
	//idが一致する物の削除
	cmd := `delete from todos where id = ?`
	//DBに対してExecでコマンドを渡す。　cmd内のdeleteを渡す
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}