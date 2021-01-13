package main

import (
	"github.com/Sovianum/hustleScrape/blocksplit"
	"github.com/Sovianum/hustleScrape/loading"
	"github.com/Sovianum/hustleScrape/parsing"
	"github.com/Sovianum/hustleScrape/parsing/jnj/category"
	"github.com/Sovianum/hustleScrape/parsing/jnj/final"
	"github.com/Sovianum/hustleScrape/parsing/jnj/phase"
	"github.com/Sovianum/hustleScrape/parsing/jnj/place"
	"github.com/Sovianum/hustleScrape/parsing/jnj/prefinal"
	"github.com/Sovianum/hustleScrape/parsing/judges"
	"github.com/Sovianum/hustleScrape/structuring"
)

func main() {
	lines, err := loading.LoadPageRaw("http://hustle-sa.ru/forum/index.php?showtopic=4909")
	if err != nil {
		panic(err)
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

	converter := structuring.NewConverter()

	var structuredData []structuring.Data
	for _, part := range dataParts {
		structuredData = append(structuredData, converter.Convert(part)...)
	}

	tables := structuring.GroupToTables(structuredData)

	if err := tables.Write("/tmp/hustle"); err != nil {
		panic(err)
	}
}
