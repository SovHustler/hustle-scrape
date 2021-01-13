package phase

import (
	"regexp"
	"strconv"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

var (
	preFinal = regexp.MustCompile(`^\s*1/(?P<Phase>\d+)\sфинала`)
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
	case line == "финал":
		p.data.Phase = 1
		return parsing.LineProcessingStatusOk, nil

	case preFinal.MatchString(line):
		return p.parsePreFinal(line)
	}

	return parsing.LineProcessingStatusAnotherBlock, nil
}

func (p *parser) parsePreFinal(line string) (parsing.LineProcessingStatus, error) {
	submatches := preFinal.FindStringSubmatch(line)

	phase, err := strconv.Atoi(submatches[1])
	if err != nil {
		return parsing.LineProcessingStatusNone, err
	}

	p.data.Phase = domain.CompetitionPhase(phase)

	return parsing.LineProcessingStatusOk, nil
}

func (p *parser) GetData() parsing.Data {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
