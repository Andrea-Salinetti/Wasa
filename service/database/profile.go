package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetProfile(userId string, toVisitUserId string) ([]Photo, error) {
	rows, err := db.c.Query("SELECT image, photoId FROM Photos WHERE userId=? AND NOT EXISTS (SELECT * FROM Bans WHERE userId=? AND bannedUserId=?)", toVisitUserId, toVisitUserId, userId)
	if err != nil {
		return nil, err
	}

	var image string
	var photoId string
	images := []Photo{}
	for rows.Next() {
		err = rows.Scan(&image, &photoId)
		comments, err := db.RetrieveComments(photoId)
		likes, err := db.RetrieveLikes(photoId)
		username, err := db.RetrieveUsername(toVisitUserId)
		if err != nil {
			return nil, err
		}
		likeId, err := db.retrieveLikeId(photoId, userId)
		if errors.Is(err, sql.ErrNoRows) {
			likeId = ""
		}

		photo := Photo{image, photoId, username, comments, likes, likeId}
		images = append(images, photo)
	}
	return images, err
}
