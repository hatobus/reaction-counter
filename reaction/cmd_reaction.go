package reaction

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/slack-go/slack"
)

var reg = regexp.MustCompile(`\.*\.slack\.com`)

func GetReactedUsers(sc *slack.Client, cmdTxt string) (map[string][]string, error) {
	u, err := url.Parse(cmdTxt)
	if err != nil {
		return nil, err
	}

	// check url from slack post
	if !reg.MatchString(u.Host) {
		return nil, fmt.Errorf("invalid arguments please check slack url")
	}

	item, err := ItemRef(u)
	if err != nil {
		return nil, err
	}

	reactions, err := sc.GetReactions(*item, slack.GetReactionsParameters{Full: true})
	if err != nil {
		return nil, err
	}

	reactionData := map[string][]string{}

	for _, reaction := range reactions {
		users, err := getUserNameFromUserID(sc, reaction.Users)
		if err != nil {
			return nil, err
		}
		reactionData[reaction.Name] = users
	}

	return reactionData, nil
}
