package comment

import (
	"fmt"

	"github.com/slack-go/slack"
)

func PostReactionCountedMessage(sc *slack.Client, channelID, urlStr string, reactions map[string]string) error {
	for emoji, users := range reactions {
		textMsg := slack.MsgOptionText(fmt.Sprintf("%v に :%v: のリアクションをつけた方は以下のようになっています", urlStr, emoji), true)
		_, _, err := sc.PostMessage(channelID, textMsg)
		if err != nil {
			return err
		}

		reactions := slack.MsgOptionBlocks(
			slack.NewSectionBlock(slack.NewTextBlockObject("plain_text", users, false, false), nil, nil))
		_, _, err = sc.PostMessage(channelID, reactions)
		if err != nil {
			return err
		}
	}
	return nil
}
