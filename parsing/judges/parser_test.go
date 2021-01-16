package judges

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	parsing.TestSuite
}

func TestJudgesTestSuite(t *testing.T) {
	suite.Run(t, &ParserTestSuite{})
}

func (s *ParserTestSuite) TestJudgeRegexp() {
	s.EqualValues([]string{
		"a",
		"катунин павел",
	}, judgeRegexp.FindStringSubmatch("1 (a) - катунин павел")[1:])

	s.Len(judgeRegexp.FindStringSubmatch("dnd beginner (нерейтинг) (партнеры). участников: 17"), 0)
}

func (s *ParserTestSuite) TestParser() {
	p := NewParser()

	s.CheckStatus(p, parsing.LineProcessingStatusOk, "главный судья:")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, " катунин павел")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "судейская бригада:")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "1 (a) - катунин павел")
	s.CheckStatus(p, parsing.LineProcessingStatusOk, "2 (в) - николаева екатерина")

	data := p.GetData()
	s.EqualValues(Data{
		MainJudge: "катунин павел",
		Judges: []Judge{
			{"a", "катунин павел"},
			{"b", "николаева екатерина"},
		},
	}, data.(Data))
}
