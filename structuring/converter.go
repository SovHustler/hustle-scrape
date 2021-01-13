package structuring

import (
	"fmt"

	"github.com/Sovianum/hustleScrape/domain"
	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/Sovianum/hustleScrape/parsing/jnj/category"
	"github.com/Sovianum/hustleScrape/parsing/jnj/final"
	"github.com/Sovianum/hustleScrape/parsing/jnj/phase"
	"github.com/Sovianum/hustleScrape/parsing/jnj/place"
	"github.com/Sovianum/hustleScrape/parsing/jnj/prefinal"
	"github.com/Sovianum/hustleScrape/parsing/judges"
)

type Converter struct {
	competitionID domain.CompetitionID // TODO fill in

	participantMatching  map[domain.CompetitionParticipantID]domain.ParticipantID
	currentJudgesMapping map[domain.JudgeLabel]domain.JudgeName

	currentJudges     judges.Data
	currentCategoryID domain.CategoryID
	currentPhase      domain.CompetitionPhase
}

func NewConverter() *Converter {
	return &Converter{
		participantMatching: map[domain.CompetitionParticipantID]domain.ParticipantID{},
	}
}

func (c *Converter) Convert(data parsing.Data) []Data {
	switch casted := data.(type) {
	case judges.Data:
		return c.consumeJudgesData(casted)

	case category.Data:
		return c.consumeCategoryData(casted)

	case final.Data:
		return nil // todo handle some way

	case phase.Data:
		return c.consumePhaseData(casted)

	case place.Data:
		return c.consumePlaceData(casted)

	case prefinal.Data:
		return c.consumePreFinalData(casted)

	default:
		panic(fmt.Sprintf("unexpected type %T", casted))
	}
}

func (c *Converter) consumeJudgesData(data judges.Data) []Data {
	c.currentJudges = data

	c.currentJudgesMapping = map[domain.JudgeLabel]domain.JudgeName{}
	for _, j := range data.Judges {
		j := j
		c.currentJudgesMapping[j.Label] = j.Name
	}

	return nil
}

func (c *Converter) consumeCategoryData(data category.Data) []Data {
	c.currentCategoryID = data.ID

	var result []Data
	result = append(result, Category{
		ID:            data.ID,
		CompetitionID: c.competitionID,
		Type:          domain.CategoryTypeJNJ,
		Sex:           data.Sex,
	})

	for _, j := range c.currentJudges.Judges {
		result = append(result, Judge{
			CompetitionID: c.competitionID,
			CategoryID:    c.currentCategoryID,
			Label:         j.Label,
			Name:          j.Name,
		})
	}

	return result
}

func (c *Converter) consumePhaseData(data phase.Data) []Data {
	c.currentPhase = data.Phase
	return nil
}

func (c *Converter) consumePlaceData(data place.Data) []Data {
	var result []Data
	for _, r := range data.Results {
		c.participantMatching[r.CompetitionParticipantID] = r.ParticipantID

		result = append(result, ParticipantResult{
			ParticipantID: r.ParticipantID,
			CategoryID:    c.currentCategoryID,
			PlaceRange: PlaceRange{
				Lower: r.PlaceRange.Lower,
				Upper: r.PlaceRange.Upper,
			},
		})
	}

	return result
}

func (c *Converter) consumePreFinalData(data prefinal.Data) []Data {
	var result []Data
	addCross := func(id domain.ParticipantID, label domain.JudgeLabel) {
		result = append(result, Cross{
			ParticipantID: id,
			CompetitionID: c.competitionID,
			JudgeName:     c.currentJudgesMapping[label],
			CategoryID:    c.currentCategoryID,
			Phase:         c.currentPhase,
		})
	}

	for _, crosses := range data.Crosses {
		id := c.participantMatching[crosses.CompetitionParticipantID]

		for _, label := range crosses.FirstDanceCrosses {
			addCross(id, label)
		}

		for _, label := range crosses.SecondDanceCrosses {
			addCross(id, label)
		}
	}

	return result
}
