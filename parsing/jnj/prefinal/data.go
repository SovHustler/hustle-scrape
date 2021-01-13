package prefinal

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	Crosses []ParticipantCrossesJNJ
}

type ParticipantCrossesJNJ struct {
	ParticipantID domain.ParticipantID

	FirstDanceCrosses  []domain.JudgeLabel
	SecondDanceCrosses []domain.JudgeLabel
}
