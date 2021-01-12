package final

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Sovianum/hustleScrape/parsers"
)

var (
	tableBoundary = regexp.MustCompile(`-+\+-+\+-+`)
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

type Parser struct {
	finalTableState finalTableState

	data BlockData
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Process(line string) (parsers.LineProcessingStatus, error) {
	if p.finalTableState == finalTableStateBodyFinished {
		return parsers.LineProcessingStatusAnotherBlock, nil
	}

	if tableBoundary.MatchString(line) {
		p.finalTableState = p.finalTableState.next()

		return parsers.LineProcessingStatusOk, nil
	}

	switch p.finalTableState {
	case finalTableStateHeaderStarted:
		return parsers.LineProcessingStatusOk, nil
	case finalTableStateNone:
		return parsers.LineProcessingStatusAnotherBlock, nil
	case finalTableStateBodyStarted:
		participantPlaces, ok, err := p.parseFinalPlaces(line)
		if err != nil {
			return parsers.LineProcessingStatusNone, err
		}

		if !ok {
			return parsers.LineProcessingStatusOk, nil
		}

		p.data.Places = append(p.data.Places, participantPlaces)
		return parsers.LineProcessingStatusOk, nil

	default:
		panic(fmt.Sprintf("unexpected state %d", p.finalTableState))
	}
}

func (p *Parser) parseFinalPlaces(line string) (ParticipantPlaces, bool, error) {
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
			JudgeLabel: parsers.JudgeLabel(judgeLables[i]),
			Place:      p,
		})
	}

	return ParticipantPlaces{
		ParticipantID: parsers.CompetitionParticipantID(submatches[1]),
		Places:        judgePlaces,
	}, true, nil
}

func (p *Parser) GetData() parsers.BlockData {
	return &p.data
}
