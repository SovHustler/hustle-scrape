package phase

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

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "финал")
	s.EqualValues(&BlockData{
		Phase: 1,
	}, p.GetData().(*BlockData))

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "1/2 финала")
	s.EqualValues(&BlockData{
		Phase: 2,
	}, p.GetData().(*BlockData))
}
