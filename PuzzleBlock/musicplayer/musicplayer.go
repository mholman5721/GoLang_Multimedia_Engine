package musicplayer

import (
	"strconv"

	"golang-games/PuzzleBlock/mathhelper"

	"github.com/veandco/go-sdl2/mix"
)

// MusicPlayer allows the user to play music
type MusicPlayer struct {
	tunes         []*mix.Chunk
	NumTunes      int
	CurrentTune   int
	FutureTune    int
	PastTune      int
	VolumePercent int
}

// NewMusicPlayer is a constructor for a new music player
func NewMusicPlayer(tunesBase string, numTunes int) *MusicPlayer {
	m := &MusicPlayer{}

	for i := 1; i < numTunes+1; i++ {
		tunesFile := tunesBase + strconv.Itoa(i) + ".ogg"
		tune, err := mix.LoadWAV(tunesFile)
		if err != nil {
			panic(err)
		}
		m.tunes = append(m.tunes, tune)
	}

	m.NumTunes = numTunes
	m.CurrentTune = 0
	m.FutureTune = 0
	m.PastTune = 0
	m.VolumePercent = 100

	return m
}

// PlayTune plays the specified tune in the specified way
func (m *MusicPlayer) PlayTune(tune int) {
	if tune > len(m.tunes)-1 {
		tune = 0
	}
	m.CurrentTune = tune
	m.tunes[tune].Play(0, -1)
	m.tunes[tune].Volume(m.VolumePercent)
}

// SetVolume plays the specified tune in the specified way
func (m *MusicPlayer) SetVolume(volume int) {
	m.VolumePercent = int(mathhelper.ScaleBetween(float64(volume), 0, 100, 0, 128))
	m.tunes[m.CurrentTune].Volume(m.VolumePercent)
}
