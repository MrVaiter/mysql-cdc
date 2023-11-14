package cdc

import (
	"database/sql"
	"reflect"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GTID Testing", func() {
	It("It checks GTID's mode", func() {
		config, err := ReadConfig()
		Expect(err).To(BeNil())

		db, err := sql.Open(
			config.Flavor,
			config.User+":"+
				config.Password+"@tcp("+
				config.Host+":"+
				strconv.Itoa(config.Port)+")/"+
				config.DB)
		Expect(err).To(BeNil())

		result, err := CheckGTID(db)
		Expect(err).To(BeNil())
		Expect(result).To(BeTrue())

		db.Close()
	})

	It("It contains GTID", func() {
		config, err := ReadConfig()
		Expect(err).To(BeNil())

		GTID, err := GetGTIDs(config)

		Expect(err).To(BeNil())
		Expect(reflect.TypeOf(GTID).String()).To(Equal("[]string"))
	})
})
