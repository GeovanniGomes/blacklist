package usecase

import(
	repository "github.com/GeovanniGomes/blacklist/internal/application/contracts/repository"
	"github.com/GeovanniGomes/blacklist/internal/application/contracts/usecase/blacklist")


var _ blacklist.RemoveBlackListInterface = (*RemoveBlackListUseCase)(nil)

type RemoveBlackListUseCase struct {
	blacklist_repository repository.BlackListRepositoryInterface
}

func NewRemoveBlacklist(blacklist_repository repository.BlackListRepositoryInterface) *RemoveBlackListUseCase {
	return &RemoveBlackListUseCase{blacklist_repository: blacklist_repository}
}

func (usecase *RemoveBlackListUseCase) Execute(userIdentifier int, eventId string)(error){
	return usecase.blacklist_repository.Remove(userIdentifier, eventId)
}