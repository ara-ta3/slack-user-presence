package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	mkr "github.com/mackerelio/mackerel-client-go"
)

type response struct {
	OK      bool   `json:"ok"`
	Members []user `json:"members"`
}

type user struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Deleted           bool        `json:"deleted"`
	Color             string      `json:"color"`
	RealName          string      `json:"real_name"`
	TZ                string      `json:"tz,omitempty"`
	TZLabel           string      `json:"tz_label"`
	TZOffset          int         `json:"tz_offset"`
	Profile           interface{} `json:"profile"`
	IsBot             bool        `json:"is_bot"`
	IsAdmin           bool        `json:"is_admin"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	Has2FA            bool        `json:"has_2fa"`
	HasFiles          bool        `json:"has_files"`
	Presence          string      `json:"presence"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	client := mkr.NewClient(os.Getenv("MACKEREL_TOKEN"))
	token := os.Getenv("SLACK_TOKEN")
	resp, err := http.Get(
		fmt.Sprintf("https://slack.com/api/users.list?token=%s&presence=1", token),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	r := response{}
	json.Unmarshal(body, &r)
	c := 0
	for _, u := range r.Members {
		if u.Presence == "active" && !u.IsBot {
			c++
		}
	}
	err = client.PostServiceMetricValues("Dark", []*mkr.MetricValue{
		&mkr.MetricValue{
			Name:  "dark.slack.active_rate",
			Time:  time.Now().Unix(),
			Value: c,
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
