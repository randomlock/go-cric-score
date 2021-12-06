package models

import "fmt"

type BallType int

const (
    BallTypeWide BallType = iota
    BallTypeNoBall
    NormalBall
)

type Ball struct {
    runScored int
    ballType BallType
}

func NewBall(runScored int, ballType BallType) *Ball {
    return &Ball{runScored: runScored, ballType: ballType}
}

func (b *Ball) RunScored() int {
    return b.runScored
}

func (b *Ball) BallType() BallType {
    return b.ballType
}

func (b *Ball) ShouldCount() bool  {
    return b.ballType == NormalBall
}

func (b *Ball) Print()  {
    println(fmt.Sprintf("Run Scored - %d - BallType - %v", b.runScored, b.ballType))
}