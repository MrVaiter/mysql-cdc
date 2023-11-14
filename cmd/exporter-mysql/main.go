package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/rs/zerolog/log"

	"github.com/advantiss/cloudreef/platform/utility/exporter/pkg/cdc"
)

func main() {
	config, err := cdc.ReadConfig()
	if err != nil {
		log.Fatal().Err(err)
	}

	// Create a binlog syncer with a unique server id, the server id must be different from other MySQL's.
	// flavor is mysql or mariadb
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   config.Flavor,
		Host:     config.Host,
		Port:     uint16(config.Port),
		User:     config.User,
		Password: config.Password,
	}
	syncer := replication.NewBinlogSyncer(cfg)

	gtidFlag := flag.String("gtid", "", "gtid to start from")
	gtid, err := mysql.ParseGTIDSet(config.Flavor, *gtidFlag)

	if gtidFlag == nil || len(*gtidFlag) == 0 {
		GTIDs, err := cdc.GetGTIDs(config)
		if err != nil {
			log.Fatal().Err(err)
		}

		lastGTID, err := mysql.ParseGTIDSet(config.Flavor, GTIDs[len(GTIDs)-1])
		if err != nil {
			log.Fatal().Err(err)
		}

		gtid = lastGTID
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigs
		cancel()
	}()

	cdc.ReadBinLogs(ctx, syncer, &gtid)
}