package structuring

import (
	"fmt"
)

func GroupToTables(dataBlocks ...[]Data) Tables {
	var result Tables

	for _, blocks := range dataBlocks {
		for _, block := range blocks {
			switch casted := block.(type) {
			case Judge:
				result.Judges = append(result.Judges, casted)

			case Category:
				result.Categories = append(result.Categories, casted)

			case ParticipantResult:
				result.Results = append(result.Results, casted)

			case Participant:
				result.Participants = append(result.Participants, casted)

			case Cross:
				result.Crosses = append(result.Crosses, casted)

			default:
				panic(fmt.Sprintf("unexpected type %T", casted))
			}
		}
	}

	return result
}

type Tables struct {
	Competitions []Competition
	Categories   []Category
	Participants []Participant
	Results      []ParticipantResult
	Crosses      []Cross
	Judges       []Judge
}
