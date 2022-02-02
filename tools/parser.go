package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// ammo parser
// doesnt include images

var (
	FileName = "types.txt"
	Output   = "ammo.json"
)

type Data struct {
	Ammo []struct {
		Name  string
		Stats struct {
			Penetration     string
			FleshDamage     string
			ArmourDamage    string
			ProjectileSpeed string
			Fragment_Chance string
			RichochetChance string
		}
		Weapons   []string
		Image_URL string
	}
}

var (
	SD Data
)

func main() {
	// read file
	f, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		linesplit := strings.Split(scanner.Text(), "	")
		log.Printf("\x1b[38;5;214m[+] Apending To Ouput | \x1b[38;5;51m%s\x1b[0;37m", linesplit[0])
		SD.Ammo = append(SD.Ammo, struct {
			Name  string
			Stats struct {
				Penetration     string
				FleshDamage     string
				ArmourDamage    string
				ProjectileSpeed string
				Fragment_Chance string
				RichochetChance string
			}
			Weapons   []string
			Image_URL string
		}{
			Name: linesplit[0],
			Stats: struct {
				Penetration     string
				FleshDamage     string
				ArmourDamage    string
				ProjectileSpeed string
				Fragment_Chance string
				RichochetChance string
			}{
				Penetration:     linesplit[1],
				FleshDamage:     linesplit[2],
				ArmourDamage:    linesplit[3],
				ProjectileSpeed: linesplit[4],
				Fragment_Chance: linesplit[5],
				RichochetChance: linesplit[6],
			},
			Weapons:   []string{},
			Image_URL: "",
		})

	}
	js, err := json.Marshal(SD)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Create(Output)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Fprint(file, string(js))
}
