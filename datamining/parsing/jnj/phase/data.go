package phase

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

type Data struct {
	parsing.Data

	Phase domain.CompetitionPhase
}
