package config

type AmmoStruct struct {
	Ammo []struct {
		Name  string `json:"Name"`
		Stats struct {
			Penetration     string `json:"penetration"`
			FleshDamage     string `json:"FleshDamage"`
			ArmourDamage    string `json:"ArmourDamage"`
			ProjectileSpeed string `json:"ProjectileSpeed"`
			FragmentChance  string `json:"Fragment_Chance"`
			RichochetChance string `json:"RichochetChance"`
		} `json:"Stats"`
		Weapons  []string `json:"weapons"`
		ImageURL string   `json:"image_url"`
	} `json:"ammo"`
}

type Settings struct {
	Token         string   `json:"token"`
	Prefix        string   `json:"prefix"`
	AuthorizedIds []string `json:"authorized_ids"`
}
