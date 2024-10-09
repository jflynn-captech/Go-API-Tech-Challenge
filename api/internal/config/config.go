package config

import (
	"fmt"
	"os"

	"jf.go.techchallenge/internal/applog"
)

type Configuration struct {
	RpcTarget string
}

type EnvProvider interface {
	Env(property string) (string, bool)
}

type OsEnvProvider struct{}

func (OsEnvProvider) Env(property string) (string, bool) {
	return os.LookupEnv(property)
}

// Creates and returns a new configuration based on environment variables passed to the executing program.
// Will return an error if any required environment variables are not set.
func New(log *applog.AppLogger) (*Configuration, error) {
	return NewWithProvider(log, OsEnvProvider{})
}

func NewWithProvider(log *applog.AppLogger, provider EnvProvider) (*Configuration, error) {
	rpcTarget, rpcTargetSet := provider.Env("RPC_TARGET")
	if !rpcTargetSet {
		return nil, fmt.Errorf("database environment variable must be set")
	}

	config := &Configuration{
		RpcTarget: rpcTarget,
	}

	return config, nil
}
