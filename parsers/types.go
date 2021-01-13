package parsers

type Parser interface {
	Process(line string) (LineProcessingStatus, error)
	GetData() BlockData
}

type LineProcessingStatus int

const (
	LineProcessingStatusOk LineProcessingStatus = iota + 1
	LineProcessingStatusAnotherBlock
	LineProcessingStatusNone
)

type JudgeLabel string

type CompetitionPhase int

type ParticipantID string

type CompetitionParticipantID string

type ClassicLevel string

type JNJLevel string

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)
