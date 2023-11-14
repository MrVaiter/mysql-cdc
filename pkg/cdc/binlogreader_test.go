package cdc

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinLog Reading", func() {
	It("It contains something", func() {
		ctx := context.Background()
		config, err := ReadConfig()
		Expect(err).To(BeNil())

		cfg := replication.BinlogSyncerConfig{
			ServerID: 100,
			Flavor:   config.Flavor,
			Host:     config.Host,
			Port:     uint16(config.Port),
			User:     config.User,
			Password: config.Password,
		}
		syncer := replication.NewBinlogSyncer(cfg)

		GTIDs, err := GetGTIDs(config)
		Expect(err).To(BeNil())
		
		lastGTID, err := mysql.ParseGTIDSet(config.Flavor, GTIDs[len(GTIDs)-1])
		Expect(err).To(BeNil())

		imitation()
		channel := ReadBinLogs(ctx, syncer, &lastGTID)

		Eventually(<-channel, "50ms").ShouldNot(BeNil())
	})
})

func imitation() {
	config, err := ReadConfig()
	Expect(err).To(BeNil())

	db, err := sql.Open(
		config.Flavor,
		config.User+":"+
			config.Password+"@tcp("+
			config.Host+":"+
			strconv.Itoa(config.Port)+")/"+
			config.DB)

	rows, err := db.Query("use mydata")
	Expect(err).To(BeNil())
	Expect(rows)

	rows, err = db.Query("insert into friends values (11, 'Katy', 25, 'f');")
	Expect(err).To(BeNil())
	
	rows, err = db.Query("update friends set age=25 where name='Anjali';")
	Expect(err).To(BeNil())

	rows, err = db.Query("delete from friends where id=11;")
	Expect(err).To(BeNil())

	db.Close()
}
