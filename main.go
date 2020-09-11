package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"native-chi-mariadb/engine"
	"native-chi-mariadb/model"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("Tambahkan sebuah argumen")
	}

	fmt.Println(args)

	switch args[0] {
	case "html":
		if args[1] != "" {
			jsonFile, err := os.Open(args[1])
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			jsonByte, _ := ioutil.ReadAll(jsonFile)
			var b []model.Pair
			json.Unmarshal(jsonByte, &b)
			for _, item := range b {
				engine.MakeHTMLForm(item)
				//engine.MakeControllerForm(item)
			}
		} else {
			panic("Mohon lakukan seperti ini html <lokasi/filejson>")
		}
	}

}
