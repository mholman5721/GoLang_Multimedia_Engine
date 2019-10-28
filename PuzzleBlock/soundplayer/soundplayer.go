package soundplayer

import (
	"golang-games/PuzzleBlock/mathhelper"

	"github.com/veandco/go-sdl2/mix"
)

// SoundPlayer allows the user to play sounds
type SoundPlayer struct {
	sounds        map[string]*mix.Chunk
	soundNames    []string
	VolumePercent int
}

// NewSoundPlayer is a constructor for a new sound player
func NewSoundPlayer(sounds []string) *SoundPlayer {
	var err error

	s := &SoundPlayer{}

	s.sounds = make(map[string]*mix.Chunk)

	s.soundNames = make([]string, 0)

	for i := range sounds {
		name := sounds[i]
		file := "assets/" + name + ".ogg"
		s.soundNames = append(s.soundNames, name)
		s.sounds[name], err = mix.LoadWAV(file)
		if err != nil {
			panic(err)
		}
	}

	s.VolumePercent = 100

	return s
}

// PlaySound plays the specified tune in the specified way
func (s *SoundPlayer) PlaySound(sound string) {
	s.sounds[sound].Play(-1, 0)
}

// SetVolume plays the specified tune in the specified way
func (s *SoundPlayer) SetVolume(volume int) {
	s.VolumePercent = int(mathhelper.ScaleBetween(float64(volume), 0, 100, 0, 128))
	for _, sound := range s.soundNames {
		s.sounds[sound].Volume(s.VolumePercent)
	}
}
