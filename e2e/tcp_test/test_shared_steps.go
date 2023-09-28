package tcp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func itShouldReturnSpecificError(err error, errorMessage string) {
	It("Should return error:"+errorMessage, func() {
		Expect(err.Error()).To(ContainSubstring(errorMessage))
	})
}

func itShouldReturnUnauthenticatedError(err error) {
	itShouldReturnSpecificError(err, "unauthenticated")
}

func itShouldNotReturnError(err error) {
	It("Should not return error", func() {
		Expect(err).To(BeNil())
	})
}

func itShouldReturnError(err error) {
	It("Should return error", func() {
		Expect(err).ToNot(BeNil())
	})
}
