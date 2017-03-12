package game

const PARTstats = "stats"

type Stats map[string]*Stat

func (s Stats) Part() string {
	return PARTstats
}

func (parts Parts) GetStats() Stats {
	stats, _ := parts[PARTstats].(Stats)
	return stats
}
