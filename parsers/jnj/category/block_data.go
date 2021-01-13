package category

import (
	"github.com/Sovianum/hustleScrape/parsers"
)

type BlockData struct {
	parsers.DataBlock

	Name             string
	Sex              parsers.Sex
	TotalCompetitors int
}
