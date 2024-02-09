package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	w.Header().Set("content-type", "image/png")
	_ = r.ParseForm()
	userId := r.URL.Query().Get("userId")
	token := r.Header.Get("Authorization")

	if !rt.authorize(userId, token) {
		w.WriteHeader(http.StatusUnauthorized)
		ctx.Logger.Info("401, unauthorized")
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Info("400, bad request")
	}

	image, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.Info("400, bad request")
		return
	}

	defer image.Close()
	photoId := generateId()
	err = rt.db.CheckPhotoExistence(photoId)
	for !errors.Is(err, sql.ErrNoRows) {
		photoId = generateId()
		err = rt.db.CheckPhotoExistence(photoId)
	}

	photoDir := "./webui/public/photos/"
	_, err = os.Stat(photoDir)
	if err != nil {
		os.Mkdir(photoDir, os.ModePerm)
	}

	newPhotoName := "./webui/public/photos/" + "image-" + photoId + ".png"
	newPhoto, err := os.Create(newPhotoName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
	}

	_, err = io.Copy(newPhoto, image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
	}

	err = rt.db.PostPhoto(photoId, userId, newPhotoName, time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.Info("500, internal server error")
	}
	w.WriteHeader(http.StatusCreated)
	ctx.Logger.Info("201, post created")
	_ = json.NewEncoder(w).Encode(photoId)

}
