package phase

import (
	"regexp"
	"strconv"

	"github.com/Sovianum/hustleScrape/parsers"
)

var (
	preFinal = regexp.MustCompile(`^\s*1/(?P<Phase>\d+)\sфинала`)
)

type parser struct {
	data BlockData
}

var _ parsers.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {

	switch {
	case line == "финал":
		p.data.Phase = 1
		return parsers.LineProcessingStatusOk, nil

	case preFinal.MatchString(line):
		return p.parsePreFinal(line)
	}

	return parsers.LineProcessingStatusAnotherBlock, nil
}

func (p *parser) parsePreFinal(line string) (parsers.LineProcessingStatus, error) {
	submatches := preFinal.FindStringSubmatch(line)

	phase, err := strconv.Atoi(submatches[1])
	if err != nil {
		return parsers.LineProcessingStatusNone, err
	}

	p.data.Phase = parsers.CompetitionPhase(phase)

	return parsers.LineProcessingStatusOk, nil
}

func (p *parser) GetData() parsers.DataBlock {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
