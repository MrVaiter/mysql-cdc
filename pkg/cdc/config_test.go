package cdc

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config reading", func() {
	It("Config read successfully", func() {
		config, err := ReadConfig()
		Expect(err).To(BeNil())
		Expect(reflect.TypeOf(config).String()).To(Equal("*cdc.Config"))
	})
})
