package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/Abeldlp/architectured-go/model"
	"github.com/Abeldlp/architectured-go/service"
)

type loggingService struct {
	next service.Service
}

func NewLoggingService(next service.Service) service.Service {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) GetCatFact(ctx context.Context) (fact *model.CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
