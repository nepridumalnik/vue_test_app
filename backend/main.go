package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type todoElement struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type todoList struct {
	arr *[]todoElement
}

func makeServer() *todoList {
	arr := new([]todoElement)

	td := &todoList{
		arr: arr,
	}

	return td
}

func (td *todoList) getAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	bytes, err := json.Marshal(&td.arr)

	if err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(bytes))
}

func (td *todoList) putOne(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var element todoElement
	err := json.NewDecoder(r.Body).Decode(&element)

	if err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	*td.arr = append(*td.arr, element)
}

func main() {
	td := makeServer()
	http.HandleFunc("/get_all", td.getAll)
	http.HandleFunc("/put_one", td.putOne)

	if err := http.ListenAndServe(":4321", nil); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s\n", err)
	}
}
