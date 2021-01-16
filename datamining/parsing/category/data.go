package category

import (
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
)

type Data struct {
	parsing.Data

	JNJ     *JNJData
	Classic *ClassicData
}

type JNJData struct {
	ID               domain.CategoryID
	Sex              domain.Sex
	TotalCompetitors int
}

type ClassicData struct {
	ID               domain.CategoryID
	TotalCompetitors int
}
