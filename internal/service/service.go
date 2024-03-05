package service

type IService interface {
	SearchLiterature(literatureList []string) (bool, error)
}

type Service struct {
}
