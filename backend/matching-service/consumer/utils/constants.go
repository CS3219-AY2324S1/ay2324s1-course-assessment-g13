package utils

type MatchCriteria string

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
