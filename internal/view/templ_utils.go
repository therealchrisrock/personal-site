package view

import "github.com/a-h/templ"

type DocumentHead struct {
	PageInfo    PageInfo    // Title, description, canonical URL
	OGMeta      OGMeta      // Open Graph metadata for social sharing
	TwitterMeta TwitterMeta // Twitter Card metadata
	SEOMeta     SEOMeta     // SEO metadata (keywords, robots)
	Assets      PageAssets  // Assets (favicon, theme color, CSS/JS links)
}

// PageInfo holds general information for the page
type PageInfo struct {
	Title        string // Page title
	Description  string // Meta description
	CanonicalURL string // Canonical URL
}

// OGMeta holds Open Graph metadata
type OGMeta struct {
	OGTitle       string // Open Graph title
	OGDescription string // Open Graph description
	OGImage       string // Open Graph image URL
	OGType        string // Open Graph type (e.g., "website" or "article")
	OGURL         string // Open Graph URL
}

// TwitterMeta holds Twitter Card metadata
type TwitterMeta struct {
	TwitterCard        string // Type of Twitter card (e.g., "summary_large_image")
	TwitterTitle       string // Twitter title
	TwitterDescription string // Twitter description
	TwitterImage       string // Twitter image URL
}

// SEOMeta holds SEO-related metadata
type SEOMeta struct {
	Keywords []string // Keywords for the page (for SEO)
	Author   string   // Author name
	Robots   string   // Robots meta tag (e.g., "index, follow")
}

// PageAssets holds assets such as favicon, CSS, and JS links
type PageAssets struct {
	ThemeColor string   // Theme color for mobile browsers
	CustomCSS  []string // Array of custom CSS URLs
	CustomJS   []string // Array of custom JS URLs
}

func DefaultHead() DocumentHead {
	return DocumentHead{
		PageInfo: PageInfo{
			Title:        "Chris Rock's Portfolio",
			Description:  "Hello, I'm a software developer specializing in building modern web applications, APIs, and data-driven solutions. Available for freelance and collaboration opportunities.",
			CanonicalURL: "https://chrisrock.ca/",
		},
		OGMeta: OGMeta{
			OGTitle:       "Chris Rock's Portfolio",
			OGDescription: "Hello, I'm a software developer specializing in building modern web applications, APIs, and data-driven solutions. Available for freelance and collaboration opportunities.",
			OGImage:       "https://example.com/static/img/default-og-image.jpg",
			OGType:        "website",
			OGURL:         "https://chrisrock.ca/",
		},
		TwitterMeta: TwitterMeta{
			TwitterCard:        "summary_large_image",
			TwitterTitle:       "Chris Rock's Portfolio",
			TwitterDescription: "Hello, I'm a software developer specializing in building modern web applications, APIs, and data-driven solutions. Available for freelance and collaboration opportunities.",
			TwitterImage:       "https://example.com/static/img/default-twitter-image.jpg",
		},
		SEOMeta: SEOMeta{
			Keywords: []string{"website", "awesome", "home"},
			Author:   "Chris Rock",
			Robots:   "index, follow",
		},
		Assets: PageAssets{
			ThemeColor: "#ffffff",
			CustomCSS:  []string{"/static/css/main.css"},
			CustomJS:   []string{"/static/js/main.js"},
		},
	}
}

// GetHeadProp returns the value if not empty or the fallback if the value is empty.
func GetHeadProp[T comparable](value *T, fallback T) T {
	if value == nil {
		return fallback
	}
	return *value
}

// Components
type CardThumbnail struct {
	Img  templ.Attributes
	Link templ.Attributes
}
