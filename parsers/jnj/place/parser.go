package place

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Sovianum/hustleScrape/parsers"
	"github.com/joomcode/errorx"
)

var (
	jnjCompetitorResultRegexp = regexp.MustCompile(`^(?P<PlaceRange>(\d+|\d+-\d+))\sместо-№\d+-(?P<Name>.+)\((?P<ID>\d+),(?P<ClubName>.+),(?P<ClassicLevel>[a-z]+),(?P<JNJLevel>[a-z]+)\)$`)
	preFinal                  = regexp.MustCompile(`^\s*1/\d\s+финала`)
)

type parser struct {
	results []JNJResult
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Process(line string) (parsers.LineProcessingStatus, error) {

	switch {
	case line == "финал" || preFinal.MatchString(line):
		return parsers.LineProcessingStatusOk, nil

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

func (p *parser) GetData() parsers.BlockData {
	return &BlockData{
		Results: p.results,
	}
}
