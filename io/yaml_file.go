package io

type YamlFile struct {
	Window           Window        `yaml:"window"`
	Scrolling        Scrolling     `yaml:"scrolling"`
	Font             Font          `yaml:"font"`
	Colors           Colors        `yaml:"colors"`
	Bell             Bell          `yaml:"bell"`
	Keybindings      []Keybindings `yaml:"key_bindings"`
	Mouse            Mouse         `yaml:"mouse"`
	Selection        Selection     `yaml:"selection"`
	MouseBindings    MouseBindings `yaml:"mouse_bindings"`
	Cursor           Cursor        `yaml:"cursor"`
	LiveConfigReload bool          `yaml:"live_config_reload"`
	Debug            Debug         `yaml:"debug"`
	Shell            Shell         `yaml:"shell"`
}

type Window struct {
	Dimensions     WindowDimensions `yaml:"dimensions"`
	Padding        Coordinates      `yaml:"padding"`
	DynamicPadding bool             `yaml:"dynamic_padding"`
	Decorations    string           `yaml:"decorations"`
	StartupMode    string           `yaml:"startup_mode"`
	Opacity        float32          `yaml:"opacity"`
}

type WindowDimensions struct {
	Columns int `yaml:"columns"`
	Lines   int `yaml:"lines"`
}

type Coordinates struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Scrolling struct {
	History    int `yaml:"history"`
	Multiplier int `yaml:"multiplier"`
}

type Font struct {
	Normal         FontParams  `yaml:"normal"`
	Bold           FontParams  `yaml:"bold"`
	Italic         FontParams  `yaml:"italic"`
	Size           float32     `yaml:"size"`
	Offset         Coordinates `yaml:"offset"`
	GlyphOffset    Coordinates `yaml:"glyph_offset"`
	UseThinStrokes bool        `yaml:"use_thin_strokes"`
}

type FontParams struct {
	Family string `yaml:"family"`
	Style  string `yaml:"style"`
}

type Colors struct {
	Primary PrimaryColor `yaml:"primary"`
	Cursor  CursorColor  `yaml:"cursor"`
	Normal  ThemeColors  `yaml:"normal"`
	Bright  ThemeColors  `yaml:"bright"`
}

type PrimaryColor struct {
	Background string `yaml:"background"`
	Foreground string `yaml:"foreground"`
}

type CursorColor struct {
	Text   string `yaml:"text"`
	Cursor string `yaml:"cursor"`
}

type ThemeColors struct {
	Black   string `yaml:"black"`
	Red     string `yaml:"red"`
	Green   string `yaml:"green"`
	Yellow  string `yaml:"yellow"`
	Blue    string `yaml:"blue"`
	Magenta string `yaml:"magenta"`
	Cyan    string `yaml:"cyan"`
	White   string `yaml:"white"`
}

type Keybindings struct {
	Key     string `yaml:"key"`
	Mods    string `yaml:"mods,omitempty"`
	Action  string `yaml:"action,omitempty"`
	Mode    string `yaml:"mode,omitempty"`
	Chars   string `yaml:"chars,omitempty"`
	Command string `yaml:"command,omitempty"`
}

type Bell struct {
	Animation string `yaml:"animation"`
	Color     string `yaml:"color"`
	Duration  string `yaml:"duration"`
}

type Mouse struct {
	DoubleClick struct {
		Threshold int `yaml:"threshold"`
	} `yaml:"double_click"`
	TripleClick struct {
		Threshold int `yaml:"threshold"`
	} `yaml:"triple_click"`
	HideWhenTyping bool `yaml:"hide_when_typing"`
}

type Selection struct {
	SemanticEscapeChars string `yaml:"semantic_escape_chars"`
	SaveToClipboard     bool   `yaml:"save_to_clipboard"`
}

type MouseBindings struct {
	Mouse  string `yaml:"mouse"`
	Action string `yaml:"action"`
	Mods   string `yaml:"mods,omitempty"`
}

type Cursor struct {
	Style           string `yaml:"style"`
	UnfocusedHollow bool   `yaml:"unfocused_hollow"`
}

type Debug struct {
	RenderTimer       bool   `yaml:"render_timer"`
	PersistentLogging bool   `yaml:"persistent_logging"`
	LogLevel          string `yaml:"log_level"`
	PrintEvents       bool   `yaml:"print_events"`
	RefTest           bool   `yaml:"ref_test"`
}

type Shell struct {
	Program string   `yaml:"program"`
	Args    []string `yaml:"args"`
}
