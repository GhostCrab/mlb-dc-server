package types

type MLBGame struct {
	UID       string `json:"uid" bson:"_id"`
	ID        int64  `json:"id" bson:"id"`
	Home      string `json:"home_team" bson:"home_team"`
	Away      string `json:"away_team" bson:"away_team"`
	HomeScore int64  `json:"home_score" bson:"home_score"`
	AwayScore int64  `json:"away_score" bson:"away_score"`
	Active    bool   `json:"active,omitempty" bson:"active,omitempty"`
	Complete  bool   `json:"complete,omitempty" bson:"complete,omitempty"`
	Canceled  bool   `json:"canceled,omitempty" bson:"canceled,omitempty"`
	GameTime  int64  `json:"game_time" bson:"game_time"`
	Status    string `json:"status,omitempty" bson:"status,omitempty"`
}