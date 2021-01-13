package domain

type CompetitionID string

type Competition struct {
	ID   CompetitionID
	Name string
}

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

type Category struct {
	ID            CategoryID
	CompetitionID CompetitionID
	Type          CategoryType
	JNJLevel      JNJLevel
	ClassicLevel  ClassicLevel
	Sex           Sex
}

type ParticipantID string

type CompetitionParticipantID string

type Competitor struct {
	ID   ParticipantID
	Name string
}

type ParticipantResult struct {
	ParticipantID ParticipantID
	Place         int
}

type CompetitionPhase int

type Cross struct {
	CompetitorID  ParticipantID
	CompetitionID CompetitionID
	JudgeLabel    JudgeLabel
	CategoryID    CategoryID
	Phase         CompetitionPhase
}

type JudgeLabel string

type Judge struct {
	CompetitionID CompetitionID
	CategoryID    CategoryID
	Label         JudgeLabel
	Name          string
}
