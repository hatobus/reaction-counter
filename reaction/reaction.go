package reaction

import "strings"

func AggregateReactions(reaction map[string][]string) map[string]string {
	reactions := map[string]string{}

	for emoji, users := range reaction {
		userStr := strings.Join(users[:], "\n- ")
		reactions[emoji] = userStr
	}

	return reactions
}
