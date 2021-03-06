package zendesk

// Page is base struct for resource pagination
type Page struct {
	PreviousPage *string `json:"previous_page"`
	NextPage     *string `json:"next_page"`
	Count        int64   `json:"count"`
}

// HasPrev checks if the Page has previous page
func (p Page) HasPrev() bool {
	return (p.PreviousPage != nil)
}

// HasNext checks if the Page has next page
func (p Page) HasNext() bool {
	return (p.NextPage != nil)
}
