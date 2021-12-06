package main

import (
    "math/rand"

    "./controller"
    "./datastore"
    "./models"
)


func main() {
    matchController := controller.NewMatchController(datastore.NewMatchDataStore(), datastore.NewInningDataStore())
    match, err := matchController.CreateNewMatch(2, 5, 2)
    if err != nil {
        println("error in creating new match")
        return
    }

    for !match.IsMatchFinished() {
        inning := match.GetCurrentInnings()
        for _, over := range inning.Overs() {
            for !over.IsOverFinished() {
                ball := models.NewBall(rand.Intn(6) + 1, models.NormalBall)
                over.AddBall(ball)
                matchController.Play(match, inning, over, ball)
            }
            matchController.PrintScore(inning)
        }
    }

    for _, inning := range match.Innings() {
        inning.Print()
    }


}

// OutPut
//
// player - player 4 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 0 - run scored - 25 ball played - 6 4s - 1 6s - 3
// player - player 1 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 2 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 3 - run scored - 0 ball played - 0 4s - 0 6s - 0
// =============================================
// player - player 3 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 4 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 0 - run scored - 33 ball played - 10 4s - 1 6s - 3
// player - player 1 - run scored - 8 ball played - 2 4s - 0 6s - 0
// player - player 2 - run scored - 0 ball played - 0 4s - 0 6s - 0
// =============================================
// player - player 4 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 0 - run scored - 4 ball played - 2 4s - 0 6s - 0
// player - player 1 - run scored - 21 ball played - 4 4s - 1 6s - 2
// player - player 2 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 3 - run scored - 0 ball played - 0 4s - 0 6s - 0
// =============================================
// player - player 0 - run scored - 16 ball played - 5 4s - 0 6s - 1
// player - player 1 - run scored - 33 ball played - 7 4s - 1 6s - 3
// player - player 2 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 3 - run scored - 0 ball played - 0 4s - 0 6s - 0
// player - player 4 - run scored - 0 ball played - 0 4s - 0 6s - 0
// =============================================
// Run Scored - 6 - BallType - 2
// Run Scored - 4 - BallType - 2
// Run Scored - 6 - BallType - 2
// Run Scored - 6 - BallType - 2
// Run Scored - 2 - BallType - 2
// Run Scored - 1 - BallType - 2
// -------------------------
// Run Scored - 2 - BallType - 2
// Run Scored - 3 - BallType - 2
// Run Scored - 5 - BallType - 2
// Run Scored - 1 - BallType - 2
// Run Scored - 3 - BallType - 2
// Run Scored - 2 - BallType - 2
// -------------------------
// ******************************
// Run Scored - 1 - BallType - 2
// Run Scored - 6 - BallType - 2
// Run Scored - 5 - BallType - 2
// Run Scored - 3 - BallType - 2
// Run Scored - 4 - BallType - 2
// Run Scored - 6 - BallType - 2
// -------------------------
// Run Scored - 6 - BallType - 2
// Run Scored - 3 - BallType - 2
// Run Scored - 6 - BallType - 2
// Run Scored - 1 - BallType - 2
// Run Scored - 3 - BallType - 2
// Run Scored - 5 - BallType - 2
// -------------------------
// ******************************

