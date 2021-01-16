package place

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/stretchr/testify/suite"
)

type ResultsTestSuite struct {
	parsing.TestSuite
}

func TestResultsTestSuite(t *testing.T) {
	suite.Run(t, &ResultsTestSuite{})
}

func (s *ResultsTestSuite) TestParser() {
	p := NewParser()

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "1 место-№366-рябов михаил александрович(10465,ivara,d,bg)")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "9-10 место-№244-федечкин сергей николаевич(11052,движение,d,bg)")

	data := p.GetData()
	s.EqualValues(Data{
		Results: []Result{
			{
				PlaceRange:               PlaceRange{1, 1},
				ParticipantID:            "10465",
				CompetitionParticipantID: "366",
			},
			{
				PlaceRange:               PlaceRange{9, 10},
				ParticipantID:            "11052",
				CompetitionParticipantID: "244",
			},
		},
	}, data.(Data))
}

func (s *ResultsTestSuite) TestClassicNotParsed() {
	p := NewParser()

	s.CheckStatus(p, parsing.LineProcessingStatusAnotherBlock, "1 место-№64-корзинин кирилл константинович(дебют,ivara,e)-даниеле даугелайте(дебют,ivara,e)")
}
