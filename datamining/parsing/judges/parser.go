package judges

import (
	"regexp"
	"sort"
	"strings"

	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

var judgeRegexp = regexp.MustCompile(`^\d+\s\((?P<JudgeLabel>.)\)\s-\s(?P<JudgeName>.*)`)

type parser struct {
	mainJudgeHeaderFlag bool

	data Data
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {
	switch {
	case strings.HasPrefix(line, "главный судья"):
		p.mainJudgeHeaderFlag = true
		return parsing.LineProcessingStatusOk, nil

	case p.mainJudgeHeaderFlag:
		p.data.MainJudge = strings.TrimSpace(line)
		p.mainJudgeHeaderFlag = false
		return parsing.LineProcessingStatusOk, nil

	case strings.HasPrefix(line, "судейская бригада"):
		return parsing.LineProcessingStatusOk, nil

	default:
		submatches := judgeRegexp.FindStringSubmatch(line)
		if len(submatches) == 0 {
			return parsing.LineProcessingStatusAnotherBlock, nil
		}

		label := domain.JudgeLabel(submatches[1])

		switch label { // substitute english letters
		case "в":
			label = "b"

		case "с":
			label = "c"
		}

		rawJudgeName := submatches[2]
		judgeNameParts := strings.Split(rawJudgeName, " ")
		sort.Strings(judgeNameParts)
		judgeName := strings.Join(judgeNameParts, " ")

		p.data.Judges = append(p.data.Judges, Judge{
			Label: label,
			Name:  domain.JudgeName(judgeName),
		})

		return parsing.LineProcessingStatusOk, nil
	}
}

func (p *parser) GetData() parsing.Data {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
