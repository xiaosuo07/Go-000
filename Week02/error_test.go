package error_homework

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/pkg/errors"
)

func TestServiceMethod(t *testing.T) {
	svr := CreateService()

	num, err := svr.GetUserCount()
	if err != nil {
		log.Println("HTTP 500")
	}
	log.Println(num)
}

type Dao struct {
}

func New() *Dao {
	return &Dao{}
}

func (dao *Dao) GetUserCountDao() (int, error) {
	return 0, errors.Wrap(sql.ErrNoRows, "dao error")
}

type Service struct {
	Dao
}

func CreateService() *Service {
	return &Service{}
}

func (svr *Service) GetUserCount() (int, error) {
	fmt.Println("service GetUserCount")
	_, err := svr.GetUserCountDao()
	if err != nil {
		if errors.Is(errors.Cause(err), sql.ErrNoRows) {
			log.Printf("发生了sql.ErrNoRows错误\n")
			log.Printf("原始错误发生信息：%T %v\n", errors.Cause(err), errors.Cause(err))
			log.Printf("堆栈信息：\n%+v\n", err)
			return 0, err
		}
	}
	return 1, nil
}
