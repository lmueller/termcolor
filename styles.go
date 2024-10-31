package termcolor

var ColorMap = map[string]string{
	"black": "\033[30m", "red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m",
	"blue": "\033[34m", "magenta": "\033[35m", "cyan": "\033[36m", "white": "\033[37m",
	"brightblack": "\033[90m", "brightred": "\033[91m", "brightgreen": "\033[92m", "brightyellow": "\033[93m",
	"brightblue": "\033[94m", "brightmagenta": "\033[95m", "brightcyan": "\033[96m", "brightwhite": "\033[97m",
}

var BColorMap = map[string]string{
	"bgblack": "\033[40m", "bgred": "\033[41m", "bggreen": "\033[42m", "bgyellow": "\033[43m",
	"bgblue": "\033[44m", "bgmagenta": "\033[45m", "bgcyan": "\033[46m", "bgwhite": "\033[47m",
	"bgbrightblack": "\033[100m", "bgbrightred": "\033[101m", "bgbrightgreen": "\033[102m", "bgbrightyellow": "\033[103m",
	"bgbrightblue": "\033[104m", "bgbrightmagenta": "\033[105m", "bgbrightcyan": "\033[106m", "bgbrightwhite": "\033[107m",
}

var StyleMap = map[string]string{
	"bold":      "\033[1m",
	"dim":       "\033[2m",
	"underline": "\033[4m",
	"reverse":   "\033[7m",
}

const (
	StyleReset = "\033[0m"
)

var CSSClassMap = map[string]string{
	// Basic colors
	"r": "color: red",
	"b": "color: blue",
	"g": "color: green",
	"y": "color: yellow",
	"m": "color: magenta",
	"c": "color: cyan",
	"w": "color: white",
	"k": "color: black",

	// Bright colors
	"br": "color: brightred",
	"bb": "color: brightblue",
	"bg": "color: brightgreen",
	"by": "color: brightyellow",
	"bm": "color: brightmagenta",
	"bc": "color: brightcyan",
	"bw": "color: brightwhite",
	"bk": "color: brightblack",

	// Background colors
	"bgr": "background-color: red",
	"bgb": "background-color: blue",
	"bgg": "background-color: green",
	"bgy": "background-color: yellow",
	"bgm": "background-color: magenta",
	"bgc": "background-color: cyan",
	"bgw": "background-color: white",
	"bgk": "background-color: black",

	// Bright background colors
	"bgbr": "background-color: brightred",
	"bgbb": "background-color: brightblue",
	"bgbg": "background-color: brightgreen",
	"bgby": "background-color: brightyellow",
	"bgbm": "background-color: brightmagenta",
	"bgbc": "background-color: brightcyan",
	"bgbw": "background-color: brightwhite",
	"bgbk": "background-color: brightblack",

	// Styles
	"bold": "font-weight: bold",
	"dim":  "opacity: 0.8",
	"u":    "text-decoration: underline",
	"rev":  "filter: invert(100%)",
}
