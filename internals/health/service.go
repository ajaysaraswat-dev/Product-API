package health

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Service interface {
	CheckHealth(ctx context.Context) map[string] interface{}
}

type svc struct {
	db *mongo.Client
}

func NewService(db *mongo.Client) Service {
	return &svc{
		db : db,
	}
}

func(s *svc) CheckHealth(ctx context.Context) map[string]interface{} {
	status := "up"
	dbstatus := "up"
	ctx,cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()
	if err := s.db.Ping(ctx,nil);err!= nil {
		status = "down"
		dbstatus = "degraded"
	}
	return map[string] interface{}{
		"status" : status,
		"database":dbstatus,
	}
	
}