package controller

import (
	"net/http"
)

func CheckAuthorize(handler func(res http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		apiKey := req.Header.Get("Authorized")
		if apiKey != "password" {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("Unauthorized"))
			return
		}
		handler(res, req)
	})
}
