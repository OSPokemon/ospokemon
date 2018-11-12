package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ospokemon "ospokemon.com"
	"ztaylor.me/cast"
)

func Editor() {
	reader := bufio.NewReader(os.Stdin)
	var line string
	var parts []string
	fmt.Println("OSPokemon edit mode. Enter 'help', or 'q' to quit")
	for {
		fmt.Print("ospokemon: ")
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-2] // remove \n

		if line == "q" {
			return
		}

		parts = strings.Split(line, " ")

		if parts[0] == "help" {
			EditorHelp(parts[1:])
		} else if parts[0] == "u" && len(parts) > 1 {
			EditorUniverse(parts[1:], reader)
		} else if parts[0] == "a" {
			EditorAccount(reader)
		} else if parts[0] == "c" {
			EditorClass(reader)
		} else {
			fmt.Println("unknown command:", line)
		}
	}
}

func EditorHelp(parts []string) {
	fmt.Print(
		`ospokemon editor help
help  --  --  --  --  -- print this help
u (universeid)           enter EditorUniverse mode
a --  --  --  --  --  -- enter EditorAccount mode
c                        enter EditorClass mode
q --  --  --  --  --  -- quit
`)
}

func EditorUniverse(parts []string, reader *bufio.Reader) {
	universeid := uint(cast.Int(parts[0]))
	universe, err := ospokemon.GetUniverse(universeid)
	if err != nil {
		fmt.Println("OSPokemon edit universe mode failed:", err.Error())
		return
	}
	fmt.Println("OSPokemon edit universe mode. Enter 'help', 'q' to quit")
	prompt := fmt.Sprintf("ospokemon-universe#%d:", universeid)
	var line string
	for {
		fmt.Print(prompt)
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-2] // remove \n

		if line == "q" {
			return
		} else if line == "help" {
			EditorUniverseHelp(universeid)
		} else if line == "print" {
			for k, e := range universe.Entities {
				fmt.Printf("%d: %v\n", k, e)
			}
		} else {
			fmt.Println("unknown command:", line)
		}
	}
}

func EditorUniverseHelp(universeid uint) {
	fmt.Printf(`ospokemon universe(%d) editor help
help  --  --  --  -- print this help
print                print the entities in this universe
`, universeid)
}

func EditorAccount(reader *bufio.Reader) {
	fmt.Println("OSPokemon edit account mode. Enter 'help', 'q' to quit")
	var line string
	for {
		fmt.Print("ospokemon-account:")
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-2] // remove \n
		parts := strings.Split(line, " ")

		if line == "q" {
			return
		} else if line == "help" {
			EditorAccountHelp()
		} else if parts[0] == "print" && len(parts) > 1 && parts[1] != "" {
			account, err := ospokemon.GetAccount(parts[1])
			if err != nil {
				fmt.Println("OSPokemon edit account mode failed:", err.Error())
				return
			}
			fmt.Printf("%+v\n", account)
		} else {
			fmt.Println("unknown command:", line)
		}
	}
}

func EditorAccountHelp() {
	fmt.Printf(`ospokemon account editor help
help  --  --  --  -- print this help
print (username)     print account object
`)
}

func EditorClass(reader *bufio.Reader) {
	fmt.Println("OSPokemon edit class mode. Enter 'help', 'q' to quit")
	var line string
	for {
		fmt.Print("ospokemon-class:")
		line, _ = reader.ReadString('\n')
		line = line[:len(line)-2] // remove \n
		parts := strings.Split(line, " ")

		if line == "q" {
			return
		} else if line == "help" {
			EditorClassHelp()
		} else if parts[0] == "print" && len(parts) > 1 && parts[1] != "" {
			classid := uint(cast.Int(parts[1]))
			class, err := ospokemon.GetClass(classid)
			if err != nil {
				fmt.Println("OSPokemon edit class mode failed:", err.Error())
			}
			fmt.Printf("%+v\n", class)
		} else if line == "print all" {
			fmt.Println("print all classes", len(ospokemon.Classes.Cache))
			for _, c := range ospokemon.Classes.Cache {
				fmt.Printf("%+v\n", c)
			}
		} else {
			fmt.Println("unknown command:", line)
		}
	}
}

func EditorClassHelp() {
	fmt.Printf(`ospokemon class editor help
help  --  --  --  --  print this help
print (classid)       print a class object
print all --  --  --  print all the class objects
`)
}
