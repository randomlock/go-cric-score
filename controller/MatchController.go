package controller

import (
    "fmt"

    "../datastore"
    "../models"
)

type MatchController struct {
    matchDataStore *datastore.MatchDataStore
    inningDataStore *datastore.InningDataStore
}

func NewMatchController(matchDataStore *datastore.MatchDataStore, inningDataStore *datastore.InningDataStore) *MatchController {
    return &MatchController{matchDataStore: matchDataStore, inningDataStore: inningDataStore}
}

func (c *MatchController) CreateNewMatch(innings int, player int, over int) (match *models.Match, err error) {

    if innings % 2 != 0 {
        return match, fmt.Errorf("no of inning shoud be even")
    }
    var inn []*models.Innings
    for i := 0; i < innings; i++ {
        var batsman []*models.BatsMan
        var bowler []*models.Bowler
        for i := 0; i < player; i++ {
            batsman = append(batsman, models.NewBatsMan(fmt.Sprintf("player %d", i), "20"))
            bowler = append(bowler, models.NewBowler(fmt.Sprintf("player %d", i), "20"))
        }


        inning := models.NewInnings(batsman, bowler)
        for i := 0; i < over; i++ {
            inning.AddOver(models.NewOver())
        }
        c.inningDataStore.AddInning(inning)
        inn = append(inn, inning)
    }
    match = models.NewMatch(inn, innings, over, player, player)
    c.matchDataStore.AddMatch(match)
    return match, nil
}


func (c *MatchController) Play(match *models.Match, inning *models.Innings, over *models.Over, ball *models.Ball) {

    currentPlayer := inning.Striker()

    c.inningDataStore.AddScore(currentPlayer.Name(), inning.Id(), ball.RunScored())

    if ball.RunScored() % 2 != 0 && !over.IsOverFinished() {
        inning.SwapStriker()
    }
    if ball.RunScored() % 2 == 0 && over.IsOverFinished() {
        inning.SwapStriker()
    }

    // c.matchDataStore.UpdateMatch(match)

}


func (c *MatchController) PrintScore(inning *models.Innings) {

    score := c.inningDataStore.GetInningScore(inning.Id())

    for player, score := range score {
        println(fmt.Sprintf("player - %s - run scored - %d ball played - %d 4s - %d 6s - %d", player, score.NoOfRuns, score.NoOfBallPlayed, score.NoOf4, score.NoOf6))
    }
    println("=============================================")
}

func (c *MatchController) GetMatchInfo(matchId string) *models.Match {
    return c.matchDataStore.GetMatch(matchId)
}