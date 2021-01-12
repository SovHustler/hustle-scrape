package prefinal

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/stretchr/testify/suite"
)

type TechTestSuite struct {
	parsers.TestSuite
}

func TestTechTestSuite(t *testing.T) {
	suite.Run(t, &TechTestSuite{})
}

func (s *TechTestSuite) TestJNJ() {
	p := NewParser()

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+---------------------------------")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, " пары   |  кресты (буквы судей)")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+---------------------------------")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "  109   | ab|cde ==> выход в 1/4 финала")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "   79   |   |   место: 23-28")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+---------------------------------")

	data := p.GetData()
	s.EqualValues(&BlockData{
		Crosses: []ParticipantCrossesJNJ{
			{
				ParticipantID:      "109",
				FirstDanceCrosses:  []parsers.JudgeLabel{"a", "b"},
				SecondDanceCrosses: []parsers.JudgeLabel{"c", "d", "e"},
			},
			{
				ParticipantID:      "79",
				FirstDanceCrosses:  []parsers.JudgeLabel{},
				SecondDanceCrosses: []parsers.JudgeLabel{},
			},
		},
	}, data.(*BlockData))
}
