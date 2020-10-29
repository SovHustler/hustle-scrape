package parser

type LineProcessingStatus int

const (
	LineProcessingStatusOk LineProcessingStatus = iota + 1
	LineProcessingStatusAnotherBlock
	LineProcessingStatusNone
)
