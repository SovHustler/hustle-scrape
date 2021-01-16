package final

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
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
