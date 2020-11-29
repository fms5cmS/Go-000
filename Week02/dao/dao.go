package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goTraining/Week02/model"
)

func GetArticleByTitle(title string) (*model.Article, error) {
	article := &model.Article{}
	err := sql.ErrNoRows
	if err != nil {
		return nil, err
	}
	return article, nil
}
