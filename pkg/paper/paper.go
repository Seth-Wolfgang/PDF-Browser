package paper

import (
	"encoding/json"
	"os"
	"strings"
)

const dataPath = "pdfbrowser/data/"

type Paper struct {
	Id          int
	Title       string
	Author      []string
	Year        int
	Venue       string
	Description string
	Tags        []string
}

func ReadFiles() []Paper {
	// temporary
	var papers []Paper
	var err error

	files, err := os.ReadDir(dataPath)
	check(err)

	for i := 0; i < len(files); i++ {
		papers = append(papers, SearchPaper(i))
	}

	return papers
}

func (p Paper) WriteFile() {
	b, err := json.Marshal(p)
	check(err)

	file, err := os.OpenFile(dataPath+string(p.Id)+".json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	check(err)
	defer file.Close()

	file.Write(b)
	file.Sync()
}

func (p *Paper) AddPaper() {

}

func (p Paper) DeletePaper() {

}

func (p Paper) UpdatePaper(id int, field string, newVar interface{}) {
	field = strings.ToLower(field)

	switch field {
	case "id":
		p.Id = converter[int](newVar)
	case "year":
		p.Year = converter[int](newVar)

	case "title":
		p.Title = converter[string](newVar)
	case "venue":
		p.Venue = converter[string](newVar)
	case "description":
		p.Description = converter[string](newVar)
	case "tags":
		p.Tags = converter[[]string](newVar)
	case "author":
		p.Author = converter[[]string](newVar)
	default:
		panic(string("Invalid Field in update of paper " + string(id) + " on field " + field + "!\n"))
	}

	p.WriteFile()

}

func SearchPaper(id int) Paper {
	var p Paper
	files, err := os.ReadDir(dataPath)
	check(err)

	for _, file := range files {
		b, err := os.ReadFile(file.Name())
		check(err)

		err = json.Unmarshal(b, &p)
		check(err)

		if p.Id == id {
			return p
		}
	}

	panic("paper not found!") // temporary
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func converter[T any](val any) T {
	if newVar, ok := val.(T); ok {
		return newVar
	} else {
		panic(ok)
	}
}

func (p Paper) ReturnField(field string) []string {
	field = strings.ToLower(field)
	var temp []string

	switch field {
	case "id":
		temp = append(temp, string(p.Id))
		return temp
	case "title":
		temp = append(temp, p.Title)
		return temp
	case "author":
		return p.Author
	case "year":
		temp = append(temp, string(p.Year))
		return temp
	case "venue":
		temp = append(temp, p.Venue)
		return temp
	case "description":
		temp = append(temp, p.Description)
		return temp
	case "tag":
		return p.Tags
	default:
		panic("Invalid field!")
	}
}
