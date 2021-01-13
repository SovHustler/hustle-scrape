package structuring

import "github.com/Sovianum/hustleScrape/domain"

type Data interface {
	dataMarker()
}

type Competition struct {
	Data

	ID   domain.CompetitionID
	Name string
}

type Category struct {
	Data

	ID            domain.CategoryID
	CompetitionID domain.CompetitionID
	Type          domain.CategoryType
	Sex           domain.Sex
}

type Participant struct {
	Data

	ID   domain.ParticipantID
	Name string
}

type ParticipantResult struct {
	Data

	ParticipantID domain.ParticipantID
	CategoryID    domain.CategoryID
	PlaceRange    PlaceRange
}

type PlaceRange struct {
	Lower int
	Upper int
}

type Cross struct {
	Data

	ParticipantID domain.ParticipantID
	CompetitionID domain.CompetitionID
	JudgeName     domain.JudgeName
	CategoryID    domain.CategoryID
	Phase         domain.CompetitionPhase
}

type Judge struct {
	Data

	CompetitionID domain.CompetitionID
	CategoryID    domain.CategoryID
	Label         domain.JudgeLabel
	Name          domain.JudgeName
}
