package judges

import (
	"regexp"
	"strings"

	"github.com/Sovianum/hustleScrape/parser"
)

var judgeRegexp = regexp.MustCompile(`^\d+\s\((?P<JudgeLabel>.)\)\s-\s(?P<JudgeName>.*)`)

type JudgeBlockParser struct {
	mainJudgeHeaderFlag bool

	data BlockData
}

func NewJudgeBlockParser() *JudgeBlockParser {
	return &JudgeBlockParser{
		data: BlockData{
			Judges: make(map[parsers.JudgeLabel]string),
		},
	}
}

func (p *JudgeBlockParser) Process(line string) (parsers.LineProcessingStatus, error) {
	switch {
	case strings.HasPrefix(line, "главный судья"):
		p.mainJudgeHeaderFlag = true
		return parsers.LineProcessingStatusOk, nil

	case p.mainJudgeHeaderFlag:
		p.data.MainJudge = strings.TrimSpace(line)
		p.mainJudgeHeaderFlag = false
		return parsers.LineProcessingStatusOk, nil

	case strings.HasPrefix(line, "судейская бригада"):
		return parsers.LineProcessingStatusOk, nil

	default:
		submatches := judgeRegexp.FindStringSubmatch(line)
		if len(submatches) == 0 {
			return parsers.LineProcessingStatusAnotherBlock, nil
		}

		p.data.Judges[parsers.JudgeLabel(submatches[1])] = submatches[2]
		return parsers.LineProcessingStatusOk, nil
	}
}

func (p *JudgeBlockParser) GetData() parsers.BlockData {
	return &p.data
}
