package utils

type MatchCriteria string

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
