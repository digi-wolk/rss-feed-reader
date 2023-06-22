package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
}

func ReadRSSFeed(url string, hoursBack int) ([]Item, error) {
	var result []Item

	// Make an HTTP GET request to fetch the RSS feed
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching RSS feed: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading RSS feed body: %s", err.Error())
		return nil, err
	}

	// Parse the XML body into the RSS struct
	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Printf("Error parsing RSS feed XML: %s", err.Error())
		return nil, err
	}

	// Get the current date and time
	now := time.Now()

	// Print the titles of the items published within the last month
	for _, item := range rss.Channel.Items {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			fmt.Printf("Error parsing pubDate: %s", err.Error())
			continue
		}

		// Calculate the difference between the current date and the publication date
		diff := now.Sub(pubDate)

		// Check if the item was published within the last x hours
		if diff < (time.Duration(hoursBack) * time.Hour) {
			// Append to result
			result = append(result, item)
		}
	}
	return result, nil
}

func FilterFeedItems(items []Item, filterWords []string, excludeWords []string, caseInsensitive bool) []Item {
	// Filter title and description based on filterWords
	// and exclude based on excludeWords
	var result []Item
	for _, item := range items {

		if len(filterWords) == 0 {
			if !containsAny(item.Title, excludeWords, caseInsensitive) &&
				!containsAny(item.Description, excludeWords, caseInsensitive) {
				result = append(result, item)
			}
			continue
		}

		if containsAny(item.Title, filterWords, caseInsensitive) ||
			containsAny(item.Description, filterWords, caseInsensitive) {
			if !containsAny(item.Title, excludeWords, caseInsensitive) &&
				!containsAny(item.Description, excludeWords, caseInsensitive) {
				result = append(result, item)
			}
		}
	}
	return result
}

// Helper function to check if a string contains any of the words in a slice
func containsAny(s string, words []string, caseInsensitive bool) bool {
	for _, word := range words {
		if caseInsensitive {
			if strings.Contains(strings.ToLower(s), strings.ToLower(word)) {
				return true
			}
		} else {
			if strings.Contains(s, word) {
				return true
			}
		}
	}
	return false
}
