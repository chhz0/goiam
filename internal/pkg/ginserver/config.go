package ginserver

import (
	"net"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	JWT             *JWTInfo
	Mode            string
	Middlewares     []string

	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

// CertKey 证书和密钥
type CertKey struct {
	Cert string
	Key  string
}

// SecureServingInfo  TLS 加密的服务器信息 即HTTPS
type SecureServingInfo struct {
	InsecureServingInfo
	CertKey CertKey
}

func (i SecureServingInfo) Address() string {
	return net.JoinHostPort(i.BindAddress, strconv.Itoa(i.BindPort))
}

// InSecureServingInfo 非 TLS 加密的服务器信息 即HTTP
type InsecureServingInfo struct {
	BindAddress string
	BindPort    int
}

func (i InsecureServingInfo) Address() string {
	return net.JoinHostPort(i.BindAddress, strconv.Itoa(i.BindPort))
}

// JWTInfo JWT 认证信息
type JWTInfo struct {
	Realm      string        // 域名
	Key        string        // 密钥
	Timeout    time.Duration // 过期时间
	MaxRefresh time.Duration // 最大刷新时间
}

func NewZeroConfig() *Config {
	return &Config{
		SecureServing:   &SecureServingInfo{},
		InsecureServing: &InsecureServingInfo{},
		JWT: &JWTInfo{
			Realm:      "",
			Key:        "",
			Timeout:    1 * time.Hour,
			MaxRefresh: 20 * time.Minute,
		},
		Mode:            gin.ReleaseMode,
		Middlewares:     make([]string, 0),
		Healthz:         false,
		EnableProfiling: false,
		EnableMetrics:   false,
	}
}

func (c *Config) Complete() complateConfig {
	return complateConfig{c}
}

type complateConfig struct {
	*Config
}

func (c complateConfig) Server() (*Server, error) {
	gin.SetMode(c.Mode)

	s := &Server{
		Conf: c.Config,

		Engine: gin.New(),
	}
	s.init()

	return s, nil
}
