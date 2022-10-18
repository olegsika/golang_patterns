package bridge

import "fmt"

type AMD struct {
}

func (n *AMD) ShowVideo() {
	fmt.Println("Show video by AMD")
}
