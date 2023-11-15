package model

type Policy struct {
	ID     int
	UserID int
	Type   int
}

func (p *Policy) Save() error {
	return db.Save(p).Error
}
