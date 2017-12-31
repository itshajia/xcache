package stats

import (
	"github.com/segmentio/stats"
	"github.com/segmentio/stats/grafana"
	"github.com/segmentio/stats/promotheus"
)

const (
	DefaultDataFormatter string = "none"
)

var (
	LOGX_STATS_DATASOURCE stats.Handler
)

func init() {

	switch DefaultDataFormatter {
	case "grafana":
		LOGX_STATS_DATASOURCE = grafana.Handler

	case "promotheus":
		LOGX_STATS_DATASOURCE = promotheus.Handler
	}

}
