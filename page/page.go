package page

import "io/ioutil"

const dataPath = "./data/"

// Page represents a page, it contains a title and a body
type Page struct {
	Title string
	Body  []byte
}

// Save saves a Page struct to a .txt file under the data/ folder
// the file name corresponds to the Title of the Page
func (p *Page) Save() error {
	filename := dataPath + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0666)
}

// Load loads a Page from the data/ directory
func Load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(dataPath + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
