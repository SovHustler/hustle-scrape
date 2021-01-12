package final

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	parsers.TestSuite
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, &ParserTestSuite{})
}

func (s *ParserTestSuite) TestJNJ() {
	p := NewParser()

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+-------------------+---------")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "        │ места за 1-й.     │итоговое")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "        │ танец             │ место")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+-------------------│ пары")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, " судьи  │ a b c d e f g h i │")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+-------------------+---------")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "  пары  │                   │")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "  109   | 2 2 1 1 1         |    1")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "  389   | 5 5 2 3 3         |    2")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "--------+-------------------+---------")

	data := p.GetData()
	s.EqualValues(&BlockData{
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
	}, data.(*BlockData))
}
