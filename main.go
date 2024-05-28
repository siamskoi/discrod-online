package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

var status = "online"
var customStatus = ""
var userToken = ""

type userData struct {
	Username      string
	Discriminator string
	Id            string
}

type Auth struct {
	Op int
	D  authDetails
	S  interface{}
	T  interface{}
}

type authDetails struct {
	Token      string
	Properties Properties
	Presence   Presence
}

type Properties struct {
	OS      string
	Browser string
	Device  string
}

type Presence struct {
	Status string
	Afk    bool
}

type CStatus struct {
	Op int
	D  cstatusDetails
}

type cstatusDetails struct {
	Since      int
	Activities []Activity
	Status     string
	Afk        bool
}

type Activity struct {
	Type  int
	State string
	Name  string
	Id    string
}

func main() {
	userToken = os.Getenv("DISCORD_USER_TOKEN")
	if len(userToken) == 0 {
		log.Panic("Not found env variable DISCORD_USER_TOKEN")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://discordapp.com/api/v9/users/@me", nil)
	req.Header.Set("Authorization", userToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("[ERROR] Your token might be invalid. Please check it again.")
	}

	var receivedData userData
	err = json.NewDecoder(resp.Body).Decode(&receivedData)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Logged in as %s#%s (%s).\n", receivedData.Username, receivedData.Discriminator, receivedData.Id)

	for {
		runOnlineWS()
		time.Sleep(50 * time.Second)
	}
}

func runOnlineWS() {
	log.Println("I'am online:)")
	c, _, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=9&encoding=json", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal(msg, &result)
	heartbeat := int(result["d"].(map[string]interface{})["heartbeat_interval"].(float64))

	auth := Auth{
		Op: 2,
		D: authDetails{
			Token: userToken,
			Properties: Properties{
				OS:      "Windows 10",
				Browser: "Google Chrome",
				Device:  "Windows",
			},
			Presence: Presence{
				Status: status,
				Afk:    false,
			},
		},
	}

	authJson, _ := json.Marshal(auth)
	err = c.WriteMessage(websocket.TextMessage, authJson)
	if err != nil {
		log.Println("write:", err)
		return
	}

	cstatus := CStatus{
		Op: 3,
		D: cstatusDetails{
			Since: 0,
			Activities: []Activity{
				{
					Type:  4,
					State: customStatus,
					Name:  "Custom Status",
					Id:    "custom",
				},
			},
			Status: status,
			Afk:    false,
		},
	}

	cstatusJson, _ := json.Marshal(cstatus)
	err = c.WriteMessage(websocket.TextMessage, cstatusJson)
	if err != nil {
		log.Println("write:", err)
		return
	}

	online := struct {
		Op int
		D  string
	}{
		Op: 1,
		D:  "None",
	}

	onlineJson, _ := json.Marshal(online)
	err = c.WriteMessage(websocket.TextMessage, onlineJson)
	if err != nil {
		log.Println("write:", err)
		return
	}

	time.Sleep(time.Duration(heartbeat) * time.Millisecond)
}
