package utils

type MatchCriteria string

// Each of these criteria will be used to instantiate a channel
// and a consumer goroutine will be listening for messages in the MQ
const (
	Easy   MatchCriteria = "Easy"
	Medium MatchCriteria = "Medium"
	Hard   MatchCriteria = "Hard"
)

var MatchCriterias = []MatchCriteria{
	Easy, Medium, Hard,
}

func ConstructResultChanIdentifier(str string) string {
	return "results/" + str
}
