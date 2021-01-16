package main

import (
	"github.com/Sovianum/hustleScrape/datamining/blocksplit"
	"github.com/Sovianum/hustleScrape/datamining/domain"
	"github.com/Sovianum/hustleScrape/datamining/loading/competition"
	"github.com/Sovianum/hustleScrape/datamining/loading/competitions"
	"github.com/Sovianum/hustleScrape/datamining/parsing"
	"github.com/Sovianum/hustleScrape/datamining/parsing/category"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/final"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/phase"
	"github.com/Sovianum/hustleScrape/datamining/parsing/jnj/place"
	"github.com/Sovianum/hustleScrape/datamining/parsing/judges"
	"github.com/Sovianum/hustleScrape/datamining/parsing/prefinal"
	"github.com/Sovianum/hustleScrape/datamining/structuring"
)

func main() {
	pageStart := competitions.FirstNewFormatEvent

	var result []structuring.Data

	for {
		cs, err := competitions.GetCompetitions(pageStart)
		if err != nil {
			panic(err)
		}

		if len(cs) == 0 {
			break
		}

		for _, c := range cs {
			dataParts, err := parseCompetition(c)
			if err != nil {
				panic(err)
			}

			result = append(result, dataParts...)
		}

		pageStart += competitions.EventPageSize
	}

	tables := structuring.GroupToTables(result)

	if err := tables.Write("notebooks/dataset"); err != nil {
		panic(err)
	}
}

func parseCompetition(c competitions.Competition) ([]structuring.Data, error) {
	lines, err := competition.LoadPageRaw(c.URL)
	if err != nil {
		panic(err)
	}

	if len(lines) == 0 {
		return nil, nil
	}

	p := blocksplit.NewProcessor([]parsing.Parser{
		judges.NewParser(),
		category.NewParser(),
		final.NewParser(),
		phase.NewParser(),
		place.NewParser(),
		prefinal.NewParser(),
	})

	for _, line := range lines {
		err = p.Process(line)
		if err != nil {
			panic(err)
		}
	}

	dataParts := p.GetData()

	converter := structuring.NewConverter(domain.CompetitionID(c.Name + "_" + c.StartDate.Format("2006-01-02")))

	var structuredData []structuring.Data
	for _, part := range dataParts {
		structuredData = append(structuredData, converter.Convert(part)...)
	}

	return structuredData, nil
}
