package dal

var client Factory

type Factory interface {
	Users() UserSDal
	Policies() PolicyDal
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
