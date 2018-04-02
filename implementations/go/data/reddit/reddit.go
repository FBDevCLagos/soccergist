package reddit

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/turnage/graw/reddit"
)

const (
	subreddit = "/r/footballhighlights"
)

var (
	redditBot     reddit.Bot
	videoURLRegex = regexp.MustCompile(`href="([\w\/:?\.-=]+|[\w\/:\.-]+)"`)
)

func init() {
	var err error
	redditBot, err = reddit.NewBotFromAgentFile("reddit.agent", 0)
	if err != nil {
		log.Fatal("Failed to create bot handle: ", err)
	}
}

type Highlight struct {
	*reddit.Post
	URLs []string
}

// GetHighlightPosts returns posts in the highlights subreddit
// Filtering only posts that have 'Premier' league in their title
func GetHighlightPosts(after string) (highlights []*Highlight) {
	params := map[string]string{"after": after}
	harvest, err := redditBot.ListingWithParams(subreddit, params)
	if err != nil {
		log.Printf("Failed to fetch: %s %s\n", subreddit, err)
		return
	}

	for _, post := range harvest.Posts {
		if !strings.Contains(post.Title, "Premier") {
			continue
		}

		highlight := &Highlight{Post: post, URLs: getHighlightVideoURLs(post)}
		if len(highlight.URLs) == 0 {
			fmt.Println(highlight.Title)
			fmt.Println(post.SelfTextHTML)
		}
		highlights = append(highlights, highlight)
	}

	return
}

func getHighlightVideoURLs(post *reddit.Post) (urls []string) {
	videoURLs := videoURLRegex.FindAllStringSubmatch(post.SelfTextHTML, -1)
	for _, url := range videoURLs {
		// TODO: handle unique videos
		urls = append(urls, url[1])
	}
	return
}
