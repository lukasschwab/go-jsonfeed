package jsonfeed

import (
	"encoding/json"
	"reflect"
	"testing"
)

// Feed parsing and validation.

func parseFromString(data string) (Feed, error) {
	return Parse(b(data))
}

func b(data string) []byte {
	return []byte(data)
}

func assertValidFeed(t *testing.T, feed string) {
	_, err := parseFromString(feed)
	if err != nil {
		t.Error("Failed to parse/validate valid JSON feed.", err)
	}
}

func assertInvalidFeed(t *testing.T, feed string, missing string) {
	_, err := parseFromString(feed)
	if err == nil {
		t.Errorf("Accepted feed without required %s.", missing)
	}
}

func TestFeedValidation(t *testing.T) {
	assertInvalidFeed(t, FeedWithoutVersion, "version")
	assertInvalidFeed(t, FeedWithoutTitle, "title")
	assertInvalidFeed(t, FeedWithoutItems, "items")
	// Hubs validation.
	assertValidFeed(t, FeedWithEmptyHubs)
	assertInvalidFeed(t, FeedWithInvalidHubs, "hub fields")
	// Feed validation.
	assertValidFeed(t, FeedWithEmptyItems)
	assertInvalidFeed(t, FeedWithInvalidItems, "item IDs")
	// Complex integration.
	assertValidFeed(t, AtlasFeed)
}

func TestHubValidation(t *testing.T) {
	var h0 Hub
	json.Unmarshal(b(`{}`), &h0)
	if err := h0.Validate(); err == nil {
		t.Error("Accepted empty hub.")
	}
	var h1 Hub
	json.Unmarshal(b(`{"type":"some"}`), &h1)
	if err := h1.Validate(); err == nil {
		t.Error("Accepted hub without URL.")
	}
	var h2 Hub
	json.Unmarshal(b(`{"url":"some"}`), &h2)
	if err := h2.Validate(); err == nil {
		t.Error("Accepted hub without Type.")
	}
	var validHub Hub
	json.Unmarshal(b(`{"url":"some","type":"some"}`), &validHub)
	if err := validHub.Validate(); err != nil {
		t.Error("Validation failed for valid Hub.", err)
	}
}

func TestItemValidation(t *testing.T) {
	var i0 Item
	json.Unmarshal(b(`{}`), &i0)
	if err := i0.Validate(); err == nil {
		t.Error("Accepted item without ID.")
	}
	var validItem Item
	json.Unmarshal(b(`{"id":"some"}`), &validItem)
	if err := validItem.Validate(); err != nil {
		t.Error("Validation failed for valid Item.", err)
	}
}

// Feed construction.

func TestFeedConstruction(t *testing.T) {
	f := NewFeed("My Test Feed", []Item{})
	if err := f.Validate(); err != nil {
		t.Error("Validation failed for minimal constructed Feed.", err)
	}
	// Make the feed invalid.
	f.Items = nil
	if err := f.Validate(); err == nil {
		t.Error("Validation succeeded for a feed with nil Items.")
	}
}

// Feed-to-JSON conversion.

func TestToJSON(t *testing.T) {
	f := NewFeed("My Test Feed", []Item{})
	j, err := f.ToJSON()
	if err != nil {
		t.Error("Failed to marshal valid feed to JSON.", err)
	}
	_, err = Parse(j)
	if err != nil {
		t.Error("Error parsing written JSON.", err)
	}
}

func TestUnchangedToJSON(t *testing.T) {
	atlas0, _ := parseFromString(AtlasFeed)
	j, _ := atlas0.ToJSON()
	atlas1, _ := Parse(j)
	if !reflect.DeepEqual(atlas0, atlas1) {
		t.Error("Original feed doesn't match parsed feed.")
	}
}
