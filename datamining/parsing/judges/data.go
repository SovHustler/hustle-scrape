package judges

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

type Data struct {
	parsing.Data

	MainJudge string
	Judges    []Judge
}

type Judge struct {
	Label domain.JudgeLabel
	Name  domain.JudgeName
}
