package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Abeldlp/architectured-go/model"
)

type Service interface {
	GetCatFact(context.Context) (*model.CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*model.CatFact, error) {
	res, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	fact := &model.CatFact{}
	if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}
