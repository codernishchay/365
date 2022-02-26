package repo

type repo struct{}

type MongoRepository interface {
	GetOne() (string, error)
	GetByFilter(filter string) (string, error)
}

func (*repo) GetByFilter(filter string) {
}
func (*repo) GetOne() {}
