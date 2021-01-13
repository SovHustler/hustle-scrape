package place

import (
	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
)

type BlockData struct {
	parsing.Data

	Results []JNJResult
}

type JNJResult struct {
	PlaceRange PlaceRange
	ID         domain.ParticipantID
}

type PlaceRange struct {
	Lower int
	Upper int
}
