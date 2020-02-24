package jsonfeed

// TODO: validate timestamp formats.

// Validate that F represents a valid JSON feed.
func (f Feed) Validate() error {
	if len(f.Version) == 0 {
		return &MissingRequiredValueError{"Feed", "version"}
	}
	if len(f.Title) == 0 {
		return &MissingRequiredValueError{"Feed", "version"}
	}
	// Validate inner values with requirements.
	if f.Items == nil {
		return &MissingRequiredValueError{"Feed", "items"}
	}
	if err := validateItems(f.Items); err != nil {
		return err
	}
	if f.Hubs != nil {
		if err := validateHubs(f.Hubs); err != nil {
			return err
		}
	}
	return nil
}

// Validate that H represents a valid JSON Feed hub.
func (h Hub) Validate() error {
	if len(h.Type) == 0 {
		return &MissingRequiredValueError{"Hub", "type"}
	}
	if len(h.URL) == 0 {
		return &MissingRequiredValueError{"Hub", "url"}
	}
	return nil
}

// Validate that I represents a valid JSON Feed item.
func (i Item) Validate() error {
	if len(i.ID) == 0 {
		return &MissingRequiredValueError{"Item", "id"}
	}
	if i.Attachments != nil {
		if err := validateAttachments(i.Attachments); err != nil {
			return err
		}
	}
	return nil
}

// Validate that A represents a valid JSON Feed attachment.
func (a Attachment) Validate() error {
	if len(a.URL) == 0 {
		return &MissingRequiredValueError{"Attachment", "url"}
	}
	if len(a.MIMEType) == 0 {
		return &MissingRequiredValueError{"Attachment", "mime_type"}
	}
	return nil
}

// Utility functions for validating arrays within standard feeds.

func validateHubs(hubs []Hub) error {
	for i, hub := range hubs {
		if err := hub.Validate(); err != nil {
			return &IndexedMissingRequiredValueError{err, i}
		}
	}
	return nil
}

func validateItems(items []Item) error {
	for i, item := range items {
		if err := item.Validate(); err != nil {
			return &IndexedMissingRequiredValueError{err, i}
		}
	}
	return nil
}

func validateAttachments(attachments []Attachment) error {
	for i, attachment := range attachments {
		if err := attachment.Validate(); err != nil {
			return &IndexedMissingRequiredValueError{err, i}
		}
	}
	return nil
}
