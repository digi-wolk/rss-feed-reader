package rss

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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

func ReadRSSFeed(url string) ([]Item, error) {
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

		// Check if the item was published within the last 2x24 hours
		if diff < (2 * 24 * time.Hour) {
			// Append to result
			result = append(result, item)
		}
	}
	return result, nil
}
