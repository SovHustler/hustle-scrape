package judges

import (
	"regexp"
	"strings"

	"github.com/Sovianum/hustleScrape/parsers"
)

var judgeRegexp = regexp.MustCompile(`^\d+\s\((?P<JudgeLabel>.)\)\s-\s(?P<JudgeName>.*)`)

type parser struct {
	mainJudgeHeaderFlag bool

	data DataBlock
}

var _ parsers.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{
		data: DataBlock{
			Judges: make(map[parsers.JudgeLabel]string),
		},
	}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {
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

func (p *parser) GetData() parsers.DataBlock {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
