package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	files, _ := filepath.Glob("*.ini")
	for _, file := range files {
		bak_file := file + ".bak"
		if err := os.Rename(file, bak_file); err == nil {
			if bb, err := os.ReadFile(bak_file); err == nil {
				f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
				if err == nil {
					for i, b := range bb {
						if b < 128 {
							f.Write(bb[i : i+1])
						} else {
							s := []byte(string(charmap.Windows1251.DecodeByte(b)))
							for _, sb := range s {
								str := fmt.Sprintf("/%X", sb)
								f.Write([]byte(str))
							}
						}
					}
					f.Close()
					fmt.Println(file, " - выполнено")
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("\nНажмите <Enter> для завершения")
	fmt.Scanln()
}
