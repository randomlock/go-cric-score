package datastore

import (
    "../models"
)

type Score struct {
    NoOfBallPlayed int
    NoOfRuns int
    NoOf4 int
    NoOf6 int
}


type InningDataStore struct {
    inningDataStore map[string]map[string]*Score
}

func NewInningDataStore() *InningDataStore {
    return &InningDataStore{
        inningDataStore: map[string]map[string]*Score{},
    }
}

func (m *InningDataStore) AddInning(Inning *models.Innings)  {
    m.inningDataStore[Inning.Id()] = map[string]*Score{}
    for _, player := range Inning.Batsman() {
        m.inningDataStore[Inning.Id()][player.Name()] = &Score{}
    }
}

func (m *InningDataStore) GetInningScore(inningId string) (InningScore map[string]*Score) {
    return m.inningDataStore[inningId]
}

func (m *InningDataStore) AddScore(playerId string, inningId string, runs int) {
    score := m.inningDataStore[inningId][playerId]
    var has4, has6 int
    if runs == 4 {
        has4 = 1
    }
    if runs == 6 {
        has6 = 1
    }

    score = &Score{
        NoOfBallPlayed: score.NoOfBallPlayed+1,
        NoOfRuns:       score.NoOfRuns+runs,
        NoOf4:          score.NoOf4+has4,
        NoOf6:          score.NoOf6+has6,
    }

    m.inningDataStore[inningId][playerId] = score
    return
}
