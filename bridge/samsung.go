package bridge

import "fmt"

type Samsung struct {
	videoCard VideoCard
}

func (a *Samsung) Show() {
	fmt.Println("Show by Samsung")
	a.videoCard.ShowVideo()
}

func (a *Samsung) SetVideoCard(videoCard VideoCard) {
	a.videoCard = videoCard
}
