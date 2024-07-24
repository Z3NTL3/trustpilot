package model

import (
	"io"

	"github.com/antchfx/htmlquery"
)

type Result struct {
	Company string // The company name
	Reviews struct {
		Total string // Total amount reviews
		Text  string // Description text
	}
	ImageSRC string // The SVG of the star rate graphical vector
}

var LookupTable = [...]string{
	"//*[@id=\"business-unit-title\"]/h1/span[1]", // Company header
	"//*[@id=\"business-unit-title\"]/span[1]/span", // Reviews Text header
	"//*[@id=\"business-unit-title\"]/div/div/p", // Reviews Total header
	"//*[@id=\"business-unit-title\"]/div/div/div/img", // ImageSRC header
}

// Parse parses the HTML content from the given io.ReadCloser and returns a Result struct and an error.
//
// It takes a parameter `body` of type `io.ReadCloser` which represents the HTML content.
//
// The function first parses the HTML content using the `htmlquery.Parse` function.
// If there is an error during the parsing, it returns the error.
//
// Then, it initializes an empty `Result` struct.
//
// Next, it iterates over the `LookupTable` and performs different operations based on the index.
//
// - If the index is 0, it sets the `Company` field of the `Result` struct to the inner text of the parsed node.
// - If the index is 1, it sets the `Text` field of the `Reviews` struct to the inner text of the parsed node.
// - If the index is 2, it sets the `Total` field of the `Reviews` struct to the inner text of the parsed node.
// - If the index is 3, it sets the `ImageSRC` field of the `Result` struct to the value of the "src" attribute of the parsed node.
//
// Finally, it returns a pointer to the `Result` struct and a `nil` error.
func Parse(body io.ReadCloser) (*Result, error) {
	doc, err := htmlquery.Parse(body)
	if err != nil {
		return nil, err
	}
	
	var result Result

	for index, lookup := range LookupTable {
		node, err := htmlquery.Query(doc, lookup)
		if err != nil {
			return nil, err
		}

		switch(index){
		case 0: // Company header
			result.Company = htmlquery.InnerText(node)
		case 1: // Reviews Text header
			result.Reviews.Text = htmlquery.InnerText(node)
		case 2: // Reviews Total header
			result.Reviews.Total = htmlquery.InnerText(node)
		case 3: // ImageSRC header
			result.ImageSRC = htmlquery.SelectAttr(node, "src")
		}
	}

	return &result, nil
}

// GetCompanyName returns the name of the company.
//
// This function takes no parameters and returns a string representing the name of the company.
func (r *Result) GetCompanyName() string {
	return r.Company
}

// GetTotalReviews returns the total number of reviews.
//
// It does not take any parameters.
// It returns a string.
func (r *Result) GetTotalReviews() string {
	return r.Reviews.Total
}

// GetReviewsText returns the text of the reviews.
//
// It does not take any parameters.
// It returns a string.
func (r *Result) GetReviewsText() string {
	return r.Reviews.Text
}
