package puppy

import (
	"fmt"

	"github.com/marius-chirila/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func DifferentTag() {
	fmt.Println("Added a tag!")
}
