package main

import (
	"net/http"
	"html/template"
	"log"
)

// 2回目の送信
func resultView(res http.ResponseWriter, req *http.Request, User *player, Cpu *player) error {

	m := map[string]interface{}{"user" : User, "cpu" : Cpu}
	t := template.Must(template.ParseFiles("./resource/result.html"))

	err := t.Execute(res, m)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

