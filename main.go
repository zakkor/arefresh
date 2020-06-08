package main

import (
	"os"

	"9fans.net/go/acme"
)

func main() {
	winfos, err := acme.Windows()
	if err != nil {
		panic(err)
	}

	for _, winfo := range winfos {
		if _, err := os.Stat(winfo.Name); err != nil {
			// skip windows that don't correspond to files
			continue
		}

		win, err := acme.Open(winfo.ID, nil)
		if err != nil {
			panic(err)
		}

		err = win.Ctl("get")
		if err != nil {
			panic(err)
		}
	}
}
