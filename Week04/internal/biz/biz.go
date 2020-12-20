package biz

import (
	"errors"
	
	xerrors "github.com/pkg/errors"
	"goTraining/Week04/internal/dao"
)

var (
	ErrEmpty = errors.New("empty string")
)

type StringService interface {
	Save(string) (int, error)
	GetAmount(s string) (int, error)
}

type stringService struct {
	dao.Repo
}

var _ StringService = (*stringService)(nil)

func NewStringService(repo dao.Repo) StringService {
	return &stringService{repo}
}

func (ss *stringService) Save(s string) (int, error) {
	if s == "" {
		return 0, ErrEmpty
	}

	return len(s), xerrors.Wrapf(ss.Repo.SaveInfo(s), "failed saving string")
}

func (ss *stringService) GetInfo(s string) (int, error) {
	return ss.Repo.GetInfo(s)
}
