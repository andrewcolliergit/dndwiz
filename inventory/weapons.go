package inventory

type (
	Weapon           int
	WeaponCategory   int
	WeaponType       int
	WeaponDamage     int
	WeaponProperties int
	WeaponMastery    int
	WeaponInfo       struct {
		Name       string
		Damage     WeaponDamage
		Properties WeaponProperties
		Mastery    WeaponMastery
		Weight     ItemWeight
		Cost       Currency
	}
)
