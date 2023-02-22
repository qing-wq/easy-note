package pack

import (
	"github.com/qing-wq/easy-note/cmd/user/dal/db"
	"github.com/qing-wq/easy-note/kitex_gen/user"
)

func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		UserId:   int64(u.ID),
		Username: u.UserName,
		Avatar:   u.PassWord,
	}
}

func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
