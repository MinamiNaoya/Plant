package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Config struct {
	TestMail struct {
		ReplyTo string `yaml:"reply_to"`
		To string `yaml:"to"`
		
	} `yaml:"testMail"`
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "mail/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}


func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func readYAML() (config Config){
	data, err := os.ReadFile("mail/config.yml")
	if err != nil {
		fmt.Println("Cannot Open File!")
		return 
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Cannot parse YAML data:", err)
		return 
	}
	return config
}

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("mail/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}
	fmt.Println("Created Gmail service", srv)
	var conf Config
	conf = readYAML()
	
	msgStr := fmt.Sprintf("From: 'me'\r\n" +
		"Reply-To: %v\r\n" + //送信元
		"To: %v\r\n" + //送信先
		"Subject:あいさつ\r\n" +
		"\r\n" + "hogeですhogehogeさんお元気ですか", conf.TestMail.ReplyTo, conf.TestMail.To)
	reader := strings.NewReader(msgStr)
	transformer := japanese.ISO2022JP.NewEncoder()
	msgISO2022JP, err := io.ReadAll(transform.NewReader(reader, transformer))
	if err != nil {
		log.Fatalf("Unable to convert to ISO2022JP: %v", err)
	}
	msg := []byte(msgISO2022JP)
	message := gmail.Message{}
	message.Raw = base64.StdEncoding.EncodeToString(msg)
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
