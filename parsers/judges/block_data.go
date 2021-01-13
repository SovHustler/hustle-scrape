package judges

import "github.com/Sovianum/hustleScrape/parsers"

type DataBlock struct {
	parsers.DataBlock

	MainJudge string
	Judges    map[parsers.JudgeLabel]string
}
