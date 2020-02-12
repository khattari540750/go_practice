package main

import (
	"log"
	"net/http"
	"html/template"
	"sync"
	"path/filepath"
	"flag"
)

// templ は１つのテンプレートを表します
type templateHandler struct {
	once		sync.Once
	filename	string
	templ		*template.Template
}

// ServerHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func(){
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates",t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーション」のアドレス")
	flag.Parse() // フラグを解釈します
	r := newRoom()
	http.Handle("/", &templateHandler{filename:"chat.html"})
	http.Handle("/room",r)
	// チャットルームを開始します
	go r.run()
	// webサーバーを起動します
	log.Println("webサーバーを開始します。ポート： ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}