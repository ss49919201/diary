package user

type User struct {
	id UserID
}

func NewUser(id UserID) User {
	return User{id}
}

func (u *User) UserID() UserID {
	return u.UserID()
}

type UserID struct {
	value string
}

func NewUserID(v string) UserID {
	return UserID{
		value: v,
	}
}

func (u *UserID) Value() string {
	return u.value
}
