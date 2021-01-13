package prefinal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Sovianum/hustleScrape/parsers"
)

var (
	techResultsBorder = regexp.MustCompile(`-+\+-+`)
	crossesJNJ        = regexp.MustCompile(`[abcdefghij]*\s*\|[abcdefghij]*`)
)

type techResultsState int

func (s techResultsState) next() techResultsState {
	switch s {
	case techResultsStateNone:
		return techResultsStateHeaderStarted
	case techResultsStateHeaderStarted:
		return techResultsStateBodyStarted
	case techResultsStateBodyStarted:
		return techResultsStateBodyFinished
	default:
		panic(fmt.Sprintf("unexpected state %d", s))
	}
}

const (
	techResultsStateNone techResultsState = iota
	techResultsStateHeaderStarted
	techResultsStateBodyStarted
	techResultsStateBodyFinished
)

type parser struct {
	techResultsState techResultsState

	data BlockData
}

var _ parsers.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {
	if p.techResultsState == techResultsStateBodyFinished {
		return parsers.LineProcessingStatusAnotherBlock, nil
	}

	if techResultsBorder.MatchString(line) {
		p.techResultsState = p.techResultsState.next()

		return parsers.LineProcessingStatusOk, nil
	}

	switch p.techResultsState {
	case techResultsStateHeaderStarted:
		return parsers.LineProcessingStatusOk, nil
	case techResultsStateNone:
		return parsers.LineProcessingStatusAnotherBlock, nil
	case techResultsStateBodyStarted:
		participantCrosses, err := p.parseJNJCrosses(line)
		if err != nil {
			return parsers.LineProcessingStatusNone, err
		}

		p.data.Crosses = append(p.data.Crosses, participantCrosses)
		return parsers.LineProcessingStatusOk, nil

	default:
		panic(fmt.Sprintf("unexpected state %d", p.techResultsState))
	}
}

func (p *parser) parseJNJCrosses(line string) (ParticipantCrossesJNJ, error) {
	parts := strings.SplitN(line, " | ", 2)
	participantID := strings.TrimSpace(parts[0])

	crosses := crossesJNJ.FindStringSubmatch(parts[1])[0]
	splitCrosses := strings.Split(crosses, "|")

	firstDanceCrosses := make([]parsers.JudgeLabel, 0, len(splitCrosses[0]))
	for _, c := range strings.TrimSpace(splitCrosses[0]) {
		firstDanceCrosses = append(firstDanceCrosses, parsers.JudgeLabel(c))
	}

	secondDanceCrosses := make([]parsers.JudgeLabel, 0, len(splitCrosses[1]))
	for _, c := range strings.TrimSpace(splitCrosses[1]) {
		secondDanceCrosses = append(secondDanceCrosses, parsers.JudgeLabel(c))
	}

	return ParticipantCrossesJNJ{
		ParticipantID:      parsers.ParticipantID(participantID),
		FirstDanceCrosses:  firstDanceCrosses,
		SecondDanceCrosses: secondDanceCrosses,
	}, nil
}

func (p *parser) GetData() parsers.DataBlock {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
