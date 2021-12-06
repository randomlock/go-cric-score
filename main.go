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

    matchMap := matchController.GetMatchInfo(match.Id())
    println(matchMap.)

}
