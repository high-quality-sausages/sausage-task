package service

import (
	"fmt"

	"github.com/zbcheng/xdays/user/internal/dao"
)

type Service struct {
	dao *dao.MainDao
}

func New(d *dao.MainDao) (s *Service) {
	fmt.Println("hello service")
	s = &Service{
		dao: d,
	}
	// cf = s.Close
	return
}

func (s *Service) Close() {
	s.dao.Close()
}

func (s *Service) GetInfo() {
	resp := s.dao.GetJob()
	fmt.Println(resp)
}
