package mysql

import (
	"fmt"
	"sync"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/pkg/options"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/store/mysql"
	"gorm.io/gorm"
)

type dbStore struct {
	db *gorm.DB
}

// PolicyAudits implements dal.Factory.
func (ds *dbStore) PolicyAudits() dal.PolicyAuditDal {
	return newPolicyAudits(ds)
}

// Secrets implements dal.Factory.
func (ds *dbStore) Secrets() dal.SecretDal {
	return newSecrets(ds)
}

func (ds *dbStore) Users() dal.UserSDal {
	return newUsers(ds)
}

func (ds *dbStore) Policies() dal.PolicyDal {
	return newPolicies(ds)
}

// Close implements dal.Factory.
func (ds *dbStore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var once sync.Once
var mysqlFactory dal.Factory

func GetMysqlFactoryOr(opts *options.MySQLOptions) (dal.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql factory")
	}

	var err error
	var db *gorm.DB
	once.Do(func() {
		mysqlOpts := mysql.Options{
			Host:            opts.Host,
			User:            opts.User,
			Password:        opts.Password,
			Databasse:       opts.Database,
			MaxIdleConns:    opts.MaxIdleConns,
			MaxOpenConns:    opts.MaxOpenConns,
			MaxConnLifeTime: opts.MaxConnLifeTime,
			LogLevel:        opts.LogLevel,
			// Logger:          logger.New(os.Stdout,&logger.Config{
			// }),

			AutoMigrateTables: []any{},
		}
		db, err = mysql.NewMySQLClient(&mysqlOpts)

		mysqlFactory = &dbStore{db}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql factory, err: %v", err)
	}

	return mysqlFactory, nil
}
