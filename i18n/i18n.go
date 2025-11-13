package i18n

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed errs.*.toml
var LocalFs embed.FS
var Local *i18n.Localizer

type LocalizeConfig = i18n.LocalizeConfig

func init() {
	bundle := i18n.NewBundle(language.SimplifiedChinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(LocalFs, "errs.zh-Hans.toml")
	Local = i18n.NewLocalizer(bundle, language.SimplifiedChinese.String())
}
