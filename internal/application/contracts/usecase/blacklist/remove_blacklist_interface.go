package blacklist

type RemoveBlackListInterface interface {
	Execute(userIndentifier int, eventId string)(error)
}