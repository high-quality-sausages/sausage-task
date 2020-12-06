package service

import (
	"log"

	"github.com/high-quality-sausages/sausage-task/task/internal/dao"
)

type Service struct {
	dao *dao.MainDao
}

func New(d *dao.MainDao) (s *Service) {
	s = &Service{
		dao: d,
	}
	// cf = s.Close
	return
}

func (s *Service) Close() {
	s.dao.Close()
	log.Println("service has been closed....")
}
