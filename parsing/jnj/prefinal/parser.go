package prefinal

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

var (
	techResultsBorder = regexp.MustCompile(`^-+\+-+$`)
	crosses           = regexp.MustCompile(`[abcdefghij]*\s*\|[abcdefghij]*`)
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

	data Data
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {
	if p.techResultsState == techResultsStateBodyFinished {
		return parsing.LineProcessingStatusAnotherBlock, nil
	}

	if techResultsBorder.MatchString(line) {
		p.techResultsState = p.techResultsState.next()

		return parsing.LineProcessingStatusOk, nil
	}

	switch p.techResultsState {
	case techResultsStateHeaderStarted:
		return parsing.LineProcessingStatusOk, nil
	case techResultsStateNone:
		return parsing.LineProcessingStatusAnotherBlock, nil
	case techResultsStateBodyStarted:
		participantCrosses, err := p.parseCrosses(line)
		if err != nil {
			return parsing.LineProcessingStatusNone, err
		}

		p.data.Crosses = append(p.data.Crosses, participantCrosses)
		return parsing.LineProcessingStatusOk, nil

	default:
		panic(fmt.Sprintf("unexpected state %d", p.techResultsState))
	}
}

func (p *parser) parseCrosses(line string) (ParticipantCrosses, error) {
	parts := strings.SplitN(line, "|", 2)
	participantID := strings.TrimSpace(parts[0])

	crosses := crosses.FindStringSubmatch(parts[1])[0]
	splitCrosses := strings.Split(crosses, "|")

	firstDanceCrosses := make([]domain.JudgeLabel, 0, len(splitCrosses[0]))
	for _, c := range strings.TrimSpace(splitCrosses[0]) {
		firstDanceCrosses = append(firstDanceCrosses, domain.JudgeLabel(c))
	}

	secondDanceCrosses := make([]domain.JudgeLabel, 0, len(splitCrosses[1]))
	for _, c := range strings.TrimSpace(splitCrosses[1]) {
		secondDanceCrosses = append(secondDanceCrosses, domain.JudgeLabel(c))
	}

	return ParticipantCrosses{
		CompetitionParticipantID: domain.CompetitionParticipantID(participantID),
		FirstDanceCrosses:        firstDanceCrosses,
		SecondDanceCrosses:       secondDanceCrosses,
	}, nil
}

func (p *parser) GetData() parsing.Data {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
