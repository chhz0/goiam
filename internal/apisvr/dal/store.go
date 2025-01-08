package dal

var client Factory

type Factory interface {
	Users() UserSDal
	Secrets() SecretDal
	Policies() PolicyDal
	PolicyAudits() PolicyAuditDal
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
