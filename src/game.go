package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (p *Player) AfficherSante() string {
	return fmt.Sprintf("%d", p.Health)
}

var (
	JoueurActuel *Player
	Lecteur      = bufio.NewReader(os.Stdin)
)

func InitJeu(p Player) {
	JoueurActuel = &p
}

func LireSaisie(reader *bufio.Reader, message string) string {
	fmt.Print(message)
	saisie, _ := reader.ReadString('\n')
	return strings.TrimSpace(saisie)
}
