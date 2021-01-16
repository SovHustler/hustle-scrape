package domain

type CompetitionID string

type CategoryID string

type CategoryType string

const (
	CategoryTypeJNJ     CategoryType = "jnj"
	CategoryTypeClassic CategoryType = "classic"
)

type JNJLevel string

type ClassicLevel string

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

type ParticipantID string

type CompetitionParticipantID string

type CompetitionPhase int

type JudgeLabel string

type JudgeName string // it is not a very good key, but...
