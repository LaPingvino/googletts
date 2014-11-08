// This package interfaces with the Google Translate TTS undocumented api as
// found on http://findingscience.com/linux/os_x/2010/11/28/os-x-say-command-on-linux.html
// The goal is to be able to let your program say something.
package tts

import (
	"os/exec"
	"html/template"
)

var Player = "cvlc"

func ttsLink(text string, language string) string {
	return "http://translate.google.com/translate_tts?tl=" + language + "&q=" + template.URLQueryEscaper(text)
}

type Text struct {
	Text string
	Language string
	Player string
}

func (t Text) Play() err {
	t.PlayWith(t.Player)
}

func (t Text) PlayWith(player string) err {
	cmd := exec.Command(player, ttsLink(t.Text, t.Language))
	err := cmd.Run()
	return err
}

func (t Text) GetURL() string {
	return ttsLink(t.Text, t.Language)
}

func NewText(text string, language string) Text {
	if len(text) < 2 {
		text = "en"
	}
	return Text{
		Text: text,
		Language: language,
		Player: Player,
	}
}
