package bridge

import "fmt"

type Apple struct {
	videoCard VideoCard
}

func (a *Apple) Show() {
	fmt.Println("Show by Apple")
	a.videoCard.ShowVideo()
}

func (a *Apple) SetVideoCard(videoCard VideoCard) {
	a.videoCard = videoCard
}
