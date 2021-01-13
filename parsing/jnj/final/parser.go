package final

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

var (
	tableBoundary = regexp.MustCompile(`^-+\+-+\+-+$`)
	placesLine    = regexp.MustCompile(`^\s*(?P<PartnerCompetitionNumber>\d+)\s+\|\s+(?P<Places>(\d\s)+)\s*\|\s+(?P<TotalPlace>\d)`)
	judgeLables   = "abcdefgji"
)

type finalTableState int

func (s finalTableState) next() finalTableState {
	switch s {
	case finalTableStateNone:
		return finalTableStateHeaderStarted
	case finalTableStateHeaderStarted:
		return finalTableStateBodyStarted
	case finalTableStateBodyStarted:
		return finalTableStateBodyFinished
	default:
		panic(fmt.Sprintf("unexpected state %d", s))
	}
}

const (
	finalTableStateNone finalTableState = iota
	finalTableStateHeaderStarted
	finalTableStateBodyStarted
	finalTableStateBodyFinished
)

type parser struct {
	finalTableState finalTableState

	data Data
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {
	if p.finalTableState == finalTableStateBodyFinished {
		return parsing.LineProcessingStatusAnotherBlock, nil
	}

	if tableBoundary.MatchString(line) {
		p.finalTableState = p.finalTableState.next()

		return parsing.LineProcessingStatusOk, nil
	}

	switch p.finalTableState {
	case finalTableStateHeaderStarted:
		return parsing.LineProcessingStatusOk, nil
	case finalTableStateNone:
		return parsing.LineProcessingStatusAnotherBlock, nil
	case finalTableStateBodyStarted:
		participantPlaces, ok, err := p.parseFinalPlaces(line)
		if err != nil {
			return parsing.LineProcessingStatusNone, err
		}

		if !ok {
			return parsing.LineProcessingStatusOk, nil
		}

		p.data.Places = append(p.data.Places, participantPlaces)
		return parsing.LineProcessingStatusOk, nil

	default:
		panic(fmt.Sprintf("unexpected state %d", p.finalTableState))
	}
}

func (p *parser) parseFinalPlaces(line string) (ParticipantPlaces, bool, error) {
	submatches := placesLine.FindStringSubmatch(line)
	if len(submatches) < 2 {
		return ParticipantPlaces{}, false, nil
	}

	var places []int
	for _, s := range strings.Split(strings.TrimSpace(submatches[2]), " ") {
		p, err := strconv.Atoi(s)
		if err != nil {
			return ParticipantPlaces{}, false, err
		}

		places = append(places, p)
	}

	judgePlaces := make([]Place, 0, len(places))
	for i, p := range places {
		judgePlaces = append(judgePlaces, Place{
			JudgeLabel: domain.JudgeLabel(judgeLables[i]),
			Place:      p,
		})
	}

	return ParticipantPlaces{
		ParticipantID: domain.CompetitionParticipantID(submatches[1]),
		Places:        judgePlaces,
	}, true, nil
}

func (p *parser) GetData() parsing.Data {
	return p.data
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
