package models

type MatchStatus int

const (
    MatchStatusInitiated MatchStatus = iota
    MatchStatusInProgress
    MatchStatusFinished
)

type Match struct {
    *BaseModel
    innings []*Innings
    status  MatchStatus
    currentInning int
    noOfInnings int
    noOfOver int
    noOfBatsman int
    noOfBowler int
}

func NewMatch(innings []*Innings, noOfInnings int, noOfOver int, noOfBatsman int, noOfBowler int) *Match {
    return &Match{
        BaseModel: NewBaseModel(),
        innings: innings,
        status: MatchStatusInitiated,
        currentInning: 1,
        noOfInnings: noOfInnings,
        noOfOver: noOfOver,
        noOfBatsman: noOfBatsman,
        noOfBowler: noOfBowler,
    }
}

func (m *Match) CurrentInning() int {
    return m.currentInning
}

func (m *Match) SetCurrentInning(currentInning int) {
    m.currentInning = currentInning
}

func (m *Match) Status() MatchStatus {
    return m.status
}

func (m *Match) SetStatus(status MatchStatus) {
    m.status = status
}

func (m *Match) IsMatchFinished() bool {

    for _, inning := range m.Innings() {
        if !inning.IsInningFinished() {
            return false
        }
    }
    return true
}

func (m Match) Innings() []*Innings {
    return m.innings
}

func (m Match) GetCurrentInnings() *Innings {
    for _, inning := range m.Innings() {
        if !inning.IsInningFinished() {
            return inning
        }
    }
    return nil
}
