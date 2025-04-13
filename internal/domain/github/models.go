package github

type Repository struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	HTMLURL     string `json:"html_url"`
}

type SearchResult struct {
	Items []Repository `json:"items"`
}
