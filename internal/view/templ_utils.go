package view

import "github.com/a-h/templ"

// Site metadata constants - single source of truth
const (
	SiteTitle = "The Real Chris Rock"
	SiteDescription = "Welcome to my website. Feel free to look around for as long as you like.."
	SiteURL = "https://chrisrock.ca/"
	SocialImageURL = "https://chrisrock.ca/static/media/social-card.png"
	SiteAuthor = "Chris Rock"
	SiteThemeColor = "#1e4877"
)

type DocumentHead struct {
	PageInfo    PageInfo    // Title, description, canonical URL
	OGMeta      OGMeta      // Open Graph metadata for social sharing
	TwitterMeta TwitterMeta // Twitter Card metadata
	SEOMeta     SEOMeta     // SEO metadata (keywords, robots)
	Assets      PageAssets  // Assets (favicon, theme color, CSS/JS links)
	Version     string      // Version string for cache busting
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
			Title:        SiteTitle,
			Description:  SiteDescription,
			CanonicalURL: SiteURL,
		},
		OGMeta: OGMeta{
			OGTitle:       SiteTitle,
			OGDescription: SiteDescription,
			OGImage:       SocialImageURL,
			OGType:        "website",
			OGURL:         SiteURL,
		},
		TwitterMeta: TwitterMeta{
			TwitterCard:        "summary_large_image",
			TwitterTitle:       SiteTitle,
			TwitterDescription: SiteDescription,
			TwitterImage:       SocialImageURL,
		},
		SEOMeta: SEOMeta{
			Keywords: []string{"creative developer", "web development", "interactive design", "three.js", "creative coding"},
			Author:   SiteAuthor,
			Robots:   "index, follow",
		},
		Assets: PageAssets{
			ThemeColor: SiteThemeColor,
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
