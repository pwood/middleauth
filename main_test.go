package main

import (
	"github.com/pwood/middleauth/check"
	"github.com/pwood/middleauth/check/iplist"
	"github.com/pwood/middleauth/check/static"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_constructServerFromEnvironmentConfig(t *testing.T) {
	t.Run("constructs a simple middleauth server with ip allow list, and then reject", func(t *testing.T) {
		cfg := Config{
			ConfigMode:        "ENV",
			ServerHost:        "127.0.0.1",
			ServerPort:        8088,
			PermittedNetworks: []string{"10.10.0.0/16"},
		}

		srv := constructServerFromEnvironmentConfig(cfg)

		assert.Equal(t, "127.0.0.1", srv.Host)
		assert.Equal(t, 8088, srv.Port)

		chkIp, ok := srv.Checks[0].(iplist.IPList)
		assert.True(t, ok)

		chkRes, _ := chkIp.Check(&http.Request{Header: http.Header{"X-Real-Ip": []string{"10.10.0.1"}}})
		assert.Equal(t, check.Accept, chkRes.Result)

		chkStatic, ok := srv.Checks[1].(static.Static)
		chkRes, _ = chkStatic.Check(nil)
		assert.Equal(t, check.Reject, chkRes.Result)

		assert.True(t, ok)
	})
}
