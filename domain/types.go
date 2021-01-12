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

type Category struct {
	ID            CategoryID
	CompetitionID CompetitionID
	Type          CategoryType
	JNJLevel      JNJLevel
	ClassicLevel  ClassicLevel
	Sex           Sex
}

type CompetitorID string

type Competitor struct {
	ID   CompetitorID
	Name string
}

type ParticipantResult struct {
	ParticipantID CompetitorID
	Place         int
}

type CompetitionPhase int

type Cross struct {
	CompetitorID  CompetitorID
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
