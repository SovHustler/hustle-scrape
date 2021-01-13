package phase

import "github.com/Sovianum/hustleScrape/parsers"

type BlockData struct {
	parsers.BlockData

	Phase parsers.CompetitionPhase
}
