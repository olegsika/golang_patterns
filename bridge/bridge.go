package bridge

func Example() {
	amdVideoCard := &AMD{}
	nVidiaVideoCard := &NVidia{}

	apple := Apple{}

	apple.SetVideoCard(amdVideoCard)
	apple.Show()

	apple.SetVideoCard(nVidiaVideoCard)
	apple.Show()

	samsung := Samsung{}

	samsung.SetVideoCard(amdVideoCard)
	samsung.Show()

	samsung.SetVideoCard(nVidiaVideoCard)
	samsung.Show()
}
