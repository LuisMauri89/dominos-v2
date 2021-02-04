package entity

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type traceLogRepository struct {
	mtx  sync.RWMutex
	conn Connection
}

func NewTraceLogRepository(conn Connection) TraceLogRepository {
	return &traceLogRepository{
		conn: conn,
	}
}

func (r *traceLogRepository) FindAll(ctx context.Context) ([]TraceLog, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	rows, err := r.conn.DB.Query("SELECT id, timestamp, service_name, caller, event, extra FROM tlogs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tlogs := []TraceLog{}

	for rows.Next() {
		var tlog TraceLog
		if err := rows.Scan(&tlog.ID, &tlog.TimeStamp, &tlog.ServiceName, &tlog.Caller, &tlog.Event, &tlog.Extra); err != nil {
			return nil, err
		}
		tlogs = append(tlogs, tlog)
	}

	return tlogs, nil
}

func (r *traceLogRepository) Create(ctx context.Context, tlog TraceLog) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	err := r.conn.DB.QueryRow("INSERT INTO tlogs(id, timestamp, service_name, caller, event, extra) VALUES($1, $2, $3, $4, $5, $6) RETURNING id",
		tlog.ID,
		tlog.TimeStamp,
		tlog.ServiceName,
		tlog.Caller,
		tlog.Event,
		tlog.Extra).Scan(&tlog.ID)

	if err != nil {
		return err
	}

	return nil
}
