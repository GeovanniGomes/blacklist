package blacklist


type CheckBlacklistInterface interface {
	Execute(userIndentifier int, evendId string) (bool, string)
}