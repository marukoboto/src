package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

// TOPページを開いた時に実行
// w http.ResponseWriterはデータを書くための部分
// r *http.Request データの読み取り
func top(w http.ResponseWriter, r *http.Request) {
	//cookieの取得
	_, err := session(w, r)
	if err != nil {
		//セッションが存在しない場合topへ
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// indexに遷移した時に実装
func index(w http.ResponseWriter, r *http.Request) {
	//sessionでのログインの可否
	sess, err := session(w, r)
	if err != nil {
		//topにリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		//sessionのuseridと一致するuserの取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		//userのtodoを取得
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

//todoの作成ページ
func todoNew(w http.ResponseWriter, r *http.Request) {
	//ログインの有無
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

//Todoの保存ページ
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		//ParseForm内に値を格納
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//ユーザー情報を取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//todoで記入した内容の取得
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/todos", 302)
	}
}

//todoの編集
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w,r, "/login", 302)
	} else {
		//ユーザーの確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//引数のIdからtodoの取得
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		//htmlページの作成
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int){
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else{
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//ユーザーの取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Panicln(err)
		}
		//編集ページのコンテントを取得
		content := r.PostFormValue("content")
		t:= &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

//todoの削除
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	//session確認
	sess, err := session(w,r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		//userの確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//idから取得
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		//エラーが無ければDeleteTodoを呼び出し削除
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}