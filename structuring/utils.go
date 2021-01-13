package structuring

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

func GroupToTables(dataBlocks ...[]Data) Tables {
	var result Tables

	for _, blocks := range dataBlocks {
		for _, block := range blocks {
			switch casted := block.(type) {
			case Judge:
				result.Judges = append(result.Judges, casted)

			case Category:
				result.Categories = append(result.Categories, casted)

			case ParticipantResult:
				result.Results = append(result.Results, casted)

			case Participant:
				result.Participants = append(result.Participants, casted)

			case Cross:
				result.Crosses = append(result.Crosses, casted)

			default:
				panic(fmt.Sprintf("unexpected type %T", casted))
			}
		}
	}

	return result
}

type Tables struct {
	Competitions []Competition
	Categories   []Category
	Participants []Participant
	Results      []ParticipantResult
	Crosses      []Cross
	Judges       []Judge
}

func (tables *Tables) Write(dir string) error {
	writeRecords := func(suffix string, data interface{}) error {
		var records [][]string

		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			item := s.Index(i).Interface().(Data)
			records = append(records, item.ToStrings())
		}

		f, err := os.Create(dir + "/" + suffix + ".csv")
		if err != nil {
			return err
		}

		defer f.Close()

		writer := csv.NewWriter(f)
		return writer.WriteAll(records)
	}

	if err := writeRecords("competitions", tables.Competitions); err != nil {
		return err
	}

	if err := writeRecords("categories", tables.Categories); err != nil {
		return err
	}

	if err := writeRecords("participants", tables.Participants); err != nil {
		return err
	}

	if err := writeRecords("results", tables.Results); err != nil {
		return err
	}

	if err := writeRecords("crosses", tables.Crosses); err != nil {
		return err
	}

	if err := writeRecords("judges", tables.Judges); err != nil {
		return err
	}

	return nil
}
