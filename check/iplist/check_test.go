package iplist

import (
	"github.com/pwood/middleauth/check"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("returns an error if an entry in the IP network list is not parsable", func(t *testing.T) {
		in := []string{"this is not a CIDR network"}

		_, err := New(in, check.ACCEPT)

		assert.Error(t, err)
	})

	t.Run("returns a new IPList with networks provided and desired result on match", func(t *testing.T) {
		in := []string{"10.0.0.0/8", "172.16.0.0/12"}

		out, err := New(in, check.ACCEPT)

		expectedIPNetworks := []net.IPNet{
			{
				IP:   []byte{10, 0, 0, 0},
				Mask: []byte{255, 0, 0, 0},
			},
			{
				IP:   []byte{172, 16, 0, 0},
				Mask: []byte{255, 240, 0, 0},
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedIPNetworks, out.nets)
		assert.Equal(t, check.ACCEPT, out.result)
	})
}

func TestIPList_Check(t *testing.T) {
	t.Run("returns NEXT and error if http header is not present", func(t *testing.T) {
		networks := []string{"10.0.0.0/8"}

		iplist, err := New(networks, check.ACCEPT)
		assert.NoError(t, err)

		in := http.Request{}

		out, err := iplist.Check(in)

		assert.Error(t, err)
		assert.Equal(t, check.NEXT, out)
	})

	t.Run("returns NEXT and error if http header is present but is not parsable", func(t *testing.T) {
		networks := []string{"10.0.0.0/8"}

		iplist, err := New(networks, check.ACCEPT)
		assert.NoError(t, err)

		in := http.Request{
			Header: http.Header{
				"X-Real-Ip": []string{"this is not a valid ip address"},
			},
		}

		out, err := iplist.Check(in)

		assert.Error(t, err)
		assert.Equal(t, check.NEXT, out)
	})

	t.Run("returns NEXT if a non proxied request does not match the IP list", func(t *testing.T) {
		networks := []string{"10.0.0.0/8"}

		iplist, err := New(networks, check.ACCEPT)
		assert.NoError(t, err)

		in := http.Request{
			Header: http.Header{
				"X-Real-Ip": []string{"192.168.0.1"},
			},
		}

		out, err := iplist.Check(in)

		assert.NoError(t, err)
		assert.Equal(t, check.NEXT, out)
	})

	t.Run("returns ACCEPT if a non proxied request matches the IP list", func(t *testing.T) {
		networks := []string{"10.0.0.0/8"}

		iplist, err := New(networks, check.ACCEPT)
		assert.NoError(t, err)

		in := http.Request{
			Header: http.Header{
				"X-Real-Ip": []string{"10.10.10.10"},
			},
		}

		out, err := iplist.Check(in)

		assert.NoError(t, err)
		assert.Equal(t, check.ACCEPT, out)
	})
}
