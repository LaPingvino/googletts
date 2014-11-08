// This package interfaces with the Google Translate TTS undocumented api as
// found on http://findingscience.com/linux/os_x/2010/11/28/os-x-say-command-on-linux.html
// The goal is to be able to let your program say something.
package tts

import (
	"os/exec"
)

var Player = "cvlc"
var PlayerAddArgs = "--play-and-exit"

func ttsLink(text string, language string) string {
	return "http://translate.google.com/translate_tts?tl=" + language + "&q=" + text
}

type Text struct {
	Text string
	Language string
	Player string
	PlayerAddArgs string
}

func (t Text) Play() error {
	err := t.PlayWith(t.Player, t.PlayerAddArgs)
	return err
}

func (t Text) PlayWith(player string, playerAddArgs string) error {
	cmd := exec.Command(player, playerAddArgs, ttsLink(t.Text, t.Language))
	err := cmd.Run()
	return err
}

func (t Text) GetURL() string {
	return ttsLink(t.Text, t.Language)
}

func NewText(text string, language string) Text {
	if len(language) < 2 {
		language = "en"
	}
	return Text{
		Text: text,
		Language: language,
		Player: Player,
		PlayerAddArgs: PlayerAddArgs,
	}
}
