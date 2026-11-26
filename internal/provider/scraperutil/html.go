package scraperutil

import (
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/html"

	"github.com/julien-noblet/download-geofabrik/pkg/catalog"
)

// ParseAnchorHrefs streams HTML from reader and invokes onHref for every <a> tag's href attribute.
func ParseAnchorHrefs(reader io.Reader, onHref func(href string)) error {
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return nil
			}

			return fmt.Errorf("cannot parse html: %w", tokenizer.Err())

		case html.StartTagToken, html.SelfClosingTagToken:
			processAnchorTag(tokenizer, onHref)

		case html.TextToken, html.EndTagToken, html.CommentToken, html.DoctypeToken:
			// Non-anchor tokens
		}
	}
}

func processAnchorTag(tokenizer *html.Tokenizer, onHref func(href string)) {
	tagName, hasAttr := tokenizer.TagName()
	if string(tagName) != "a" || !hasAttr {
		return
	}

	for {
		key, val, more := tokenizer.TagAttr()
		if string(key) == "href" {
			onHref(string(val))

			return
		}

		if !more {
			return
		}
	}
}

// HrefParser extracts an element ID and format ID from an anchor href.
type HrefParser func(href string) (elementID, formatID string)

// ParseFlatDirectory parses an HTML directory listing and populates cat using shouldSkip and parseHref.
func ParseFlatDirectory(reader io.Reader, cat *catalog.Catalog, shouldSkip func(href string) bool, parseHref HrefParser) error {
	return ParseAnchorHrefs(reader, func(href string) {
		if shouldSkip(href) {
			return
		}

		elementID, formatID := parseHref(href)
		if elementID == "" || formatID == "" {
			return
		}

		elem := catalog.Element{
			ID:   elementID,
			Name: elementID,
		}

		_ = cat.MergeElement(&elem)
		cat.AddExtension(elementID, formatID)
	})
}
