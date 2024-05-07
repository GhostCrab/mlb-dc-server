package types

type MLBAPIResponse struct {
	Copyright            string           `json:"copyright"`
	TotalItems           int64            `json:"totalItems"`
	TotalEvents          int64            `json:"totalEvents"`
	TotalGames           int64            `json:"totalGames"`
	TotalGamesInProgress int64            `json:"totalGamesInProgress"`
	Dates                []MLBAPIDate     `json:"dates"`
	ReqHeaders           MLBAPIReqHeaders `json:"reqHeaders"`
	Headers              MLBAPIHeaders    `json:"headers"`
	Status               int64            `json:"status"`
}

type MLBAPIDate struct {
	Date                 string        `json:"date"`
	TotalItems           int64         `json:"totalItems"`
	TotalEvents          int64         `json:"totalEvents"`
	TotalGames           int64         `json:"totalGames"`
	TotalGamesInProgress int64         `json:"totalGamesInProgress"`
	Games                []MLBAPIGame  `json:"games"`
	Events               []interface{} `json:"events"`
}

type MLBAPIGame struct {
	SortIndex1             int64                        `json:"sortIndex1"`
	SortIndex2             int64                        `json:"sortIndex2"`
	SortIndex3             int64                        `json:"sortIndex3"`
	SortIndex4             int64                        `json:"sortIndex4"`
	SortIndex5             int64                        `json:"sortIndex5"`
	Teams                  *MLBAPIGameTeams             `json:"teams,omitempty"`
	GamePk                 int64                        `json:"gamePk"`
	GameGUID               string                       `json:"gameGuid"`
	Link                   string                       `json:"link"`
	GameType               MLBAPIGameTypeEnum           `json:"gameType"`
	Season                 string                       `json:"season"`
	GameDate               string                       `json:"gameDate"`
	OfficialDate           string                       `json:"officialDate"`
	Status                 MLBAPIStatus                 `json:"status"`
	Linescore              *MLBAPILinescore             `json:"linescore,omitempty"`
	Decisions              *MLBAPIDecisions             `json:"decisions,omitempty"`
	Venue                  MLBAPIVenue                  `json:"venue"`
	Broadcasts             []MLBAPIBroadcast            `json:"broadcasts,omitempty"`
	Content                MLBAPIContent                `json:"content"`
	SeriesStatus           MLBAPISeriesStatus           `json:"seriesStatus"`
	IsTie                  *bool                        `json:"isTie,omitempty"`
	XrefIDS                []MLBAPIXrefID               `json:"xrefIds"`
	GameNumber             int64                        `json:"gameNumber"`
	PublicFacing           bool                         `json:"publicFacing"`
	DoubleHeader           MLBAPIDoubleHeader           `json:"doubleHeader"`
	GamedayType            MLBAPIGamedayType            `json:"gamedayType"`
	Tiebreaker             MLBAPIDoubleHeader           `json:"tiebreaker"`
	CalendarEventID        string                       `json:"calendarEventID"`
	SeasonDisplay          string                       `json:"seasonDisplay"`
	DayNight               MLBAPIDayNight               `json:"dayNight"`
	Description            *string                      `json:"description,omitempty"`
	ScheduledInnings       int64                        `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool                         `json:"reverseHomeAwayStatus"`
	InningBreakLength      int64                        `json:"inningBreakLength"`
	GamesInSeries          int64                        `json:"gamesInSeries"`
	SeriesGameNumber       int64                        `json:"seriesGameNumber"`
	SeriesDescription      MLBAPISeriesDescriptionEnum  `json:"seriesDescription"`
	Review                 *MLBAPIReview                `json:"review,omitempty"`
	Flags                  *MLBAPIFlags                 `json:"flags,omitempty"`
	HomeRuns               []MLBAPIHomeRunElement       `json:"homeRuns,omitempty"`
	RecordSource           MLBAPIGameTypeEnum           `json:"recordSource"`
	IfNecessary            MLBAPIDoubleHeader           `json:"ifNecessary"`
	IfNecessaryDescription MLBAPIIfNecessaryDescription `json:"ifNecessaryDescription"`
	GameUtils              map[string]bool              `json:"gameUtils"`
	RescheduleDate         *string                      `json:"rescheduleDate,omitempty"`
	RescheduleGameDate     *string                      `json:"rescheduleGameDate,omitempty"`
	RescheduledFrom        *string                      `json:"rescheduledFrom,omitempty"`
	RescheduledFromDate    *string                      `json:"rescheduledFromDate,omitempty"`
	Tickets                []MLBAPITicket               `json:"tickets,omitempty"`
}

type MLBAPIBroadcast struct {
	ID                    int64                   `json:"id"`
	Name                  string                  `json:"name"`
	Type                  MLBAPIBroadcastType     `json:"type"`
	Language              MLBAPIBroadcastLanguage `json:"language"`
	IsNational            bool                    `json:"isNational"`
	CallSign              string                  `json:"callSign"`
	VideoResolution       *MLBAPIVideoResolution  `json:"videoResolution,omitempty"`
	MediaState            MLBAPIMediaState        `json:"mediaState"`
	BroadcastDate         string                  `json:"broadcastDate"`
	MediaID               string                  `json:"mediaId"`
	GameDateBroadcastGUID string                  `json:"gameDateBroadcastGuid"`
	HomeAway              MLBAPIHomeAway          `json:"homeAway"`
	FreeGame              bool                    `json:"freeGame"`
	AvailableForStreaming bool                    `json:"availableForStreaming"`
	PostGameShow          bool                    `json:"postGameShow"`
	MvpdAuthRequired      bool                    `json:"mvpdAuthRequired"`
	PreGameShow           *string                 `json:"preGameShow,omitempty"`
	Availability          *MLBAPIAvailability     `json:"availability,omitempty"`
	SourceComment         *string                 `json:"sourceComment,omitempty"`
}

type MLBAPIAvailability struct {
	AvailabilityID   int64                  `json:"availabilityId"`
	AvailabilityCode MLBAPIAvailabilityCode `json:"availabilityCode"`
	AvailabilityText MLBAPIAvailabilityText `json:"availabilityText"`
}

type MLBAPIMediaState struct {
	MediaStateID   int64                    `json:"mediaStateId"`
	MediaStateCode MLBAPIMediaStateCodeEnum `json:"mediaStateCode"`
	MediaStateText MLBAPIMediaStateText     `json:"mediaStateText"`
}

type MLBAPIVideoResolution struct {
	Code            MLBAPICode            `json:"code"`
	ResolutionShort MLBAPIResolutionShort `json:"resolutionShort"`
	ResolutionFull  MLBAPIResolutionFull  `json:"resolutionFull"`
}

type MLBAPIContent struct {
	Link       string           `json:"link"`
	Editorial  *MLBAPIEditorial `json:"editorial,omitempty"`
	Media      MLBAPIMedia      `json:"media"`
	Highlights *MLBAPIEditorial `json:"highlights,omitempty"`
	Summary    *MLBAPISummary   `json:"summary,omitempty"`
	GameNotes  *MLBAPIEditorial `json:"gameNotes,omitempty"`
}

type MLBAPIEditorial struct {
}

type MLBAPIMedia struct {
	Epg           []MLBAPIEpg          `json:"epg,omitempty"`
	FeaturedMedia *MLBAPIFeaturedMedia `json:"featuredMedia,omitempty"`
	FreeGame      *bool                `json:"freeGame,omitempty"`
	EnhancedGame  *bool                `json:"enhancedGame,omitempty"`
}

type MLBAPIEpg struct {
	Title MLBAPITitle  `json:"title"`
	Items []MLBAPIItem `json:"items"`
}

type MLBAPIItem struct {
	CallLetters       *string                   `json:"callLetters,omitempty"`
	EspnAuthRequired  *bool                     `json:"espnAuthRequired,omitempty"`
	TbsAuthRequired   *bool                     `json:"tbsAuthRequired,omitempty"`
	Espn2AuthRequired *bool                     `json:"espn2AuthRequired,omitempty"`
	GameDate          *string                   `json:"gameDate,omitempty"`
	ContentID         *string                   `json:"contentId,omitempty"`
	Fs1AuthRequired   *bool                     `json:"fs1AuthRequired,omitempty"`
	MediaID           *string                   `json:"mediaId,omitempty"`
	MediaFeedType     *MLBAPIMediaFeedType      `json:"mediaFeedType,omitempty"`
	MlbnAuthRequired  *bool                     `json:"mlbnAuthRequired,omitempty"`
	FoxAuthRequired   *bool                     `json:"foxAuthRequired,omitempty"`
	MediaFeedSubType  *string                   `json:"mediaFeedSubType,omitempty"`
	FreeGame          *bool                     `json:"freeGame,omitempty"`
	ID                int64                     `json:"id"`
	MediaState        *MLBAPIMediaStateCodeEnum `json:"mediaState,omitempty"`
	AbcAuthRequired   *bool                     `json:"abcAuthRequired,omitempty"`
	RenditionName     *MLBAPIRenditionNameEnum  `json:"renditionName,omitempty"`
	Description       *string                   `json:"description,omitempty"`
	Language          *MLBAPIItemLanguage       `json:"language,omitempty"`
	Type              *MLBAPIItemType           `json:"type,omitempty"`
	Appletv           *MLBAPIAppletv            `json:"appletv,omitempty"`
}

type MLBAPIAppletv struct {
	Fgow    bool   `json:"fgow"`
	VideoID string `json:"videoId"`
	PageID  string `json:"pageId"`
}

type MLBAPIFeaturedMedia struct {
	ID string `json:"id"`
}

type MLBAPISummary struct {
	HasPreviewArticle  bool `json:"hasPreviewArticle"`
	HasRecapArticle    bool `json:"hasRecapArticle"`
	HasWrapArticle     bool `json:"hasWrapArticle"`
	HasHighlightsVideo bool `json:"hasHighlightsVideo"`
}

type MLBAPIDecisions struct {
	Winner MLBAPILoser  `json:"winner"`
	Loser  MLBAPILoser  `json:"loser"`
	Save   *MLBAPILoser `json:"save,omitempty"`
}

type MLBAPILoser struct {
	ID                 int64                     `json:"id"`
	FullName           string                    `json:"fullName"`
	Link               string                    `json:"link"`
	FirstName          string                    `json:"firstName"`
	LastName           string                    `json:"lastName"`
	PrimaryNumber      *string                   `json:"primaryNumber,omitempty"`
	BirthDate          string                    `json:"birthDate"`
	CurrentAge         int64                     `json:"currentAge"`
	BirthCity          *string                   `json:"birthCity,omitempty"`
	BirthStateProvince *MLBAPIBirthStateProvince `json:"birthStateProvince,omitempty"`
	BirthCountry       MLBAPICountry             `json:"birthCountry"`
	Height             MLBAPIHeight              `json:"height"`
	Weight             int64                     `json:"weight"`
	Active             bool                      `json:"active"`
	PrimaryPosition    MLBAPIPrimaryPosition     `json:"primaryPosition"`
	UseName            string                    `json:"useName"`
	UseLastName        string                    `json:"useLastName"`
	MiddleName         *string                   `json:"middleName,omitempty"`
	BoxscoreName       string                    `json:"boxscoreName"`
	Gender             MLBAPIGender              `json:"gender"`
	IsPlayer           bool                      `json:"isPlayer"`
	IsVerified         bool                      `json:"isVerified"`
	DraftYear          *int64                    `json:"draftYear,omitempty"`
	Stats              []MLBAPIStat              `json:"stats,omitempty"`
	MlbDebutDate       string                    `json:"mlbDebutDate"`
	BatSide            MLBAPIBatSide             `json:"batSide"`
	PitchHand          MLBAPIBatSide             `json:"pitchHand"`
	NameFirstLast      string                    `json:"nameFirstLast"`
	NameSlug           string                    `json:"nameSlug"`
	FirstLastName      string                    `json:"firstLastName"`
	LastFirstName      string                    `json:"lastFirstName"`
	LastInitName       string                    `json:"lastInitName"`
	InitLastName       string                    `json:"initLastName"`
	FullFMLName        string                    `json:"fullFMLName"`
	FullLFMName        string                    `json:"fullLFMName"`
	StrikeZoneTop      float64                   `json:"strikeZoneTop"`
	StrikeZoneBottom   float64                   `json:"strikeZoneBottom"`
	NickName           *string                   `json:"nickName,omitempty"`
	Pronunciation      *string                   `json:"pronunciation,omitempty"`
	NameMatrilineal    *string                   `json:"nameMatrilineal,omitempty"`
	NameTitle          *MLBAPIName               `json:"nameTitle,omitempty"`
	NameSuffix         *MLBAPIName               `json:"nameSuffix,omitempty"`
}

type MLBAPIBatSide struct {
	Code        MLBAPIGameTypeEnum       `json:"code"`
	Description MLBAPIBatSideDescription `json:"description"`
}

type MLBAPIPrimaryPosition struct {
	Code         string                    `json:"code"`
	Name         MLBAPIPrimaryPositionName `json:"name"`
	Type         MLBAPIPrimaryPositionType `json:"type"`
	Abbreviation MLBAPIGamedayType         `json:"abbreviation"`
}

type MLBAPIStat struct {
	Type       MLBAPIGroup   `json:"type"`
	Group      MLBAPIGroup   `json:"group"`
	Exemptions []interface{} `json:"exemptions"`
	Stats      MLBAPIStats   `json:"stats"`
}

type MLBAPIGroup struct {
	DisplayName MLBAPIDisplayName `json:"displayName"`
}

type MLBAPIStats struct {
	Note                   *string                     `json:"note,omitempty"`
	Summary                *string                     `json:"summary,omitempty"`
	GamesPlayed            *int64                      `json:"gamesPlayed,omitempty"`
	GamesStarted           *int64                      `json:"gamesStarted,omitempty"`
	GroundOuts             *int64                      `json:"groundOuts,omitempty"`
	AirOuts                *int64                      `json:"airOuts,omitempty"`
	Runs                   *int64                      `json:"runs,omitempty"`
	Doubles                *int64                      `json:"doubles,omitempty"`
	Triples                *int64                      `json:"triples,omitempty"`
	HomeRuns               *int64                      `json:"homeRuns,omitempty"`
	StrikeOuts             *int64                      `json:"strikeOuts,omitempty"`
	BaseOnBalls            *int64                      `json:"baseOnBalls,omitempty"`
	IntentionalWalks       *int64                      `json:"intentionalWalks,omitempty"`
	Hits                   *int64                      `json:"hits,omitempty"`
	HitByPitch             *int64                      `json:"hitByPitch,omitempty"`
	AtBats                 *int64                      `json:"atBats,omitempty"`
	CaughtStealing         *int64                      `json:"caughtStealing,omitempty"`
	StolenBases            *int64                      `json:"stolenBases,omitempty"`
	StolenBasePercentage   *MLBAPIPercentage           `json:"stolenBasePercentage,omitempty"`
	NumberOfPitches        *int64                      `json:"numberOfPitches,omitempty"`
	Era                    *string                     `json:"era,omitempty"`
	InningsPitched         *string                     `json:"inningsPitched,omitempty"`
	WINS                   *int64                      `json:"wins,omitempty"`
	Losses                 *int64                      `json:"losses,omitempty"`
	Saves                  *int64                      `json:"saves,omitempty"`
	SaveOpportunities      *int64                      `json:"saveOpportunities,omitempty"`
	Holds                  *int64                      `json:"holds,omitempty"`
	BlownSaves             *int64                      `json:"blownSaves,omitempty"`
	EarnedRuns             *int64                      `json:"earnedRuns,omitempty"`
	BattersFaced           *int64                      `json:"battersFaced,omitempty"`
	Outs                   *int64                      `json:"outs,omitempty"`
	GamesPitched           *int64                      `json:"gamesPitched,omitempty"`
	CompleteGames          *int64                      `json:"completeGames,omitempty"`
	Shutouts               *int64                      `json:"shutouts,omitempty"`
	PitchesThrown          *int64                      `json:"pitchesThrown,omitempty"`
	Balls                  *int64                      `json:"balls,omitempty"`
	Strikes                *int64                      `json:"strikes,omitempty"`
	StrikePercentage       *MLBAPIStrikePercentageEnum `json:"strikePercentage,omitempty"`
	HitBatsmen             *int64                      `json:"hitBatsmen,omitempty"`
	Balks                  *int64                      `json:"balks,omitempty"`
	WildPitches            *int64                      `json:"wildPitches,omitempty"`
	Pickoffs               *int64                      `json:"pickoffs,omitempty"`
	Rbi                    *int64                      `json:"rbi,omitempty"`
	GamesFinished          *int64                      `json:"gamesFinished,omitempty"`
	RunsScoredPer9         *string                     `json:"runsScoredPer9,omitempty"`
	HomeRunsPer9           *string                     `json:"homeRunsPer9,omitempty"`
	InheritedRunners       *int64                      `json:"inheritedRunners,omitempty"`
	InheritedRunnersScored *int64                      `json:"inheritedRunnersScored,omitempty"`
	CatchersInterference   *int64                      `json:"catchersInterference,omitempty"`
	SacBunts               *int64                      `json:"sacBunts,omitempty"`
	SacFlies               *int64                      `json:"sacFlies,omitempty"`
	PassedBall             *int64                      `json:"passedBall,omitempty"`
	FlyOuts                *int64                      `json:"flyOuts,omitempty"`
	GroundIntoDoublePlay   *int64                      `json:"groundIntoDoublePlay,omitempty"`
	GroundIntoTriplePlay   *int64                      `json:"groundIntoTriplePlay,omitempty"`
	PlateAppearances       *int64                      `json:"plateAppearances,omitempty"`
	LeftOnBase             *int64                      `json:"leftOnBase,omitempty"`
	AtBatsPerHomeRun       *string                     `json:"atBatsPerHomeRun,omitempty"`
	Avg                    *string                     `json:"avg,omitempty"`
	Obp                    *string                     `json:"obp,omitempty"`
	Whip                   *string                     `json:"whip,omitempty"`
	GroundOutsToAirouts    *string                     `json:"groundOutsToAirouts,omitempty"`
	WinPercentage          *MLBAPIPercentage           `json:"winPercentage,omitempty"`
	PitchesPerInning       *string                     `json:"pitchesPerInning,omitempty"`
	StrikeoutWalkRatio     *string                     `json:"strikeoutWalkRatio,omitempty"`
	StrikeoutsPer9Inn      *string                     `json:"strikeoutsPer9Inn,omitempty"`
	WalksPer9Inn           *string                     `json:"walksPer9Inn,omitempty"`
	HitsPer9Inn            *string                     `json:"hitsPer9Inn,omitempty"`
	Slg                    *string                     `json:"slg,omitempty"`
	Ops                    *string                     `json:"ops,omitempty"`
	Babip                  *string                     `json:"babip,omitempty"`
	TotalBases             *int64                      `json:"totalBases,omitempty"`
}

type MLBAPIFlags struct {
	NoHitter            bool `json:"noHitter"`
	PerfectGame         bool `json:"perfectGame"`
	AwayTeamNoHitter    bool `json:"awayTeamNoHitter"`
	AwayTeamPerfectGame bool `json:"awayTeamPerfectGame"`
	HomeTeamNoHitter    bool `json:"homeTeamNoHitter"`
	HomeTeamPerfectGame bool `json:"homeTeamPerfectGame"`
}

type MLBAPIHomeRunElement struct {
	Result      MLBAPIResult  `json:"result"`
	About       MLBAPIAbout   `json:"about"`
	Count       MLBAPICount   `json:"count"`
	Matchup     MLBAPIMatchup `json:"matchup"`
	PitchIndex  []interface{} `json:"pitchIndex"`
	ActionIndex []interface{} `json:"actionIndex"`
	RunnerIndex []interface{} `json:"runnerIndex"`
	Runners     []interface{} `json:"runners"`
	PlayEvents  []interface{} `json:"playEvents"`
}

type MLBAPIAbout struct {
	HalfInning MLBAPIHalfInning `json:"halfInning"`
	Inning     int64            `json:"inning"`
}

type MLBAPICount struct {
	Balls   int64 `json:"balls"`
	Strikes int64 `json:"strikes"`
	Outs    int64 `json:"outs"`
}

type MLBAPIMatchup struct {
	Batter              MLBAPILoser     `json:"batter"`
	Pitcher             MLBAPICatcher   `json:"pitcher"`
	BatterHotColdZones  []interface{}   `json:"batterHotColdZones"`
	PitcherHotColdZones []interface{}   `json:"pitcherHotColdZones"`
	Splits              MLBAPIEditorial `json:"splits"`
}

type MLBAPICatcher struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

type MLBAPIResult struct {
	Type        MLBAPIResultType `json:"type"`
	Event       MLBAPIEvent      `json:"event"`
	Description string           `json:"description"`
	Rbi         int64            `json:"rbi"`
	AwayScore   int64            `json:"awayScore"`
	HomeScore   int64            `json:"homeScore"`
}

type MLBAPILinescore struct {
	CurrentInning        *int64                `json:"currentInning,omitempty"`
	CurrentInningOrdinal *string               `json:"currentInningOrdinal,omitempty"`
	InningState          *string               `json:"inningState,omitempty"`
	InningHalf           *string               `json:"inningHalf,omitempty"`
	IsTopInning          *bool                 `json:"isTopInning,omitempty"`
	ScheduledInnings     *int64                `json:"scheduledInnings,omitempty"`
	Innings              []MLBAPIInning        `json:"innings,omitempty"`
	Teams                *MLBAPILinescoreTeams `json:"teams,omitempty"`
	Defense              MLBAPIDefense         `json:"defense"`
	Offense              MLBAPIOffense         `json:"offense"`
	Balls                *int64                `json:"balls,omitempty"`
	Strikes              *int64                `json:"strikes,omitempty"`
	Outs                 *int64                `json:"outs,omitempty"`
	Note                 *MLBAPINote           `json:"note,omitempty"`
}

type MLBAPIDefense struct {
	Pitcher      *MLBAPILoser   `json:"pitcher,omitempty"`
	Catcher      *MLBAPICatcher `json:"catcher,omitempty"`
	First        *MLBAPICatcher `json:"first,omitempty"`
	Second       *MLBAPICatcher `json:"second,omitempty"`
	Third        *MLBAPICatcher `json:"third,omitempty"`
	Shortstop    *MLBAPICatcher `json:"shortstop,omitempty"`
	Left         *MLBAPICatcher `json:"left,omitempty"`
	Center       *MLBAPICatcher `json:"center,omitempty"`
	Right        *MLBAPICatcher `json:"right,omitempty"`
	BattingOrder *int64         `json:"battingOrder,omitempty"`
}

type MLBAPIInning struct {
	Num        int64            `json:"num"`
	OrdinalNum string           `json:"ordinalNum"`
	Home       MLBAPIInningAway `json:"home"`
	Away       MLBAPIInningAway `json:"away"`
}

type MLBAPIInningAway struct {
	Runs       *int64 `json:"runs,omitempty"`
	Hits       *int64 `json:"hits,omitempty"`
	Errors     *int64 `json:"errors,omitempty"`
	LeftOnBase *int64 `json:"leftOnBase,omitempty"`
	IsWinner   *bool  `json:"isWinner,omitempty"`
}

type MLBAPIOffense struct {
	Batter       *MLBAPILoser             `json:"batter,omitempty"`
	OnDeck       *MLBAPILoser             `json:"onDeck,omitempty"`
	InHole       *MLBAPILoser             `json:"inHole,omitempty"`
	BattingOrder *int64                   `json:"battingOrder,omitempty"`
	Team         *MLBAPISpringLeagueClass `json:"team,omitempty"`
	Second       *MLBAPILoser             `json:"second,omitempty"`
	Third        *MLBAPILoser             `json:"third,omitempty"`
	First        *MLBAPILoser             `json:"first,omitempty"`
}

type MLBAPISpringLeagueClass struct {
	ID           int64                   `json:"id"`
	Name         *MLBAPITeamVenueName    `json:"name,omitempty"`
	Link         MLBAPITeamLink          `json:"link"`
	Abbreviation *MLBAPITeamAbbreviation `json:"abbreviation,omitempty"`
}

type MLBAPILinescoreTeams struct {
	Home MLBAPIInningAway `json:"home"`
	Away MLBAPIInningAway `json:"away"`
}

type MLBAPIReview struct {
	HasChallenges bool             `json:"hasChallenges"`
	Away          MLBAPIReviewAway `json:"away"`
	Home          MLBAPIReviewAway `json:"home"`
}

type MLBAPIReviewAway struct {
	Used      int64 `json:"used"`
	Remaining int64 `json:"remaining"`
}

type MLBAPISeriesStatus struct {
	GameNumber       int64                          `json:"gameNumber"`
	TotalGames       int64                          `json:"totalGames"`
	IsTied           bool                           `json:"isTied"`
	IsOver           bool                           `json:"isOver"`
	WINS             int64                          `json:"wins"`
	Losses           int64                          `json:"losses"`
	WinningTeam      *MLBAPIIngTeam                 `json:"winningTeam,omitempty"`
	LosingTeam       *MLBAPIIngTeam                 `json:"losingTeam,omitempty"`
	Description      MLBAPISeriesDescriptionEnum    `json:"description"`
	ShortDescription MLBAPIShort                    `json:"shortDescription"`
	Result           *string                        `json:"result,omitempty"`
	ShortName        MLBAPIShort                    `json:"shortName"`
	Abbreviation     MLBAPISeriesStatusAbbreviation `json:"abbreviation"`
}

type MLBAPIIngTeam struct {
	SpringLeague  MLBAPISpringLeagueClass `json:"springLeague"`
	AllStarStatus MLBAPIDoubleHeader      `json:"allStarStatus"`
	ID            int64                   `json:"id"`
	Name          MLBAPITeamVenueName     `json:"name"`
	Link          MLBAPITeamLink          `json:"link"`
}

type MLBAPIStatus struct {
	AbstractGameState MLBAPIAbstractGameState `json:"abstractGameState"`
	CodedGameState    MLBAPICodedGameState    `json:"codedGameState"`
	DetailedState     string                  `json:"detailedState"`
	StatusCode        MLBAPICodedGameState    `json:"statusCode"`
	StartTimeTBD      bool                    `json:"startTimeTBD"`
	AbstractGameCode  MLBAPIAbstractGameCode  `json:"abstractGameCode"`
	Reason            *MLBAPIReason           `json:"reason,omitempty"`
}

type MLBAPIGameTeams struct {
	Home MLBAPIPurpleAway `json:"home"`
	Away MLBAPIPurpleAway `json:"away"`
}

type MLBAPIPurpleAway struct {
	IsContextTeam   bool                    `json:"isContextTeam"`
	IsFollowed      bool                    `json:"isFollowed"`
	IsFavorite      bool                    `json:"isFavorite"`
	LeagueRecord    MLBAPILeagueRecord      `json:"leagueRecord"`
	Score           *int64                  `json:"score,omitempty"`
	Team            MLBAPIPurpleTeam        `json:"team"`
	IsWinner        *bool                   `json:"isWinner,omitempty"`
	ProbablePitcher *MLBAPILoser            `json:"probablePitcher,omitempty"`
	SplitSquad      bool                    `json:"splitSquad"`
	SeriesNumber    int64                   `json:"seriesNumber"`
	SpringLeague    MLBAPISpringLeagueClass `json:"springLeague"`
}

type MLBAPILeagueRecord struct {
	WINS   int64  `json:"wins"`
	Losses int64  `json:"losses"`
	Pct    string `json:"pct"`
}

type MLBAPIPurpleTeam struct {
	SpringLeague    MLBAPISpringLeagueClass `json:"springLeague"`
	AllStarStatus   MLBAPIDoubleHeader      `json:"allStarStatus"`
	ID              int64                   `json:"id"`
	Name            MLBAPITeamVenueName     `json:"name"`
	Link            MLBAPITeamLink          `json:"link"`
	Season          int64                   `json:"season"`
	Venue           MLBAPISpringLeagueClass `json:"venue"`
	SpringVenue     MLBAPISpringVenue       `json:"springVenue"`
	TeamCode        string                  `json:"teamCode"`
	FileCode        string                  `json:"fileCode"`
	Abbreviation    string                  `json:"abbreviation"`
	TeamName        string                  `json:"teamName"`
	LocationName    MLBAPICity              `json:"locationName"`
	FirstYearOfPlay string                  `json:"firstYearOfPlay"`
	League          MLBAPISpringLeagueClass `json:"league"`
	Division        MLBAPISpringLeagueClass `json:"division"`
	Sport           MLBAPISpringLeagueClass `json:"sport"`
	ShortName       MLBAPICity              `json:"shortName"`
	FranchiseName   MLBAPICity              `json:"franchiseName"`
	ClubName        string                  `json:"clubName"`
	Active          bool                    `json:"active"`
	TeamLeaders     []MLBAPITeamLeader      `json:"teamLeaders,omitempty"`
}

type MLBAPISpringVenue struct {
	ID   int64                 `json:"id"`
	Link MLBAPISpringVenueLink `json:"link"`
}

type MLBAPITeamLeader struct {
	LeaderCategory MLBAPILeaderCategory    `json:"leaderCategory"`
	Season         string                  `json:"season"`
	GameType       MLBAPIGameType          `json:"gameType"`
	Leaders        []MLBAPILeader          `json:"leaders"`
	StatGroup      MLBAPIDisplayName       `json:"statGroup"`
	Team           MLBAPISpringLeagueClass `json:"team"`
	TotalSplits    int64                   `json:"totalSplits"`
}

type MLBAPIGameType struct {
	ID          MLBAPIGameTypeEnum          `json:"id"`
	Description MLBAPISeriesDescriptionEnum `json:"description"`
}

type MLBAPILeader struct {
	Rank   int64                   `json:"rank"`
	Value  string                  `json:"value"`
	Team   MLBAPISpringLeagueClass `json:"team"`
	League MLBAPISpringLeagueClass `json:"league"`
	Person MLBAPIPerson            `json:"person"`
	Sport  MLBAPISpringLeagueClass `json:"sport"`
	Season string                  `json:"season"`
}

type MLBAPIPerson struct {
	ID        int64  `json:"id"`
	FullName  string `json:"fullName"`
	Link      string `json:"link"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type MLBAPITicket struct {
	TicketType  MLBAPITicketType  `json:"ticketType"`
	TicketLinks MLBAPITicketLinks `json:"ticketLinks"`
}

type MLBAPITicketLinks struct {
	Home string `json:"home"`
}

type MLBAPIVenue struct {
	ID       int64               `json:"id"`
	Name     MLBAPITeamVenueName `json:"name"`
	Link     MLBAPITeamLink      `json:"link"`
	Location MLBAPILocation      `json:"location"`
	Active   bool                `json:"active"`
	Season   string              `json:"season"`
}

type MLBAPILocation struct {
	Address1           *string                   `json:"address1,omitempty"`
	City               MLBAPICity                `json:"city"`
	State              *MLBAPIState              `json:"state,omitempty"`
	StateAbbrev        *MLBAPIBirthStateProvince `json:"stateAbbrev,omitempty"`
	PostalCode         *string                   `json:"postalCode,omitempty"`
	DefaultCoordinates *MLBAPIDefaultCoordinates `json:"defaultCoordinates,omitempty"`
	AzimuthAngle       *float64                  `json:"azimuthAngle,omitempty"`
	Elevation          *int64                    `json:"elevation,omitempty"`
	Country            MLBAPICountry             `json:"country"`
	Phone              *string                   `json:"phone,omitempty"`
}

type MLBAPIDefaultCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MLBAPIXrefID struct {
	XrefID    string         `json:"xrefId"`
	XrefType  MLBAPIXrefType `json:"xrefType"`
	XrefIDTwo *string        `json:"xrefIdTwo,omitempty"`
}

type MLBAPIHeaders struct {
	XStitchSHA   string `json:"X-STITCH-SHA"`
	XStitchCache string `json:"X-STITCH-CACHE"`
	CacheControl string `json:"Cache-Control"`
	ContentType  string `json:"Content-Type"`
}

type MLBAPIReqHeaders struct {
	XStitchPRef string `json:"X-STITCH-P-REF"`
}

type MLBAPIAvailabilityCode string

const (
	LocalInMarket    MLBAPIAvailabilityCode = "local_in_market"
	LocalOutOfMarket MLBAPIAvailabilityCode = "local_out_of_market"
	National         MLBAPIAvailabilityCode = "national"
)

type MLBAPIAvailabilityText string

const (
	AvailabilityTextLocalInMarket    MLBAPIAvailabilityText = "Local (In Market)"
	AvailabilityTextLocalOutOfMarket MLBAPIAvailabilityText = "Local (Out of Market)"
	AvailabilityTextNational         MLBAPIAvailabilityText = "National"
)

type MLBAPIHomeAway string

const (
	Away MLBAPIHomeAway = "away"
	Home MLBAPIHomeAway = "home"
)

type MLBAPIBroadcastLanguage string

const (
	Fr       MLBAPIBroadcastLanguage = "fr"
	PurpleEn MLBAPIBroadcastLanguage = "en"
	PurpleEs MLBAPIBroadcastLanguage = "es"
)

type MLBAPIMediaStateCodeEnum string

const (
	MediaArchive MLBAPIMediaStateCodeEnum = "MEDIA_ARCHIVE"
	MediaOff     MLBAPIMediaStateCodeEnum = "MEDIA_OFF"
	MediaOn      MLBAPIMediaStateCodeEnum = "MEDIA_ON"
)

type MLBAPIMediaStateText string

const (
	MediaStateTextMediaArchive MLBAPIMediaStateText = "Media Archive"
	MediaStateTextMediaOff     MLBAPIMediaStateText = "Media Off"
	MediaStateTextMediaOn      MLBAPIMediaStateText = "Media On"
)

type MLBAPIBroadcastType string

const (
	Am MLBAPIBroadcastType = "AM"
	Fm MLBAPIBroadcastType = "FM"
	Tv MLBAPIBroadcastType = "TV"
)

type MLBAPICode string

const (
	H MLBAPICode = "H"
)

type MLBAPIResolutionFull string

const (
	HighResolution MLBAPIResolutionFull = "High Resolution"
)

type MLBAPIResolutionShort string

const (
	HD MLBAPIResolutionShort = "HD"
)

type MLBAPIItemLanguage string

const (
	En       MLBAPIItemLanguage = "EN"
	Es       MLBAPIItemLanguage = "ES"
	FluffyEn MLBAPIItemLanguage = "en"
	FluffyEs MLBAPIItemLanguage = "es"
)

type MLBAPIMediaFeedType string

const (
	MediaFeedTypeAWAY     MLBAPIMediaFeedType = "AWAY"
	MediaFeedTypeHOME     MLBAPIMediaFeedType = "HOME"
	MediaFeedTypeNATIONAL MLBAPIMediaFeedType = "NATIONAL"
)

type MLBAPIRenditionNameEnum string

const (
	English       MLBAPIRenditionNameEnum = "English"
	EnglishRadio  MLBAPIRenditionNameEnum = "English Radio"
	RadioEspañola MLBAPIRenditionNameEnum = "Radio Española"
	RenditionName MLBAPIRenditionNameEnum = ""
)

type MLBAPIItemType string

const (
	Type     MLBAPIItemType = ""
	TypeAWAY MLBAPIItemType = "AWAY"
	TypeHOME MLBAPIItemType = "HOME"
)

type MLBAPITitle string

const (
	Audio      MLBAPITitle = "Audio"
	MLBTVAudio MLBAPITitle = "MLBTV-Audio"
	Mlbtv      MLBAPITitle = "MLBTV"
)

type MLBAPIDayNight string

const (
	Day   MLBAPIDayNight = "day"
	Night MLBAPIDayNight = "night"
)

type MLBAPIGameTypeEnum string

const (
	GameTypeL MLBAPIGameTypeEnum = "L"
	GameTypeS MLBAPIGameTypeEnum = "S"
	R         MLBAPIGameTypeEnum = "R"
)

type MLBAPIBatSideDescription string

const (
	Left   MLBAPIBatSideDescription = "Left"
	Right  MLBAPIBatSideDescription = "Right"
	Switch MLBAPIBatSideDescription = "Switch"
)

type MLBAPICountry string

const (
	Aruba             MLBAPICountry = "Aruba"
	Australia         MLBAPICountry = "Australia"
	Bahamas           MLBAPICountry = "Bahamas"
	Brazil            MLBAPICountry = "Brazil"
	Canada            MLBAPICountry = "Canada"
	Colombia          MLBAPICountry = "Colombia"
	Cuba              MLBAPICountry = "Cuba"
	Curacao           MLBAPICountry = "Curacao"
	DominicanRepublic MLBAPICountry = "Dominican Republic"
	Germany           MLBAPICountry = "Germany"
	Honduras          MLBAPICountry = "Honduras"
	Japan             MLBAPICountry = "Japan"
	Mexico            MLBAPICountry = "Mexico"
	Nicaragua         MLBAPICountry = "Nicaragua"
	Panama            MLBAPICountry = "Panama"
	PanamaCanalZone   MLBAPICountry = "Panama Canal Zone"
	Peru              MLBAPICountry = "Peru"
	Portugal          MLBAPICountry = "Portugal"
	PuertoRico        MLBAPICountry = "Puerto Rico"
	RepublicOfKorea   MLBAPICountry = "Republic of Korea"
	SouthAfrica       MLBAPICountry = "South Africa"
	Usa               MLBAPICountry = "USA"
	Venezuela         MLBAPICountry = "Venezuela"
)

type MLBAPIBirthStateProvince string

const (
	Ab                   MLBAPIBirthStateProvince = "AB"
	Ar                   MLBAPIBirthStateProvince = "AR"
	Az                   MLBAPIBirthStateProvince = "AZ"
	BajaCalifornia       MLBAPIBirthStateProvince = "Baja California"
	Bc                   MLBAPIBirthStateProvince = "BC"
	BirthStateProvinceAL MLBAPIBirthStateProvince = "AL"
	CA                   MLBAPIBirthStateProvince = "CA"
	CT                   MLBAPIBirthStateProvince = "CT"
	Co                   MLBAPIBirthStateProvince = "CO"
	Dc                   MLBAPIBirthStateProvince = "DC"
	De                   MLBAPIBirthStateProvince = "DE"
	FL                   MLBAPIBirthStateProvince = "FL"
	Ga                   MLBAPIBirthStateProvince = "GA"
	Hi                   MLBAPIBirthStateProvince = "HI"
	IL                   MLBAPIBirthStateProvince = "IL"
	Ia                   MLBAPIBirthStateProvince = "IA"
	In                   MLBAPIBirthStateProvince = "IN"
	Ks                   MLBAPIBirthStateProvince = "KS"
	Ky                   MLBAPIBirthStateProvince = "KY"
	La                   MLBAPIBirthStateProvince = "LA"
	MS                   MLBAPIBirthStateProvince = "MS"
	Ma                   MLBAPIBirthStateProvince = "MA"
	Md                   MLBAPIBirthStateProvince = "MD"
	Mi                   MLBAPIBirthStateProvince = "MI"
	Mn                   MLBAPIBirthStateProvince = "MN"
	Mo                   MLBAPIBirthStateProvince = "MO"
	Nayarit              MLBAPIBirthStateProvince = "Nayarit"
	Nc                   MLBAPIBirthStateProvince = "NC"
	Nd                   MLBAPIBirthStateProvince = "ND"
	Ne                   MLBAPIBirthStateProvince = "NE"
	Nh                   MLBAPIBirthStateProvince = "NH"
	Nj                   MLBAPIBirthStateProvince = "NJ"
	Nm                   MLBAPIBirthStateProvince = "NM"
	Nv                   MLBAPIBirthStateProvince = "NV"
	Ny                   MLBAPIBirthStateProvince = "NY"
	Oh                   MLBAPIBirthStateProvince = "OH"
	Ok                   MLBAPIBirthStateProvince = "OK"
	On                   MLBAPIBirthStateProvince = "ON"
	Or                   MLBAPIBirthStateProvince = "OR"
	Pa                   MLBAPIBirthStateProvince = "PA"
	Qc                   MLBAPIBirthStateProvince = "QC"
	Sc                   MLBAPIBirthStateProvince = "SC"
	Sinaloa              MLBAPIBirthStateProvince = "Sinaloa"
	So                   MLBAPIBirthStateProvince = "SO"
	Sonora               MLBAPIBirthStateProvince = "Sonora"
	Tn                   MLBAPIBirthStateProvince = "TN"
	Tx                   MLBAPIBirthStateProvince = "TX"
	Ut                   MLBAPIBirthStateProvince = "UT"
	Va                   MLBAPIBirthStateProvince = "VA"
	Wa                   MLBAPIBirthStateProvince = "WA"
	Wi                   MLBAPIBirthStateProvince = "WI"
	Wy                   MLBAPIBirthStateProvince = "WY"
)

type MLBAPIGender string

const (
	M MLBAPIGender = "M"
)

type MLBAPIHeight string

const (
	The510 MLBAPIHeight = "5' 10\""
	The511 MLBAPIHeight = "5' 11\""
	The55  MLBAPIHeight = "5' 5\""
	The56  MLBAPIHeight = "5' 6\""
	The57  MLBAPIHeight = "5' 7\""
	The58  MLBAPIHeight = "5' 8\""
	The59  MLBAPIHeight = "5' 9\""
	The60  MLBAPIHeight = "6' 0\""
	The61  MLBAPIHeight = "6' 1\""
	The62  MLBAPIHeight = "6' 2\""
	The63  MLBAPIHeight = "6' 3\""
	The64  MLBAPIHeight = "6' 4\""
	The65  MLBAPIHeight = "6' 5\""
	The66  MLBAPIHeight = "6' 6\""
	The67  MLBAPIHeight = "6' 7\""
	The68  MLBAPIHeight = "6' 8\""
	The69  MLBAPIHeight = "6' 9\""
)

type MLBAPIName string

const (
	Ii  MLBAPIName = "II"
	Iii MLBAPIName = "III"
	Jr  MLBAPIName = "Jr."
)

type MLBAPIGamedayType string

const (
	C            MLBAPIGamedayType = "C"
	CF           MLBAPIGamedayType = "CF"
	Dh           MLBAPIGamedayType = "DH"
	GamedayTypeP MLBAPIGamedayType = "P"
	LF           MLBAPIGamedayType = "LF"
	PR           MLBAPIGamedayType = "PR"
	Ph           MLBAPIGamedayType = "PH"
	RF           MLBAPIGamedayType = "RF"
	Ss           MLBAPIGamedayType = "SS"
	The1B        MLBAPIGamedayType = "1B"
	The2B        MLBAPIGamedayType = "2B"
	The3B        MLBAPIGamedayType = "3B"
)

type MLBAPIPrimaryPositionName string

const (
	DesignatedHitter MLBAPIPrimaryPositionName = "Designated Hitter"
	FirstBase        MLBAPIPrimaryPositionName = "First Base"
	NameCatcher      MLBAPIPrimaryPositionName = "Catcher"
	NameOutfielder   MLBAPIPrimaryPositionName = "Outfielder"
	NamePitcher      MLBAPIPrimaryPositionName = "Pitcher"
	PinchHitter      MLBAPIPrimaryPositionName = "Pinch Hitter"
	PinchRunner      MLBAPIPrimaryPositionName = "Pinch Runner"
	SecondBase       MLBAPIPrimaryPositionName = "Second Base"
	Shortstop        MLBAPIPrimaryPositionName = "Shortstop"
	ThirdBase        MLBAPIPrimaryPositionName = "Third Base"
)

type MLBAPIPrimaryPositionType string

const (
	Hitter         MLBAPIPrimaryPositionType = "Hitter"
	Infielder      MLBAPIPrimaryPositionType = "Infielder"
	Runner         MLBAPIPrimaryPositionType = "Runner"
	TypeCatcher    MLBAPIPrimaryPositionType = "Catcher"
	TypeOutfielder MLBAPIPrimaryPositionType = "Outfielder"
	TypePitcher    MLBAPIPrimaryPositionType = "Pitcher"
)

type MLBAPIDisplayName string

const (
	GameLog           MLBAPIDisplayName = "gameLog"
	Hitting           MLBAPIDisplayName = "hitting"
	Pitching          MLBAPIDisplayName = "pitching"
	StatsSingleSeason MLBAPIDisplayName = "statsSingleSeason"
)

type MLBAPIPercentage string

const (
	Empty          MLBAPIPercentage = ".---"
	Percentage000  MLBAPIPercentage = ".000"
	Percentage1000 MLBAPIPercentage = "1.000"
	Percentage250  MLBAPIPercentage = ".250"
	Percentage400  MLBAPIPercentage = ".400"
	Percentage500  MLBAPIPercentage = ".500"
	Percentage600  MLBAPIPercentage = ".600"
	Percentage700  MLBAPIPercentage = ".700"
	Percentage750  MLBAPIPercentage = ".750"
	Percentage800  MLBAPIPercentage = ".800"
	Percentage900  MLBAPIPercentage = ".900"
	The200         MLBAPIPercentage = ".200"
	The333         MLBAPIPercentage = ".333"
	The444         MLBAPIPercentage = ".444"
	The571         MLBAPIPercentage = ".571"
	The667         MLBAPIPercentage = ".667"
	The714         MLBAPIPercentage = ".714"
	The727         MLBAPIPercentage = ".727"
	The778         MLBAPIPercentage = ".778"
	The833         MLBAPIPercentage = ".833"
	The857         MLBAPIPercentage = ".857"
	The875         MLBAPIPercentage = ".875"
	The889         MLBAPIPercentage = ".889"
)

type MLBAPIStrikePercentageEnum string

const (
	StrikePercentage     MLBAPIStrikePercentageEnum = "-.--"
	StrikePercentage000  MLBAPIStrikePercentageEnum = ".000"
	StrikePercentage1000 MLBAPIStrikePercentageEnum = "1.000"
	StrikePercentage250  MLBAPIStrikePercentageEnum = ".250"
	StrikePercentage400  MLBAPIStrikePercentageEnum = ".400"
	StrikePercentage500  MLBAPIStrikePercentageEnum = ".500"
	StrikePercentage510  MLBAPIStrikePercentageEnum = ".510"
	StrikePercentage600  MLBAPIStrikePercentageEnum = ".600"
	StrikePercentage700  MLBAPIStrikePercentageEnum = ".700"
	StrikePercentage750  MLBAPIStrikePercentageEnum = ".750"
	StrikePercentage800  MLBAPIStrikePercentageEnum = ".800"
	StrikePercentage900  MLBAPIStrikePercentageEnum = ".900"
	The290               MLBAPIStrikePercentageEnum = ".290"
	The350               MLBAPIStrikePercentageEnum = ".350"
	The390               MLBAPIStrikePercentageEnum = ".390"
	The410               MLBAPIStrikePercentageEnum = ".410"
	The420               MLBAPIStrikePercentageEnum = ".420"
	The430               MLBAPIStrikePercentageEnum = ".430"
	The450               MLBAPIStrikePercentageEnum = ".450"
	The460               MLBAPIStrikePercentageEnum = ".460"
	The470               MLBAPIStrikePercentageEnum = ".470"
	The480               MLBAPIStrikePercentageEnum = ".480"
	The490               MLBAPIStrikePercentageEnum = ".490"
	The520               MLBAPIStrikePercentageEnum = ".520"
	The530               MLBAPIStrikePercentageEnum = ".530"
	The540               MLBAPIStrikePercentageEnum = ".540"
	The550               MLBAPIStrikePercentageEnum = ".550"
	The560               MLBAPIStrikePercentageEnum = ".560"
	The570               MLBAPIStrikePercentageEnum = ".570"
	The580               MLBAPIStrikePercentageEnum = ".580"
	The590               MLBAPIStrikePercentageEnum = ".590"
	The610               MLBAPIStrikePercentageEnum = ".610"
	The620               MLBAPIStrikePercentageEnum = ".620"
	The630               MLBAPIStrikePercentageEnum = ".630"
	The640               MLBAPIStrikePercentageEnum = ".640"
	The650               MLBAPIStrikePercentageEnum = ".650"
	The660               MLBAPIStrikePercentageEnum = ".660"
	The670               MLBAPIStrikePercentageEnum = ".670"
	The680               MLBAPIStrikePercentageEnum = ".680"
	The690               MLBAPIStrikePercentageEnum = ".690"
	The710               MLBAPIStrikePercentageEnum = ".710"
	The720               MLBAPIStrikePercentageEnum = ".720"
	The730               MLBAPIStrikePercentageEnum = ".730"
	The740               MLBAPIStrikePercentageEnum = ".740"
	The760               MLBAPIStrikePercentageEnum = ".760"
	The770               MLBAPIStrikePercentageEnum = ".770"
	The780               MLBAPIStrikePercentageEnum = ".780"
	The790               MLBAPIStrikePercentageEnum = ".790"
	The810               MLBAPIStrikePercentageEnum = ".810"
	The820               MLBAPIStrikePercentageEnum = ".820"
	The830               MLBAPIStrikePercentageEnum = ".830"
	The840               MLBAPIStrikePercentageEnum = ".840"
	The850               MLBAPIStrikePercentageEnum = ".850"
	The860               MLBAPIStrikePercentageEnum = ".860"
	The870               MLBAPIStrikePercentageEnum = ".870"
	The880               MLBAPIStrikePercentageEnum = ".880"
	The890               MLBAPIStrikePercentageEnum = ".890"
)

type MLBAPIDoubleHeader string

const (
	DoubleHeaderS MLBAPIDoubleHeader = "S"
	N             MLBAPIDoubleHeader = "N"
	Y             MLBAPIDoubleHeader = "Y"
)

type MLBAPIHalfInning string

const (
	Bottom MLBAPIHalfInning = "bottom"
	Top    MLBAPIHalfInning = "top"
)

type MLBAPIEvent string

const (
	HomeRun MLBAPIEvent = "Home Run"
)

type MLBAPIResultType string

const (
	AtBat MLBAPIResultType = "atBat"
)

type MLBAPIIfNecessaryDescription string

const (
	NormalGame MLBAPIIfNecessaryDescription = "Normal Game"
)

type MLBAPINote string

const (
	GameCalledRainAfterTheBottomOfThe8Th MLBAPINote = "Game called (rain) after the bottom of the 8th."
	NoneOutWhenWinningRunScored          MLBAPINote = "None out when winning run scored."
	OneOutWhenWinningRunScored           MLBAPINote = "One out when winning run scored."
	TwoOutWhenWinningRunScored           MLBAPINote = "Two out when winning run scored."
)

type MLBAPITeamAbbreviation string

const (
	Cl  MLBAPITeamAbbreviation = "CL"
	Gl  MLBAPITeamAbbreviation = "GL"
	Mlb MLBAPITeamAbbreviation = "MLB"
)

type MLBAPITeamLink string

const (
	APIV1Divisions200 MLBAPITeamLink = "/api/v1/divisions/200"
	APIV1Divisions201 MLBAPITeamLink = "/api/v1/divisions/201"
	APIV1Divisions202 MLBAPITeamLink = "/api/v1/divisions/202"
	APIV1Divisions203 MLBAPITeamLink = "/api/v1/divisions/203"
	APIV1Divisions204 MLBAPITeamLink = "/api/v1/divisions/204"
	APIV1Divisions205 MLBAPITeamLink = "/api/v1/divisions/205"
	APIV1League103    MLBAPITeamLink = "/api/v1/league/103"
	APIV1League104    MLBAPITeamLink = "/api/v1/league/104"
	APIV1League114    MLBAPITeamLink = "/api/v1/league/114"
	APIV1League115    MLBAPITeamLink = "/api/v1/league/115"
	APIV1Sports1      MLBAPITeamLink = "/api/v1/sports/1"
	APIV1Teams108     MLBAPITeamLink = "/api/v1/teams/108"
	APIV1Teams109     MLBAPITeamLink = "/api/v1/teams/109"
	APIV1Teams110     MLBAPITeamLink = "/api/v1/teams/110"
	APIV1Teams111     MLBAPITeamLink = "/api/v1/teams/111"
	APIV1Teams112     MLBAPITeamLink = "/api/v1/teams/112"
	APIV1Teams113     MLBAPITeamLink = "/api/v1/teams/113"
	APIV1Teams114     MLBAPITeamLink = "/api/v1/teams/114"
	APIV1Teams115     MLBAPITeamLink = "/api/v1/teams/115"
	APIV1Teams116     MLBAPITeamLink = "/api/v1/teams/116"
	APIV1Teams117     MLBAPITeamLink = "/api/v1/teams/117"
	APIV1Teams118     MLBAPITeamLink = "/api/v1/teams/118"
	APIV1Teams119     MLBAPITeamLink = "/api/v1/teams/119"
	APIV1Teams120     MLBAPITeamLink = "/api/v1/teams/120"
	APIV1Teams121     MLBAPITeamLink = "/api/v1/teams/121"
	APIV1Teams133     MLBAPITeamLink = "/api/v1/teams/133"
	APIV1Teams134     MLBAPITeamLink = "/api/v1/teams/134"
	APIV1Teams135     MLBAPITeamLink = "/api/v1/teams/135"
	APIV1Teams136     MLBAPITeamLink = "/api/v1/teams/136"
	APIV1Teams137     MLBAPITeamLink = "/api/v1/teams/137"
	APIV1Teams138     MLBAPITeamLink = "/api/v1/teams/138"
	APIV1Teams139     MLBAPITeamLink = "/api/v1/teams/139"
	APIV1Teams140     MLBAPITeamLink = "/api/v1/teams/140"
	APIV1Teams141     MLBAPITeamLink = "/api/v1/teams/141"
	APIV1Teams142     MLBAPITeamLink = "/api/v1/teams/142"
	APIV1Teams143     MLBAPITeamLink = "/api/v1/teams/143"
	APIV1Teams144     MLBAPITeamLink = "/api/v1/teams/144"
	APIV1Teams145     MLBAPITeamLink = "/api/v1/teams/145"
	APIV1Teams146     MLBAPITeamLink = "/api/v1/teams/146"
	APIV1Teams147     MLBAPITeamLink = "/api/v1/teams/147"
	APIV1Teams158     MLBAPITeamLink = "/api/v1/teams/158"
	APIV1Venues1      MLBAPITeamLink = "/api/v1/venues/1"
	APIV1Venues10     MLBAPITeamLink = "/api/v1/venues/10"
	APIV1Venues12     MLBAPITeamLink = "/api/v1/venues/12"
	APIV1Venues14     MLBAPITeamLink = "/api/v1/venues/14"
	APIV1Venues15     MLBAPITeamLink = "/api/v1/venues/15"
	APIV1Venues17     MLBAPITeamLink = "/api/v1/venues/17"
	APIV1Venues19     MLBAPITeamLink = "/api/v1/venues/19"
	APIV1Venues2      MLBAPITeamLink = "/api/v1/venues/2"
	APIV1Venues22     MLBAPITeamLink = "/api/v1/venues/22"
	APIV1Venues2392   MLBAPITeamLink = "/api/v1/venues/2392"
	APIV1Venues2394   MLBAPITeamLink = "/api/v1/venues/2394"
	APIV1Venues2395   MLBAPITeamLink = "/api/v1/venues/2395"
	APIV1Venues2602   MLBAPITeamLink = "/api/v1/venues/2602"
	APIV1Venues2680   MLBAPITeamLink = "/api/v1/venues/2680"
	APIV1Venues2681   MLBAPITeamLink = "/api/v1/venues/2681"
	APIV1Venues2889   MLBAPITeamLink = "/api/v1/venues/2889"
	APIV1Venues3      MLBAPITeamLink = "/api/v1/venues/3"
	APIV1Venues31     MLBAPITeamLink = "/api/v1/venues/31"
	APIV1Venues32     MLBAPITeamLink = "/api/v1/venues/32"
	APIV1Venues3289   MLBAPITeamLink = "/api/v1/venues/3289"
	APIV1Venues3309   MLBAPITeamLink = "/api/v1/venues/3309"
	APIV1Venues3312   MLBAPITeamLink = "/api/v1/venues/3312"
	APIV1Venues3313   MLBAPITeamLink = "/api/v1/venues/3313"
	APIV1Venues4      MLBAPITeamLink = "/api/v1/venues/4"
	APIV1Venues4169   MLBAPITeamLink = "/api/v1/venues/4169"
	APIV1Venues4705   MLBAPITeamLink = "/api/v1/venues/4705"
	APIV1Venues5      MLBAPITeamLink = "/api/v1/venues/5"
	APIV1Venues5325   MLBAPITeamLink = "/api/v1/venues/5325"
	APIV1Venues5340   MLBAPITeamLink = "/api/v1/venues/5340"
	APIV1Venues680    MLBAPITeamLink = "/api/v1/venues/680"
	APIV1Venues7      MLBAPITeamLink = "/api/v1/venues/7"
)

type MLBAPITeamVenueName string

const (
	AmericanFamilyField     MLBAPITeamVenueName = "American Family Field"
	AmericanLeague          MLBAPITeamVenueName = "American League"
	AmericanLeagueCentral   MLBAPITeamVenueName = "American League Central"
	AmericanLeagueEast      MLBAPITeamVenueName = "American League East"
	AmericanLeagueWest      MLBAPITeamVenueName = "American League West"
	AngelStadium            MLBAPITeamVenueName = "Angel Stadium"
	ArizonaDiamondbacks     MLBAPITeamVenueName = "Arizona Diamondbacks"
	AtlantaBraves           MLBAPITeamVenueName = "Atlanta Braves"
	BaltimoreOrioles        MLBAPITeamVenueName = "Baltimore Orioles"
	BostonRedSox            MLBAPITeamVenueName = "Boston Red Sox"
	BuschStadium            MLBAPITeamVenueName = "Busch Stadium"
	CactusLeague            MLBAPITeamVenueName = "Cactus League"
	ChaseField              MLBAPITeamVenueName = "Chase Field"
	ChicagoCubs             MLBAPITeamVenueName = "Chicago Cubs"
	ChicagoWhiteSox         MLBAPITeamVenueName = "Chicago White Sox"
	CincinnatiReds          MLBAPITeamVenueName = "Cincinnati Reds"
	CitiField               MLBAPITeamVenueName = "Citi Field"
	CitizensBankPark        MLBAPITeamVenueName = "Citizens Bank Park"
	ClevelandGuardians      MLBAPITeamVenueName = "Cleveland Guardians"
	ColoradoRockies         MLBAPITeamVenueName = "Colorado Rockies"
	ComericaPark            MLBAPITeamVenueName = "Comerica Park"
	CoorsField              MLBAPITeamVenueName = "Coors Field"
	DetroitTigers           MLBAPITeamVenueName = "Detroit Tigers"
	DodgerStadium           MLBAPITeamVenueName = "Dodger Stadium"
	EstadioAlfredoHarpHelu  MLBAPITeamVenueName = "Estadio Alfredo Harp Helu"
	FenwayPark              MLBAPITeamVenueName = "Fenway Park"
	GlobeLifeField          MLBAPITeamVenueName = "Globe Life Field"
	GrapefruitLeague        MLBAPITeamVenueName = "Grapefruit League"
	GreatAmericanBallPark   MLBAPITeamVenueName = "Great American Ball Park"
	GuaranteedRateField     MLBAPITeamVenueName = "Guaranteed Rate Field"
	HoustonAstros           MLBAPITeamVenueName = "Houston Astros"
	KansasCityRoyals        MLBAPITeamVenueName = "Kansas City Royals"
	KauffmanStadium         MLBAPITeamVenueName = "Kauffman Stadium"
	LoanDepotPark           MLBAPITeamVenueName = "loanDepot park"
	LosAngelesAngels        MLBAPITeamVenueName = "Los Angeles Angels"
	LosAngelesDodgers       MLBAPITeamVenueName = "Los Angeles Dodgers"
	MajorLeagueBaseball     MLBAPITeamVenueName = "Major League Baseball"
	MiamiMarlins            MLBAPITeamVenueName = "Miami Marlins"
	MilwaukeeBrewers        MLBAPITeamVenueName = "Milwaukee Brewers"
	MinnesotaTwins          MLBAPITeamVenueName = "Minnesota Twins"
	MinuteMaidPark          MLBAPITeamVenueName = "Minute Maid Park"
	NameAL                  MLBAPITeamVenueName = "AL"
	NationalLeague          MLBAPITeamVenueName = "National League"
	NationalLeagueCentral   MLBAPITeamVenueName = "National League Central"
	NationalLeagueEast      MLBAPITeamVenueName = "National League East"
	NationalLeagueWest      MLBAPITeamVenueName = "National League West"
	NationalsPark           MLBAPITeamVenueName = "Nationals Park"
	NewYorkMets             MLBAPITeamVenueName = "New York Mets"
	NewYorkYankees          MLBAPITeamVenueName = "New York Yankees"
	Nl                      MLBAPITeamVenueName = "NL"
	OaklandAthletics        MLBAPITeamVenueName = "Oakland Athletics"
	OaklandColiseum         MLBAPITeamVenueName = "Oakland Coliseum"
	OraclePark              MLBAPITeamVenueName = "Oracle Park"
	OrioleParkAtCamdenYards MLBAPITeamVenueName = "Oriole Park at Camden Yards"
	PNCPark                 MLBAPITeamVenueName = "PNC Park"
	PetcoPark               MLBAPITeamVenueName = "Petco Park"
	PhiladelphiaPhillies    MLBAPITeamVenueName = "Philadelphia Phillies"
	PittsburghPirates       MLBAPITeamVenueName = "Pittsburgh Pirates"
	ProgressiveField        MLBAPITeamVenueName = "Progressive Field"
	RogersCentre            MLBAPITeamVenueName = "Rogers Centre"
	SANDiegoPadres          MLBAPITeamVenueName = "San Diego Padres"
	SANFranciscoGiants      MLBAPITeamVenueName = "San Francisco Giants"
	SeattleMariners         MLBAPITeamVenueName = "Seattle Mariners"
	StLouisCardinals        MLBAPITeamVenueName = "St. Louis Cardinals"
	TMobilePark             MLBAPITeamVenueName = "T-Mobile Park"
	TampaBayRays            MLBAPITeamVenueName = "Tampa Bay Rays"
	TargetField             MLBAPITeamVenueName = "Target Field"
	TexasRangers            MLBAPITeamVenueName = "Texas Rangers"
	TorontoBlueJays         MLBAPITeamVenueName = "Toronto Blue Jays"
	TropicanaField          MLBAPITeamVenueName = "Tropicana Field"
	TruistPark              MLBAPITeamVenueName = "Truist Park"
	WashingtonNationals     MLBAPITeamVenueName = "Washington Nationals"
	WrigleyField            MLBAPITeamVenueName = "Wrigley Field"
	YankeeStadium           MLBAPITeamVenueName = "Yankee Stadium"
)

type MLBAPISeriesDescriptionEnum string

const (
	RegularSeason MLBAPISeriesDescriptionEnum = "Regular Season"
)

type MLBAPISeriesStatusAbbreviation string

const (
	Rs MLBAPISeriesStatusAbbreviation = "RS"
)

type MLBAPIShort string

const (
	Season MLBAPIShort = "Season"
)

type MLBAPIAbstractGameCode string

const (
	AbstractGameCodeF MLBAPIAbstractGameCode = "F"
	AbstractGameCodeL MLBAPIAbstractGameCode = "L"
	AbstractGameCodeP MLBAPIAbstractGameCode = "P"
)

type MLBAPIAbstractGameState string

const (
	AbstractGameStateFinal MLBAPIAbstractGameState = "Final"
	Live                   MLBAPIAbstractGameState = "Live"
	Preview                MLBAPIAbstractGameState = "Preview"
)

type MLBAPICodedGameState string

const (
	CodedGameStateF  MLBAPICodedGameState = "F"
	CodedGameStateFR MLBAPICodedGameState = "FR"
	CodedGameStateP  MLBAPICodedGameState = "P"
	CodedGameStateS  MLBAPICodedGameState = "S"
	D                MLBAPICodedGameState = "D"
	DR               MLBAPICodedGameState = "DR"
	Di               MLBAPICodedGameState = "DI"
	I                MLBAPICodedGameState = "I"
)

type MLBAPIReason string

const (
	InclementWeather MLBAPIReason = "Inclement Weather"
	Rain             MLBAPIReason = "Rain"
)

type MLBAPICity string

const (
	Anaheim        MLBAPICity = "Anaheim"
	Arlington      MLBAPICity = "Arlington"
	Atlanta        MLBAPICity = "Atlanta"
	Baltimore      MLBAPICity = "Baltimore"
	Boston         MLBAPICity = "Boston"
	Bronx          MLBAPICity = "Bronx"
	ChiCubs        MLBAPICity = "Chi Cubs"
	ChiWhiteSox    MLBAPICity = "Chi White Sox"
	Chicago        MLBAPICity = "Chicago"
	Cincinnati     MLBAPICity = "Cincinnati"
	CityArizona    MLBAPICity = "Arizona"
	CityColorado   MLBAPICity = "Colorado"
	CityMinnesota  MLBAPICity = "Minnesota"
	CityNewYork    MLBAPICity = "New York"
	CityTexas      MLBAPICity = "Texas"
	CityWashington MLBAPICity = "Washington"
	Cleveland      MLBAPICity = "Cleveland"
	Denver         MLBAPICity = "Denver"
	Detroit        MLBAPICity = "Detroit"
	Flushing       MLBAPICity = "Flushing"
	Houston        MLBAPICity = "Houston"
	KansasCity     MLBAPICity = "Kansas City"
	LAAngels       MLBAPICity = "LA Angels"
	LADodgers      MLBAPICity = "LA Dodgers"
	LosAngeles     MLBAPICity = "Los Angeles"
	MexicoCity     MLBAPICity = "Mexico City"
	Miami          MLBAPICity = "Miami"
	Milwaukee      MLBAPICity = "Milwaukee"
	Minneapolis    MLBAPICity = "Minneapolis"
	NYMets         MLBAPICity = "NY Mets"
	NYYankees      MLBAPICity = "NY Yankees"
	Oakland        MLBAPICity = "Oakland"
	Philadelphia   MLBAPICity = "Philadelphia"
	Phoenix        MLBAPICity = "Phoenix"
	Pittsburgh     MLBAPICity = "Pittsburgh"
	SANDiego       MLBAPICity = "San Diego"
	SANFrancisco   MLBAPICity = "San Francisco"
	Seattle        MLBAPICity = "Seattle"
	StLouis        MLBAPICity = "St. Louis"
	StPetersburg   MLBAPICity = "St. Petersburg"
	TampaBay       MLBAPICity = "Tampa Bay"
	Toronto        MLBAPICity = "Toronto"
)

type MLBAPISpringVenueLink string

const (
	APIV1Venues2500 MLBAPISpringVenueLink = "/api/v1/venues/2500"
	APIV1Venues2507 MLBAPISpringVenueLink = "/api/v1/venues/2507"
	APIV1Venues2508 MLBAPISpringVenueLink = "/api/v1/venues/2508"
	APIV1Venues2511 MLBAPISpringVenueLink = "/api/v1/venues/2511"
	APIV1Venues2518 MLBAPISpringVenueLink = "/api/v1/venues/2518"
	APIV1Venues2520 MLBAPISpringVenueLink = "/api/v1/venues/2520"
	APIV1Venues2523 MLBAPISpringVenueLink = "/api/v1/venues/2523"
	APIV1Venues2526 MLBAPISpringVenueLink = "/api/v1/venues/2526"
	APIV1Venues2530 MLBAPISpringVenueLink = "/api/v1/venues/2530"
	APIV1Venues2532 MLBAPISpringVenueLink = "/api/v1/venues/2532"
	APIV1Venues2534 MLBAPISpringVenueLink = "/api/v1/venues/2534"
	APIV1Venues2536 MLBAPISpringVenueLink = "/api/v1/venues/2536"
	APIV1Venues2603 MLBAPISpringVenueLink = "/api/v1/venues/2603"
	APIV1Venues2700 MLBAPISpringVenueLink = "/api/v1/venues/2700"
	APIV1Venues2856 MLBAPISpringVenueLink = "/api/v1/venues/2856"
	APIV1Venues2862 MLBAPISpringVenueLink = "/api/v1/venues/2862"
	APIV1Venues3809 MLBAPISpringVenueLink = "/api/v1/venues/3809"
	APIV1Venues3834 MLBAPISpringVenueLink = "/api/v1/venues/3834"
	APIV1Venues4249 MLBAPISpringVenueLink = "/api/v1/venues/4249"
	APIV1Venues4309 MLBAPISpringVenueLink = "/api/v1/venues/4309"
	APIV1Venues4629 MLBAPISpringVenueLink = "/api/v1/venues/4629"
	APIV1Venues5000 MLBAPISpringVenueLink = "/api/v1/venues/5000"
	APIV1Venues5380 MLBAPISpringVenueLink = "/api/v1/venues/5380"
)

type MLBAPILeaderCategory string

const (
	BattingAverage MLBAPILeaderCategory = "battingAverage"
	HomeRuns       MLBAPILeaderCategory = "homeRuns"
	RunsBattedIn   MLBAPILeaderCategory = "runsBattedIn"
)

type MLBAPITicketType string

const (
	Mobile MLBAPITicketType = "mobile"
	Wired  MLBAPITicketType = "wired"
)

type MLBAPIState string

const (
	California         MLBAPIState = "California"
	DistrictOfColumbia MLBAPIState = "District of Columbia"
	Florida            MLBAPIState = "Florida"
	Georgia            MLBAPIState = "Georgia"
	Illinois           MLBAPIState = "Illinois"
	Maryland           MLBAPIState = "Maryland"
	Massachusetts      MLBAPIState = "Massachusetts"
	Michigan           MLBAPIState = "Michigan"
	Missouri           MLBAPIState = "Missouri"
	Ohio               MLBAPIState = "Ohio"
	Ontario            MLBAPIState = "Ontario"
	Pennsylvania       MLBAPIState = "Pennsylvania"
	StateArizona       MLBAPIState = "Arizona"
	StateColorado      MLBAPIState = "Colorado"
	StateMinnesota     MLBAPIState = "Minnesota"
	StateNewYork       MLBAPIState = "New York"
	StateTexas         MLBAPIState = "Texas"
	StateWashington    MLBAPIState = "Washington"
	Wisconsin          MLBAPIState = "Wisconsin"
)

type MLBAPIXrefType string

const (
	BAMAppletvFree   MLBAPIXrefType = "bam_appletv_free"
	BAMMlbtvFree     MLBAPIXrefType = "bam_mlbtv_free"
	Bis              MLBAPIXrefType = "bis"
	PostgameShowAway MLBAPIXrefType = "postgame_show_away"
	PostgameShowHome MLBAPIXrefType = "postgame_show_home"
	PregameShowAway  MLBAPIXrefType = "pregame_show_away"
	PregameShowHome  MLBAPIXrefType = "pregame_show_home"
)
