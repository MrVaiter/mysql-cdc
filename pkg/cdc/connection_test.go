package cdc

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SQL Connection", func() {
	It("Can connect", func() {
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

		rows, err := db.Query("SELECT @@global.gtid_executed;")
		Expect(err).To(BeNil())
		Expect(rows).NotTo(BeNil())

		db.Close()
	})
})