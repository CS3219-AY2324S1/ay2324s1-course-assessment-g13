package utils

type MatchCriteria string

const (
	Easy   MatchCriteria = "easy"
	Medium MatchCriteria = "medium"
	Hard   MatchCriteria = "hard"
)

var MatchCriterias = []MatchCriteria{
	Easy, Medium, Hard,
}

func ConstructResultChanIdentifier(str string) string {
	return "results/" + str
}
