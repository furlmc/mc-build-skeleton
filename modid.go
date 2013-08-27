package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ForgeInfo []struct {
	Modid string
	Name  string
}

func main() {
	mcmod, err := os.Open("src/mcmod.info")
	if err != nil {
		panic(err)
	}
	var infos ForgeInfo
	err = json.NewDecoder(mcmod).Decode(&infos)
	if err != nil {
		panic(err)
	}
	fmt.Println("TARGET=" + infos[0].Modid)
	fmt.Println("TARGET_NAME='" + infos[0].Name + "'")
}
