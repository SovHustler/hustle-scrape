package category

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

var (
	categoryLine = regexp.MustCompile(`^\s*(?P<Name>.+)\s+\((?P<Sex>партнеры|девушки)\)\.\s+участников:\s+(?P<Total>\d+)`)
)

type parser struct {
	data Data
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {

	if categoryLine.MatchString(line) {
		return p.parseCategory(line)
	}

	return parsing.LineProcessingStatusAnotherBlock, nil
}

func (p *parser) parseCategory(line string) (parsing.LineProcessingStatus, error) {
	submatches := categoryLine.FindStringSubmatch(line)

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
		Name:             submatches[1],
		Sex:              sex,
		TotalCompetitors: totalCompetitors,
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
