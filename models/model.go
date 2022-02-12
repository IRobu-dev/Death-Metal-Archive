package models

type Album struct {
	Id    int    `json:"id"`
	Band  int    `json:"band"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type Band struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Status   string `json:"status"`
	FormedIn int    `json:"formed-in"`
	Genre    string `json:"genre"`
	Theme    string `json:"theme"`
	Active   string `json:"active"`
}

type Reviews struct {
	Id      int     `json:"id"`
	Album   string  `json:"album"`
	Title   string  `json:"title"`
	Score   float32 `json:"score"`
	Content string  `json:"content"`
}

type Pagination struct {
	First int    `json:"first"`
	Last  int    `json:"second"`
	CSV   string `json:"csv"`
}

type SearchResult struct {
	Bands  []Band
	Albums []Album
}
