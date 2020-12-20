package dao

import (
	"database/sql"
	
	xerrors "github.com/pkg/errors"
)

type Repo interface {
	SaveInfo(s string) error
	GetInfo(s string) (string, error)
}

type repo struct {
	db *sql.DB
}

var _ Repo = (*repo)(nil)

func NewRepo(db *sql.DB) Repo {
	return &repo{db: db}
}

func (r *repo) SaveInfo(info string) error {
	_, err := r.db.Query("insert into message(info) values (?)", info)
	return xerrors.Wrapf(err, "failed insert")
}

func (r *repo) GetInfo(topic string) (string, error) {
	var info string
	rows, err := r.db.Query("select info from message where topic = ?", topic)
	if err != nil {
		return "", xerrors.Wrapf(err, "failed select")
	}
	defer rows.Close()
	
	for rows.Next() {
		if err := rows.Scan(&info); err != nil {
			return "", xerrors.Wrapf(err, "failed scanning")
		}
	}
	
	return info, nil
}
