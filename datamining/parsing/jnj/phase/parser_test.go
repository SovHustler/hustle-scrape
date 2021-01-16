package phase

import (
	"testing"

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

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "финал")
	s.EqualValues(Data{
		Phase: 1,
	}, p.GetData().(Data))

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "1/2 финала")
	s.EqualValues(Data{
		Phase: 2,
	}, p.GetData().(Data))
}
