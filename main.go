package main

import (
	"fmt"

	"github.com/Rocky-6/danny-go/instrument"
)

func main() {
	var bass instrument.Bass = instrument.Bass{}
	var i instrument.Inst
	i = bass
	err := instrument.GetSMF(i)
	if err != nil {
		fmt.Printf("error")
		return
	}
}
