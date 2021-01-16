package competitions

import (
	"fmt"
	"regexp"
	"time"

	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
)

var (
	eventURL   = regexp.MustCompile(`http://hustle-sa.ru/forum/index.php?.*showtopic=\d+`)
	eventTitle = regexp.MustCompile(`\((?P<FirstDate>\d{4}-\d{2}-\d{2}).*\)\s(?P<Name>.+)`)
)

const (
	FirstNewFormatEvent = 550 // min start is 550 - first competition with new format
	EventPageSize       = 15
)

func GetCompetitions(start int) ([]Competition, error) {
	resp, err := soup.Get(fmt.Sprintf("http://hustle-sa.ru/forum/index.php?showforum=6&prune_day=100&sort_by=A-Z&sort_key=title&st=%d", start))
	if err != nil {
		return nil, err
	}

	doc := soup.HTMLParse(resp)
	return getEventDataList(doc)
}

type Competition struct {
	Name      string
	StartDate time.Time
	URL       string
}

func getEventDataList(node soup.Root) ([]Competition, error) {
	hrefs := node.FindAll("a")

	result := make([]Competition, 0, 15)
	for _, root := range hrefs {
		root := root

		url := root.Attrs()["href"]
		if !eventURL.MatchString(url) {
			continue
		}

		textNode := root.Pointer.LastChild
		if textNode.Type != html.TextNode {
			continue
		}

		text := textNode.Data

		submatches := eventTitle.FindStringSubmatch(text)
		if len(submatches) == 0 {
			continue
		}

		date, err := time.Parse("2006-01-02", submatches[1])
		if err != nil {
			return nil, err
		}

		result = append(result, Competition{
			Name:      submatches[2],
			StartDate: date,
			URL:       url,
		})
	}

	return result, nil
}
