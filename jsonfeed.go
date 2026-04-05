// jsonfeed is a package for parsing and constructing JSON Feeds: https://jsonfeed.org/version/1.1. It explicitly supports JSON Feed Version 1.1.
package jsonfeed

import (
	"bytes"
	"encoding/json"
)

// Version is the URL of the JSON Feed spec implemented here.
const Version = "https://jsonfeed.org/version/1.1"

// Parse constructs and validates a JSON feed from DATA.
func Parse(data []byte) (Feed, error) {
	var feed Feed
	err := json.Unmarshal(data, &feed)
	if err != nil {
		return feed, err
	}
	err = feed.Validate()
	return feed, err
}

// ToJSON converts a Feed to its JSON representation.
// TODO: enforce order
func (f Feed) ToJSON() ([]byte, error) {
	if err := f.Validate(); err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "\t")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// A Feed is a JSON Feed.
type Feed struct {
	Version     string     `json:"version"`
	Title       string     `json:"title"`
	HomePageURL *string  `json:"home_page_url,omitempty"`
	FeedURL     *string  `json:"feed_url,omitempty"`
	Description *string  `json:"description,omitempty"`
	UserComment *string  `json:"user_comment,omitempty"`
	NextURL     *string  `json:"next_url,omitempty"`
	Icon        *string  `json:"icon,omitempty"`
	Favicon     *string  `json:"favicon,omitempty"`
	Author      *Author  `json:"author,omitempty"` // Deprecated
	Authors     []Author `json:"authors,omitempty"`
	Language    *string  `json:"language,omitempty"`
	Expired     *bool    `json:"expired,omitempty"`
	Hubs        []Hub      `json:"hubs,omitempty"`
	Items       []Item     `json:"items"`
}

// NewFeed constructs a minimal Feed.
func NewFeed(title string, items []Item) Feed {
	return Feed{
		Version: Version,
		Title:   title,
		Items:   items,
	}
}

// An Author is a JSON Feed structure identifying an Author. Feeds and Items
// both have single authors.
type Author struct {
	Name   *string `json:"name,omitempty"`
	URL    *string `json:"url,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
}

// NewAuthor constructs a minimal Author. Because there are no required fields
// for a JSON Feed author, this returns an empty Author struct.
func NewAuthor() Author {
	return Author{}
}

// A Hub describes an endpoint that can be used to subscribe to real-time
// notifications from the publisher of this feed.
type Hub struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// NewHub constructs a minimal Hub.
func NewHub(t string, url string) Hub {
	return Hub{Type: t, URL: url}
}

// An Item is an item in a JSON Feed.
type Item struct {
	ID            string       `json:"id"`
	URL           *string      `json:"url,omitempty"`
	ExternalURL   *string      `json:"external_url,omitempty"`
	Title         *string      `json:"title,omitempty"`
	ContentHTML   *string      `json:"content_html,omitempty"`
	ContentText   *string      `json:"content_text,omitempty"`
	Summary       *string      `json:"summary,omitempty"`
	Image         *string      `json:"image,omitempty"`
	BannerImage   *string      `json:"banner_image,omitempty"`
	DatePublished *string      `json:"date_published,omitempty"`
	DateModified  *string      `json:"date_modified,omitempty"`
	Author        *Author      `json:"author,omitempty"` // Deprecated
	Authors       []Author     `json:"authors,omitempty"`
	Tags          []string     `json:"tags,omitempty"`
	Language      *string      `json:"language,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

// NewItem constructs a minimal Item.
func NewItem(id string) Item {
	return Item{ID: id}
}

// An Attachment is an attachment on an item in a JSON Feed.
type Attachment struct {
	URL               string     `json:"url"`
	MIMEType          string     `json:"mime_type"`
	Title             *string `json:"title,omitempty"`
	SizeInBytes       *int    `json:"size_in_bytes,omitempty"`
	DurationInSeconds *int    `json:"duration_in_seconds,omitempty"`
}

// NewAttachment constructs a minimal Attachment.
func NewAttachment(url string, mimeType string) Attachment {
	return Attachment{URL: url, MIMEType: mimeType}
}
