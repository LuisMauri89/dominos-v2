package service

import (
	"context"
	"strconv"
	"time"

	"logs/internal/entity"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

type TraceLogService interface {
	FindAll(ctx context.Context) ([]entity.TraceLog, error)
	Create(ctx context.Context, tlog entity.TraceLog) error
}

type traceLogService struct {
	repository entity.TraceLogRepository
	logger     log.Logger
}

func NewTraceLogService(repository entity.TraceLogRepository, logger log.Logger) TraceLogService {
	return &traceLogService{
		repository: repository,
		logger:     logger,
	}
}

func (s *traceLogService) FindAll(ctx context.Context) ([]entity.TraceLog, error) {
	logs, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return logs, nil
}

func (s *traceLogService) Create(ctx context.Context, tlog entity.TraceLog) error {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	tlog.ID = id

	now := time.Now()
	timestamp := now.Unix()
	tlog.TimeStamp = strconv.FormatInt(timestamp, 10)

	if err := s.repository.Create(ctx, tlog); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")
	return nil
}
