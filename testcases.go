package termcolor

import "fmt"

var TestCases = []string{
	"[src1] This is normal text <span style=\"color: yellow\">Hello World</span> and normal text again",
	"[src2] This text is <span style=\"color: red; font-weight: bold\">bold and red</span> but this is not",
	"[src3] Here we have <span style=\"color: blue\">blue text with some <span style=\"color: green\">green</span> inside</span>",
	"[src4] This is <span style=\"color: blue; font-weight: bold\">This is <u>blue and bold</u> with <span style=\"color: green\">green nested text</span></span> and normal text",
	"[src5] Start <span style=\"color: red\">red <span style=\"color: yellow\">yellow <span style=\"color: green\">green</span> yellow</span> red</span> end",
	"[src6] Testing <span style=\"color: blue; font-weight: bold\">blue and bold</span> and <span style=\"color: magenta; text-decoration: underline\">magenta with underline</span>",
	"[src7] This <span style=\"color: red\">should work</span but this won't",
	"[src8] This <span style=\"color: red\">should work</span but <u>this</u> won't - can you see the `but`?",
	"[src9] This has <span style=\"color: white; background-color: red\">white text on red</span> and <span style=\"color: yellow; background-color: blue\">yellow on blue</span>",
	"[src10] <span style=\"font-weight: bold; opacity: 0.8; text-decoration: underline; filter: invert(100%)\">All styles at once</span>",
	"[src11] <span style=\"background-color: blue\">Blue bg <span style=\"color: yellow\">yellow text <span style=\"background-color: red\">on red</span> back to blue</span> white on blue</span>",
	"[src12] <span style=\"color: red\">Test</span <span style=\"color: green\">another</span <span style=\"color: blue\">third</span> normal text",
	"[src13] <span style=\"color: brightred\">Bright red <span style=\"color: brightblue\">and blue <span style=\"color: brightyellow\">and yellow</span></span></span>",
	"[src14] <span style=\"color: red; font-weight: bold\">Red bold <span style=\"color: green; text-decoration: underline\">green test</span> end</span>",
	"[src15] <span></span><span style=\"color: red\">Text</span></span> and <span>middle</span> and <span>end</span>",
	"[src16] <span style=\"color: blue\"><span style=\"color: green\"><span style=\"color: yellow\">Deep</span><span style=\"color: red\">middle</span></span></span>",
	"[src17] Normal text with <span style=\"filter: invert(100%)\">reversed video</span> and back to normal",
	"[src18] <span style=\"background-color: red\">Red bg <span style=\"background-color: blue\">Blue bg <span style=\"background-color: green\">Green bg</span> back to blue</span> back to red</span>",
	"[src19] Normal text with <span style=\"opacity: 0.8\">dimmed text</span> and normal again",
	"[src20] <span style=\"color: blue; font-weight: bold\">Bold blue <span style=\"text-decoration: underline\">with underline <span style=\"color: red\">and red</span></span></span>",
	"[src21] <span style=\"background-color: brightred; color: black\">Bright red bg <span style=\"background-color: brightblue\">bright blue bg</span></span>",
	"[src22] <span style=\"font-weight: bold; color: yellow\">Bold yellow <span style=\"text-decoration: underline; color: cyan\">underlined cyan <span style=\"filter: invert(100%); color: green\">reversed green</span></span></span>",
	"[src23] <span style=\"opacity: 0.8\">Dimmed text with <span style=\"filter: invert(100%)\">reversed video</span> and back to dim</span>",
	"[src24] <span style=\"background-color: brightmagenta; color: white\">White on bright magenta <span style=\"background-color: brightcyan; color: black\">Black on bright cyan</span></span>",
	"[src25] <span style=\"color: red; font-weight: bold\">Red bold <span style=malformed>should stay red bold</span> still red bold</span>",
	"[src26] <span style=\"\">Normal <span style=\"color: green\">Green</span> <span style=\"   \">Still normal</span></span>",
	"[src27] <span style=\"color: blue; font-weight: bold\">Bold blue <span style=\"text-decoration: underline\">and underlined <span style=\"filter: invert(100%)\">and reversed <span style=\"opacity: 0.8\">and dimmed</span></span></span></span>",
	"[src28] Normal text with <span class=\"r\">red text</span> and <span class=\"b\">blue text</span>",
	"[src29] <span class=\"r bold\">Bold red text</span> and <span class=\"b u\">underlined blue</span>",
	"[src30] <span class=\"b\">Blue text with <span class=\"r bold\">bold red</span> inside</span>",
	"[src31] <span class=\"g bold\" style=\"text-decoration: underline\">Green bold underlined</span>",
	"[src32] <span class=\"r\">Red <span class=\"bold\">bold red <span class=\"y\">yellow bold</span></span></span>",
	"[src33] <span class=\"br\">Bright red</span> and <span class=\"by\">bright yellow</span> and <span class=\"bc\">bright cyan</span>",
	"[src34] <span class=\"w bgb\">White on blue</span> and <span class=\"k bgbr\">Black on bright red</span>",
	"[src35] <span class=\"br bgby bold\">Bright red on bright yellow</span> <span class=\"bm bgbc u\">Bright magenta on bright cyan</span>",
	"[src36] <span class=\"bw bgbk\">Bright white on bright black <span class=\"bc bgbm bold\">bright cyan on bright magenta</span></span>",
	"[src37] <span class=\"br bgb bold u\">Bright red bold underlined on blue</span>",
	"[src38] <span class=\"g bgk bold\">Green on black bold <span class=\"y bgbr rev\">Yellow reversed on bright red</span></span>",
	"[src39] <span class=\"bw bgb dim\">Bright white dimmed on blue <span class=\"br bgby u\">bright red underlined on bright yellow</span></span>",
	"[src40] <span class=\"br bold bgk\">Error:</span>|<span class=\"by bgb\">Important section</span>",
	"[src41] Normal text <error>Critical error message</error> followed by <warning>Warning message</warning>",
}

// Helper function to print test cases and usage examples
func doTestCasesAndUsage() {
	customTags := NewCustomTagManager()
	customTags.Add("error", "br bold bgk") // bright red bold on black
	customTags.Add("warning", "y bold")    // yellow bold
	customTags.Add("success", "bg bold")   // bright green bold
	customTags.Add("info", "bc u")         // bright cyan underlined
	s := ""
	for i, testCase := range TestCases {
		s = testCase
		s = fmt.Sprintf("[%d]:%s", i+1, EncodeHTMLToTerm(customTags, s)) // termColEncodeCustom(s))
		fmt.Println(s)
	}
}
