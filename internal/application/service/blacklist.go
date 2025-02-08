package service

import (
	"context"
	"fmt"
	"log"

	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/application/interfaces"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/producer"
)

type BlacklistService struct {
	usecaseCreateBlacklist blacklist.IAddBlacklist
	usecaseCheckBlacklist  blacklist.ICheckBlacklist
	usecaseRemoveBlacklist blacklist.IRemoveBlackList
	register_audit         contracts.IAuditLogger
	persistence_cache      contracts.ICache
	producer               *producer.BlacklistProducer
}

func NewBlackListService(
	usecaseCreateBlacklist blacklist.IAddBlacklist,
	usecaseCheckBlacklist blacklist.ICheckBlacklist,
	usecaseRemoveBlacklist blacklist.IRemoveBlackList,
	register_audit contracts.IAuditLogger,
	persistence_cache contracts.ICache,
	producer *producer.BlacklistProducer,
) *BlacklistService {
	return &BlacklistService{
		usecaseCreateBlacklist: usecaseCreateBlacklist,
		usecaseCheckBlacklist:  usecaseCheckBlacklist,
		usecaseRemoveBlacklist: usecaseRemoveBlacklist,
		register_audit:         register_audit,
		persistence_cache:      persistence_cache,
		producer:               producer,
	}
}

func (s *BlacklistService) AddBlacklist(requestInput interfaces.BlacklistInput) error {
	log.Printf("Start add blacklist witth data: %v %v", requestInput.UserIdentifier, requestInput.EventId)
	ctx := context.Background()

	result, err := s.CheckBlacklist(interfaces.BlacklistInputCheck{
		UserIdentifier: requestInput.UserIdentifier,
		EventId:        requestInput.EventId,
	})
	if err != nil {
		return err
	}
	if result.IsBlocked {
		return fmt.Errorf("there is already a blacklist with the user and the event entered")
	}

	blacklistEntitty, err := s.usecaseCreateBlacklist.Execute(
		requestInput.UserIdentifier,
		requestInput.EventId,
		requestInput.Reason,
		requestInput.Document,
		requestInput.Scope,
		requestInput.BlockedUntil,
	)
	if err != nil {
		return fmt.Errorf("the operation to add to the blacklist could not be completed, please try again later")
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

func (s *BlacklistService) CheckBlacklist(requestInput interfaces.BlacklistInputCheck) (interfaces.BlacklistOutputCheck, error) {
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
		return interfaces.BlacklistOutputCheck{IsBlocked: is_blocked, Reason: reason}, nil
	}

	reason, err := s.usecaseCheckBlacklist.Execute(userIdentifier, eventId)

	if err != nil {
		return interfaces.BlacklistOutputCheck{}, err
	}
	var isBlocked = false

	if reason != "" {
		isBlocked = true
	}
	return interfaces.BlacklistOutputCheck{IsBlocked: isBlocked, Reason: reason}, nil

}

func (s *BlacklistService) RemoveBlacklist(requestInput interfaces.BlacklistInputRemove) error {
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
		log.Printf("errore in delete cache %v", err)
	}
	return nil

}

func (s *BlacklistService) generate_key_cache(userIdentifier int, eventId string) string {
	return fmt.Sprintf("%v_%v", userIdentifier, eventId)
}
