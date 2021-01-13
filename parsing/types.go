package parsing

type Parser interface {
	Process(line string) (LineProcessingStatus, error)
	GetData() Data
	Reset()
}

type Data interface {
	dataMarker()
}

type LineProcessingStatus int

const (
	LineProcessingStatusOk LineProcessingStatus = iota + 1
	LineProcessingStatusAnotherBlock
	LineProcessingStatusNone
)
