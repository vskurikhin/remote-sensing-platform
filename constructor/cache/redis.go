package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/remote-sensing-platform/constructor/config"
	"github.com/vskurikhin/remote-sensing-platform/constructor/domain"
)

var ctx = context.Background()

type Redis struct {
	Cache *redis.Client
}

func CreateRedisCacheClient(cfg *config.Config) *Redis {
	client, err := newRedisClient(cfg)
	if err != nil {
		panic(err.Error())
	}
	return &Redis{Cache: client}
}

func newRedisClient(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port)
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Username:   cfg.Cache.Username,
		Password:   cfg.Cache.Password,
		DB:         0,
		MaxRetries: 2,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		_ = client.Close()
		return nil, err
	}
	logger.Debugf("connected to: %s", addr)
	return client, nil
}

func (r *Redis) PutEPoll(pollId int64, a *domain.EPoll) {
	result, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
		return
	}
	key := keyEPoll(pollId)
	err = r.Cache.Set(ctx, key, result, redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutEPoll: %v", err)
		return
	}
}

func (r *Redis) GetEPoll(pollId int64) (*domain.EPoll, error) {
	key := keyEPoll(pollId)
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.GetEPoll: pollId: %s %v", pollId, cmd.Err())
		return nil, cmd.Err()
	} else {
		s := cmd.Val()
		var a domain.EPoll
		err := json.Unmarshal([]byte(s), &a)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return &a, nil
	}
}

func keyEPoll(pollId int64) string {
	return fmt.Sprintf("EPoll-%d", pollId)
}

func (r *Redis) PutEPollSettings(pollId int64, a *domain.EPollSettings) {
	result, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
		return
	}
	key := keyEPollSettings(pollId)
	err = r.Cache.Set(ctx, key, result, redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutEPollSettings: %v", err)
		return
	}
}

func (r *Redis) GetEPollSettings(pollId int64) (*domain.EPollSettings, error) {
	key := keyEPollSettings(pollId)
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.PutEPollSettings: pollId: %s %v", pollId, cmd.Err())
		return nil, cmd.Err()
	} else {
		s := cmd.Val()
		var a domain.EPollSettings
		err := json.Unmarshal([]byte(s), &a)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return &a, nil
	}
}

func keyEPollSettings(pollId int64) string {
	return fmt.Sprintf("EPollSettings-%d", pollId)
}

func (r *Redis) PutEPollDesign(pollId int64, a *domain.EPollDesign) {
	result, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
		return
	}
	key := keyEPollDesign(pollId)
	err = r.Cache.Set(ctx, key, result, redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutEPollDesign: %v", err)
		return
	}
}

func (r *Redis) GetEPollDesign(pollId int64) (*domain.EPollDesign, error) {
	key := keyEPollDesign(pollId)
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.PutEPollDesign: pollId: %s %v", pollId, cmd.Err())
		return nil, cmd.Err()
	} else {
		s := cmd.Val()
		var a domain.EPollDesign
		err := json.Unmarshal([]byte(s), &a)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return &a, nil
	}
}

func keyEPollDesign(pollId int64) string {
	return fmt.Sprintf("EPollDesign-%d", pollId)
}

func (r *Redis) PutEPollChannel(pollId int64, a []domain.EPollChannel) {
	result, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
		return
	}
	key := keyEPollChannel(pollId)
	err = r.Cache.Set(ctx, key, result, redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutEPollChannel: %v", err)
		return
	}
}

func (r *Redis) GetEPollChannel(pollId int64) ([]domain.EPollChannel, error) {
	key := keyEPollChannel(pollId)
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.GetEPollChannel: pollId: %s %v", pollId, cmd.Err())
		return nil, cmd.Err()
	} else {
		s := cmd.Val()
		var a []domain.EPollChannel
		err := json.Unmarshal([]byte(s), &a)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return a, nil
	}
}

func keyEPollChannel(pollId int64) string {
	return fmt.Sprintf("EPollChannel-%d", pollId)
}

func (r *Redis) PutArrayOfFScreenMain(pollId int64, a []domain.FScreenMain) {
	result, err := json.Marshal(a)
	if err != nil {
		logger.Error(err)
		return
	}
	key := keyArrayOfFScreenMain(pollId)
	err = r.Cache.Set(ctx, key, result, redis.KeepTTL).Err()
	if err != nil {
		logger.Errorf("cache.PutArrayOfFScreenMain: %v", err)
		return
	}
	value := key
	for _, e := range a {
		key = keyInvalidateArrayOfFScreenMain(e.PollItem.PollId)
		err = r.Cache.Set(ctx, key, value, redis.KeepTTL).Err()
		if err != nil {
			logger.Errorf("cache.PutArrayOfFScreenMain: %v", err)
			return
		}
	}
}

func (r *Redis) GetArrayOfFScreenMain(pollId int64) ([]domain.FScreenMain, error) {
	key := keyArrayOfFScreenMain(pollId)
	cmd := r.Cache.Get(ctx, key)
	if cmd != nil && cmd.Err() != nil {
		logger.Errorf("cache.GetArrayOfFScreenMain: pollId: %s %v", pollId, cmd.Err())
		return nil, cmd.Err()
	} else {
		s := cmd.Val()
		var a []domain.FScreenMain
		err := json.Unmarshal([]byte(s), &a)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		return a, nil
	}
}

func keyArrayOfFScreenMain(pollId int64) string {
	return fmt.Sprintf("ArrayOfFScreenMain-%d", pollId)
}

func keyInvalidateArrayOfFScreenMain(pollItemId int64) string {
	return fmt.Sprintf("InvalidateArrayOfFScreenMain-pollItemId-%d", pollItemId)
}
