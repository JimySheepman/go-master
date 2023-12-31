package middleware

import "net/http"

func CommonHeader(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")
		writer.Header().Add("Content-Type", "application/json")
		nextFunc(writer, request)
	}
}
