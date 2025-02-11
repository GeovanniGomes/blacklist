package blacklist

type ICheckBlacklist interface {
	Execute(userIndentifier int, evendId string) (string, error)
}
