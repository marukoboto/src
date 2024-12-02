package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//sessionの可否
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			//sessionが存在する→todosにリダイレクト
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		//入力フォームの解析
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		//models.Userを作成しuserに突っ込む
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		//新しいユーザーの作成
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		//TOPへ
		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)

	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//logout
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err :=r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	//NoCookieではない場合　UUIDに取得したcookieの値を入れる
	if err != http.ErrNoCookie {
		session:= models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}