package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/JuandeJuni/discordhooks"
	"github.com/Pallinder/go-randomdata"
)

const waffleUrl = "https://www.wafflehouse.com/regulars-club/"

const wafflePost = "https://www.wafflehouse.com/doprocess"

const webhook = "" //enter ur webhook link or youll get a goofy error lol

var fn = randomdata.FirstName(randomdata.Male)
var ln = randomdata.LastName()
var email = ""
var zip = randomdata.PostalCode("SE")

func main() {
	fmt.Println("Initializing... ")

	fmt.Print("Enter Email: ")
	fmt.Scan(&email)

	fmt.Println("Submitting Form...")

	data := url.Values{
		"fn":        {fn},
		"ln":        {ln},
		"em":        {email},
		"emcon":     {email},
		"bmon":      {"11"},
		"bday":      {"16"},
		"zip":       {zip},
		"theaction": {"regulars-club-two"},
	}
	req, err := http.PostForm(wafflePost, data) //posts waffle house necessary data
	if err != nil {
		panic(err)
	}
	if req.StatusCode == 200 {
		fmt.Println("Successfully Generated Free Hasbrowns!")
	} else {
		fmt.Println("Error Generating Hashbrowns")
	}
	embed := discordhooks.Embed{
		Title:       "Waffle House",
		Description: "Successfully Generated Free Hashbrowns!",
		Color:       0x008FFF,
		Thumbnail: discordhooks.Thumbnail{
			Url: "https://static.wikia.nocookie.net/logopedia/images/e/e0/Wreujwv2.png/revision/latest?cb=20151126181627",
		},
		Fields: []discordhooks.Field{
			discordhooks.Field{
				Name:  "Email",
				Value: email,
			},
		},
	}

	discordhooks.SendEmbed(webhook, embed) //this send the webhook

	// read response body
	/*body, error := ioutil.ReadAll(req.Body)
	if error != nil {
		fmt.Println(error)
	}

	req.Body.Close()

	fmt.Println(string(body))
	//ngl this shit is not necessary just needed to see if html body response was 1 */
}
