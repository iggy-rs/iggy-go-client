package tcp_test

import (
	. "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LOGIN FEATURE:", func() {
	When("user is not logged in", func() {
		Context("and tries to log with correct data", func() {
			client := createMessageStream()
			user, err := client.LogIn(LogInRequest{
				Username: "iggy",
				Password: "iggy",
			})

			itShouldNotReturnError(err)
			itShouldReturnUserId(user, 1)
		})

		Context("and tries to log with invalid credentials", func() {
			client := createMessageStream()
			user, err := client.LogIn(LogInRequest{
				Username: "incorrect",
				Password: "random",
			})

			itShouldReturnError(err)
			itShouldNotReturnUser(user)
		})
	})

	When("user is already logged in", func() {
		Context("and tries to log with correct data", func() {
			client := createAuthorizedStream()
			user, err := client.LogIn(LogInRequest{
				Username: "iggy",
				Password: "iggy",
			})

			itShouldNotReturnError(err)
			itShouldReturnUserId(user, 1)
		})

		Context("and tries to log with invalid credentials", func() {
			client := createAuthorizedStream()
			user, err := client.LogIn(LogInRequest{
				Username: "incorrect",
				Password: "random",
			})

			itShouldReturnError(err)
			itShouldNotReturnUser(user)
		})
	})
})

func itShouldReturnUserId(user *LogInResponse, id uint32) {
	It("should return user id", func() {
		Expect(user.UserId).To(Equal(id))
	})
}

func itShouldNotReturnUser(user *LogInResponse) {
	It("should return user id", func() {
		Expect(user).To(BeNil())
	})
}
