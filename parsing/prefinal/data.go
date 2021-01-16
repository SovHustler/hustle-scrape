package prefinal

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	Crosses []ParticipantCrosses
}

type ParticipantCrosses struct {
	CompetitionParticipantID domain.CompetitionParticipantID

	Crosses []domain.JudgeLabel
}
