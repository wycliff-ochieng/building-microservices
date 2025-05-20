package data

import (
	"time"
)

type Post struct {
	ID          int
	Author      string
	Title       string
	Body        string
	TimeCreated time.Time
}

var PostList = []*Post{
	&Post{
		ID:          1,
		Author:      "Willis Raburu",
		Title:       "Europa Glory",
		Body:        "MAn united sweeep the Europa championship by beating Tottenham by 10-2 in pealtiees.With this great news Man u is bigger than Arsenal",
		TimeCreated: time.Now().UTC(),
	},
	&Post{
		ID:          2,
		Author:      "Jane Smith",
		Title:       "Go Programming Best Practices",
		Body:        "Understanding Go concurrency patterns and error handling is crucial for building robust applications...",
		TimeCreated: time.Now().Add(-24 * time.Hour).UTC(),
	},
	&Post{
		ID:          3,
		Author:      "John Doe",
		Title:       "Tech Trends 2025",
		Body:        "AI and machine learning continue to reshape how we approach software development...",
		TimeCreated: time.Now().Add(-48 * time.Hour).UTC(),
	},
	&Post{
		ID:          4,
		Author:      "Alice Johnson",
		Title:       "Microservices Architecture",
		Body:        "Building scalable systems requires careful consideration of service boundaries and communication patterns...",
		TimeCreated: time.Now().Add(-72 * time.Hour).UTC(),
	},
}

func GetPost() []*Post {
	return PostList
}

func AddPost(p *Post) {
	p.ID = getNextID()
	PostList = append(PostList, p)
}

func getNextID() int {
	PL := PostList[len(PostList)-1]
	return PL.ID + 1
}
