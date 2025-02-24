package demo

import "github.com/jesseduffield/lazygit/pkg/config"

// Gives us nicer colours when we generate a git repo history with `shell.CreateRepoHistory()`
func setGeneratedAuthorColours(config *config.AppConfig) {
	config.GetUserConfig().Gui.AuthorColors = map[string]string{
		"Fredrica Greenhill": "#fb5aa3",
		"Oscar Reuenthal":    "#86c82f",
		"Paul Oberstein":     "#ffd500",
		"Siegfried Kircheis": "#fe7e11",
		"Yang Wen-li":        "#8e3ccb",
	}
}

func setDefaultDemoConfig(config *config.AppConfig) {
	// demos look much nicer with icons shown
	config.GetUserConfig().Gui.NerdFontsVersion = "3"
}
