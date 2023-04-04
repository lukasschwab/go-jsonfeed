// jsonfeed is a package for parsing and constructing JSON Feeds: https://jsonfeed.org/version/1.1. It explicitly supports JSON Feed Version 1.1.
package jsonfeed

import (
	"encoding/json"

	opt "github.com/lukasschwab/optional"
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
	return json.Marshal(f)
}

// A Feed is a JSON Feed.
type Feed struct {
	Version     string     `json:"version"`
	Title       string     `json:"title"`
	HomePageURL opt.String `json:"home_page_url,omitempty"`
	FeedURL     opt.String `json:"feed_url,omitempty"`
	Description opt.String `json:"description,omitempty"`
	UserComment opt.String `json:"user_comment,omitempty"`
	NextURL     opt.String `json:"next_url,omitempty"`
	Icon        opt.String `json:"icon,omitempty"`
	Favicon     opt.String `json:"favicon,omitempty"`
	Author      *Author    `json:"author,omitempty"` // Deprecated
	Authors     []Author   `json:"authors,omitempty"`
  Language opt.String `json:"language,omitempty"`
	Expired     opt.Bool   `json:"expired,omitempty"`
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
	Name   opt.String `json:"name,omitempty"`
	URL    opt.String `json:"url,omitempty"`
	Avatar opt.String `json:"avatar,omitempty"`
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
	URL           opt.String   `json:"url,omitempty"`
	ExternalURL   opt.String   `json:"external_url,omitempty"`
	Title         opt.String   `json:"title,omitempty"`
	ContentHTML   opt.String   `json:"content_html,omitempty"`
	ContentText   opt.String   `json:"content_text,omitempty"`
	Summary       opt.String   `json:"summary,omitempty"`
	Image         opt.String   `json:"image,omitempty"`
	BannerImage   opt.String   `json:"banner_image,omitempty"`
	DatePublished opt.String   `json:"date_published,omitempty"`
	DateModified  opt.String   `json:"date_modified,omitempty"`
	Author        *Author      `json:"author,omitempty"` // Deprecated
	Authors       []Author     `json:"authors,omitempty"`
	Tags          []string     `json:"tags,omitempty"`
  Language opt.String `json:"language,omitempty"`
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
	Title             opt.String `json:"title,omitempty"`
	SizeInBytes       opt.Int    `json:"size_in_bytes,omitempty"`
	DurationInSeconds opt.Int    `json:"duration_in_seconds,omitempty"`
}

// NewAttachment constructs a minimal Attachment.
func NewAttachment(url string, mimeType string) Attachment {
	return Attachment{URL: url, MIMEType: mimeType}
}
