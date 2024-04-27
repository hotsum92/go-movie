package main

import (
	"log"
	"net/http"
)

func main() {
	//ディレクトリを指定する
	fs := http.FileServer(http.Dir("public"))
	//ルーティング設定。"/"というアクセスがきたらstaticディレクトリのコンテンツを表示させる
	http.Handle("/", fs)

	log.Println("Listening...")
	// 3000ポートでサーバーを立ち上げる
	http.ListenAndServe(":8080", nil)
}
