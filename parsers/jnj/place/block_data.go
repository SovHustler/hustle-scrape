package place

import (
	"github.com/Sovianum/hustleScrape/parsers"
)

type BlockData struct {
	parsers.BlockData

	Results []JNJResult
}

type JNJResult struct {
	PlaceRange PlaceRange
	ID         parsers.ParticipantID
}

type PlaceRange struct {
	Lower int
	Upper int
}
