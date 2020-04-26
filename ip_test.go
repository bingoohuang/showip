package ip_test

import (
	"testing"

	"github.com/bingoohuang/ip"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestListAllIPv4(t *testing.T) {
	ips, err := ip.ListAllIPv4()

	assert.Nil(t, err)
	logrus.Infof("ListAllIPv4 %+v", ips)
}

func TestGetOutboundIP(t *testing.T) {
	logrus.Infof("Outbound:%s", ip.Outbound())
	mainIP, ipList := ip.MainIP()
	logrus.Infof("MainIP:%s", mainIP)
	logrus.Infof("ipList:%+v", ipList)
}
