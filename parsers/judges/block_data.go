package judges

import "github.com/Sovianum/hustleScrape/parser"

type BlockData struct {
	parsers.BlockData

	MainJudge string
	Judges    map[parsers.JudgeLabel]string
}
