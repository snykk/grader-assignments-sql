package main_test

import (
	main "a21hc3NpZ25tZW50"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	Describe("read output.txt", func() {
		It("contains CAMP_ID and hash CAMP_ID", func() {
			b, err := ioutil.ReadFile("output.txt")

			Expect(err).To(BeNil())

			data := strings.Split(string(b), " ")
			Expect(data[0]).To(Equal(main.CAMP_ID))

			Expect(data[1]).To(Equal(fmt.Sprintf("%x", sha256.Sum256([]byte(data[0])))))
		})
	})
})
