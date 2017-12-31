package stats

import (
	"github.com/segmentio/stats/datadog"
	"github.com/segmentio/stats/influxdb"
)

const (
	DefaultEngineClient = "influxdb"
)

var (
	LOGX_STATS_CLIENT_INFLUXDB *influxdb.Client
	LOGX_STATS_CLIENT_DATADOG  *datadog.Client
)

func init() {

	switch DefaultEngineClient {
	case "influxdb":
		LOGX_STATS_CLIENT_INFLUXDB = influxdb.NewClient("localhost:8125")
		stats.Register(LOGX_STATS_CLIENT_INFLUXDB)

	case "datadog":
		LOGX_STATS_CLIENT_DATADOG = datadog.NewClient("localhost:8125")
		stats.Register(LOGX_STATS_CLIENT_DATADOG)

	}
	// defer stats.Flush()
}
