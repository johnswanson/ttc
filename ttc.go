package ttc

type Seed int64
type Gap int64
type Token string
type Endpoint string
type PingTime int64
type PingTags string

type Ping struct {
	Timestamp PingTime `json:"ping/ts"`
	Tags      PingTags `json:"ping/tags"`
}

type Config struct {
	Seed int64 `json:"tagtime-seed"`
	Gap  int64 `json:"tagtime-gap"`
}

type PingAPI interface {
	Save(p *Ping) error
}

type Client interface {
	WebService()
}

type API struct {
	URL   string
	Token string
}
