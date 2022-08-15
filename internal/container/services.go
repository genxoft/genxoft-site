package container

import (
	"genxoft.dev/internal/store"
	"github.com/sirupsen/logrus"
)

type services struct {
	logger *logrus.Logger
	store  store.Store
}

type Services interface {
	GetLogger() *logrus.Logger
	GetStore() store.Store
}

func New(l *logrus.Logger, s store.Store) Services {
	return &services{
		l,
		s,
	}
}

func (s *services) GetLogger() *logrus.Logger {
	return s.logger
}

func (s *services) GetStore() store.Store {
	return s.store
}
