package controller

import "net/http"

func HandleRecipe(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello"))
}
