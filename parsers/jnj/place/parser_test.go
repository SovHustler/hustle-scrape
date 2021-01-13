package place

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/stretchr/testify/suite"
)

type ResultsTestSuite struct {
	parsers.TestSuite
}

func TestResultsTestSuite(t *testing.T) {
	suite.Run(t, &ResultsTestSuite{})
}

func (s *ResultsTestSuite) TestJNJ() {
	p := NewParser()

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "1 место-№366-рябов михаил александрович(10465,ivara,d,bg)")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "9-10 место-№244-федечкин сергей николаевич(11052,движение,d,bg)")

	data := p.GetData()
	s.EqualValues(BlockData{
		Results: []JNJResult{
			{
				PlaceRange: PlaceRange{1, 1},
				ID:         "10465",
			},
			{
				PlaceRange: PlaceRange{9, 10},
				ID:         "11052",
			},
		},
	}, data.(BlockData))
}
