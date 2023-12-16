package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/mangohow/mangokit-template/internal/model"
)

type GreeterDao struct {
	client *sqlx.DB
}

func NewGreeterDao(db *sqlx.DB) *GreeterDao {
	return &GreeterDao{client: db}
}

func (u *GreeterDao) FindById(id int) (user *model.User, err error) {
	sql := `SELECT * FROM t_user WHERE id = ?`
	user = new(model.User)
	err = u.client.Get(user, sql, id)

	return
}

func (u *GreeterDao) String() string {
	return "GreeterDao"
}
