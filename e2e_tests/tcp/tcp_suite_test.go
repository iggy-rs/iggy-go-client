package tcp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTcp(t *testing.T) {
	//this assumes there is a running iggy server
	RegisterFailHandler(Fail)
	RunSpecs(t, "My Feature Suite")
}
