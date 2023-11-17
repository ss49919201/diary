package user

type User struct {
	id UserID
}

type UserID struct {
	value uint
}

func NewUserID(v uint) UserID {
	return UserID{
		value: v,
	}
}

func (u *UserID) Value() uint {
	return u.value
}
