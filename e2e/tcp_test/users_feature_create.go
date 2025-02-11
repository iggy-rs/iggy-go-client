package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo/v2"
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

		Context("tries to create user with correct data and custom permissions", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream("ss", client)
			topicId, _ := successfullyCreateTopic(streamId, client)

			topicPermissionRequest := iggcon.TopicPermissions{
				ManageTopic:  false,
				ReadTopic:    true,
				PollMessages: true,
				SendMessages: true,
			}
			streamPermissionRequest := iggcon.StreamPermissions{
				ManageStream: false,
				ReadStream:   true,
				ManageTopics: true,
				ReadTopics:   true,
				PollMessages: false,
				SendMessages: true,
				Topics: map[int]*iggcon.TopicPermissions{
					int(topicId): &topicPermissionRequest,
				},
			}

			userStreamPermissions := map[int]*iggcon.StreamPermissions{
				int(streamId): &streamPermissionRequest,
			}

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
					Streams: userStreamPermissions,
				},
			}

			err := client.CreateUser(request)
			defer deleteUserAfterTests(request.Username, client)
			defer deleteStreamAfterTests(streamId, client)

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateUserWithPermissions(request.Username, client, userStreamPermissions)
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
