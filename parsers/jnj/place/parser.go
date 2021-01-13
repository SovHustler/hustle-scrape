package place

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/joomcode/errorx"
)

var (
	jnjCompetitorResultRegexp = regexp.MustCompile(`^(?P<PlaceRange>(\d+|\d+-\d+))\sместо-№\d+-(?P<Name>.+)\((?P<ID>(\d+|дебют)),(?P<ClubName>.+),(?P<ClassicLevel>[a-z]+),(?P<JNJLevel>[a-z]+)\)$`)
)

type parser struct {
	results []JNJResult
}

var _ parsers.Parser = (*parser)(nil)

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {

	switch {
	case jnjCompetitorResultRegexp.MatchString(line):
		return p.parseJNJCompetitor(line)

	default:
		return parsers.LineProcessingStatusAnotherBlock, nil
	}
}

func (p *parser) parseJNJCompetitor(line string) (parsers.LineProcessingStatus, error) {
	submatches := jnjCompetitorResultRegexp.FindStringSubmatch(line)

	placeRange, err := p.parsePlaceRange(submatches[1])
	if err != nil {
		return parsers.LineProcessingStatusNone, errorx.IllegalArgument.Wrap(err, "failed to parse participant place \"%s\"", submatches[1])
	}

	p.results = append(p.results, JNJResult{
		PlaceRange: placeRange,
		ID:         parsers.ParticipantID(submatches[4]),
	})

	return parsers.LineProcessingStatusOk, nil
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

func (p *parser) GetData() parsers.DataBlock {
	return BlockData{
		Results: p.results,
	}
}

func (p *parser) Reset() {
	newParser := NewParser()
	*p = *newParser
}
