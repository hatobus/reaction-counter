package reaction_counter

import (
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"

	"github.com/hatobus/reaction-counter/comment"
	"github.com/hatobus/reaction-counter/reaction"
)

var s *slack.Client

func init() {
	s = slack.New(os.Getenv("ACCESS_TOKEN"))
}

func ReactionCounter(w http.ResponseWriter, r *http.Request) {
	commands, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid commands", http.StatusBadRequest)
		return
	}

	if !commands.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
		log.Println("invalid validation token")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	switch commands.Command {
	case "/reactions":
		reacted, err := reaction.GetReactedUsers(s, commands.Text)
		if err != nil {
			log.Printf("invalid request: %v", err)
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		reactions := reaction.AggregateReactions(reacted)

		comment.PostReactionCountedMessage(commands.ChannelID, commands.Text, reactions)
	default:
	}
}
