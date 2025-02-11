package depedence_injector

import (
	"fmt"

	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist"
	"github.com/GeovanniGomes/blacklist/internal/application/service"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/contracts"
	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue"

	"github.com/GeovanniGomes/blacklist/internal/infrastructure/queue/consumer"
)

func getDependency[T any](c *ContainerInjection) (T, error) {
	var instance T
	err := c.Invoke(func(dep T) {
		instance = dep
	})
	if err != nil {
		return instance, fmt.Errorf("erro ao obter dependência: %v", err)
	}
	return instance, nil
}

func (container *ContainerInjection) GetUsecaseAddBlacklist() (blacklist.IAddBlacklist, error) {
	return getDependency[blacklist.IAddBlacklist](container)
}

func (container *ContainerInjection) GetUsecaseCheckBlacklist() (blacklist.ICheckBlacklist, error) {
	return getDependency[blacklist.ICheckBlacklist](container)
}

func (container *ContainerInjection) GetUsecaseRemoveBlacklist() (blacklist.IRemoveBlackList, error) {
	return getDependency[blacklist.IRemoveBlackList](container)
}

func (container *ContainerInjection) GetBlacklistService() (*service.BlacklistService, error) {
	return getDependency[*service.BlacklistService](container)
}

func (container *ContainerInjection) GetBlacklistConsumer() (*consumer.BlacklistConsumer, error) {
	return getDependency[*consumer.BlacklistConsumer](container)
}

func (container *ContainerInjection) GetBlacklistReportConsumer() (*consumer.BlacklistReportConsumer, error) {
	return getDependency[*consumer.BlacklistReportConsumer](container)
}

func (container *ContainerInjection) GetNewRabbitMQ() (contracts.IQueue, error) {
	return getDependency[contracts.IQueue](container)
}
func (container *ContainerInjection) GetDispatcher() (*queue.Dispatcher, error) {
	return getDependency[*queue.Dispatcher](container)
}
