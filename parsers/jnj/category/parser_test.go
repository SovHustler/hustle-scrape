package category

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

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "dnd beginner (нерейтинг) (партнеры). участников: 17")

	data := p.GetData()
	s.EqualValues(BlockData{
		Name:             "dnd beginner (нерейтинг)",
		Sex:              parsers.SexMale,
		TotalCompetitors: 17,
	}, data.(BlockData))
}
