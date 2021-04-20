package reaction_counter

import (
	"log"
	"net/http"
	"os"

	"github.com/hatobus/reaction-counter/reaction"
	"github.com/slack-go/slack"
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

		log.Println(reacted)
	default:
	}
}
