package main

import (
	"fmt"

	"github.com/levinion/xtts-go"
)

func main() {
	client := xtts.C("Bearer 5mrNIhoEcHo6kLRRStxsEItz0E69MCwKrRdxedGq9t1yzpSxOZ5ZI8B5HhfeabB9")
	voice, err := client.CreateVoiceFromFiles("test_create_voice_from_file","./tests/create_voice_from_file/sample.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(voice)
}
