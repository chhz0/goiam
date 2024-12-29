package options

import (
	"time"

	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/chhz0/goiam/pkg/store/mysql"
	"github.com/spf13/pflag"
	"gorm.io/gorm"
)

const (
	DefaultMySQLHost       = "127.0.0.1:3306"
	DefaultMySQLUser       = "root"
	DefaultMySQLPass       = ""
	DefaultMySQLDB         = ""
	DefaultMaxIdleConns    = 100
	DefaultMaxOpenConns    = 100
	DefaultMaxConnLifeTime = 60 * time.Second
	DefaultMySQLLogLevel   = 4
)

type MySQLOptions struct {
	Host            string        `json:"host" mapstructure:"host"`
	User            string        `json:"user" mapstructure:"user"`
	Password        string        `json:"password" mapstructure:"password"`
	Database        string        `json:"database" mapstructure:"database"`
	MaxIdleConns    int           `json:"max-idle-conns" mapstructure:"max-idle-conns"`
	MaxOpenConns    int           `json:"max-open-conns" mapstructure:"max-open-conns"`
	MaxConnLifeTime time.Duration `json:"max-conn-life-time" mapstructure:"max-conn-life-time"`

	LogLevel int `json:"log-level" mapstructure:"log-level"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (o *MySQLOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("mysql", pflag.ExitOnError)

	fs.StringVar(&o.Host, "mysql.host", o.Host, ""+
		"MySQL service host address. If left blank, the following related mysql options will be ignored.")

	fs.StringVar(&o.User, "mysql.username", o.User, ""+
		"Username for access to mysql service.")

	fs.StringVar(&o.Password, "mysql.password", o.Password, ""+
		"Password for access to mysql, should be used pair with password.")

	fs.StringVar(&o.Database, "mysql.database", o.Database, ""+
		"Database name for the logicServer to use.")

	fs.IntVar(&o.MaxIdleConns, "mysql.max-idle-connections", o.MaxOpenConns, ""+
		"Maximum idle connections allowed to connect to mysql.")

	fs.IntVar(&o.MaxOpenConns, "mysql.max-open-connections", o.MaxOpenConns, ""+
		"Maximum open connections allowed to connect to mysql.")

	fs.DurationVar(&o.MaxConnLifeTime, "mysql.max-connection-life-time", o.MaxConnLifeTime, ""+
		"Maximum connection life time allowed to connect to mysql.")

	fs.IntVar(&o.LogLevel, "mysql.log-mode", o.LogLevel, ""+
		"Specify gorm log level.")

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (m *MySQLOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*MySQLOptions)(nil)

func (o *MySQLOptions) Client() (*gorm.DB, error) {
	return mysql.NewMySQLClient(&mysql.Options{
		Host:            o.Host,
		User:            o.User,
		Password:        o.Password,
		Databasse:       o.Database,
		MaxIdleConns:    o.MaxIdleConns,
		MaxOpenConns:    o.MaxOpenConns,
		MaxConnLifeTime: o.MaxConnLifeTime,
		LogLevel:        o.LogLevel,
		Logger:          nil,
	})
}

func NewDefaultMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Host:            DefaultMySQLHost,
		User:            DefaultMySQLUser,
		Password:        DefaultMySQLPass,
		Database:        DefaultMySQLDB,
		MaxIdleConns:    DefaultMaxIdleConns,
		MaxOpenConns:    DefaultMaxOpenConns,
		MaxConnLifeTime: DefaultMaxConnLifeTime,
		LogLevel:        DefaultMySQLLogLevel,
	}
}
