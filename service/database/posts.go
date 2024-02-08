package database

func (db *appdbimpl) PostPhoto(photoId string, userId string, image string, date string) error {
	_, err := db.c.Exec(`INSERT INTO Photos (photoId, userId, image, date) VALUES (?, ?, ?, ?)`, photoId, userId, image, date)
	return err
}

func (db *appdbimpl) DeletePhoto(photoId string, userId string) error {
	_, err := db.c.Exec(`DELETE FROM Photos WHERE photoId=? AND userId=?`, photoId, userId)
	return err
}

func (db *appdbimpl) CheckPhotoOwnership(photoId string, userId string) error {
	var placeholder string
	err := db.c.QueryRow(`SELECT photoId FROM Photos WHERE photoId=? AND userId=?`, photoId, userId).Scan(&placeholder)
	return err
}
