package models

type Player struct {
    name string
    age string
}

func (p Player) Name() string {
    return p.name
}

func NewPlayer(name string, age string) *Player {
    return &Player{name: name, age: age}
}


type BatsMan struct {
    *Player
    out bool
}

func (b *BatsMan) SetOut(out bool) {
    b.out = out
}

func NewBatsMan(name string, age string) *BatsMan {
    return &BatsMan{Player: NewPlayer(name, age)}
}



type Bowler struct {
    *Player
}

func NewBowler(name string, age string) *Bowler {
    return &Bowler{Player: NewPlayer(name, age)}
}

