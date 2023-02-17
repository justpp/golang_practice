package services

import (
	"github.com/pkg/errors"
)

type AuthRequest struct {
	AppKey    string `json:"app_key" form:"app_key" binding:"required"`
	AppSecret string `json:"app_secret" form:"app_secret" binding:"required"`
}

func (s *Service) CheckAuth(param *AuthRequest) error {
	auth, err := s.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exists")
}
