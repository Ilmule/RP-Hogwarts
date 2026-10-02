package asciiart

const (
	Reset      = "\033[0m"                // Reset
	Red        = "\033[38;2;217;45;11m"   // Rouge
	Green      = "\033[38;2;31;181;22m"   // Vert
	Yellow     = "\033[38;2;203;217;72m"  // Jaune
	Blue       = "\033[38;2;0;148;219m"   // Bleu
	Purple     = "\033[38;2;131;0;219m"   // Violet
	Cyan       = "\033[38;2;0;255;251m"   // Cyan
	White      = "\033[38;2;255;255;255m" // Blanc
	Black      = "\033[38;2;0;0;0m"       // Noir basique
	DarkGray   = "\033[38;2;84;84;84m"    // Gris foncé (Noir brillant)
	LightGray  = "\033[38;2;196;196;196m" // Gris clair
	Orange     = "\033[38;2;255;145;0m"   // Orange
	Pink       = "\033[38;2;238;108;245m" // Rose fluo
	SoftPink   = "\033[38;2;249;161;255m" // Rose pastel
	YellowPale = "\033[38;2;255;255;107m"
	BrownDark  = "\033[38;2;87;67;49m"
	BrownLight = "\033[38;2;138;86;50m"
	Beige      = "\033[38;2;184;155;108m"
	BeigeLight = "\033[38;2;224;197;148m"

	Green1 = "\033[38;2;38;112;44m"   // Vert forêt / Sombre
	Green2 = "\033[38;2;7;217;23m"    // Vert émeraude / Standard
	Green3 = "\033[38;2;0;255;20m"    // Vert vif / Brillant
	Green4 = "\033[38;2;0;255;20m"    // Vert clair / Lime
	Green5 = "\033[38;2;110;255;123m" // Vert fluo / Néon

	Welcometo = "\033[38;2;255;215;0m" // jaune or
	Hogwarts = "\033[38;2;220;20;60m" // Rouge cramoisi
)


var Nomscraftcolores = map[string]string{
	"Armure en cuir":          BrownLight + "Armure en cuir" + Reset,
	"Armure en fer":           LightGray + "Armure en fer" + Reset,
	"Armure en fer améliorée": DarkGray + "Armure en fer améliorée" + Reset,
	"Armure magique":          Purple + "Armure magique" + Reset,

	"Epée en fer":           DarkGray + "Epée en fer" + Reset,
	"Epée en fer améliorée": SoftPink + "Epée en fer améliorée" + Reset,

	"Baguette magique": Pink + "Baguette magique" + Reset,
	"Bâton de sorcier": BrownLight + "Bâton de sorcier" + Reset,

	"Arc en bois":   BrownDark + "Arc en bois" + Reset,
	"Arc en bambou": Green4 + "Arc en bambou" + Reset,

	"Bois":                BrownDark + "Bois" + Reset,
	"Bambou":              Green + "Bambou" + Reset,
	"Cuir":                BrownLight + "Cuir" + Reset,
	"Fer":                 LightGray + "Fer" + Reset,
	"Ficelle":             White + "Ficelle" + Reset,
	"Crystal de niveau 1": Pink + "Crystal de niveau 1" + Reset,
	"Crystal de niveau 2": Purple + "Crystal de niveau 2" + Reset,
}

var NomsArmesColores = map[string]string{
	"Mains":                 Beige + "Mains" + Reset,
	"Epée en bois":          BrownDark + "Epée en bois" + Reset,
	"Lance pierre":          LightGray + "Lance pierre" + Reset,
	"Epée en fer":           DarkGray + "Epée en fer" + Reset,
	"Epée en fer améliorée": SoftPink + "Epée en fer améliorée" + Reset,
	"Baguette magique":      Pink + "Baguette magique" + Reset,
	"Bâton de sorcier":      BrownLight + "Bâton de sorcier" + Reset,
	"Arc en bois":           BrownDark + "Arc en bois" + Reset,
	"Arc en bambou":         Green4 + "Arc en bambou" + Reset,
}

var NomsArmuresColores = map[string]string{
	"Armure en cuir":          BrownLight + "Armure en cuir" + Reset,
	"Armure en fer":           LightGray + "Armure en fer" + Reset,
	"Armure en fer améliorée": DarkGray + "Armure en fer améliorée" + Reset,
	"Armure magique":          Purple + "Armure magique" + Reset,
}

var NomsAttaquesColores = map[string]string{
	"Boule de feu":         Red + "Boule de feu" + Reset,
	"Coup de baguette":     Orange + "Coup de baguette" + Reset,
	"Coup d'épée":          DarkGray + "Coup d'épée" + Reset,
	"Lancer de poussière":  LightGray + "Lancer de poussière" + Reset,
	"Tir":                  BrownDark + "Tir" + Reset,
	"Lancer de projectile": BrownLight + "Lancer de projectile" + Reset,
	"Coup de poing":        BeigeLight + "Coup de poing" + Reset,
	"Coup de pieds":        Beige + "Coup de pieds" + Reset,
}
