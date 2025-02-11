package blacklist

type IRemoveBlackList interface {
	Execute(userIndentifier int, eventId string) error
}
