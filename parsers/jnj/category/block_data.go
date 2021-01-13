package category

import (
	"github.com/Sovianum/hustleScrape/parsers"
)

type BlockData struct {
	parsers.BlockData

	Name             string
	Sex              parsers.Sex
	TotalCompetitors int
}
