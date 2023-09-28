package tcp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTcp(t *testing.T) {
	//TODO: Refactor this so it's not just tcp, it's all protocols
	//there should be no conflicts, but clean up iggy/local_data from time to time
	//this assumes there is a running iggy server
	RegisterFailHandler(Fail)
	RunSpecs(t, "My Feature Suite")
}
