package main

import "fmt"

type FootballPlayer struct {
	name      string
	team      string
	goalCount int
}

func main() {
	salah := FootballPlayer{"Salah", "Liverpool", 18}
	score(&salah)
	fmt.Printf("%#v\n", salah)
	salah.score1()
	fmt.Printf("%#v\n", salah)
	salah.score2()
	fmt.Printf("%#v\n", salah)
}

func score(fp *FootballPlayer) {
	fp.goalCount++
}

func (fp FootballPlayer) score1() {
	fp.goalCount++
}

func (fp *FootballPlayer) score2() {
	fp.goalCount++
}

// output :

// main.FootballPlayer{name:"Salah", team:"Liverpool", goalCount:19}
// main.FootballPlayer{name:"Salah", team:"Liverpool", goalCount:19}
// main.FootballPlayer{name:"Salah", team:"Liverpool", goalCount:20}
