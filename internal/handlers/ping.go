package handlers

import (
	"fmt"
	"net/http"
)

// ping хендлер для проверки работоспособности
func ping(w http.ResponseWriter, r *http.Request) {
	res := "pong"

	fmt.Fprintln(w, res)
}
