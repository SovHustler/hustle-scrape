package judges

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type DataBlock struct {
	parsing.Data

	MainJudge string
	Judges    map[domain.JudgeLabel]string
}
