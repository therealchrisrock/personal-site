// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

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
			Title:        "My Awesome Site",
			Description:  "Welcome to my awesome site! Explore our features.",
			CanonicalURL: "https://chrisrock.ca/",
		},
		OGMeta: OGMeta{
			OGTitle:       "Chris Rock",
			OGDescription: "Discover the amazing features of our website.",
			OGImage:       "https://example.com/static/img/default-og-image.jpg",
			OGType:        "website",
			OGURL:         "https://chrisrock.ca/",
		},
		TwitterMeta: TwitterMeta{
			TwitterCard:        "summary_large_image",
			TwitterTitle:       "My Awesome Site",
			TwitterDescription: "Discover the amazing features of our website.",
			TwitterImage:       "https://example.com/static/img/default-twitter-image.jpg",
		},
		SEOMeta: SEOMeta{
			Keywords: []string{"website", "awesome", "home"},
			Author:   "Site Admin",
			Robots:   "index, follow",
		},
		Assets: PageAssets{
			ThemeColor: "#ffffff",
			CustomCSS:  []string{"/static/css/main.css"},
			CustomJS:   []string{"/static/js/main.js"},
		},
	}
}

func Base(title string, head *DocumentHead) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link rel=\"stylesheet\" href=\"/static/css/output.css\"><script type=\"module\" src=\"/static/js/app.js\"></script><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/template/template.templ`, Line: 88, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</title></head><body class=\"font-mono\"><main class=\"p-6 grid gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
