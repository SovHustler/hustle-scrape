package parsers

import (
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) CheckStatus(p Parser, expectedStatus LineProcessingStatus, line string) {
	status, err := p.Process(line)
	s.Require().NoError(err)
	s.EqualValues(expectedStatus, status)
}
