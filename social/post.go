package social

func PostToAll(identifier string, warning_text string) {
	PostToBluesky(identifier, warning_text)
	// PostToFediverse(identifier, warning_text)
}