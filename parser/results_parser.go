package parser

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/joomcode/errorx"
)

var (
	dndCategoryRegexp     = regexp.MustCompile(`^(?P<CategoryName>.*)\.\sучастников:\s(?P<TotalParticipants>\d+)$`)
	classicCategoryRegexp = regexp.MustCompile(`^.*участвовало пар:\s\d+$`)
	competitorRegexp      = regexp.MustCompile(`^(?P<Place>\d+)\sместо-№\d+-(?P<Name>.+)\((?P<ID>\d+),(?P<ClubName>.+),(?P<ClassicClass>[a-z]+),(?P<DNDClass>[a-z]+)\)$`)
)

type BlockDataResults struct {
}

type BlockDataCategoryDND struct {
	CategoryName      string
	TotalParticipants int
}

type BlockDataCategoryClassic struct {
	CategoryName string
	TotalPairs   string
}

type PlaceParser struct {
	dndCategoryFlag bool

	totalParticipants int
	category          string
}

func (p *PlaceParser) Process(line string) (LineProcessingStatus, error) {
	switch {
	case dndCategoryRegexp.MatchString(line):
		submatches := dndCategoryRegexp.FindStringSubmatch(line)

		p.category = submatches[1]

		totalParticipants, err := strconv.Atoi(submatches[2])
		if err != nil {
			return LineProcessingStatusNone, errorx.IllegalArgument.Wrap(err, "failed to parse total participants %s", submatches[2])
		}
		p.totalParticipants = totalParticipants

	}
}

func (p *PlaceParser) fillCategoryData() {

}

//type JudgeBlockParser struct {
//	mainJudgeHeaderFlag bool
//
//	mainJudge string
//	judges    map[JudgeLabel]string
//}
//
//func NewJudgeBlockParser() *JudgeBlockParser {
//	return &JudgeBlockParser{
//		judges: make(map[JudgeLabel]string),
//	}
//}
//
//func (p *JudgeBlockParser) Process(line string) (LineProcessingStatus, error) {
//	switch {
//	case strings.HasPrefix(line, "главный судья"):
//		// главный судья
//		p.mainJudgeHeaderFlag = true
//		return LineProcessingStatusOk, nil
//
//	case p.mainJudgeHeaderFlag:
//		p.mainJudge = strings.TrimSpace(line)
//		p.mainJudgeHeaderFlag = false
//		return LineProcessingStatusOk, nil
//
//	case strings.HasPrefix(line, "судейская бригада"):
//		return LineProcessingStatusOk, nil
//
//	default:
//		submatches := judgeRegexp.FindStringSubmatch(line)
//		if len(submatches) == 0 {
//			return LineProcessingStatusAnotherBlock, nil
//		}
//
//		p.judges[JudgeLabel(submatches[1])] = submatches[2]
//		return LineProcessingStatusOk, nil
//	}
//}
//
//func (p *JudgeBlockParser) GetData() BlockData {
//	return BlockData{
//		Judges: &BlockDataJudges{
//			MainJudge: p.mainJudge,
//			Judges:    p.judges,
//		},
//	}
//}
