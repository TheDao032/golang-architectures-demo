package kafka

import (
	"context"
	"time"

	gemmessages "github.com/TheDao032/proto/gemservice/gopb"
	creategemsource "github.com/TheDao032/golang-architectures-demo/internal/application/gem/commands/create_gem_source"
	"github.com/google/uuid"

	v "github.com/TheDao032/go-backend-utils-architecture/validation"
	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

var (
	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
)

func (c *Consumer) processGemSourceCreated(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	msg := &gemmessages.NewGemTransaction{}
	if err := proto.Unmarshal(m.Value, msg); err != nil {
		c.logger.Error(ctx, "proto.Unmarshal", zap.Error(err))
		c.commitErrMessage(ctx, r, m)
		return
	}

	createGemSourceCommand := &creategemsource.CreateGemSourceCommand{
		Id:          uuid.New().String(),
		UserId:      msg.GetUserId(),
		SourceId:    msg.GetSourceId(),
		Gems:        msg.GetGems(),
		Type:        msg.GetType(),
		Reason:      msg.GetReason(),
		Metadata:    msg.GetMetadata(),
		CollectedAt: msg.GetCollectedAt().AsTime(),
		CreatedBy:   msg.GetCreatedBy(),
	}
	if err := v.ValidateStruct(createGemSourceCommand); err != nil {
		c.logger.Error(ctx, "Validate", zap.Error(err))
		c.commitErrMessage(ctx, r, m)
		return
	}

	if err := retry.Do(func() error {
		// Add logic here
		_, err := c.service.GemService.CreateGemSourceHandler.Handle(ctx, createGemSourceCommand)

		c.logger.Info(ctx, "Consume kafka: "+time.Now().String())
		return err
	}, append(retryOptions, retry.Context(ctx))...); err != nil {
		c.logger.Error(ctx, "UpdateGemHandler.Handle", zap.Error(err))
		return
	}

	c.commitMessage(ctx, r, m)
}
