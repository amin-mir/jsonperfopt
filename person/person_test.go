package person

import (
	"encoding/json"
	"reflect"
	"testing"
)

var fixture = []byte(`{
	"id":"d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
	"name":{
		"fullName":"Leonid Bugaev",
		"givenName":"Leonid",
		"familyName":"Bugaev"
	},
	"email":"leonsbox@gmail.com",
	"gender":"male",
	"location":"Saint Petersburg, Saint Petersburg, RU",
	"geo":{
		"city":"Saint Petersburg",
		"state":"Saint Petersburg",
		"country":"Russia",
		"lat":59.9342802,
		"lng":30.3350986
	},
	"bio":"Senior engineer at Granify.com",
	"site":"http://flickfaver.com",
	"avatar":"https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
	"employment":{
		"name":"www.latera.ru",
		"title":"Software Engineer",
		"domain":"gmail.com"
	},
	"facebook":{
		"handle":"leonid.bugaev"
	},
	"github":{
		"handle":"buger",
		"id":14009,
		"avatar":"https://avatars.githubusercontent.com/u/14009?v=3",
		"company":"Granify",
		"blog":"http://leonsbox.com",
		"followers":95,
		"following":10
	},
	"twitter":{
		"handle":"flickfaver",
		"id":77004410,
		"bio":null,
		"followers":2,
		"following":1,
		"statuses":5,
		"favorites":0,
		"location":"",
		"site":"http://flickfaver.com",
		"avatar":null
	},
	"linkedin":{
		"handle":"in/leonidbugaev"
	},
	"googleplus":{
		"handle":null
	},
	"angellist":{
		"handle":"leonid-bugaev",
		"id":61541,
		"bio":"Senior engineer at Granify.com",
		"blog":"http://buger.github.com",
		"site":"http://buger.github.com",
		"followers":41,
		"avatar":"https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
	},
	"klout":{
		"handle":null,
		"score":null
	},
	"foursquare":{
		"handle":null
	},
	"aboutme":{
		"handle":"leonid.bugaev",
		"bio":null,
		"avatar":null
	},
	"gravatar":{
		"handle":"buger",
		"urls":[

		],
		"avatar":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
		"avatars":[
			{
				"url":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
				"type":"thumbnail"
			}
		]
	},
	"fuzzy":false
}`)

func TestPersonUnmarshal(t *testing.T) {
	actualPerson := Person{}
	if err := json.Unmarshal(fixture, &actualPerson); err != nil {
		t.Fatalf("error unmarshalling person: %v", err)
	}

	expectedPerson := Person{
		ID: "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
		Name: Name{
			FullName:   "Leonid Bugaev",
			GivenName:  "Leonid",
			FamilyName: "Bugaev",
		},
		Email:    "leonsbox@gmail.com",
		Gender:   "male",
		Location: "Saint Petersburg, Saint Petersburg, RU",
		Geo: Geo{
			City:    "Saint Petersburg",
			State:   "Saint Petersburg",
			Country: "Russia",
			Lat:     59.9342802,
			Lng:     30.3350986,
		},
		Bio:    "Senior engineer at Granify.com",
		Site:   "http://flickfaver.com",
		Avatar: "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
		Employment: Employment{
			Name:   "www.latera.ru",
			Title:  "Software Engineer",
			Domain: "gmail.com",
		},
		Facebook: Facebook{
			Handle: "leonid.bugaev",
		},
		Github: Github{
			Handle:    "buger",
			ID:        14009,
			Avatar:    "https://avatars.githubusercontent.com/u/14009?v=3",
			Company:   "Granify",
			Blog:      "http://leonsbox.com",
			Followers: 95,
			Following: 10,
		},
		Twitter: Twitter{
			Handle:    "flickfaver",
			ID:        77004410,
			Bio:       "",
			Followers: 2,
			Following: 1,
			Statuses:  5,
			Favorites: 0,
			Location:  "",
			Site:      "http://flickfaver.com",
			Avatar:    "",
		},
		LinkedIn: LinkedIn{
			Handle: "in/leonidbugaev",
		},
		GooglePlus: GooglePlus{
			Handle: "",
		},
		AngelList: AngelList{
			Handle:    "leonid-bugaev",
			ID:        61541,
			Bio:       "Senior engineer at Granify.com",
			Blog:      "http://buger.github.com",
			Site:      "http://buger.github.com",
			Followers: 41,
			Avatar:    "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390",
		},
		Klout: Klout{
			Handle: "",
			Score:  0,
		},
		FourSquare: FourSquare{
			Handle: "",
		},
		AboutMe: AboutMe{
			Handle: "leonid.bugaev",
			Bio:    "",
			Avatar: "",
		},
		Gravatar: Gravatar{
			Handle: "buger",
			Urls:   []string{},
			Avatar: "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
			Avatars: []GravatarAvatar{
				{
					Url:  "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
					Type: "thumbnail",
				},
			},
		},
		Fuzzy: false,
	}

	if !reflect.DeepEqual(expectedPerson, actualPerson) {
		t.Fatalf("expected person %+v,\ngot %+v", expectedPerson, actualPerson)
	}
}
