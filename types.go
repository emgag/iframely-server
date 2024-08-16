package main

type Response struct {
	Provider string `json:"provider"`
	URL      string `json:"url"`
	HTML     string `json:"html"`
	Image    string `json:"image"`
	Error    string `json:"error,omitempty"`
}

type IframelyResponse struct {
	URL             string `json:"url"`
	Type            string `json:"type"`
	Version         string `json:"version"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ProviderName    string `json:"provider_name"`
	Description     string `json:"description"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	HTML            string `json:"html"`
	CacheAge        int    `json:"cache_age"`
	Status          int    `json:"status"`
	Error           string `json:"error"`
}
