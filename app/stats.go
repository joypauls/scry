package app

// want to
// - track favorites
// - track traversal
// - write to json
// - read json at beginning
// - operate in memory

type StatsTracker struct {
	favorites []string
	counts    map[string]int
}

func (stats *StatsTracker) Read() {
	stats.favorites = append(stats.favorites, "a")
}

func NewStatsTracker() *StatsTracker {
	stats := new(StatsTracker)
	return stats
}
