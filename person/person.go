package person

type Person struct {
	ID         string     `json:"id"`
	Name       Name       `json:"name"`
	Email      string     `json:"email"`
	Gender     string     `json:"gender"`
	Location   string     `json:"location"`
	Geo        Geo        `json:"geo"`
	Bio        string     `json:"bio"`
	Site       string     `json:"site"`
	Avatar     string     `json:"avatar"`
	Employment Employment `json:"employment"`
	Facebook   Facebook   `json:"facebook"`
	Github     Github     `json:"github"`
	Twitter    Twitter    `json:"twitter"`
	LinkedIn   LinkedIn   `json:"linkedin"`
	GooglePlus GooglePlus `json:"googleplus"`
	AngelList  AngelList  `json:"angellist"`
	Klout      Klout      `json:"klout"`
	FourSquare FourSquare `json:"foursquare"`
	AboutMe    AboutMe    `json:"aboutme"`
	Gravatar   Gravatar   `json:"gravatar"`
	Fuzzy      bool       `json:"fuzzy"`
}

type Name struct {
	FullName   string `json:"fullName"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type Geo struct {
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type Employment struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Domain string `json:"domain"`
}

type Facebook struct {
	Handle string `json:"handle"`
}

type Github struct {
	Handle    string `json:"handle"`
	ID        int64  `json:"id"`
	Avatar    string `json:"avatar"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
}

type Twitter struct {
	Handle    string `json:"handle"`
	ID        int64  `json:"id"`
	Bio       string `json:"bio"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
	Statuses  int64  `json:"statuses"`
	Favorites int64  `json:"favorites"`
	Location  string `json:"location"`
	Site      string `json:"site"`
	Avatar    string `json:"avatar"`
}

type LinkedIn struct {
	Handle string `json:"handle"`
}

type GooglePlus struct {
	Handle string `json:"handle"`
}

type AngelList struct {
	Handle    string `json:"handle"`
	ID        int64  `json:"id"`
	Bio       string `json:"bio"`
	Blog      string `json:"blog"`
	Site      string `json:"site"`
	Followers int64  `json:"followers"`
	Avatar    string `json:"avatar"`
}

type Klout struct {
	Handle string `json:"handle"`
	Score  int64  `json:"score"`
}

type FourSquare struct {
	Handle string `json:"handle"`
}

type AboutMe struct {
	Handle string `json:"handle"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}

type Gravatar struct {
	Handle  string           `json:"handle"`
	Urls    []string         `json:"urls"`
	Avatar  string           `json:"avatar"`
	Avatars []GravatarAvatar `json:"avatars"`
}
type GravatarAvatar struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}
