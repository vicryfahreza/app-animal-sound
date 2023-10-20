package main

import (
	"fmt"

	go_animal_sound "github.com/vicryfahreza/go-animal-sound"
)

func main() {
	fmt.Println(go_animal_sound.DogSound())
	fmt.Println(go_animal_sound.CatSound())
	fmt.Println(go_animal_sound.DragonSound())
	fmt.Println(go_animal_sound.DuckSound())
}
