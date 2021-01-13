package final

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	parsing.TestSuite
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, &ParserTestSuite{})
}

func (s *ParserTestSuite) TestParser() {
	p := NewParser()

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+-------------------+---------")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "        │ места за 1-й.     │итоговое")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "        │ танец             │ место")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+-------------------│ пары")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, " судьи  │ a b c d e f g h i │")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+-------------------+---------")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "  пары  │                   │")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "  109   | 2 2 1 1 1         |    1")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "  389   | 5 5 2 3 3         |    2")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "--------+-------------------+---------")

	data := p.GetData()
	s.EqualValues(Data{
		Places: []ParticipantPlaces{
			{
				ParticipantID: "109",
				Places: []Place{
					{"a", 2},
					{"b", 2},
					{"c", 1},
					{"d", 1},
					{"e", 1},
				},
			},
			{
				ParticipantID: "389",
				Places: []Place{
					{"a", 5},
					{"b", 5},
					{"c", 2},
					{"d", 3},
					{"e", 3},
				},
			},
		},
	}, data.(Data))
}
