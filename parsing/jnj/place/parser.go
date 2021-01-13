package place

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/joomcode/errorx"
)

var (
	jnjCompetitorResultRegexp = regexp.MustCompile(`^(?P<PlaceRange>(\d+|\d+-\d+))\sместо-№\d+-(?P<Name>.+)\((?P<ID>(\d+|дебют)),(?P<ClubName>.+),(?P<ClassicLevel>[a-z]+),(?P<JNJLevel>[a-z]+)\)$`)
)

type parser struct {
	results []JNJResult
}

var _ parsing.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsing.LineProcessingStatus, error) {

	switch {
	case jnjCompetitorResultRegexp.MatchString(line):
		return p.parseJNJCompetitor(line)

	default:
		return parsing.LineProcessingStatusAnotherBlock, nil
	}
}

func (p *parser) parseJNJCompetitor(line string) (parsing.LineProcessingStatus, error) {
	submatches := jnjCompetitorResultRegexp.FindStringSubmatch(line)

	placeRange, err := p.parsePlaceRange(submatches[1])
	if err != nil {
		return parsing.LineProcessingStatusNone, errorx.IllegalArgument.Wrap(err, "failed to parse participant place \"%s\"", submatches[1])
	}

	p.results = append(p.results, JNJResult{
		PlaceRange: placeRange,
		ID:         domain.ParticipantID(submatches[4]),
	})

	return parsing.LineProcessingStatusOk, nil
}

func (p *parser) parsePlaceRange(placeRange string) (PlaceRange, error) {
	parts := strings.Split(placeRange, "-")

	if len(parts) == 1 {
		place, err := strconv.Atoi(parts[0])
		if err != nil {
			return PlaceRange{}, err
		}

		return PlaceRange{
			Lower: place,
			Upper: place,
		}, nil
	}

	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return PlaceRange{}, err
	}

	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return PlaceRange{}, err
	}

	return PlaceRange{
		Lower: lower,
		Upper: upper,
	}, err
}

func (p *parser) GetData() parsing.Data {
	return BlockData{
		Results: p.results,
	}
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
