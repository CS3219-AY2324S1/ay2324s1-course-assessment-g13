package utils

type MatchCriteria string

// Each of these criteria will be used to instantiate a channel
// and a consumer goroutine will be listening for messages in the MQ
const (
	None   MatchCriteria = "none"
	Easy   MatchCriteria = "easy"
	Medium MatchCriteria = "medium"
	Hard   MatchCriteria = "hard"
)

var MatchCriterias = []MatchCriteria{
	None, Easy, Medium, Hard,
}

func ConstructResultChanIdentifier(str string) string {
	return "results/" + str
}
