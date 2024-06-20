package handler

import (
	"context"
	"main/views"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	sendMain(w, r)
}

func sendMain(w http.ResponseWriter, r *http.Request) {
	views.Main().Render(context.TODO(), w)
}
