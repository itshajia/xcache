package loggerstats

import (
	"github.com/segmentio/stats/procstats"
)

var (
	LOGX_STATS_COLLECTOR_PROCSTATS         *procstats.Collector
	LOGX_STATS_COLLECTOR_PROCSTATS_METRICS *procstats.Metrics
)

func init() {

	// Start a new collector for the current process, reporting Go metrics.
	LOGX_STATS_COLLECTOR_PROCSTATS_METRICS = procstats.NewGoMetrics()
	LOGX_STATS_COLLECTOR_PROCSTATS = procstats.StartCollector(LOGX_STATS_COLLECTOR_PROCSTATS_METRICS)

	// Gracefully stops stats collection.
	defer LOGX_STATS_COLLECTOR_PROCSTATS.Close()

}
