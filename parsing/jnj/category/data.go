package category

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type Data struct {
	parsing.Data

	ID               domain.CategoryID
	Sex              domain.Sex
	TotalCompetitors int
}
