package judges

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	MainJudge string
	Judges    []Judge
}

type Judge struct {
	Label domain.JudgeLabel
	Name  string
}
