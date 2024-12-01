package models

type Setting struct {
	Token          string   `json:"token"`
	ShellLocation  string   `json:"shell_location"`
	Buttons        []string `json:"-"`
	RowButtonCount int      `json:"row_button_count"`
	IsDebug        bool     `json:"is_debug"`
}
