package models

type Innings struct {
	*BaseModel
    overs []*Over
    batsman []*BatsMan
    bowler []*Bowler
    striker *BatsMan
    nonStriker *BatsMan
}

func (i *Innings) Striker() *BatsMan {
	return i.striker
}

func (i *Innings) SetStriker(striker *BatsMan) {
	i.striker = striker
}

func (i *Innings) NonStriker() *BatsMan {
	return i.nonStriker
}

func (i *Innings) SetNonStriker(nonStriker *BatsMan) {
	i.nonStriker = nonStriker
}

func (i *Innings) Overs() []*Over {
	return i.overs
}

func (i *Innings) Batsman() []*BatsMan {
	return i.batsman
}

func (i *Innings) Bowler() []*Bowler {
	return i.bowler
}



func NewInnings(batsman []*BatsMan, bowler []*Bowler) *Innings {
    return &Innings{batsman: batsman, bowler: bowler, striker: batsman[0], nonStriker: batsman[1], BaseModel: NewBaseModel()}
}

func (i *Innings) AddOver(over *Over)  {
	i.overs = append(i.overs, over)
}

func (i *Innings) SwapStriker()  {
	i.striker, i.nonStriker = i.nonStriker, i.striker
}

func (i *Innings) IsInningFinished() bool  {
	for _, over := range i.overs {
		if !over.IsOverFinished() {
			return false
		}
	}
	// TODO add all out scenario
	// for _, batsman := range i.batsman {
	// 	if !batsman.IsOverFinished() {
	// 		return false
	// 	}
	// }
	return true
}

func (i *Innings) Print()  {
	for _, over := range i.Overs() {
		over.Print()
	}
	println("******************************")
}