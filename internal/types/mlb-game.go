package types

type MLBGame struct {
	UID       string `bson:"_id"`
	ID        int64  `bson:"id"`
	Home      string `bson:"home_team"`
	Away      string `bson:"away_team"`
	HomeScore int64  `bson:"home_score"`
	AwayScore int64  `bson:"away_score"`
	Active    bool   `bson:"active,omitempty"`
	Complete  bool   `bson:"complete,omitempty"`
	Canceled  bool   `bson:"canceled,omitempty"`
	GameTime  int64  `bson:"game_time"`
	Status    string `bson:"status,omitempty"`
}