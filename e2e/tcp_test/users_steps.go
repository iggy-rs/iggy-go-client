package tcp_test

import (
	"github.com/iggy-rs/iggy-go-client"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// OPERATIONS

func successfullyCreateUser(name string, client iggy.MessageStream) uint32 {
	request := iggcon.CreateUserRequest{
		Username: name,
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
	itShouldNotReturnError(err)
	user, err := client.GetUser(iggcon.NewIdentifier(name))
	itShouldNotReturnError(err)

	return user.Id
}

// ASSERTIONS

func itShouldSuccessfullyCreateUser(name string, client iggy.MessageStream) {
	user, err := client.GetUser(iggcon.NewIdentifier(name))

	itShouldNotReturnError(err)

	It("should create user with name "+name, func() {
		Expect(user.Username).To(Equal(name))
	})
}

func itShouldSuccessfullyUpdateUser(id uint32, name string, client iggy.MessageStream) {
	user, err := client.GetUser(iggcon.NewIdentifier(name))

	itShouldNotReturnError(err)

	It("should update user with id "+string(rune(id)), func() {
		Expect(user.Id).To(Equal(id))
	})

	It("should update user with name "+name, func() {
		Expect(user.Username).To(Equal(name))
	})
}

func itShouldSuccessfullyDeleteUser(userId int, client iggy.MessageStream) {
	user, err := client.GetUser(iggcon.NewIdentifier(userId))

	itShouldReturnSpecificError(err, "cannot_load_resource")
	It("should not return user", func() {
		Expect(user).To(BeNil())
	})
}

func itShouldSuccessfullyUpdateUserPermissions(userId uint32, client iggy.MessageStream) {
	user, err := client.GetUser(iggcon.NewIdentifier(int(userId)))

	itShouldNotReturnError(err)

	It("should update user permissions with id "+string(rune(userId)), func() {
		Expect(user.Permissions.Global.ManageServers).To(BeFalse())
		Expect(user.Permissions.Global.ReadServers).To(BeFalse())
		Expect(user.Permissions.Global.ManageUsers).To(BeFalse())
		Expect(user.Permissions.Global.ReadUsers).To(BeFalse())
		Expect(user.Permissions.Global.ManageStreams).To(BeFalse())
		Expect(user.Permissions.Global.ReadStreams).To(BeFalse())
		Expect(user.Permissions.Global.ManageTopics).To(BeFalse())
		Expect(user.Permissions.Global.ReadTopics).To(BeFalse())
		Expect(user.Permissions.Global.PollMessages).To(BeFalse())
		Expect(user.Permissions.Global.SendMessages).To(BeFalse())
	})
}

func itShouldBePossibleToLogInWithCredentials(username string, password string) {
	ms := createConnection()

	userId, err := ms.LogIn(iggcon.LogInRequest{
		Username: username,
		Password: password,
	})

	itShouldNotReturnError(err)
	It("should return userId", func() {
		Expect(userId).NotTo(BeNil())
	})
}

//CLEANUP

func deleteUserAfterTests(name any, client iggy.MessageStream) {
	_ = client.DeleteUser(iggcon.NewIdentifier(name))
}
