package termcolor

//CustomTagManager is a struct to manage custom tags and their styles
//It allows adding custom tags and their styles, and retrieving styles for a given tag.
// Example usage:
// customTags := NewCustomTagManager()
// customTags.Add("error", "br bold bgk") // bright red bold on black
// customTags.Add("warning", "y bold")    // yellow bold
// customTags.Add("success", "bg bold")   // bright green bold
// customTags.Add("info", "bc u")         // bright cyan underlined

type CustomTagManager struct {
	tags map[string]string
}

func NewCustomTagManager() *CustomTagManager {
	return &CustomTagManager{tags: make(map[string]string)}
}

func (ctm *CustomTagManager) Add(tag, style string) {
	ctm.tags[tag] = style
}

func (ctm *CustomTagManager) Get(tag string) (string, bool) {
	style, exists := ctm.tags[tag]
	return style, exists
}
