package termcolor

import (
	"golang.org/x/net/html"
	"strings"
)

func EncodeHTMLToTerm(customTags *CustomTagManager, src string) string {
	var result strings.Builder
	var styleStack []string

	doc, err := parseHTML(src)
	if err != nil {
		return src
	}

	// Process the document body
	if body := findBody(doc); body != nil {
		processNode(customTags, body, &result, &styleStack)
	}

	return result.String()
}

// Parse the HTML and return the root node
func parseHTML(src string) (*html.Node, error) {
	wrappedSrc := "<div>" + src + "</div>"
	doc, err := html.Parse(strings.NewReader(wrappedSrc))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// Helper function to find the body node
func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "div" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findBody(c); result != nil {
			return result
		}
	}
	return nil
}

func processNode(customTags *CustomTagManager, n *html.Node, result *strings.Builder, styleStack *[]string) {
	if n.Type == html.TextNode {
		result.WriteString(n.Data)
	} else if n.Type == html.ElementNode {
		var styles strings.Builder

		// Check if the node is a custom tag
		if classes, ok := customTags.Get(n.Data); ok {
			n.Data = "span"
			n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: classes})
		}

		// Process style and class attributes
		for _, attr := range n.Attr {
			if attr.Key == "style" {
				styles.WriteString(parseStyleAttribute(attr.Val))
			} else if attr.Key == "class" {
				styles.WriteString(parseClassAttribute(attr.Val))
			}
		}

		if styles.Len() > 0 {
			*styleStack = append(*styleStack, styles.String())
			result.WriteString(styles.String())
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(customTags, c, result, styleStack)
		}

		if styles.Len() > 0 {
			*styleStack = (*styleStack)[:len(*styleStack)-1]
			if len(*styleStack) > 0 {
				result.WriteString((*styleStack)[len(*styleStack)-1])
			} else {
				result.WriteString(StyleReset)
			}
		}
	}
}

func mapStyleToCode(style string) string {
	style = strings.TrimSpace(style)
	if strings.HasPrefix(style, "color:") {
		color := strings.TrimSpace(strings.TrimPrefix(style, "color:"))
		if code, ok := ColorMap[color]; ok {
			return code
		}
	} else if strings.HasPrefix(style, "background-color:") {
		bgColor := strings.TrimSpace(strings.TrimPrefix(style, "background-color:"))
		if code, ok := BColorMap["bg"+bgColor]; ok {
			return code
		}
	} else if strings.HasPrefix(style, "font-weight: bold") {
		return StyleMap["bold"]
	} else if strings.HasPrefix(style, "text-decoration: underline") {
		return StyleMap["underline"]
	} else if strings.HasPrefix(style, "filter: invert") {
		return StyleMap["reverse"]
	} else if strings.HasPrefix(style, "opacity: 0.8") {
		return StyleMap["dim"]
	}
	return ""
}

func parseStyles(styles []string) string {
	var result strings.Builder
	for _, style := range styles {
		result.WriteString(mapStyleToCode(style))
	}
	return result.String()
}

func parseStyleAttribute(style string) string {
	styleAttrs := strings.Split(style, ";")
	return parseStyles(styleAttrs)
}

func parseClassAttribute(class string) string {
	var styles []string
	classes := strings.Fields(class)
	for _, class := range classes {
		if styleStr, ok := CSSClassMap[class]; ok {
			styles = append(styles, strings.Split(styleStr, ";")...)
		}
	}
	return parseStyles(styles)
}
