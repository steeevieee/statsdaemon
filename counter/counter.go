package counter

import (
	"github.com/vimeo/statsdaemon/common"
)

func Add(counters map[string]float64, metric *common.Metric) {
	_, ok := counters[metric.Bucket]
	if !ok {
		counters[metric.Bucket] = 0
	}
	counters[metric.Bucket] += metric.Value * float64(1/metric.Sampling)
}