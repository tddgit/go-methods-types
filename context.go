package main

import (
	"context"
	"net/http"
)

func main() {

}

func logic(ctx context.Context, info string) (string, error) {
	return "", nil
}

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			req = req.WithContext(ctx)
			handler.ServeHTTP(rw, req)
		})
}
