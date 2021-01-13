package final

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	Places []ParticipantPlaces
}

type ParticipantPlaces struct {
	ParticipantID domain.CompetitionParticipantID

	Places []Place
}

type Place struct {
	JudgeLabel domain.JudgeLabel
	Place      int
}
