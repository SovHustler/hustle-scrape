package category

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	Name             string
	Sex              domain.Sex
	TotalCompetitors int
}
