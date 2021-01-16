package place

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

type Data struct {
	parsing.Data

	Results []Result
}

type Result struct {
	PlaceRange               PlaceRange
	ParticipantID            domain.ParticipantID
	CompetitionParticipantID domain.CompetitionParticipantID
}

type PlaceRange struct {
	Lower int
	Upper int
}
