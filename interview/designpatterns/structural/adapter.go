package structural

import "fmt"

type MediaPlayer interface {
	Play(fileType, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(fileName string) {
	fmt.Println("Playing vlc file. Name:", fileName)
}

func (v *VlcPlayer) PlayMp4(fileName string) {
	fmt.Println("PlayMp4: ", fileName)
}

type Mp4Player struct{}

func (m *Mp4Player) PlayVlc(fileName string) {
	fmt.Println("PlayVlc: ", fileName)
}

func (m *Mp4Player) PlayMp4(fileName string) {
	fmt.Println("Playing mp4 file. Name:", fileName)
}

type MediaAdapter struct {
	advancedMediaPlayer AdvancedMediaPlayer
}

func (m *MediaAdapter) Play(fileType, fileName string) {
	if fileType == "vlc" {
		m.advancedMediaPlayer.PlayVlc(fileName)
	} else if fileType == "mp4" {
		m.advancedMediaPlayer.PlayMp4(fileName)
	}
}

type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func (a *AudioPlayer) Play(fileType, fileName string) {
	if fileType == "mp3" {
		fmt.Println("Playing mp3 file. Name:", fileName)
	} else if fileType == "vlc" || fileType == "mp4" {
		a.mediaAdapter = &MediaAdapter{}
		if fileType == "vlc" {
			a.mediaAdapter.advancedMediaPlayer = &VlcPlayer{}
		} else {
			a.mediaAdapter.advancedMediaPlayer = &Mp4Player{}
		}
		a.mediaAdapter.Play(fileType, fileName)
	} else {
		fmt.Println("Invalid media. ", fileType, " format not supported")
	}
}

func Adapter() {
	audioPlayer := &AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
