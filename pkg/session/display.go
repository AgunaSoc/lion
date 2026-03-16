package session

import (
	"github.com/spf13/viper"

	"lion/pkg/guacd"
)

type valueType string

const (
	Boolean valueType = "boolean"
	String  valueType = "string"
	Integer valueType = "integer"
)

type DisplayParameter struct {
	Key          string
	DefaultValue string
	valueType    valueType
}

var (
	colorDepth               = DisplayParameter{Key: guacd.RDPColorDepth, DefaultValue: "24", valueType: Integer}
	dpi                      = DisplayParameter{Key: guacd.RDPDpi, DefaultValue: "", valueType: Integer}
	disableAudio             = DisplayParameter{Key: guacd.RDPDisableAudio, DefaultValue: "", valueType: Boolean}
	enableWallpaper          = DisplayParameter{Key: guacd.RDPEnableWallpaper, DefaultValue: "", valueType: Boolean}
	enableTheming            = DisplayParameter{Key: guacd.RDPEnableTheming, DefaultValue: "", valueType: Boolean}
	enableFontSmoothing      = DisplayParameter{Key: guacd.RDPEnableFontSmoothing, DefaultValue: "", valueType: Boolean}
	enableFullWindowDrag     = DisplayParameter{Key: guacd.RDPEnableFullWindowDrag, DefaultValue: "", valueType: Boolean}
	enableDesktopComposition = DisplayParameter{Key: guacd.RDPEnableDesktopComposition, DefaultValue: "", valueType: Boolean}
	enableMenuAnimations     = DisplayParameter{Key: guacd.RDPEnableMenuAnimations, DefaultValue: "", valueType: Boolean}
	disableBitmapCaching     = DisplayParameter{Key: guacd.RDPDisableBitmapCaching, DefaultValue: "", valueType: Boolean}
	disableOffscreenCaching  = DisplayParameter{Key: guacd.RDPDisableOffscreenCaching, DefaultValue: "", valueType: Boolean}
	vncCursorRender          = DisplayParameter{Key: guacd.VNCCursor, DefaultValue: "", valueType: String}
	enableConsoleAudio       = DisplayParameter{Key: guacd.RDPConsoleAudio, DefaultValue: "", valueType: Boolean}
	enableAudioInput         = DisplayParameter{Key: guacd.RDPEnableAudioInput, DefaultValue: "", valueType: Boolean}
)

type Display struct {
	data map[string]DisplayParameter
}

func (d Display) GetDisplayParams() map[string]string {
	res := make(map[string]string)
	for envKey, displayParam := range d.data {
		res[displayParam.Key] = displayParam.DefaultValue
		if value := viper.GetString(envKey); value != "" {
			switch displayParam.valueType {
			case Boolean:
				booleanValue := viper.GetBool(envKey)
				res[displayParam.Key] = ConvertBoolToString(booleanValue)
			default:
				res[displayParam.Key] = value
			}
		}
	}
	return res
}

var RDPDisplay = Display{data: map[string]DisplayParameter{
	"ATHERLOCK_COLOR_DEPTH":                colorDepth,
	"ATHERLOCK_DPI":                        dpi,
	"ATHERLOCK_DISABLE_AUDIO":              disableAudio,
	"ATHERLOCK_ENABLE_WALLPAPER":           enableWallpaper,
	"ATHERLOCK_ENABLE_THEMING":             enableTheming,
	"ATHERLOCK_ENABLE_FONT_SMOOTHING":      enableFontSmoothing,
	"ATHERLOCK_ENABLE_FULL_WINDOW_DRAG":    enableFullWindowDrag,
	"ATHERLOCK_ENABLE_DESKTOP_COMPOSITION": enableDesktopComposition,
	"ATHERLOCK_ENABLE_MENU_ANIMATIONS":     enableMenuAnimations,
	"ATHERLOCK_DISABLE_BITMAP_CACHING":     disableBitmapCaching,
	"ATHERLOCK_DISABLE_OFFSCREEN_CACHING":  disableOffscreenCaching,
	"ATHERLOCK_ENABLE_CONSOLE_AUDIO":       enableConsoleAudio,
	"ATHERLOCK_ENABLE_AUDIO_INPUT":         enableAudioInput,
}}

var VNCDisplay = Display{data: map[string]DisplayParameter{
	"ATHERLOCK_COLOR_DEPTH":       colorDepth,
	"ATHERLOCK_VNC_CURSOR_RENDER": vncCursorRender,
}}

var RDPBuiltIn = map[string]string{
	guacd.RDPDisableGlyphCaching: BoolTrue,
}
