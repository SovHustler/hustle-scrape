package processor

import (
	"fmt"

	"github.com/Sovianum/hustleScrape/parsing"
)

type Processor struct {
	parsers       []parsing.Parser
	currentParser parsing.Parser

	dataBlocks []parsing.Data
}

func NewProcessor(parsers []parsing.Parser) *Processor {
	return &Processor{
		parsers: parsers,
	}
}

func (p *Processor) Process(line string) error {
	if p.currentParser == nil {
		currentParser, err := p.selectNewParser(line)
		if err != nil {
			return err
		}

		p.currentParser = currentParser
		return nil
	}

	status, err := p.currentParser.Process(line)
	if err != nil {
		return err
	}

	switch status {
	case parsing.LineProcessingStatusOk:
		return nil

	case parsing.LineProcessingStatusAnotherBlock:
		p.dataBlocks = append(p.dataBlocks, p.currentParser.GetData())
		p.currentParser.Reset()

		currentParser, err := p.selectNewParser(line)
		if err != nil {
			return err
		}

		p.currentParser = currentParser
		return nil

	default:
		panic(fmt.Sprintf("unexpected status %d", status))
	}
}

func (p *Processor) GetData() []parsing.Data {
	return p.dataBlocks
}

func (p *Processor) selectNewParser(line string) (parsing.Parser, error) {
	for _, parser := range p.parsers {
		parser := parser

		status, err := parser.Process(line)
		if err != nil {
			return nil, err
		}

		if status == parsing.LineProcessingStatusOk {
			return parser, nil
		}
	}

	return nil, nil
}
