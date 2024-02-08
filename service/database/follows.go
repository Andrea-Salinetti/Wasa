package database

func (db *appdbimpl) FollowUser(userId string, followingUserId string) error {

	_, err := db.c.Exec(`INSERT INTO Followers (userId, followingUserId) VALUES (?, ?)`, userId, followingUserId)
	return err
}

func (db *appdbimpl) CheckFollow(userId string, followingUserId string) error {
	var placeholder string
	err := db.c.QueryRow("SELECT userId FROM Followers WHERE userId=? AND followingUserId=?", userId, followingUserId).Scan(&placeholder)
	return err
}

func (db *appdbimpl) DeleteFollow(userId string, toUnfollowUserId string) error {
	_, err := db.c.Exec(`DELETE FROM Followers WHERE userId=? AND followingUserId=?`, userId, toUnfollowUserId)
	return err
}
