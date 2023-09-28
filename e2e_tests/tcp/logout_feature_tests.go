package tcp

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("LOGOUT FEATURE:", func() {
	When("User is logged in", func() {
		Context("and tries to log out", func() {
			client := createAuthorizedStream()
			err := client.LogOut()

			itShouldNotReturnError(err)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to log out", func() {
			client := createMessageStream()
			err := client.LogOut()

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
