package tcp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("STATE FEATURE:", func() {
	When("user is logged in", func() {
		Context("and tries to log with correct data", func() {
			client := createAuthorizedStream()
			stats, err := client.GetStats()

			itShouldNotReturnError(err)
			It("should return stats", func() {
				Expect(stats).ToNot(BeNil())
			})
		})
	})

	When("user is not logged in", func() {
		Context("and tries get iggy statistics", func() {
			client := createMessageStream()
			stats, err := client.GetStats()

			itShouldReturnUnauthenticatedError(err)
			It("should not return stats", func() {
				Expect(stats).To(BeNil())
			})
		})
	})
})
