package prefinal

import (
	"testing"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/stretchr/testify/suite"
)

type TechTestSuite struct {
	parsing.TestSuite
}

func TestTechTestSuite(t *testing.T) {
	suite.Run(t, &TechTestSuite{})
}

func (s *TechTestSuite) TestParser() {
	p := NewParser()

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+---------------------------------")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, " пары   |  кресты (буквы судей)")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+---------------------------------")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "  109   | ab|cde ==> выход в 1/4 финала")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "   79   |   |   место: 23-28")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+---------------------------------")

	data := p.GetData()
	s.EqualValues(Data{
		Crosses: []ParticipantCrosses{
			{
				ParticipantID:      "109",
				FirstDanceCrosses:  []domain.JudgeLabel{"a", "b"},
				SecondDanceCrosses: []domain.JudgeLabel{"c", "d", "e"},
			},
			{
				ParticipantID:      "79",
				FirstDanceCrosses:  []domain.JudgeLabel{},
				SecondDanceCrosses: []domain.JudgeLabel{},
			},
		},
	}, data.(Data))
}
