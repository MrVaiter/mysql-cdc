package cdc

import (
	"bytes"
	"context"
	"time"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/rs/zerolog/log"
)

func ReadBinLogs(ctx context.Context, syncer *replication.BinlogSyncer, gtid *mysql.GTIDSet) chan string {
	config, err := ReadConfig()
	if err != nil {
		log.Fatal().Err(err)
	}

	channel := make(chan string)
	var buffer bytes.Buffer
	streamer, _ := syncer.StartSyncGTID(*gtid)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ev, err := streamer.GetEvent(context.Background())

				if err == context.DeadlineExceeded {
					time.Sleep(config.Sleep)
					continue
				}

				buffer.Reset()
				ev.Dump(&buffer)
				channel <- buffer.String()
			}
		}
	}()

	return channel
}
