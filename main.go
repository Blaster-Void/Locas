package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"parsa/FileDetector"
)

var vlog = "v1.1.0 -> fix Some Bug And Remove `-1` you can use just `lo`\nv1.0.10 -> `-1a` chenged to `-a`\nv1.1.0 -> fix bug and some erorrs and remove `-n` "
var verion = "v1.1.0"

const help = "Locas :\nlo : show items in path\n -a  for show path and hiden items in path.\n -rm for remove(name).\n -rn for rename (oldname , newname).\nfor install use --install\nnote: for install use sudo temp `sudo ./lo --install`"

func main() {
	//check we have arg?
	if len(os.Args) == 2 || len(os.Args) == 3 || len(os.Args) == 4 {
		// yes (ID1)
		f := os.Args[1]
		if f == "--help" { /* ======================(--help)=======================  */
			fmt.Print(help)

		} else if f == "-rm" { /* ======================remove(-rm)=======================  */

			namefile := os.Args[2]
			err := os.Remove(namefile)
			if err != nil {
				fmt.Print(err)
			} else {
				fmt.Print("✅ File Removed")
			}
		} else if f == "-rn" { /* ======================rename(-rn)=======================  */
			if len(os.Args) == 4 {
				old := os.Args[2]
				new := os.Args[3]
				err := os.Rename(old, new)
				if err != nil {
					fmt.Print(err)
				} else {
					fmt.Print("✅ renamed")
				}
			} else if len(os.Args) < 3 {
				fmt.Println("Missing arg")
			} else if len(os.Args) > 3 {
				fmt.Println(os.Args[4], "] Extra arg")
			}
		} else if f == "--install" {

			os.Rename("lo", "/usr/bin/lo")
			per := isroot()
			if per {
				fmt.Print("✅ installed")
			} else {
				fmt.Print("Root Needed")
			}
		} else if f == "-a" { /* ======================hidenls(-1a)=======================  */
			var opt string
			if len(os.Args) == 3 {
				opt = os.Args[2]
			} else {
				opt = ""
			}
			path, _ := os.Getwd()
			path += "/" + opt
			list, err := os.ReadDir(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Path : %v\n", path)
			for _, i := range list {

				if i.IsDir() {

					fmt.Printf("📂 %v\n", i.Name())

				} else {

					continue

				}
			}
			for _, f := range list {
				if f.IsDir() {
					continue
				} else {
					fmt.Printf("📄 %v\n", f.Name())
				}
			}
		} else if f == "-v" {
			fmt.Printf("%s\nBlaster_Void\n", verion)
			fmt.Println(vlog)
		} else {
			opt := os.Args[1]
			path, _ := os.Getwd()
			path += "/" + opt
			list, err := os.ReadDir(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Path : %v\n", path)
			for _, i := range list {
				filename := i.Name()
				if string(filename[0]) == "." {
					continue
				} else {
					if i.IsDir() { //is Dir? yes Print Dir No Dont Print and (.) is hiddin and dont Print
						fmt.Printf("📂 %v\n", i.Name())
					} else {
						continue
					}
				}
			}
			for _, f := range list {
				filename := f.Name()
				if string(filename[0]) == "." {
					continue
				} else {
					if f.IsDir() {
						continue
					} else {
						xt := FileDetector.Scan_File(f.Name())
						fmt.Printf("%v %v\n", FileDetector.Type_File(xt), f.Name())
					}
				}
			}

		}
	} else {
		var opt string
		if len(os.Args) == 2 {
			opt = os.Args[2]
		}
		path, _ := os.Getwd()
		path += "/" + opt
		list, err := os.ReadDir(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Path : %v\n", path)
		for _, i := range list {
			filename := i.Name()
			if string(filename[0]) == "." {
				continue
			} else {
				if i.IsDir() { //is Dir? yes Print Dir No Dont Print and (.) is hiddin and dont Print
					fmt.Printf("📂 %v\n", i.Name())
				} else {
					continue
				}
			}
		}
	}
}
func isroot() bool {
	cuusr, err := user.Current()
	if err != nil {
		log.Fatal("Erorr User not root :", err)
	}
	return cuusr.Username == "root"
}
