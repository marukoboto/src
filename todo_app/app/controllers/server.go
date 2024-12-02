package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"todo_app/app/models"
	"todo_app/config"
)

// generateHTML　指定されたテンプレートファイルを解析、レスポンスライターにHTMLを生成
// filenames ...string でstring型を受け取る
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		//fileで受け取った部分にapp/views/templates/をつけ、htmlのパス化してfilesに入れる
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	//mustにすることで、エラー状態になった際、パニック状態を発生させる
	templates := template.Must(template.ParseFiles(files...))
	//defineで定義しているlayoutを実行し、作成したHTMLをwに書き込む
	templates.ExecuteTemplate(w, "layout", data)
}

//cookieの取得・sessionの有無
func session(w http.ResponseWriter, r * http.Request) (sess models.Session, err error) {
	//cookieの取得
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		//valuse内のUUIDを取得
		sess = models.Session{UUID: cookie.Value}
		//UUIDとの合致
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

//URLパスの検証
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

//parseURL URLのパスから特定の整数を取り出し、それを引数として渡す
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// /todo/edit/1
		//URLのチェック　FindStringSubmatchで文字列の一致を確認
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w,r)
			return
		}
		//数字型に変換
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w,r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	//http.StripPrefixで　staticを取り除く
	//http.Handleで特定のURL、pattern（今回の場合/static/）が来た際そのリクエストを処理するhandler（http.StripPrefix("/static/", files））が機能する
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// /に関するリクエストが来た際、topを呼び出す
	//　第一引数　/ 第二引数　top
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	//parseURLでURLパスの末尾を解析→todoEditに渡される
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	//ListenAndServe:指定されたポートでサーバーを起動
	//":"+config.Config.Port で、ポート番号を指定
	//nilを入れておくことで、登録していないのを出すとpage not fundになる
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
