package services

type RepoInterfacefyc interface {
	Push(PATH, filename, content string) (string, string, string)
	Del(filepath, sha string) string
}