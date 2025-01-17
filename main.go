package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"parsa/FileDetector"
)

var vlog = "v1.0.10 -> `-1a` chenged to `-a`\nv1.1.0 -> fix bug and some erorrs and remove `-n`"

const help = "Locas :\n -1 and -1a  for show path and item in path.\n -h  for show hostname.\n -n  for create new file(name).\n -rm for remove(name).\n -rn for rename (oldname , newname).\nfor install use --install\nnote: for install use sudo temp `sudo ./lo --install`"

func main() {
	if len(os.Args) == 2 /*|| len(os.Args) == 4*/ {
		f := os.Args[1]
		if f == "-1" { /* ======================ls(-1)=======================  */
			path, _ := os.Getwd()
			list, _ := os.ReadDir(path)
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

		} else if f == "-h" { /* ======================hostname(-h)=======================  */
			hostname, _ := os.Hostname()
			fmt.Printf("👤 %v", hostname)

		} else if f == "--help" { /* ======================(--help)=======================  */
			fmt.Print("Locas :\n -1 and -1a  for show path and item in path.\n -h  for show hostname.\n -n  for create new file(name).\n -rm for remove(name).\n -rn for rename (oldname , newname).\nfor install use --install\nnote: for install use sudo temp `sudo ./lo --install`")

			// } else if f == "-n" { /* ======================newfile(-n)=======================  */
			// 	if len(os.Args) == 3 {
			// 		namefile := os.Args[2]
			// 		_, err := os.Create(namefile)
			// 		if err != nil {
			// 			fmt.Print(err)
			// 		}
			// 		fmt.Print("✅ File Created")
			// 	} else if len(os.Args) == 2 {
			// 		fmt.Println("Missng a arg")
			// 	} else if len(os.Args) > 3 {
			// 		fmt.Println("Erorr Extra arg")
			// 	}
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

			path, _ := os.Getwd()
			list, _ := os.ReadDir(path)
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
			fmt.Print("v1.1.0\nBlaster_Void\n")
			fmt.Println(vlog)
		} else {
			fmt.Print(help)
		}

	} else if len(os.Args) > 4 {
		fmt.Print(help)
	} else {
		/* ======================ls(-1)=======================  */
		path, _ := os.Getwd()
		list, _ := os.ReadDir(path)
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
}

func isroot() bool {
	cuusr, err := user.Current()
	if err != nil {
		log.Fatal("Erorr User not root :", err)
	}
	return cuusr.Username == "root"
}
