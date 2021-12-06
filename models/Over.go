package models


const NoOfBallsPerOver = 6


type Over struct {
    balls []*Ball
}

func (o *Over) Balls() []*Ball {
    return o.balls
}

func NewOver() *Over {
    return &Over{}
}

func (o *Over) AddBall(ball *Ball)  {
    o.balls = append(o.balls, ball)
}

func (o *Over) IsOverFinished() bool {
    count := 0
    for _, b := range o.Balls() {
        if b.ShouldCount() {
            count++
        }
    }
    return count == NoOfBallsPerOver
}

func (o *Over) TotalRun() int {
    score := 0
    for _, b := range o.Balls() {
        score += b.runScored
    }
    return score
}

func (o *Over) Print()  {
    for _, ball := range o.Balls() {
        ball.Print()
    }
    println("-------------------------")
}