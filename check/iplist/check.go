package iplist

import (
	"fmt"
	"github.com/pwood/middleauth/check"
	"net"
	"net/http"
)

var _ check.Checker = (*IPList)(nil)

type IPList struct {
	nets   []net.IPNet
	result check.Result
}

func New(stringNets []string, result check.Result) (IPList, error) {
	i := IPList{result: result}

	for _, n := range stringNets {
		_, parsedNet, err := net.ParseCIDR(n)
		if err != nil {
			return IPList{}, fmt.Errorf("cidr parse fail: %s: %w", n, err)
		}

		i.nets = append(i.nets, *parsedNet)
	}

	return i, nil
}

const header string = "X-Real-Ip"

func (i IPList) Check(r http.Request) (check.Result, error) {
	v := r.Header.Values(header)

	if len(v) <= 0 {
		return check.NEXT, fmt.Errorf("remote addr header not found: %s", header)
	}

	parsedIp := net.ParseIP(v[0])
	if parsedIp == nil {
		return check.NEXT, fmt.Errorf("remote addr unparsable: %s", v[0])
	}

	for _, n := range i.nets {
		if n.Contains(parsedIp) {
			return i.result, nil
		}
	}

	return check.NEXT, nil
}
