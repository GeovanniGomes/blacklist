package blacklist

type ICheckBlacklist interface {
	Execute(userIndentifier int, evendId string) (bool, string)
}
