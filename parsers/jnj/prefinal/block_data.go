package prefinal

import (
	"github.com/Sovianum/hustleScrape/parsers"
)

type BlockData struct {
	parsers.DataBlock

	Crosses []ParticipantCrossesJNJ
}

type ParticipantCrossesJNJ struct {
	ParticipantID parsers.ParticipantID

	FirstDanceCrosses  []parsers.JudgeLabel
	SecondDanceCrosses []parsers.JudgeLabel
}
