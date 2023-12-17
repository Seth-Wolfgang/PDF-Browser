package app

import (
	. "pdfbrowser/pkg/paper"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	papers         map[int]Paper
	numberOfPapers int
}

func MakeMenu() Menu {
	var result Menu 
	var papers []Paper
	
	papers = ReadFiles()

	for _, paper := range papers {
		result.numberOfPapers++
		result.papers[paper.Id] = paper		
	}

	return result
}

func (m Menu) MainMenu() {
	var option int
	var id int

	fmt.Println("1. Add Paper")
	fmt.Println("2. Delete Paper")
	fmt.Println("3. Update Paper")
	fmt.Println("4. List Papers")
	fmt.Println("5. Search Paper")
	fmt.Println("6. Exit")

	fmt.Scanln(&option)

	switch option {
	case 1:
		m.AddPaper()
	case 2:
		fmt.Println("Type ID of Paper:")
		fmt.Scanln(&id)
		m.DeletePaper(id)
	case 3:
		var field string
		var value string
		fmt.Println("Type ID of Paper:")
		fmt.Scanln(&id)

		fmt.Println("What field?: ")
		fmt.Scanln(&field)

		fmt.Println("Type new value. If tags or authors use commas to seperate them: ")
		fmt.Println("Current: ", m.papers[id].ReturnField(field))
		fmt.Scanln(&value)

		m.papers[id].UpdatePaper(id, field, value)
	case 4:
		m.ListPapers()
	case 5:
		fmt.Println("Type ID of Paper:")
		fmt.Scanln(&id)
		m.SearchPaper()
	case 6:
		m.ExitPaperBrowser()
	}
}

func (m Menu) AddPaper() {
	var p Paper
	var tags string

	p.Id = m.numberOfPapers
	m.numberOfPapers++

	fmt.Println("Enter the title of the paper: ")
	fmt.Scanln(&p.Title)

	fmt.Println("Enter the author of the paper: ")
	fmt.Scanln(&p.Author)

	fmt.Println("Enter the year of the paper: ")
	fmt.Scanln(&p.Year)

	fmt.Println("Enter the venue of the paper: ")
	fmt.Scanln(&p.Venue)

	fmt.Println("Enter the description of the paper: ")
	fmt.Scanln(&p.Description)

	fmt.Println("Enter the tags of the paper seperated by ',': ")
	fmt.Scanln(&tags)

	for _, tag := range strings.Split(tags, ",") {
		p.Tags = append(p.Tags, tag)
	}

	p.WriteFile()
	m.papers[p.Id] = p
}

func (Menu) DeletePaper(id int) {

}

func (m Menu) SearchPaper() {
	var field string
	var value string
	var findings []Paper
	var found bool = false

	fmt.Println("Seach by what field?:")
	fmt.Println("ID")
	fmt.Println("Title")
	fmt.Println("Author")
	fmt.Println("Year")
	fmt.Println("Venue")
	fmt.Println("Description")
	fmt.Println("Tags")
	fmt.Scanln(&field)

	fmt.Println("What value?: ")
	fmt.Scanln(&value)

	// temporary search algorithm
	field = strings.ToLower(field)
	for _, paper := range m.papers {
		found = false
		for _, attr := range paper.ReturnField(field) {
			if strings.Contains(attr, value) {
				found = true
				break
			}
			if found {
				findings = append(findings, paper)
			}
		}
	}

}

func (m Menu) ListPapers() {
	for _, p := range m.papers {
		fmt.Printf("%3d | %s\n", p.Id, p.Title)
	}
}

func (Menu) ExitPaperBrowser() {
	os.Exit(0)
}
