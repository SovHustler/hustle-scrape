package judges

import (
	"testing"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/stretchr/testify/suite"
)

type JudgesTestSuite struct {
	parsers.TestSuite
}

func TestJudgesTestSuite(t *testing.T) {
	suite.Run(t, &JudgesTestSuite{})
}

func (s *JudgesTestSuite) TestJudgeRegexp() {
	s.EqualValues([]string{
		"a",
		"катунин павел",
	}, judgeRegexp.FindStringSubmatch("1 (a) - катунин павел")[1:])

	s.Len(judgeRegexp.FindStringSubmatch("dnd beginner (нерейтинг) (партнеры). участников: 17"), 0)
}

func (s *JudgesTestSuite) TestJudgeParser() {
	p := NewParser()

	s.CheckStatus(p, parsers.LineProcessingStatusOk, "главный судья:")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, " катунин павел")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "судейская бригада:")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "1 (a) - катунин павел")
	s.CheckStatus(p, parsers.LineProcessingStatusOk, "2 (в) - николаева екатерина")
	s.CheckStatus(p, parsers.LineProcessingStatusAnotherBlock, "dnd beginner (нерейтинг) (партнеры). участников: 17")

	data := p.GetData()
	s.EqualValues(DataBlock{
		MainJudge: "катунин павел",
		Judges: map[parsers.JudgeLabel]string{
			"a": "катунин павел",
			"в": "николаева екатерина",
		},
	}, data.(DataBlock))
}
