package loader

import (
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	prefix     = "prefix-"
	propA      = "prop.a.key"
	propB      = "prop.b.key"
	propAvalue = "prop.a.value"
	propBvalue = "prop.b.value"
)

func TestEnvSource(t *testing.T) {
	os.Setenv(prefix+propA, propAvalue)
	os.Setenv(propB, propBvalue)
	assertSource(t, Env)
}

func TestArgsSource(t *testing.T) {
	os.Args = append(os.Args, prefix+propA+"="+propAvalue)
	os.Args = append(os.Args, propB+"="+propBvalue)
	assertSource(t, Args)
}

func assertSource(t *testing.T, loader func(string) config.Loader) {
	cfg, err := loader(prefix).Load()
	if assert.NotNil(t, cfg) && assert.NoError(t, err) {
		assert.Equal(t, propAvalue, cfg.String(propA))
		assert.False(t, cfg.IsSet(propB))
		assert.Len(t, cfg.AllSettings(), 1)
	}
}
