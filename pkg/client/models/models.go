package models

const (
	Mode                 = "articles"
	Size                 = 1
	CyberleninkaEndpoint = "/api/search"
	ElibraryEndpoint     = ""
)

type CyberleninkaRequest struct {
	Mode string `json:"mode"`
	Size int    `json:"size"`
	Q    string `json:"q"`
}

type CyberleninkaArticle struct {
	Name       string   `json:"name"`
	Annotation string   `json:"annotation"`
	Link       string   `json:"link"`
	Authors    []string `json:"authors"`
	Year       int      `json:"year"`
}

type CyberleninkaResponse struct {
	Found    int                   `json:"found"`
	Articles []CyberleninkaArticle `json:"articles"`
}

type LiteratureList struct {
	Title       string
	IsContained bool
}
