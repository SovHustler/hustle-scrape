package parser

import (
	"regexp"
	"strings"
)

var judgeRegexp = regexp.MustCompile(`^\d+\s\((?P<JudgeLabel>.)\)\s-\s(?P<JudgeName>.*)`)

type JudgeLabel string

type BlockDataJudges struct {
	MainJudge string
	Judges    map[JudgeLabel]string
}

type JudgeBlockParser struct {
	mainJudgeHeaderFlag bool

	data BlockDataJudges
}

func NewJudgeBlockParser() *JudgeBlockParser {
	return &JudgeBlockParser{
		data: BlockDataJudges{
			Judges: make(map[JudgeLabel]string),
		},
	}
}

func (p *JudgeBlockParser) Process(line string) (LineProcessingStatus, error) {
	switch {
	case strings.HasPrefix(line, "главный судья"):
		p.mainJudgeHeaderFlag = true
		return LineProcessingStatusOk, nil

	case p.mainJudgeHeaderFlag:
		p.data.MainJudge = strings.TrimSpace(line)
		p.mainJudgeHeaderFlag = false
		return LineProcessingStatusOk, nil

	case strings.HasPrefix(line, "судейская бригада"):
		return LineProcessingStatusOk, nil

	default:
		submatches := judgeRegexp.FindStringSubmatch(line)
		if len(submatches) == 0 {
			return LineProcessingStatusAnotherBlock, nil
		}

		p.data.Judges[JudgeLabel(submatches[1])] = submatches[2]
		return LineProcessingStatusOk, nil
	}
}

func (p *JudgeBlockParser) GetData() BlockData {
	return BlockData{
		Judges: &p.data,
	}
}
