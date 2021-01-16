package structuring

import (
	"strconv"

	"github.com/Sovianum/hustleScrape/datamining/domain"
)

type Data interface {
	ToStrings() []string
}

type Competition struct {
	ID   domain.CompetitionID
	Name string
}

func (c Competition) ToStrings() []string {
	return []string{
		string(c.ID), c.Name,
	}
}

type Category struct {
	ID            domain.CategoryID
	CompetitionID domain.CompetitionID
	Type          domain.CategoryType
	Sex           domain.Sex
}

func (c Category) ToStrings() []string {
	return []string{
		string(c.ID),
		string(c.CompetitionID),
		string(c.Type),
		string(c.Sex),
	}
}

type Participant struct {
	ID   domain.ParticipantID
	Name string
}

func (p Participant) ToStrings() []string {
	return []string{
		string(p.ID),
		p.Name,
	}
}

type ParticipantResult struct {
	ParticipantID domain.ParticipantID
	CategoryID    domain.CategoryID
	PlaceRange    PlaceRange
}

func (p ParticipantResult) ToStrings() []string {
	return []string{
		string(p.ParticipantID),
		string(p.CategoryID),
		strconv.Itoa(p.PlaceRange.Lower),
		strconv.Itoa(p.PlaceRange.Upper),
	}
}

type PlaceRange struct {
	Lower int
	Upper int
}

type Cross struct {
	ParticipantID domain.ParticipantID
	CompetitionID domain.CompetitionID
	JudgeName     domain.JudgeName
	CategoryID    domain.CategoryID
	Phase         domain.CompetitionPhase
	Passed        bool
}

func (c Cross) ToStrings() []string {
	passedStr := "0"
	if c.Passed {
		passedStr = "1"
	}

	return []string{
		string(c.ParticipantID),
		string(c.CompetitionID),
		string(c.JudgeName),
		string(c.CategoryID),
		strconv.Itoa(int(c.Phase)),
		passedStr,
	}
}

type Judge struct {
	CompetitionID domain.CompetitionID
	CategoryID    domain.CategoryID
	Name          domain.JudgeName
}

func (j Judge) ToStrings() []string {
	return []string{
		string(j.CompetitionID),
		string(j.CategoryID),
		string(j.Name),
	}
}
