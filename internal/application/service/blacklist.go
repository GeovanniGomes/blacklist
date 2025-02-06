package service

import (
	"context"
	"fmt"
	"log"

	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/application/dto"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
)

type BlacklistService struct {
	usecaseCreateBlacklist blacklist.AddBlacklistInterface
	usecaseCheckBlacklist  blacklist.CheckBlacklistInterface
	usecaseRemoveBlacklist blacklist.RemoveBlackListInterface
	register_audit         contracts.AuditLoggerInterface
	persistence_cache      contracts.CacheInterface
	producer      *producer.BlacklistProducer
}

func NewBlackListService(
	usecaseCreateBlacklist blacklist.AddBlacklistInterface,
	usecaseCheckBlacklist blacklist.CheckBlacklistInterface,
	usecaseRemoveBlacklist blacklist.RemoveBlackListInterface,
	register_audit contracts.AuditLoggerInterface,
	persistence_cache contracts.CacheInterface,
	producer      *producer.BlacklistProducer,
) *BlacklistService {
	return &BlacklistService{
		usecaseCreateBlacklist: usecaseCreateBlacklist,
		usecaseCheckBlacklist:  usecaseCheckBlacklist,
		usecaseRemoveBlacklist: usecaseRemoveBlacklist,
		register_audit:         register_audit,
		persistence_cache:      persistence_cache,
		producer: producer,
	}
}

func (s *BlacklistService) AddBlacklist(requestInput dto.BlacklistInput) error {
	log.Printf("Start add blacklist witth data: %v %v", requestInput.UserIdentifier, requestInput.EventId)
	ctx := context.Background()
	blacklistEntitty, err := s.usecaseCreateBlacklist.Execute(
		requestInput.UserIdentifier,
		requestInput.EventId,
		requestInput.Reason,
		requestInput.Document,
		requestInput.Scope,
		requestInput.BlockedUntil,
	)
	if err != nil {
		return err
	}
	logDetails := map[string]interface{}{
		"scope":         blacklistEntitty.GetScope(),
		"blocked_type":  blacklistEntitty.GetBlockedType(),
		"blocked_until": blacklistEntitty.GetBlockedUntil(),
		"reason":        blacklistEntitty.GetReason(),
		"is_blocked":    true,
	}
	err = s.register_audit.LogAction(blacklistEntitty.GetUserIdentifier(), blacklistEntitty.GetEventId(), contracts.ADD_BLACKLIST, logDetails)
	if err != nil {
		log.Println(err.Error())
	}
	key_cache := s.generate_key_cache(blacklistEntitty.GetUserIdentifier(), blacklistEntitty.GetEventId())
	err = s.persistence_cache.SetCache(ctx, key_cache, logDetails, nil)
	log.Println(fmt.Printf("result action set cache %v", err))

	s.producer.NotifyBlacklist(*blacklistEntitty)

	log.Println("Finish completed add blacklist")
	return nil
}

func (s *BlacklistService) CheckBlacklist(requestInput dto.BlacklistInputCheck) (dto.BlacklistOutputCheck, error) {
	userIdentifier, eventId := requestInput.UserIdentifier, requestInput.EventId
	log.Printf("Start check blacklist witth data: %v %v", requestInput.UserIdentifier, requestInput.EventId)

	ctx := context.Background()
	key_cache := s.generate_key_cache(userIdentifier, eventId)

	detail_cache, _ := s.persistence_cache.GetCache(ctx, key_cache)

	if detail_cache != nil {
		reason, _ := detail_cache["reason"].(string)
		is_blocked, _ := detail_cache["is_blocked"].(bool)

		logDetails := map[string]interface{}{
			"blocked_type":  detail_cache["blocked_type"].(string),
			"blocked_until": detail_cache["blocked_until"],
			"reason":        reason,
			"is_blocked":    is_blocked,
		}

		err := s.register_audit.LogAction(userIdentifier, eventId, contracts.CHECK_BLACKLIST, logDetails)
		if err != nil {
			log.Println(err.Error())
		}
		return dto.BlacklistOutputCheck{IsBlocked: is_blocked, Reason: reason}, nil
	}

	result, reason := s.usecaseCheckBlacklist.Execute(userIdentifier, eventId)

	return dto.BlacklistOutputCheck{IsBlocked: result, Reason: reason}, nil

}

func (s *BlacklistService) RemoveBlacklist(requestInput dto.BlacklistInputRemove) error {
	userIdentifier, eventId := requestInput.UserIdentifier, requestInput.EventId

	log.Printf("Start remove blacklist witth data: %v %v", userIdentifier, eventId)
	ctx := context.Background()
	key_cache := s.generate_key_cache(userIdentifier, eventId)

	err := s.usecaseRemoveBlacklist.Execute(userIdentifier, eventId)
	if err != nil {
		return err
	}
	err = s.persistence_cache.DeleteCache(ctx, key_cache)
	if err != nil {
		panic(err)
	}
	return nil

}

func (s *BlacklistService) generate_key_cache(userIdentifier int, eventId string) string {
	return fmt.Sprintf("%v_%v", userIdentifier, eventId)
}
