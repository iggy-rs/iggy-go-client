package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CREATE USER:", func() {
	When("User is logged in", func() {
		Context("tries to create user with correct data", func() {
			client := createAuthorizedConnection()

			request := iggcon.CreateUserRequest{
				Username: createRandomString(16),
				Password: createRandomString(16),
				Status:   iggcon.Active,
				Permissions: &iggcon.Permissions{
					Global: iggcon.GlobalPermissions{
						ManageServers: true,
						ReadServers:   true,
						ManageUsers:   true,
						ReadUsers:     true,
						ManageStreams: true,
						ReadStreams:   true,
						ManageTopics:  true,
						ReadTopics:    true,
						PollMessages:  true,
						SendMessages:  true,
					},
				},
			}

			err := client.CreateUser(request)
			defer deleteUserAfterTests(request.Username, client)

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateUser(request.Username, client)
			//itShouldBePossibleToLogInWithCredentials(request.Username, request.Password)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to create user", func() {
			client := createConnection()
			request := iggcon.CreateUserRequest{
				Username: createRandomString(16),
				Password: createRandomString(16),
				Status:   iggcon.Active,
				Permissions: &iggcon.Permissions{
					Global: iggcon.GlobalPermissions{
						ManageServers: true,
						ReadServers:   true,
						ManageUsers:   true,
						ReadUsers:     true,
						ManageStreams: true,
						ReadStreams:   true,
						ManageTopics:  true,
						ReadTopics:    true,
						PollMessages:  true,
						SendMessages:  true,
					},
				},
			}

			err := client.CreateUser(request)
			itShouldReturnUnauthenticatedError(err)
		})
	})
})
