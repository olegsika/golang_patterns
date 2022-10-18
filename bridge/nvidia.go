package bridge

import "fmt"

type NVidia struct {
}

func (n *NVidia) ShowVideo() {
	fmt.Println("Show video by nVidia")
}
