package jsonfeed

// TODO: validate timestamp formats.

// Validate that F represents a valid JSON Feed.
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

// Validate that H represents a valid JSON Feed Hub.
func (h Hub) Validate() *MissingRequiredValueError {
	if len(h.Type) == 0 {
		return &MissingRequiredValueError{"Hub", "type"}
	}
	if len(h.URL) == 0 {
		return &MissingRequiredValueError{"Hub", "url"}
	}
	return nil
}

// Validate that I represents a valid JSON Feed item.
func (i Item) Validate() *MissingRequiredValueError {
	if len(i.ID) == 0 {
		return &MissingRequiredValueError{"Item", "id"}
	}
	return nil
}

// Utility functions for validating arrays within standard feeds.

func validateHubs(hubs []Hub) *IndexedMissingRequiredValueError {
	for i, hub := range hubs {
		if err := hub.Validate(); err != nil {
			return &IndexedMissingRequiredValueError{*err, i}
		}
	}
	return nil
}

func validateItems(items []Item) *IndexedMissingRequiredValueError {
	for i, item := range items {
		if err := item.Validate(); err != nil {
			return &IndexedMissingRequiredValueError{*err, i}
		}
	}
	return nil
}
