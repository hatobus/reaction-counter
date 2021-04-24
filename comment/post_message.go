package comment

import (
	"fmt"

	"github.com/slack-go/slack"
)

func PostReactionCountedMessage(channelID, urlStr string, reactions map[string]string) {
	for emoji, users := range reactions {
		block := slack.NewTextBlockObject("plain_text", users, false, false)

		msg := &slack.Message{
			Msg: slack.Msg{Channel: channelID,
				Text:   fmt.Sprintf("%v に %v のリアクションをつけた方は以下のようになっています", urlStr, emoji),
				Blocks: slack.Blocks{BlockSet: []slack.Block{block}},
			},
		}

		slack.NewMessageItem(channelID, msg)
	}
}
