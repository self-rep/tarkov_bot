package search_engine

import (
	"TarkovBot/config"
	"strings"
)

type nwam struct {
	Name  string
	Stats struct {
		Penetration     string
		FleshDamage     string
		ArmourDamage    string
		ProjectileSpeed string
		FragmentChance  string
		RichochetChance string
	}
	Weapons  []string
	ImageUrl string
}

func Search(input string) (ammo *nwam, searchres []nwam) {

	// check if ammo in array

	found, ammo := CheckAmmo(input)
	if !found {
		// search for other results
		var otherAmmo []nwam
		for index, v := range config.AmmoData.Ammo {
			if strings.Contains(strings.ToLower(v.Name), input) {
				otherAmmo = append(otherAmmo, nwam{
					Name: config.AmmoData.Ammo[index].Name,
					Stats: struct {
						Penetration     string
						FleshDamage     string
						ArmourDamage    string
						ProjectileSpeed string
						FragmentChance  string
						RichochetChance string
					}(config.AmmoData.Ammo[index].Stats),
					Weapons:  config.AmmoData.Ammo[index].Weapons,
					ImageUrl: config.AmmoData.Ammo[index].ImageURL,
				})
			}
		}
		return nil, otherAmmo
	}
	return ammo, nil
}

func CheckAmmo(ammo string) (bool, *nwam) {
	for _, v := range config.AmmoData.Ammo {
		if strings.ToLower(v.Name) == ammo {
			return true, &nwam{Name: v.Name, Stats: struct {
				Penetration     string
				FleshDamage     string
				ArmourDamage    string
				ProjectileSpeed string
				FragmentChance  string
				RichochetChance string
			}(v.Stats), Weapons: v.Weapons, ImageUrl: v.ImageURL}
		}
	}
	return false, nil
}
