package reaction

import "github.com/slack-go/slack"

func getUserNameFromUserID(sc *slack.Client, users []string) ([]string, error) {
	username := make([]string, len(users))

	for i, uid := range users {
		user, err := sc.GetUserInfo(uid)
		if err != nil {
			return nil, err
		}
		username[i] = user.Name
	}

	return username, nil
}
