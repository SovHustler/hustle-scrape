package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJudgeRegexp(t *testing.T) {
	require.EqualValues(t, []string{
		"a",
		"катунин павел",
	}, judgeRegexp.FindStringSubmatch("1 (a) - катунин павел")[1:])

	require.Len(t, judgeRegexp.FindStringSubmatch("dnd beginner (нерейтинг) (партнеры). участников: 17"), 0)
}

func TestJudgeParser(t *testing.T) {
	parser := NewJudgeBlockParser()

	checkStatus := func(expectedStatus LineProcessingStatus, line string) {
		status, err := parser.Process(line)
		require.NoError(t, err)
		require.EqualValues(t, expectedStatus, status)
	}

	checkStatus(LineProcessingStatusOk, "главный судья:")
	checkStatus(LineProcessingStatusOk, " катунин павел")
	checkStatus(LineProcessingStatusOk, "судейская бригада:")
	checkStatus(LineProcessingStatusOk, "1 (a) - катунин павел")
	checkStatus(LineProcessingStatusOk, "2 (в) - николаева екатерина")
	checkStatus(LineProcessingStatusAnotherBlock, "dnd beginner (нерейтинг) (партнеры). участников: 17")

	data := parser.GetData()
	require.EqualValues(t, &BlockDataJudges{
		MainJudge: "катунин павел",
		Judges: map[JudgeLabel]string{
			"a": "катунин павел",
			"в": "николаева екатерина",
		},
	}, data.Judges)
}
