package category

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Sovianum/hustleScrape/parsers"
)

var (
	categoryLine = regexp.MustCompile(`^\s*(?P<Name>.+)\s+\((?P<Sex>партнеры|девушки)\)\.\s+участников:\s+(?P<Total>\d+)`)
)

type parser struct {
	data BlockData
}

var _ parsers.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {

	if categoryLine.MatchString(line) {
		return p.parseCategory(line)
	}

	return parsers.LineProcessingStatusAnotherBlock, nil
}

func (p *parser) parseCategory(line string) (parsers.LineProcessingStatus, error) {
	submatches := categoryLine.FindStringSubmatch(line)

	var sex parsers.Sex
	switch submatches[2] {
	case "партнеры":
		sex = parsers.SexMale
	case "девушки":
		sex = parsers.SexFemale
	default:
		panic(fmt.Sprintf("unexpected sex %s", submatches[2]))
	}

	totalCompetitors, err := strconv.Atoi(submatches[3])
	if err != nil {
		return parsers.LineProcessingStatusNone, err
	}

	p.data = BlockData{
		Name:             submatches[1],
		Sex:              sex,
		TotalCompetitors: totalCompetitors,
	}

	return parsers.LineProcessingStatusOk, nil
}

func (p *parser) GetData() parsers.DataBlock {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
