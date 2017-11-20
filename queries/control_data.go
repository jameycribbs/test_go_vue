package queries

type ControlData struct {
	Total       int    `json:"total"`
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	NextPageUrl string `json:"next_page_url"`
	PrevPageUrl string `json:"prev_page_url"`
	From        int    `json:"from"`
	To          int    `json:"to"`
}
