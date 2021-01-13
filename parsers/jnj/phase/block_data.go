package phase

import "github.com/Sovianum/hustleScrape/parsers"

type BlockData struct {
	parsers.DataBlock

	Phase parsers.CompetitionPhase
}
