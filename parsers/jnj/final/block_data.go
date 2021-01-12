package final

import (
	"github.com/Sovianum/hustleScrape/parsers"
)

type BlockData struct {
	parsers.BlockData

	Places []ParticipantPlaces
}

type ParticipantPlaces struct {
	ParticipantID parsers.CompetitionParticipantID

	Places []Place
}

type Place struct {
	JudgeLabel parsers.JudgeLabel
	Place      int
}
