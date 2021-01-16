package blocksplit_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Sovianum/hustleScrape/datamining/blocksplit"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
	"github.com/Sovianum/hustleScrape/datamining/parsing/category"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/final"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/phase"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/place"
	"github.com/Sovianum/hustleScrape/datamining/parsing/judges"
	"github.com/Sovianum/hustleScrape/datamining/parsing/prefinal"
	"github.com/stretchr/testify/suite"
)

type ProcessorTestSuite struct {
	suite.Suite

	lines []string
}

func TestProcessorTestSuite(t *testing.T) {
	suite.Run(t, &ProcessorTestSuite{})
}

func (s *ProcessorTestSuite) SetupSuite() {
	f, err := os.Open("./test_data.txt")
	s.Require().NoError(err)

	data, err := ioutil.ReadAll(f)
	s.Require().NoError(err)

	stringData := string(data)

	s.lines = strings.Split(stringData, "\n")
}

func (s *ProcessorTestSuite) TestProcessWholePage() {
	p := blocksplit.NewProcessor([]parsing.Parser{
		judges.NewParser(),
		category.NewParser(),
		final.NewParser(),
		phase.NewParser(),
		place.NewParser(),
		prefinal.NewParser(),
	})

	for _, l := range s.lines {
		s.Require().NoError(p.Process(l), l)
	}

	data := p.GetData()
	s.Len(data, 57)
}
