package reaction

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"

	"github.com/slack-go/slack"
)

func ItemRef(u *url.URL) (*slack.ItemRef, error) {
	queries := u.Query()
	params := strings.Split(u.Path, "/")

	item := new(slack.ItemRef)

	if queries.Get("thread_ts") != "" {
		item.Timestamp = queries.Get("thread_ts")
	} else {
		pf, err := strconv.ParseFloat(params[len(params)-1], 64)
		if err != nil {
			return nil, err
		}
		item.Timestamp = fmt.Sprintf("%16f", pf / math.Pow(10, 6))
	}

	item.Channel = params[len(params)-2]

	return item, nil
}
