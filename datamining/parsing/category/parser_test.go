package category

import (
	"testing"

	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
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

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "dnd beginner (нерейтинг) (партнеры). участников: 17")

	data := p.GetData()
	s.EqualValues(Data{
		JNJ: &JNJData{
			ID:               "dnd beginner (нерейтинг)",
			Sex:              domain.SexMale,
			TotalCompetitors: 17,
		},
	}, data.(Data))
}
