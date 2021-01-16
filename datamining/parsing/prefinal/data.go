package prefinal

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

type Data struct {
	parsing.Data

	Crosses []ParticipantCrosses
}

type ParticipantCrosses struct {
	CompetitionParticipantID domain.CompetitionParticipantID

	Crosses []domain.JudgeLabel
}

func (c *ParticipantCrosses) GetLabelSet() map[domain.JudgeLabel]struct{} {
	result := make(map[domain.JudgeLabel]struct{}, len(c.Crosses))
	for _, label := range c.Crosses {
		result[label] = struct{}{}
	}

	return result
}
