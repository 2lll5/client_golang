package prometheus

import (
	"testing"
	"time"
)

func TestHistogramVecMetricIdleTTLEvictsStaleEntries(t *testing.T) {
	now := time.Unix(0, 0)
	histogram := V2.NewHistogramVec(HistogramVecOpts{
		HistogramOpts: HistogramOpts{
			Name: "ttl_eviction_test_histogram",
			Help: "Histogram used to verify MetricIdleTTL eviction.",
			now:  func() time.Time { return now },
		},
		VariableLabels: UnconstrainedLabels([]string{"label"}),
		MetricIdleTTL:  50 * time.Millisecond,
	})

	for _, label := range []string{"a", "b", "c"} {
		histogram.WithLabelValues(label).Observe(1)
	}
	if got := metricVecSize(histogram.MetricVec); got != 3 {
		t.Fatalf("expected 3 metrics before eviction, got %d", got)
	}

	now = now.Add(51 * time.Millisecond)
	ch := make(chan Metric, 3)
	histogram.Collect(ch)

	if got := metricVecSize(histogram.MetricVec); got != 0 {
		t.Fatalf("expected all metrics to be evicted, got %d", got)
	}
	if got := len(ch); got != 0 {
		t.Fatalf("expected no stale metrics to be collected, got %d", got)
	}
}

func TestHistogramVecMetricIdleTTLKeepsActiveCachedObserver(t *testing.T) {
	now := time.Unix(0, 0)
	histogram := V2.NewHistogramVec(HistogramVecOpts{
		HistogramOpts: HistogramOpts{
			Name: "ttl_active_test_histogram",
			Help: "Histogram used to verify active MetricIdleTTL entries are kept.",
			now:  func() time.Time { return now },
		},
		VariableLabels: UnconstrainedLabels([]string{"label"}),
		MetricIdleTTL:  50 * time.Millisecond,
	})

	active := histogram.WithLabelValues("active")
	stale := histogram.WithLabelValues("stale")
	active.Observe(1)
	stale.Observe(1)

	now = now.Add(40 * time.Millisecond)
	active.Observe(2) // Refresh through a cached observer handle.
	now = now.Add(20 * time.Millisecond)

	ch := make(chan Metric, 2)
	histogram.Collect(ch)

	if got := metricVecSize(histogram.MetricVec); got != 1 {
		t.Fatalf("expected only active metric to remain, got %d", got)
	}
	if got := len(ch); got != 1 {
		t.Fatalf("expected only active metric to be collected, got %d", got)
	}
}

func metricVecSize(vec *MetricVec) int {
	vec.metricMap.mtx.RLock()
	defer vec.metricMap.mtx.RUnlock()

	var total int
	for _, metrics := range vec.metricMap.metrics {
		total += len(metrics)
	}
	return total
}
