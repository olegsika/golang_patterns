package bridge

type Computer interface {
	Show()
	SetVideoCard(VideoCard)
}

type VideoCard interface {
	ShowVideo()
}
