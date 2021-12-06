package datastore

import (
    "fmt"

    "../models"
)

type MatchDataStore struct {
    matchDataStore map[string]*models.Match
}

func NewMatchDataStore() *MatchDataStore {
    return &MatchDataStore{
        matchDataStore: map[string]*models.Match{},
    }
}

func (m *MatchDataStore) AddMatch(match *models.Match)  {
    m.matchDataStore[match.Id()] = match
}

func (m *MatchDataStore) UpdateMatch(match *models.Match)  {
    m.matchDataStore[match.Id()] = match
}

func (m *MatchDataStore) GetMatch(matchId string) *models.Match {
    return m.matchDataStore[matchId]
}

func (m MatchDataStore) TotalRunScored(matchId string, inning int) (int, error) {
    innings := m.GetMatch(matchId).Innings()
    if inning > len(innings) {
        return 0, fmt.Errorf("inning doesn't exist")
    }
    score := 0
    for _, over := range innings[inning+1].Overs() {
        score += over.TotalRun()
    }
    return score, nil
}