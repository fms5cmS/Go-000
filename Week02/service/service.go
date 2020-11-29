package service

import (
	"database/sql"
	"errors"
	"goTraining/Week02/dao"
	"goTraining/Week02/model"
)

func GetArticleByTitle(title string) (*model.Article, error) {
	article, err := dao.GetArticleByTitle(title)
	// 检查是否是 “查询结果为空” 类型的错误，如果是，则吞掉该错误
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	// 其他错误
	// return nil, err
	return article, nil
}
