package log

import (
	"strconv"
	"strings"

	color "github.com/fatih/color"
)

// [snippet source](https://github.com/gookit/color/blob/d4a4cd982d2be4df844a5d25c9fce141ab8c70b4/convert.go#L437)
// HexToRgb convert hex color string to RGB numbers
//
// Usage:
//
//	rgb := HexToRgb("ccc") // rgb: [204 204 204]
//	rgb := HexToRgb("aabbcc") // rgb: [170 187 204]
//	rgb := HexToRgb("#aabbcc") // rgb: [170 187 204]
//	rgb := HexToRgb("0xad99c0") // rgb: [170 187 204]
func HexToRgb(hex string) (rgb []int) {
	hex = strings.TrimSpace(hex)
	if hex == "" {
		return
	}

	// like from css. eg "#ccc" "#ad99c0"
	if hex[0] == '#' {
		hex = hex[1:]
	}

	hex = strings.ToLower(hex)
	switch len(hex) {
	case 3: // "ccc"
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 8: // "0xad99c0"
		hex = strings.TrimPrefix(hex, "0x")
	}

	// recheck
	if len(hex) != 6 {
		return
	}

	// convert string to int64
	if i64, err := strconv.ParseInt(hex, 16, 32); err == nil {
		color := int(i64)
		// parse int
		rgb = make([]int, 3)
		rgb[0] = color >> 16
		rgb[1] = (color & 0x00FF00) >> 8
		rgb[2] = color & 0x0000FF
	}
	return
}

func HexToRGB2(hex string) (r, g, b int) {
	o := HexToRgb(hex)
	return o[0], o[1], o[2]
}

type LikeSprintf interface {
	Sprintf(format string, a ...any) string
}

type colors struct {
	PrefixStyle      LikeSprintf
	MessageStyle     LikeSprintf
	SlogAttrKeyStyle LikeSprintf

	// Styles are in order, DEBUG,INFO,WARN,ERROR,FATAL
	LogLevelStyles [5]LikeSprintf
}

type Theme int

const (
	ThemeNoColor Theme = 0
	ThemeLight   Theme = 1
	ThemeDark    Theme = 2
)

func (t Theme) GetColors() *colors {
	switch t {
	// case ThemeLight:
	// 	{
	// 		fg := lipgloss.Color("#4a84ad")
	// 		bg := lipgloss.Color("#f4f7fb")
	//
	// 		style := lipgloss.NewStyle().Background(bg).Foreground(fg)
	//
	// 		return &colors{
	// 			PrefixStyle:      style,
	// 			SlogAttrKeyStyle: style.Bold(true),
	// 			MessageStyle:     style.UnsetBackground().Foreground(lipgloss.Color("#6d787d")),
	// 			LogLevelStyles: [5]lipgloss.Style{
	// 				style.Foreground(lipgloss.Color("#bdbfbe")).Faint(true),                   // DEBUG
	// 				style.UnsetBackground().Foreground(lipgloss.Color("#099dd6")).Faint(true), // INFO
	// 				style.UnsetBackground().Foreground(lipgloss.Color("#d6d609")).Faint(true), // WARN
	// 				style.UnsetBackground().Foreground(lipgloss.Color("#c76975")),             // ERROR
	// 				style.UnsetBackground().Foreground(lipgloss.Color("#d6d609")).Bold(true),  // FATAL
	// 			},
	// 		}
	// 	}
	case ThemeDark:
		{
			// fg := lipgloss.Color("#9addfc")
			// bg := lipgloss.Color("#172830")

			// style := lipgloss.NewStyle().Foreground(fg)

			// return &colors{
			// 	PrefixStyle:      color2.Hex("#9addfc"),
			// 	MessageStyle:     color2.Hex("#bdbfbe"),
			// 	SlogAttrKeyStyle: color2.Hex("#85d3d4"),
			// 	LogLevelStyles: [5]LikeSprintf{
			// 		color2.Hex("#bdbfbe"),
			// 		color2.Hex("#099dd6"),
			// 		color2.Hex("#d6d609"),
			// 		color2.Hex("#c76975"),
			// 		color2.Hex("#d6d609"),
			// 	},
			// }

			return &colors{
				PrefixStyle:      color.RGB(HexToRGB2("#9addfc")),
				MessageStyle:     color.RGB(HexToRGB2("#bdbfbe")),
				SlogAttrKeyStyle: color.RGB(HexToRGB2("#85d3d4")),
				LogLevelStyles: [5]LikeSprintf{
					color.RGB(HexToRGB2("#bdbfbe")),
					color.RGB(HexToRGB2("#099dd6")),
					color.RGB(HexToRGB2("#d6d609")),
					color.RGB(HexToRGB2("#c76975")),
					color.RGB(HexToRGB2("#d6d609")),
				},
			}
		}
	default: // ThemeNoColor
		{
			return &colors{}
		}
	}
}
