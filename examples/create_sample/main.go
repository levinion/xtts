package main

import (
	"fmt"

	"github.com/levinion/xtts-go"
)

func main() {
	client := xtts.C("Bearer 5mrNIhoEcHo6kLRRStxsEItz0E69MCwKrRdxedGq9t1yzpSxOZ5ZI8B5HhfeabB9")
	sample, err := client.CreateSample(xtts.DefaultSampleConf("a619f3f4-aabc-4201-b02e-99dba90afb5c", "this is a beautiful dream, and you open your arms to hug them"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sample)
}
