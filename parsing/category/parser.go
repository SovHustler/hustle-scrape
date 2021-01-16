package category

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

var (
	jnjCategoryLine     = regexp.MustCompile(`^\s*(?P<ID>.+)\s+\((?P<Sex>партнеры|девушки)\)\.\s+участников:\s+(?P<Total>\d+)`)
	classicCategoryLine = regexp.MustCompile(`^\s*(?P<ID>.+)\.\s+участвовало пар:\s+(?P<Total>\d+)`)
)

type parser struct {
	data Data
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {

	switch {
	case jnjCategoryLine.MatchString(line):
		return p.parseJnjCategory(line)

	case classicCategoryLine.MatchString(line):
		return p.parseClassicCategory(line)

	default:
		return parsing.LineProcessingStatusAnotherBlock, nil
	}
}

func (p *parser) parseClassicCategory(line string) (parsing.LineProcessingStatus, error) {
	submatches := classicCategoryLine.FindStringSubmatch(line)

	totalCompetitors, err := strconv.Atoi(submatches[2])
	if err != nil {
		return parsing.LineProcessingStatusNone, err
	}

	p.data = Data{
		Classic: &ClassicData{
			ID:               domain.CategoryID(submatches[1]),
			TotalCompetitors: totalCompetitors,
		},
	}

	return parsing.LineProcessingStatusOk, nil
}

func (p *parser) parseJnjCategory(line string) (parsing.LineProcessingStatus, error) {
	submatches := jnjCategoryLine.FindStringSubmatch(line)

	var sex domain.Sex
	switch submatches[2] {
	case "партнеры":
		sex = domain.SexMale
	case "девушки":
		sex = domain.SexFemale
	default:
		panic(fmt.Sprintf("unexpected sex %s", submatches[2]))
	}

	totalCompetitors, err := strconv.Atoi(submatches[3])
	if err != nil {
		return parsing.LineProcessingStatusNone, err
	}

	p.data = Data{
		JNJ: &JNJData{
			ID:               domain.CategoryID(submatches[1]),
			Sex:              sex,
			TotalCompetitors: totalCompetitors,
		},
	}

	return parsing.LineProcessingStatusOk, nil
}

func (p *parser) GetData() parsing.Data {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
