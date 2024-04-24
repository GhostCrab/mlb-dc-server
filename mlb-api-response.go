package main

type MLBAPIResponse struct {
  Copyright            string     `json:"copyright"`           
  TotalItems           int64      `json:"totalItems"`          
  TotalEvents          int64      `json:"totalEvents"`         
  TotalGames           int64      `json:"totalGames"`          
  TotalGamesInProgress int64      `json:"totalGamesInProgress"`
  Dates                []Date     `json:"dates"`               
  ReqHeaders           ReqHeaders `json:"reqHeaders"`          
  Headers              Headers    `json:"headers"`             
  Status               int64      `json:"status"`              
}

type Date struct {
  Date                 string        `json:"date"`                
  TotalItems           int64         `json:"totalItems"`          
  TotalEvents          int64         `json:"totalEvents"`         
  TotalGames           int64         `json:"totalGames"`          
  TotalGamesInProgress int64         `json:"totalGamesInProgress"`
  Games                []Game        `json:"games"`               
  Events               []interface{} `json:"events"`              
}

type Game struct {
  SortIndex1             int64                  `json:"sortIndex1"`                   
  SortIndex2             int64                  `json:"sortIndex2"`                   
  SortIndex3             int64                  `json:"sortIndex3"`                   
  SortIndex4             int64                  `json:"sortIndex4"`                   
  SortIndex5             int64                  `json:"sortIndex5"`                   
  Teams                  GameTeams              `json:"teams"`                        
  GamePk                 int64                  `json:"gamePk"`                       
  GameGUID               string                 `json:"gameGuid"`                     
  Link                   string                 `json:"link"`                         
  GameType               GameTypeEnum           `json:"gameType"`                     
  Season                 string                 `json:"season"`                       
  GameDate               string                 `json:"gameDate"`                     
  OfficialDate           string                 `json:"officialDate"`                 
  Status                 Status                 `json:"status"`                       
  Linescore              Linescore              `json:"linescore"`                    
  Decisions              *Decisions             `json:"decisions,omitempty"`          
  Venue                  Venue                  `json:"venue"`                        
  Broadcasts             []Broadcast            `json:"broadcasts,omitempty"`         
  Content                Content                `json:"content"`                      
  SeriesStatus           SeriesStatus           `json:"seriesStatus"`                 
  IsTie                  *bool                  `json:"isTie,omitempty"`              
  XrefIDS                []XrefID               `json:"xrefIds"`                      
  GameNumber             int64                  `json:"gameNumber"`                   
  PublicFacing           bool                   `json:"publicFacing"`                 
  DoubleHeader           DoubleHeader           `json:"doubleHeader"`                 
  GamedayType            GamedayType            `json:"gamedayType"`                  
  Tiebreaker             DoubleHeader           `json:"tiebreaker"`                   
  CalendarEventID        string                 `json:"calendarEventID"`              
  SeasonDisplay          string                 `json:"seasonDisplay"`                
  DayNight               DayNight               `json:"dayNight"`                     
  Description            *string                `json:"description,omitempty"`        
  ScheduledInnings       int64                  `json:"scheduledInnings"`             
  ReverseHomeAwayStatus  bool                   `json:"reverseHomeAwayStatus"`        
  InningBreakLength      int64                  `json:"inningBreakLength"`            
  GamesInSeries          int64                  `json:"gamesInSeries"`                
  SeriesGameNumber       int64                  `json:"seriesGameNumber"`             
  SeriesDescription      SeriesDescriptionEnum  `json:"seriesDescription"`            
  Review                 *Review                `json:"review,omitempty"`             
  Flags                  *Flags                 `json:"flags,omitempty"`              
  HomeRuns               []HomeRunElement       `json:"homeRuns,omitempty"`           
  RecordSource           GameTypeEnum           `json:"recordSource"`                 
  IfNecessary            DoubleHeader           `json:"ifNecessary"`                  
  IfNecessaryDescription IfNecessaryDescription `json:"ifNecessaryDescription"`       
  GameUtils              map[string]bool        `json:"gameUtils"`                    
  RescheduleDate         *string                `json:"rescheduleDate,omitempty"`     
  RescheduleGameDate     *string                `json:"rescheduleGameDate,omitempty"` 
  RescheduledFrom        *string                `json:"rescheduledFrom,omitempty"`    
  RescheduledFromDate    *string                `json:"rescheduledFromDate,omitempty"`
  Tickets                []Ticket               `json:"tickets,omitempty"`            
}

type Broadcast struct {
  ID                    int64             `json:"id"`                       
  Name                  string            `json:"name"`                     
  Type                  BroadcastType     `json:"type"`                     
  Language              BroadcastLanguage `json:"language"`                 
  IsNational            bool              `json:"isNational"`               
  CallSign              string            `json:"callSign"`                 
  VideoResolution       *VideoResolution  `json:"videoResolution,omitempty"`
  MediaState            MediaState        `json:"mediaState"`               
  BroadcastDate         string            `json:"broadcastDate"`            
  MediaID               string            `json:"mediaId"`                  
  GameDateBroadcastGUID string            `json:"gameDateBroadcastGuid"`    
  HomeAway              HomeAway          `json:"homeAway"`                 
  FreeGame              bool              `json:"freeGame"`                 
  AvailableForStreaming bool              `json:"availableForStreaming"`    
  PostGameShow          bool              `json:"postGameShow"`             
  MvpdAuthRequired      bool              `json:"mvpdAuthRequired"`         
  PreGameShow           *string           `json:"preGameShow,omitempty"`    
  Availability          *Availability     `json:"availability,omitempty"`   
  SourceComment         *string           `json:"sourceComment,omitempty"`  
}

type Availability struct {
  AvailabilityID   int64            `json:"availabilityId"`  
  AvailabilityCode AvailabilityCode `json:"availabilityCode"`
  AvailabilityText AvailabilityText `json:"availabilityText"`
}

type MediaState struct {
  MediaStateID   int64              `json:"mediaStateId"`  
  MediaStateCode MediaStateCodeEnum `json:"mediaStateCode"`
  MediaStateText MediaStateText     `json:"mediaStateText"`
}

type VideoResolution struct {
  Code            Code            `json:"code"`           
  ResolutionShort ResolutionShort `json:"resolutionShort"`
  ResolutionFull  ResolutionFull  `json:"resolutionFull"` 
}

type Content struct {
  Link       string     `json:"link"`                
  Editorial  *Editorial `json:"editorial,omitempty"` 
  Media      Media      `json:"media"`               
  Highlights *Editorial `json:"highlights,omitempty"`
  Summary    *Summary   `json:"summary,omitempty"`   
  GameNotes  *Editorial `json:"gameNotes,omitempty"` 
}

type Editorial struct {
}

type Media struct {
  Epg           []Epg          `json:"epg,omitempty"`          
  FeaturedMedia *FeaturedMedia `json:"featuredMedia,omitempty"`
  FreeGame      *bool          `json:"freeGame,omitempty"`     
  EnhancedGame  *bool          `json:"enhancedGame,omitempty"` 
}

type Epg struct {
  Title Title  `json:"title"`
  Items []Item `json:"items"`
}

type Item struct {
  CallLetters       *string             `json:"callLetters,omitempty"`      
  EspnAuthRequired  *bool               `json:"espnAuthRequired,omitempty"` 
  TbsAuthRequired   *bool               `json:"tbsAuthRequired,omitempty"`  
  Espn2AuthRequired *bool               `json:"espn2AuthRequired,omitempty"`
  GameDate          *string             `json:"gameDate,omitempty"`         
  ContentID         *string             `json:"contentId,omitempty"`        
  Fs1AuthRequired   *bool               `json:"fs1AuthRequired,omitempty"`  
  MediaID           *string             `json:"mediaId,omitempty"`          
  MediaFeedType     *MediaFeedType      `json:"mediaFeedType,omitempty"`    
  MlbnAuthRequired  *bool               `json:"mlbnAuthRequired,omitempty"` 
  FoxAuthRequired   *bool               `json:"foxAuthRequired,omitempty"`  
  MediaFeedSubType  *string             `json:"mediaFeedSubType,omitempty"` 
  FreeGame          *bool               `json:"freeGame,omitempty"`         
  ID                int64               `json:"id"`                         
  MediaState        *MediaStateCodeEnum `json:"mediaState,omitempty"`       
  AbcAuthRequired   *bool               `json:"abcAuthRequired,omitempty"`  
  RenditionName     *RenditionNameEnum  `json:"renditionName,omitempty"`    
  Description       *string             `json:"description,omitempty"`      
  Language          *ItemLanguage       `json:"language,omitempty"`         
  Type              *ItemType           `json:"type,omitempty"`             
  Appletv           *Appletv            `json:"appletv,omitempty"`          
}

type Appletv struct {
  Fgow    bool   `json:"fgow"`   
  VideoID string `json:"videoId"`
  PageID  string `json:"pageId"` 
}

type FeaturedMedia struct {
  ID string `json:"id"`
}

type Summary struct {
  HasPreviewArticle  bool `json:"hasPreviewArticle"` 
  HasRecapArticle    bool `json:"hasRecapArticle"`   
  HasWrapArticle     bool `json:"hasWrapArticle"`    
  HasHighlightsVideo bool `json:"hasHighlightsVideo"`
}

type Decisions struct {
  Winner Loser  `json:"winner"`        
  Loser  Loser  `json:"loser"`         
  Save   *Loser `json:"save,omitempty"`
}

type Loser struct {
  ID                 int64               `json:"id"`                          
  FullName           string              `json:"fullName"`                    
  Link               string              `json:"link"`                        
  FirstName          string              `json:"firstName"`                   
  LastName           string              `json:"lastName"`                    
  PrimaryNumber      *string             `json:"primaryNumber,omitempty"`     
  BirthDate          string              `json:"birthDate"`                   
  CurrentAge         int64               `json:"currentAge"`                  
  BirthCity          *string             `json:"birthCity,omitempty"`         
  BirthStateProvince *BirthStateProvince `json:"birthStateProvince,omitempty"`
  BirthCountry       Country             `json:"birthCountry"`                
  Height             Height              `json:"height"`                      
  Weight             int64               `json:"weight"`                      
  Active             bool                `json:"active"`                      
  PrimaryPosition    PrimaryPosition     `json:"primaryPosition"`             
  UseName            string              `json:"useName"`                     
  UseLastName        string              `json:"useLastName"`                 
  MiddleName         *string             `json:"middleName,omitempty"`        
  BoxscoreName       string              `json:"boxscoreName"`                
  Gender             Gender              `json:"gender"`                      
  IsPlayer           bool                `json:"isPlayer"`                    
  IsVerified         bool                `json:"isVerified"`                  
  DraftYear          *int64              `json:"draftYear,omitempty"`         
  Stats              []Stat              `json:"stats,omitempty"`             
  MlbDebutDate       string              `json:"mlbDebutDate"`                
  BatSide            BatSide             `json:"batSide"`                     
  PitchHand          BatSide             `json:"pitchHand"`                   
  NameFirstLast      string              `json:"nameFirstLast"`               
  NameSlug           string              `json:"nameSlug"`                    
  FirstLastName      string              `json:"firstLastName"`               
  LastFirstName      string              `json:"lastFirstName"`               
  LastInitName       string              `json:"lastInitName"`                
  InitLastName       string              `json:"initLastName"`                
  FullFMLName        string              `json:"fullFMLName"`                 
  FullLFMName        string              `json:"fullLFMName"`                 
  StrikeZoneTop      float64             `json:"strikeZoneTop"`               
  StrikeZoneBottom   float64             `json:"strikeZoneBottom"`            
  NickName           *string             `json:"nickName,omitempty"`          
  Pronunciation      *string             `json:"pronunciation,omitempty"`     
  NameMatrilineal    *string             `json:"nameMatrilineal,omitempty"`   
  NameTitle          *Name               `json:"nameTitle,omitempty"`         
  NameSuffix         *Name               `json:"nameSuffix,omitempty"`        
}

type BatSide struct {
  Code        GameTypeEnum       `json:"code"`       
  Description BatSideDescription `json:"description"`
}

type PrimaryPosition struct {
  Code         string              `json:"code"`        
  Name         PrimaryPositionName `json:"name"`        
  Type         PrimaryPositionType `json:"type"`        
  Abbreviation GamedayType         `json:"abbreviation"`
}

type Stat struct {
  Type       Group         `json:"type"`      
  Group      Group         `json:"group"`     
  Exemptions []interface{} `json:"exemptions"`
  Stats      Stats         `json:"stats"`     
}

type Group struct {
  DisplayName DisplayName `json:"displayName"`
}

type Stats struct {
  Note                   *string               `json:"note,omitempty"`                  
  Summary                *string               `json:"summary,omitempty"`               
  GamesPlayed            *int64                `json:"gamesPlayed,omitempty"`           
  GamesStarted           *int64                `json:"gamesStarted,omitempty"`          
  GroundOuts             *int64                `json:"groundOuts,omitempty"`            
  AirOuts                *int64                `json:"airOuts,omitempty"`               
  Runs                   *int64                `json:"runs,omitempty"`                  
  Doubles                *int64                `json:"doubles,omitempty"`               
  Triples                *int64                `json:"triples,omitempty"`               
  HomeRuns               *int64                `json:"homeRuns,omitempty"`              
  StrikeOuts             *int64                `json:"strikeOuts,omitempty"`            
  BaseOnBalls            *int64                `json:"baseOnBalls,omitempty"`           
  IntentionalWalks       *int64                `json:"intentionalWalks,omitempty"`      
  Hits                   *int64                `json:"hits,omitempty"`                  
  HitByPitch             *int64                `json:"hitByPitch,omitempty"`            
  AtBats                 *int64                `json:"atBats,omitempty"`                
  CaughtStealing         *int64                `json:"caughtStealing,omitempty"`        
  StolenBases            *int64                `json:"stolenBases,omitempty"`           
  StolenBasePercentage   *Percentage           `json:"stolenBasePercentage,omitempty"`  
  NumberOfPitches        *int64                `json:"numberOfPitches,omitempty"`       
  Era                    *string               `json:"era,omitempty"`                   
  InningsPitched         *string               `json:"inningsPitched,omitempty"`        
  WINS                   *int64                `json:"wins,omitempty"`                  
  Losses                 *int64                `json:"losses,omitempty"`                
  Saves                  *int64                `json:"saves,omitempty"`                 
  SaveOpportunities      *int64                `json:"saveOpportunities,omitempty"`     
  Holds                  *int64                `json:"holds,omitempty"`                 
  BlownSaves             *int64                `json:"blownSaves,omitempty"`            
  EarnedRuns             *int64                `json:"earnedRuns,omitempty"`            
  BattersFaced           *int64                `json:"battersFaced,omitempty"`          
  Outs                   *int64                `json:"outs,omitempty"`                  
  GamesPitched           *int64                `json:"gamesPitched,omitempty"`          
  CompleteGames          *int64                `json:"completeGames,omitempty"`         
  Shutouts               *int64                `json:"shutouts,omitempty"`              
  PitchesThrown          *int64                `json:"pitchesThrown,omitempty"`         
  Balls                  *int64                `json:"balls,omitempty"`                 
  Strikes                *int64                `json:"strikes,omitempty"`               
  StrikePercentage       *StrikePercentageEnum `json:"strikePercentage,omitempty"`      
  HitBatsmen             *int64                `json:"hitBatsmen,omitempty"`            
  Balks                  *int64                `json:"balks,omitempty"`                 
  WildPitches            *int64                `json:"wildPitches,omitempty"`           
  Pickoffs               *int64                `json:"pickoffs,omitempty"`              
  Rbi                    *int64                `json:"rbi,omitempty"`                   
  GamesFinished          *int64                `json:"gamesFinished,omitempty"`         
  RunsScoredPer9         *string               `json:"runsScoredPer9,omitempty"`        
  HomeRunsPer9           *string               `json:"homeRunsPer9,omitempty"`          
  InheritedRunners       *int64                `json:"inheritedRunners,omitempty"`      
  InheritedRunnersScored *int64                `json:"inheritedRunnersScored,omitempty"`
  CatchersInterference   *int64                `json:"catchersInterference,omitempty"`  
  SacBunts               *int64                `json:"sacBunts,omitempty"`              
  SacFlies               *int64                `json:"sacFlies,omitempty"`              
  PassedBall             *int64                `json:"passedBall,omitempty"`            
  FlyOuts                *int64                `json:"flyOuts,omitempty"`               
  GroundIntoDoublePlay   *int64                `json:"groundIntoDoublePlay,omitempty"`  
  GroundIntoTriplePlay   *int64                `json:"groundIntoTriplePlay,omitempty"`  
  PlateAppearances       *int64                `json:"plateAppearances,omitempty"`      
  LeftOnBase             *int64                `json:"leftOnBase,omitempty"`            
  AtBatsPerHomeRun       *string               `json:"atBatsPerHomeRun,omitempty"`      
  Avg                    *string               `json:"avg,omitempty"`                   
  Obp                    *string               `json:"obp,omitempty"`                   
  Whip                   *string               `json:"whip,omitempty"`                  
  GroundOutsToAirouts    *string               `json:"groundOutsToAirouts,omitempty"`   
  WinPercentage          *Percentage           `json:"winPercentage,omitempty"`         
  PitchesPerInning       *string               `json:"pitchesPerInning,omitempty"`      
  StrikeoutWalkRatio     *string               `json:"strikeoutWalkRatio,omitempty"`    
  StrikeoutsPer9Inn      *string               `json:"strikeoutsPer9Inn,omitempty"`     
  WalksPer9Inn           *string               `json:"walksPer9Inn,omitempty"`          
  HitsPer9Inn            *string               `json:"hitsPer9Inn,omitempty"`           
  Slg                    *string               `json:"slg,omitempty"`                   
  Ops                    *string               `json:"ops,omitempty"`                   
  Babip                  *string               `json:"babip,omitempty"`                 
  TotalBases             *int64                `json:"totalBases,omitempty"`            
}

type Flags struct {
  NoHitter            bool `json:"noHitter"`           
  PerfectGame         bool `json:"perfectGame"`        
  AwayTeamNoHitter    bool `json:"awayTeamNoHitter"`   
  AwayTeamPerfectGame bool `json:"awayTeamPerfectGame"`
  HomeTeamNoHitter    bool `json:"homeTeamNoHitter"`   
  HomeTeamPerfectGame bool `json:"homeTeamPerfectGame"`
}

type HomeRunElement struct {
  Result      Result        `json:"result"`     
  About       About         `json:"about"`      
  Count       Count         `json:"count"`      
  Matchup     Matchup       `json:"matchup"`    
  PitchIndex  []interface{} `json:"pitchIndex"` 
  ActionIndex []interface{} `json:"actionIndex"`
  RunnerIndex []interface{} `json:"runnerIndex"`
  Runners     []interface{} `json:"runners"`    
  PlayEvents  []interface{} `json:"playEvents"` 
}

type About struct {
  HalfInning HalfInning `json:"halfInning"`
  Inning     int64      `json:"inning"`    
}

type Count struct {
  Balls   int64 `json:"balls"`  
  Strikes int64 `json:"strikes"`
  Outs    int64 `json:"outs"`   
}

type Matchup struct {
  Batter              Loser         `json:"batter"`             
  Pitcher             Catcher       `json:"pitcher"`            
  BatterHotColdZones  []interface{} `json:"batterHotColdZones"` 
  PitcherHotColdZones []interface{} `json:"pitcherHotColdZones"`
  Splits              Editorial     `json:"splits"`             
}

type Catcher struct {
  ID       int64  `json:"id"`      
  FullName string `json:"fullName"`
  Link     string `json:"link"`    
}

type Result struct {
  Type        ResultType `json:"type"`       
  Event       Event      `json:"event"`      
  Description string     `json:"description"`
  Rbi         int64      `json:"rbi"`        
  AwayScore   int64      `json:"awayScore"`  
  HomeScore   int64      `json:"homeScore"`  
}

type Linescore struct {
  CurrentInning        *int64                `json:"currentInning,omitempty"`       
  CurrentInningOrdinal *CurrentInningOrdinal `json:"currentInningOrdinal,omitempty"`
  InningState          *InningHalfEnum       `json:"inningState,omitempty"`         
  InningHalf           *InningHalfEnum       `json:"inningHalf,omitempty"`          
  IsTopInning          *bool                 `json:"isTopInning,omitempty"`         
  ScheduledInnings     *int64                `json:"scheduledInnings,omitempty"`    
  Innings              []Inning              `json:"innings,omitempty"`             
  Teams                *LinescoreTeams       `json:"teams,omitempty"`               
  Defense              Defense               `json:"defense"`                       
  Offense              Offense               `json:"offense"`                       
  Balls                *int64                `json:"balls,omitempty"`               
  Strikes              *int64                `json:"strikes,omitempty"`             
  Outs                 *int64                `json:"outs,omitempty"`                
  Note                 *Note                 `json:"note,omitempty"`                
}

type Defense struct {
  Pitcher      *Loser   `json:"pitcher,omitempty"`     
  Catcher      *Catcher `json:"catcher,omitempty"`     
  First        *Catcher `json:"first,omitempty"`       
  Second       *Catcher `json:"second,omitempty"`      
  Third        *Catcher `json:"third,omitempty"`       
  Shortstop    *Catcher `json:"shortstop,omitempty"`   
  Left         *Catcher `json:"left,omitempty"`        
  Center       *Catcher `json:"center,omitempty"`      
  Right        *Catcher `json:"right,omitempty"`       
  BattingOrder *int64   `json:"battingOrder,omitempty"`
}

type Inning struct {
  Num        int64                `json:"num"`       
  OrdinalNum CurrentInningOrdinal `json:"ordinalNum"`
  Home       InningAway           `json:"home"`      
  Away       InningAway           `json:"away"`      
}

type InningAway struct {
  Runs       *int64 `json:"runs,omitempty"`      
  Hits       *int64 `json:"hits,omitempty"`      
  Errors     *int64 `json:"errors,omitempty"`    
  LeftOnBase *int64 `json:"leftOnBase,omitempty"`
  IsWinner   *bool  `json:"isWinner,omitempty"`  
}

type Offense struct {
  Batter       *Loser             `json:"batter,omitempty"`      
  OnDeck       *Loser             `json:"onDeck,omitempty"`      
  InHole       *Loser             `json:"inHole,omitempty"`      
  BattingOrder *int64             `json:"battingOrder,omitempty"`
  Team         *SpringLeagueClass `json:"team,omitempty"`        
  Second       *Loser             `json:"second,omitempty"`      
  Third        *Loser             `json:"third,omitempty"`       
  First        *Loser             `json:"first,omitempty"`       
}

type SpringLeagueClass struct {
  ID           int64             `json:"id"`                    
  Name         *TeamName         `json:"name,omitempty"`        
  Link         TeamLink          `json:"link"`                  
  Abbreviation *TeamAbbreviation `json:"abbreviation,omitempty"`
}

type LinescoreTeams struct {
  Home InningAway `json:"home"`
  Away InningAway `json:"away"`
}

type Review struct {
  HasChallenges bool       `json:"hasChallenges"`
  Away          ReviewAway `json:"away"`         
  Home          ReviewAway `json:"home"`         
}

type ReviewAway struct {
  Used      int64 `json:"used"`     
  Remaining int64 `json:"remaining"`
}

type SeriesStatus struct {
  GameNumber       int64                    `json:"gameNumber"`           
  TotalGames       int64                    `json:"totalGames"`           
  IsTied           bool                     `json:"isTied"`               
  IsOver           bool                     `json:"isOver"`               
  WINS             int64                    `json:"wins"`                 
  Losses           int64                    `json:"losses"`               
  WinningTeam      *IngTeam                 `json:"winningTeam,omitempty"`
  LosingTeam       *IngTeam                 `json:"losingTeam,omitempty"` 
  Description      SeriesDescriptionEnum    `json:"description"`          
  ShortDescription Short                    `json:"shortDescription"`     
  Result           *string                  `json:"result,omitempty"`     
  ShortName        Short                    `json:"shortName"`            
  Abbreviation     SeriesStatusAbbreviation `json:"abbreviation"`         
}

type IngTeam struct {
  SpringLeague  SpringLeagueClass `json:"springLeague"` 
  AllStarStatus DoubleHeader      `json:"allStarStatus"`
  ID            int64             `json:"id"`           
  Name          TeamName          `json:"name"`         
  Link          TeamLink          `json:"link"`         
}

type Status struct {
  AbstractGameState AbstractGameState `json:"abstractGameState"`
  CodedGameState    CodedGameState    `json:"codedGameState"`   
  DetailedState     DetailedState     `json:"detailedState"`    
  StatusCode        CodedGameState    `json:"statusCode"`       
  StartTimeTBD      bool              `json:"startTimeTBD"`     
  AbstractGameCode  AbstractGameCode  `json:"abstractGameCode"` 
  Reason            *Reason           `json:"reason,omitempty"` 
}

type GameTeams struct {
  Home PurpleAway `json:"home"`
  Away PurpleAway `json:"away"`
}

type PurpleAway struct {
  IsContextTeam   bool              `json:"isContextTeam"`            
  IsFollowed      bool              `json:"isFollowed"`               
  IsFavorite      bool              `json:"isFavorite"`               
  LeagueRecord    LeagueRecord      `json:"leagueRecord"`             
  Score           *int64            `json:"score,omitempty"`          
  Team            PurpleTeam        `json:"team"`                     
  IsWinner        *bool             `json:"isWinner,omitempty"`       
  ProbablePitcher *Loser            `json:"probablePitcher,omitempty"`
  SplitSquad      bool              `json:"splitSquad"`               
  SeriesNumber    int64             `json:"seriesNumber"`             
  SpringLeague    SpringLeagueClass `json:"springLeague"`             
}

type LeagueRecord struct {
  WINS   int64  `json:"wins"`  
  Losses int64  `json:"losses"`
  Pct    string `json:"pct"`   
}

type PurpleTeam struct {
  SpringLeague    SpringLeagueClass `json:"springLeague"`         
  AllStarStatus   DoubleHeader      `json:"allStarStatus"`        
  ID              int64             `json:"id"`                   
  Name            TeamName          `json:"name"`                 
  Link            TeamLink          `json:"link"`                 
  Season          int64             `json:"season"`               
  Venue           SpringLeagueClass `json:"venue"`                
  SpringVenue     SpringVenue       `json:"springVenue"`          
  TeamCode        string            `json:"teamCode"`             
  FileCode        string            `json:"fileCode"`             
  Abbreviation    string            `json:"abbreviation"`         
  TeamName        string            `json:"teamName"`             
  LocationName    City              `json:"locationName"`         
  FirstYearOfPlay string            `json:"firstYearOfPlay"`      
  League          SpringLeagueClass `json:"league"`               
  Division        SpringLeagueClass `json:"division"`             
  Sport           SpringLeagueClass `json:"sport"`                
  ShortName       City              `json:"shortName"`            
  FranchiseName   City              `json:"franchiseName"`        
  ClubName        string            `json:"clubName"`             
  Active          bool              `json:"active"`               
  TeamLeaders     []TeamLeader      `json:"teamLeaders,omitempty"`
}

type SpringVenue struct {
  ID   int64           `json:"id"`  
  Link SpringVenueLink `json:"link"`
}

type TeamLeader struct {
  LeaderCategory LeaderCategory    `json:"leaderCategory"`
  Season         string            `json:"season"`        
  GameType       GameType          `json:"gameType"`      
  Leaders        []Leader          `json:"leaders"`       
  StatGroup      DisplayName       `json:"statGroup"`     
  Team           SpringLeagueClass `json:"team"`          
  TotalSplits    int64             `json:"totalSplits"`   
}

type GameType struct {
  ID          GameTypeEnum          `json:"id"`         
  Description SeriesDescriptionEnum `json:"description"`
}

type Leader struct {
  Rank   int64             `json:"rank"`  
  Value  string            `json:"value"` 
  Team   SpringLeagueClass `json:"team"`  
  League SpringLeagueClass `json:"league"`
  Person Person            `json:"person"`
  Sport  SpringLeagueClass `json:"sport"` 
  Season string            `json:"season"`
}

type Person struct {
  ID        int64  `json:"id"`       
  FullName  string `json:"fullName"` 
  Link      string `json:"link"`     
  FirstName string `json:"firstName"`
  LastName  string `json:"lastName"` 
}

type Ticket struct {
  TicketType  TicketType  `json:"ticketType"` 
  TicketLinks TicketLinks `json:"ticketLinks"`
}

type TicketLinks struct {
  Home string `json:"home"`
}

type Venue struct {
  ID       int64    `json:"id"`      
  Name     TeamName `json:"name"`    
  Link     TeamLink `json:"link"`    
  Location Location `json:"location"`
  Active   bool     `json:"active"`  
  Season   string   `json:"season"`  
}

type Location struct {
  Address1           *string             `json:"address1,omitempty"`          
  City               City                `json:"city"`                        
  State              *State              `json:"state,omitempty"`             
  StateAbbrev        *BirthStateProvince `json:"stateAbbrev,omitempty"`       
  PostalCode         *string             `json:"postalCode,omitempty"`        
  DefaultCoordinates *DefaultCoordinates `json:"defaultCoordinates,omitempty"`
  AzimuthAngle       *float64            `json:"azimuthAngle,omitempty"`      
  Elevation          *int64              `json:"elevation,omitempty"`         
  Country            Country             `json:"country"`                     
  Phone              *string             `json:"phone,omitempty"`             
}

type DefaultCoordinates struct {
  Latitude  float64 `json:"latitude"` 
  Longitude float64 `json:"longitude"`
}

type XrefID struct {
  XrefID    string   `json:"xrefId"`             
  XrefType  XrefType `json:"xrefType"`           
  XrefIDTwo *string  `json:"xrefIdTwo,omitempty"`
}

type Headers struct {
  XStitchSHA   string `json:"X-STITCH-SHA"`  
  XStitchCache string `json:"X-STITCH-CACHE"`
  CacheControl string `json:"Cache-Control"` 
  ContentType  string `json:"Content-Type"`  
}

type ReqHeaders struct {
  XStitchPRef string `json:"X-STITCH-P-REF"`
}

type AvailabilityCode string
const (
  LocalInMarket AvailabilityCode = "local_in_market"
  LocalOutOfMarket AvailabilityCode = "local_out_of_market"
  National AvailabilityCode = "national"
)

type AvailabilityText string
const (
  AvailabilityTextLocalInMarket AvailabilityText = "Local (In Market)"
  AvailabilityTextLocalOutOfMarket AvailabilityText = "Local (Out of Market)"
  AvailabilityTextNational AvailabilityText = "National"
)

type HomeAway string
const (
  Away HomeAway = "away"
  Home HomeAway = "home"
)

type BroadcastLanguage string
const (
  Fr BroadcastLanguage = "fr"
  PurpleEn BroadcastLanguage = "en"
  PurpleEs BroadcastLanguage = "es"
)

type MediaStateCodeEnum string
const (
  MediaArchive MediaStateCodeEnum = "MEDIA_ARCHIVE"
  MediaOff MediaStateCodeEnum = "MEDIA_OFF"
  MediaOn MediaStateCodeEnum = "MEDIA_ON"
)

type MediaStateText string
const (
  MediaStateTextMediaArchive MediaStateText = "Media Archive"
  MediaStateTextMediaOff MediaStateText = "Media Off"
  MediaStateTextMediaOn MediaStateText = "Media On"
)

type BroadcastType string
const (
  Am BroadcastType = "AM"
  Fm BroadcastType = "FM"
  Tv BroadcastType = "TV"
)

type Code string
const (
  H Code = "H"
)

type ResolutionFull string
const (
  HighResolution ResolutionFull = "High Resolution"
)

type ResolutionShort string
const (
  HD ResolutionShort = "HD"
)

type ItemLanguage string
const (
  En ItemLanguage = "EN"
  Es ItemLanguage = "ES"
  FluffyEn ItemLanguage = "en"
  FluffyEs ItemLanguage = "es"
)

type MediaFeedType string
const (
  MediaFeedTypeAWAY MediaFeedType = "AWAY"
  MediaFeedTypeHOME MediaFeedType = "HOME"
  MediaFeedTypeNATIONAL MediaFeedType = "NATIONAL"
)

type RenditionNameEnum string
const (
  English RenditionNameEnum = "English"
  EnglishRadio RenditionNameEnum = "English Radio"
  RadioEspañola RenditionNameEnum = "Radio Española"
  RenditionName RenditionNameEnum = ""
)

type ItemType string
const (
  Type ItemType = ""
  TypeAWAY ItemType = "AWAY"
  TypeHOME ItemType = "HOME"
)

type Title string
const (
  Audio Title = "Audio"
  MLBTVAudio Title = "MLBTV-Audio"
  Mlbtv Title = "MLBTV"
)

type DayNight string
const (
  Day DayNight = "day"
  Night DayNight = "night"
)

type GameTypeEnum string
const (
  GameTypeL GameTypeEnum = "L"
  GameTypeS GameTypeEnum = "S"
  R GameTypeEnum = "R"
)

type BatSideDescription string
const (
  Left BatSideDescription = "Left"
  Right BatSideDescription = "Right"
  Switch BatSideDescription = "Switch"
)

type Country string
const (
  Aruba Country = "Aruba"
  Australia Country = "Australia"
  Bahamas Country = "Bahamas"
  Brazil Country = "Brazil"
  Canada Country = "Canada"
  Colombia Country = "Colombia"
  Cuba Country = "Cuba"
  Curacao Country = "Curacao"
  DominicanRepublic Country = "Dominican Republic"
  Germany Country = "Germany"
  Honduras Country = "Honduras"
  Japan Country = "Japan"
  Mexico Country = "Mexico"
  Nicaragua Country = "Nicaragua"
  Panama Country = "Panama"
  PanamaCanalZone Country = "Panama Canal Zone"
  Peru Country = "Peru"
  Portugal Country = "Portugal"
  PuertoRico Country = "Puerto Rico"
  RepublicOfKorea Country = "Republic of Korea"
  SouthAfrica Country = "South Africa"
  Usa Country = "USA"
  Venezuela Country = "Venezuela"
)

type BirthStateProvince string
const (
  Ab BirthStateProvince = "AB"
  Ar BirthStateProvince = "AR"
  Az BirthStateProvince = "AZ"
  BajaCalifornia BirthStateProvince = "Baja California"
  Bc BirthStateProvince = "BC"
  BirthStateProvinceAL BirthStateProvince = "AL"
  CA BirthStateProvince = "CA"
  CT BirthStateProvince = "CT"
  Co BirthStateProvince = "CO"
  Dc BirthStateProvince = "DC"
  De BirthStateProvince = "DE"
  FL BirthStateProvince = "FL"
  Ga BirthStateProvince = "GA"
  Hi BirthStateProvince = "HI"
  IL BirthStateProvince = "IL"
  Ia BirthStateProvince = "IA"
  In BirthStateProvince = "IN"
  Ks BirthStateProvince = "KS"
  Ky BirthStateProvince = "KY"
  La BirthStateProvince = "LA"
  MS BirthStateProvince = "MS"
  Ma BirthStateProvince = "MA"
  Md BirthStateProvince = "MD"
  Mi BirthStateProvince = "MI"
  Mn BirthStateProvince = "MN"
  Mo BirthStateProvince = "MO"
  Nayarit BirthStateProvince = "Nayarit"
  Nc BirthStateProvince = "NC"
  Nd BirthStateProvince = "ND"
  Ne BirthStateProvince = "NE"
  Nh BirthStateProvince = "NH"
  Nj BirthStateProvince = "NJ"
  Nm BirthStateProvince = "NM"
  Nv BirthStateProvince = "NV"
  Ny BirthStateProvince = "NY"
  Oh BirthStateProvince = "OH"
  Ok BirthStateProvince = "OK"
  On BirthStateProvince = "ON"
  Or BirthStateProvince = "OR"
  Pa BirthStateProvince = "PA"
  Qc BirthStateProvince = "QC"
  Sc BirthStateProvince = "SC"
  Sinaloa BirthStateProvince = "Sinaloa"
  So BirthStateProvince = "SO"
  Sonora BirthStateProvince = "Sonora"
  Tn BirthStateProvince = "TN"
  Tx BirthStateProvince = "TX"
  Ut BirthStateProvince = "UT"
  Va BirthStateProvince = "VA"
  Wa BirthStateProvince = "WA"
  Wi BirthStateProvince = "WI"
  Wy BirthStateProvince = "WY"
)

type Gender string
const (
  M Gender = "M"
)

type Height string
const (
  The510 Height = "5' 10\""
  The511 Height = "5' 11\""
  The55 Height = "5' 5\""
  The56 Height = "5' 6\""
  The57 Height = "5' 7\""
  The58 Height = "5' 8\""
  The59 Height = "5' 9\""
  The60 Height = "6' 0\""
  The61 Height = "6' 1\""
  The62 Height = "6' 2\""
  The63 Height = "6' 3\""
  The64 Height = "6' 4\""
  The65 Height = "6' 5\""
  The66 Height = "6' 6\""
  The67 Height = "6' 7\""
  The68 Height = "6' 8\""
  The69 Height = "6' 9\""
)

type Name string
const (
  Ii Name = "II"
  Iii Name = "III"
  Jr Name = "Jr."
)

type GamedayType string
const (
  C GamedayType = "C"
  CF GamedayType = "CF"
  Dh GamedayType = "DH"
  GamedayTypeP GamedayType = "P"
  LF GamedayType = "LF"
  PR GamedayType = "PR"
  Ph GamedayType = "PH"
  RF GamedayType = "RF"
  Ss GamedayType = "SS"
  The1B GamedayType = "1B"
  The2B GamedayType = "2B"
  The3B GamedayType = "3B"
)

type PrimaryPositionName string
const (
  DesignatedHitter PrimaryPositionName = "Designated Hitter"
  FirstBase PrimaryPositionName = "First Base"
  NameCatcher PrimaryPositionName = "Catcher"
  NameOutfielder PrimaryPositionName = "Outfielder"
  NamePitcher PrimaryPositionName = "Pitcher"
  PinchHitter PrimaryPositionName = "Pinch Hitter"
  PinchRunner PrimaryPositionName = "Pinch Runner"
  SecondBase PrimaryPositionName = "Second Base"
  Shortstop PrimaryPositionName = "Shortstop"
  ThirdBase PrimaryPositionName = "Third Base"
)

type PrimaryPositionType string
const (
  Hitter PrimaryPositionType = "Hitter"
  Infielder PrimaryPositionType = "Infielder"
  Runner PrimaryPositionType = "Runner"
  TypeCatcher PrimaryPositionType = "Catcher"
  TypeOutfielder PrimaryPositionType = "Outfielder"
  TypePitcher PrimaryPositionType = "Pitcher"
)

type DisplayName string
const (
  GameLog DisplayName = "gameLog"
  Hitting DisplayName = "hitting"
  Pitching DisplayName = "pitching"
  StatsSingleSeason DisplayName = "statsSingleSeason"
)

type Percentage string
const (
  Empty Percentage = ".---"
  Percentage000 Percentage = ".000"
  Percentage1000 Percentage = "1.000"
  Percentage250 Percentage = ".250"
  Percentage400 Percentage = ".400"
  Percentage500 Percentage = ".500"
  Percentage600 Percentage = ".600"
  Percentage700 Percentage = ".700"
  Percentage750 Percentage = ".750"
  Percentage800 Percentage = ".800"
  Percentage900 Percentage = ".900"
  The200 Percentage = ".200"
  The333 Percentage = ".333"
  The444 Percentage = ".444"
  The571 Percentage = ".571"
  The667 Percentage = ".667"
  The714 Percentage = ".714"
  The727 Percentage = ".727"
  The778 Percentage = ".778"
  The833 Percentage = ".833"
  The857 Percentage = ".857"
  The875 Percentage = ".875"
  The889 Percentage = ".889"
)

type StrikePercentageEnum string
const (
  StrikePercentage StrikePercentageEnum = "-.--"
  StrikePercentage000 StrikePercentageEnum = ".000"
  StrikePercentage1000 StrikePercentageEnum = "1.000"
  StrikePercentage250 StrikePercentageEnum = ".250"
  StrikePercentage400 StrikePercentageEnum = ".400"
  StrikePercentage500 StrikePercentageEnum = ".500"
  StrikePercentage510 StrikePercentageEnum = ".510"
  StrikePercentage600 StrikePercentageEnum = ".600"
  StrikePercentage700 StrikePercentageEnum = ".700"
  StrikePercentage750 StrikePercentageEnum = ".750"
  StrikePercentage800 StrikePercentageEnum = ".800"
  StrikePercentage900 StrikePercentageEnum = ".900"
  The290 StrikePercentageEnum = ".290"
  The350 StrikePercentageEnum = ".350"
  The390 StrikePercentageEnum = ".390"
  The410 StrikePercentageEnum = ".410"
  The420 StrikePercentageEnum = ".420"
  The430 StrikePercentageEnum = ".430"
  The450 StrikePercentageEnum = ".450"
  The460 StrikePercentageEnum = ".460"
  The470 StrikePercentageEnum = ".470"
  The480 StrikePercentageEnum = ".480"
  The490 StrikePercentageEnum = ".490"
  The520 StrikePercentageEnum = ".520"
  The530 StrikePercentageEnum = ".530"
  The540 StrikePercentageEnum = ".540"
  The550 StrikePercentageEnum = ".550"
  The560 StrikePercentageEnum = ".560"
  The570 StrikePercentageEnum = ".570"
  The580 StrikePercentageEnum = ".580"
  The590 StrikePercentageEnum = ".590"
  The610 StrikePercentageEnum = ".610"
  The620 StrikePercentageEnum = ".620"
  The630 StrikePercentageEnum = ".630"
  The640 StrikePercentageEnum = ".640"
  The650 StrikePercentageEnum = ".650"
  The660 StrikePercentageEnum = ".660"
  The670 StrikePercentageEnum = ".670"
  The680 StrikePercentageEnum = ".680"
  The690 StrikePercentageEnum = ".690"
  The710 StrikePercentageEnum = ".710"
  The720 StrikePercentageEnum = ".720"
  The730 StrikePercentageEnum = ".730"
  The740 StrikePercentageEnum = ".740"
  The760 StrikePercentageEnum = ".760"
  The770 StrikePercentageEnum = ".770"
  The780 StrikePercentageEnum = ".780"
  The790 StrikePercentageEnum = ".790"
  The810 StrikePercentageEnum = ".810"
  The820 StrikePercentageEnum = ".820"
  The830 StrikePercentageEnum = ".830"
  The840 StrikePercentageEnum = ".840"
  The850 StrikePercentageEnum = ".850"
  The860 StrikePercentageEnum = ".860"
  The870 StrikePercentageEnum = ".870"
  The880 StrikePercentageEnum = ".880"
  The890 StrikePercentageEnum = ".890"
)

type DoubleHeader string
const (
  DoubleHeaderS DoubleHeader = "S"
  N DoubleHeader = "N"
  Y DoubleHeader = "Y"
)

type HalfInning string
const (
  Bottom HalfInning = "bottom"
  Top HalfInning = "top"
)

type Event string
const (
  HomeRun Event = "Home Run"
)

type ResultType string
const (
  AtBat ResultType = "atBat"
)

type IfNecessaryDescription string
const (
  NormalGame IfNecessaryDescription = "Normal Game"
)

type CurrentInningOrdinal string
const (
  The10Th CurrentInningOrdinal = "10th"
  The11Th CurrentInningOrdinal = "11th"
  The12Th CurrentInningOrdinal = "12th"
  The13Th CurrentInningOrdinal = "13th"
  The1St CurrentInningOrdinal = "1st"
  The2Nd CurrentInningOrdinal = "2nd"
  The3RD CurrentInningOrdinal = "3rd"
  The4Th CurrentInningOrdinal = "4th"
  The5Th CurrentInningOrdinal = "5th"
  The6Th CurrentInningOrdinal = "6th"
  The7Th CurrentInningOrdinal = "7th"
  The8Th CurrentInningOrdinal = "8th"
  The9Th CurrentInningOrdinal = "9th"
)

type InningHalfEnum string
const (
  InningBottom InningHalfEnum = "Bottom"
  InningTop InningHalfEnum = "Top"
)

type Note string
const (
  GameCalledRainAfterTheBottomOfThe8Th Note = "Game called (rain) after the bottom of the 8th."
  NoneOutWhenWinningRunScored Note = "None out when winning run scored."
  OneOutWhenWinningRunScored Note = "One out when winning run scored."
  TwoOutWhenWinningRunScored Note = "Two out when winning run scored."
)

type TeamAbbreviation string
const (
  Cl TeamAbbreviation = "CL"
  Gl TeamAbbreviation = "GL"
  Mlb TeamAbbreviation = "MLB"
)

type TeamLink string
const (
  APIV1Divisions200 TeamLink = "/api/v1/divisions/200"
  APIV1Divisions201 TeamLink = "/api/v1/divisions/201"
  APIV1Divisions202 TeamLink = "/api/v1/divisions/202"
  APIV1Divisions203 TeamLink = "/api/v1/divisions/203"
  APIV1Divisions204 TeamLink = "/api/v1/divisions/204"
  APIV1Divisions205 TeamLink = "/api/v1/divisions/205"
  APIV1League103 TeamLink = "/api/v1/league/103"
  APIV1League104 TeamLink = "/api/v1/league/104"
  APIV1League114 TeamLink = "/api/v1/league/114"
  APIV1League115 TeamLink = "/api/v1/league/115"
  APIV1Sports1 TeamLink = "/api/v1/sports/1"
  APIV1Teams108 TeamLink = "/api/v1/teams/108"
  APIV1Teams109 TeamLink = "/api/v1/teams/109"
  APIV1Teams110 TeamLink = "/api/v1/teams/110"
  APIV1Teams111 TeamLink = "/api/v1/teams/111"
  APIV1Teams112 TeamLink = "/api/v1/teams/112"
  APIV1Teams113 TeamLink = "/api/v1/teams/113"
  APIV1Teams114 TeamLink = "/api/v1/teams/114"
  APIV1Teams115 TeamLink = "/api/v1/teams/115"
  APIV1Teams116 TeamLink = "/api/v1/teams/116"
  APIV1Teams117 TeamLink = "/api/v1/teams/117"
  APIV1Teams118 TeamLink = "/api/v1/teams/118"
  APIV1Teams119 TeamLink = "/api/v1/teams/119"
  APIV1Teams120 TeamLink = "/api/v1/teams/120"
  APIV1Teams121 TeamLink = "/api/v1/teams/121"
  APIV1Teams133 TeamLink = "/api/v1/teams/133"
  APIV1Teams134 TeamLink = "/api/v1/teams/134"
  APIV1Teams135 TeamLink = "/api/v1/teams/135"
  APIV1Teams136 TeamLink = "/api/v1/teams/136"
  APIV1Teams137 TeamLink = "/api/v1/teams/137"
  APIV1Teams138 TeamLink = "/api/v1/teams/138"
  APIV1Teams139 TeamLink = "/api/v1/teams/139"
  APIV1Teams140 TeamLink = "/api/v1/teams/140"
  APIV1Teams141 TeamLink = "/api/v1/teams/141"
  APIV1Teams142 TeamLink = "/api/v1/teams/142"
  APIV1Teams143 TeamLink = "/api/v1/teams/143"
  APIV1Teams144 TeamLink = "/api/v1/teams/144"
  APIV1Teams145 TeamLink = "/api/v1/teams/145"
  APIV1Teams146 TeamLink = "/api/v1/teams/146"
  APIV1Teams147 TeamLink = "/api/v1/teams/147"
  APIV1Teams158 TeamLink = "/api/v1/teams/158"
  APIV1Venues1 TeamLink = "/api/v1/venues/1"
  APIV1Venues10 TeamLink = "/api/v1/venues/10"
  APIV1Venues12 TeamLink = "/api/v1/venues/12"
  APIV1Venues14 TeamLink = "/api/v1/venues/14"
  APIV1Venues15 TeamLink = "/api/v1/venues/15"
  APIV1Venues17 TeamLink = "/api/v1/venues/17"
  APIV1Venues19 TeamLink = "/api/v1/venues/19"
  APIV1Venues2 TeamLink = "/api/v1/venues/2"
  APIV1Venues22 TeamLink = "/api/v1/venues/22"
  APIV1Venues2392 TeamLink = "/api/v1/venues/2392"
  APIV1Venues2394 TeamLink = "/api/v1/venues/2394"
  APIV1Venues2395 TeamLink = "/api/v1/venues/2395"
  APIV1Venues2602 TeamLink = "/api/v1/venues/2602"
  APIV1Venues2680 TeamLink = "/api/v1/venues/2680"
  APIV1Venues2681 TeamLink = "/api/v1/venues/2681"
  APIV1Venues2889 TeamLink = "/api/v1/venues/2889"
  APIV1Venues3 TeamLink = "/api/v1/venues/3"
  APIV1Venues31 TeamLink = "/api/v1/venues/31"
  APIV1Venues32 TeamLink = "/api/v1/venues/32"
  APIV1Venues3289 TeamLink = "/api/v1/venues/3289"
  APIV1Venues3309 TeamLink = "/api/v1/venues/3309"
  APIV1Venues3312 TeamLink = "/api/v1/venues/3312"
  APIV1Venues3313 TeamLink = "/api/v1/venues/3313"
  APIV1Venues4 TeamLink = "/api/v1/venues/4"
  APIV1Venues4169 TeamLink = "/api/v1/venues/4169"
  APIV1Venues4705 TeamLink = "/api/v1/venues/4705"
  APIV1Venues5 TeamLink = "/api/v1/venues/5"
  APIV1Venues5325 TeamLink = "/api/v1/venues/5325"
  APIV1Venues5340 TeamLink = "/api/v1/venues/5340"
  APIV1Venues680 TeamLink = "/api/v1/venues/680"
  APIV1Venues7 TeamLink = "/api/v1/venues/7"
)

type TeamName string
const (
  AmericanFamilyField TeamName = "American Family Field"
  AmericanLeague TeamName = "American League"
  AmericanLeagueCentral TeamName = "American League Central"
  AmericanLeagueEast TeamName = "American League East"
  AmericanLeagueWest TeamName = "American League West"
  AngelStadium TeamName = "Angel Stadium"
  ArizonaDiamondbacks TeamName = "Arizona Diamondbacks"
  AtlantaBraves TeamName = "Atlanta Braves"
  BaltimoreOrioles TeamName = "Baltimore Orioles"
  BostonRedSox TeamName = "Boston Red Sox"
  BuschStadium TeamName = "Busch Stadium"
  CactusLeague TeamName = "Cactus League"
  ChaseField TeamName = "Chase Field"
  ChicagoCubs TeamName = "Chicago Cubs"
  ChicagoWhiteSox TeamName = "Chicago White Sox"
  CincinnatiReds TeamName = "Cincinnati Reds"
  CitiField TeamName = "Citi Field"
  CitizensBankPark TeamName = "Citizens Bank Park"
  ClevelandGuardians TeamName = "Cleveland Guardians"
  ColoradoRockies TeamName = "Colorado Rockies"
  ComericaPark TeamName = "Comerica Park"
  CoorsField TeamName = "Coors Field"
  DetroitTigers TeamName = "Detroit Tigers"
  DodgerStadium TeamName = "Dodger Stadium"
  EstadioAlfredoHarpHelu TeamName = "Estadio Alfredo Harp Helu"
  FenwayPark TeamName = "Fenway Park"
  GlobeLifeField TeamName = "Globe Life Field"
  GrapefruitLeague TeamName = "Grapefruit League"
  GreatAmericanBallPark TeamName = "Great American Ball Park"
  GuaranteedRateField TeamName = "Guaranteed Rate Field"
  HoustonAstros TeamName = "Houston Astros"
  KansasCityRoyals TeamName = "Kansas City Royals"
  KauffmanStadium TeamName = "Kauffman Stadium"
  LoanDepotPark TeamName = "loanDepot park"
  LosAngelesAngels TeamName = "Los Angeles Angels"
  LosAngelesDodgers TeamName = "Los Angeles Dodgers"
  MajorLeagueBaseball TeamName = "Major League Baseball"
  MiamiMarlins TeamName = "Miami Marlins"
  MilwaukeeBrewers TeamName = "Milwaukee Brewers"
  MinnesotaTwins TeamName = "Minnesota Twins"
  MinuteMaidPark TeamName = "Minute Maid Park"
  NameAL TeamName = "AL"
  NationalLeague TeamName = "National League"
  NationalLeagueCentral TeamName = "National League Central"
  NationalLeagueEast TeamName = "National League East"
  NationalLeagueWest TeamName = "National League West"
  NationalsPark TeamName = "Nationals Park"
  NewYorkMets TeamName = "New York Mets"
  NewYorkYankees TeamName = "New York Yankees"
  Nl TeamName = "NL"
  OaklandAthletics TeamName = "Oakland Athletics"
  OaklandColiseum TeamName = "Oakland Coliseum"
  OraclePark TeamName = "Oracle Park"
  OrioleParkAtCamdenYards TeamName = "Oriole Park at Camden Yards"
  PNCPark TeamName = "PNC Park"
  PetcoPark TeamName = "Petco Park"
  PhiladelphiaPhillies TeamName = "Philadelphia Phillies"
  PittsburghPirates TeamName = "Pittsburgh Pirates"
  ProgressiveField TeamName = "Progressive Field"
  RogersCentre TeamName = "Rogers Centre"
  SANDiegoPadres TeamName = "San Diego Padres"
  SANFranciscoGiants TeamName = "San Francisco Giants"
  SeattleMariners TeamName = "Seattle Mariners"
  StLouisCardinals TeamName = "St. Louis Cardinals"
  TMobilePark TeamName = "T-Mobile Park"
  TampaBayRays TeamName = "Tampa Bay Rays"
  TargetField TeamName = "Target Field"
  TexasRangers TeamName = "Texas Rangers"
  TorontoBlueJays TeamName = "Toronto Blue Jays"
  TropicanaField TeamName = "Tropicana Field"
  TruistPark TeamName = "Truist Park"
  WashingtonNationals TeamName = "Washington Nationals"
  WrigleyField TeamName = "Wrigley Field"
  YankeeStadium TeamName = "Yankee Stadium"
)

type SeriesDescriptionEnum string
const (
  RegularSeason SeriesDescriptionEnum = "Regular Season"
)

type SeriesStatusAbbreviation string
const (
  Rs SeriesStatusAbbreviation = "RS"
)

type Short string
const (
  Season Short = "Season"
)

type AbstractGameCode string
const (
  AbstractGameCodeF AbstractGameCode = "F"
  AbstractGameCodeL AbstractGameCode = "L"
  AbstractGameCodeP AbstractGameCode = "P"
)

type AbstractGameState string
const (
  AbstractGameStateFinal AbstractGameState = "Final"
  Live AbstractGameState = "Live"
  Preview AbstractGameState = "Preview"
)

type CodedGameState string
const (
  CodedGameStateF CodedGameState = "F"
  CodedGameStateFR CodedGameState = "FR"
  CodedGameStateP CodedGameState = "P"
  CodedGameStateS CodedGameState = "S"
  D CodedGameState = "D"
  DR CodedGameState = "DR"
  Di CodedGameState = "DI"
  I CodedGameState = "I"
)

type DetailedState string
const (
  CompletedEarlyRain DetailedState = "Completed Early: Rain"
  DetailedStateFinal DetailedState = "Final"
  InProgress DetailedState = "In Progress"
  Postponed DetailedState = "Postponed"
  PreGame DetailedState = "Pre-Game"
  Scheduled DetailedState = "Scheduled"
)

type Reason string
const (
  InclementWeather Reason = "Inclement Weather"
  Rain Reason = "Rain"
)

type City string
const (
  Anaheim City = "Anaheim"
  Arlington City = "Arlington"
  Atlanta City = "Atlanta"
  Baltimore City = "Baltimore"
  Boston City = "Boston"
  Bronx City = "Bronx"
  ChiCubs City = "Chi Cubs"
  ChiWhiteSox City = "Chi White Sox"
  Chicago City = "Chicago"
  Cincinnati City = "Cincinnati"
  CityArizona City = "Arizona"
  CityColorado City = "Colorado"
  CityMinnesota City = "Minnesota"
  CityNewYork City = "New York"
  CityTexas City = "Texas"
  CityWashington City = "Washington"
  Cleveland City = "Cleveland"
  Denver City = "Denver"
  Detroit City = "Detroit"
  Flushing City = "Flushing"
  Houston City = "Houston"
  KansasCity City = "Kansas City"
  LAAngels City = "LA Angels"
  LADodgers City = "LA Dodgers"
  LosAngeles City = "Los Angeles"
  MexicoCity City = "Mexico City"
  Miami City = "Miami"
  Milwaukee City = "Milwaukee"
  Minneapolis City = "Minneapolis"
  NYMets City = "NY Mets"
  NYYankees City = "NY Yankees"
  Oakland City = "Oakland"
  Philadelphia City = "Philadelphia"
  Phoenix City = "Phoenix"
  Pittsburgh City = "Pittsburgh"
  SANDiego City = "San Diego"
  SANFrancisco City = "San Francisco"
  Seattle City = "Seattle"
  StLouis City = "St. Louis"
  StPetersburg City = "St. Petersburg"
  TampaBay City = "Tampa Bay"
  Toronto City = "Toronto"
)

type SpringVenueLink string
const (
  APIV1Venues2500 SpringVenueLink = "/api/v1/venues/2500"
  APIV1Venues2507 SpringVenueLink = "/api/v1/venues/2507"
  APIV1Venues2508 SpringVenueLink = "/api/v1/venues/2508"
  APIV1Venues2511 SpringVenueLink = "/api/v1/venues/2511"
  APIV1Venues2518 SpringVenueLink = "/api/v1/venues/2518"
  APIV1Venues2520 SpringVenueLink = "/api/v1/venues/2520"
  APIV1Venues2523 SpringVenueLink = "/api/v1/venues/2523"
  APIV1Venues2526 SpringVenueLink = "/api/v1/venues/2526"
  APIV1Venues2530 SpringVenueLink = "/api/v1/venues/2530"
  APIV1Venues2532 SpringVenueLink = "/api/v1/venues/2532"
  APIV1Venues2534 SpringVenueLink = "/api/v1/venues/2534"
  APIV1Venues2536 SpringVenueLink = "/api/v1/venues/2536"
  APIV1Venues2603 SpringVenueLink = "/api/v1/venues/2603"
  APIV1Venues2700 SpringVenueLink = "/api/v1/venues/2700"
  APIV1Venues2856 SpringVenueLink = "/api/v1/venues/2856"
  APIV1Venues2862 SpringVenueLink = "/api/v1/venues/2862"
  APIV1Venues3809 SpringVenueLink = "/api/v1/venues/3809"
  APIV1Venues3834 SpringVenueLink = "/api/v1/venues/3834"
  APIV1Venues4249 SpringVenueLink = "/api/v1/venues/4249"
  APIV1Venues4309 SpringVenueLink = "/api/v1/venues/4309"
  APIV1Venues4629 SpringVenueLink = "/api/v1/venues/4629"
  APIV1Venues5000 SpringVenueLink = "/api/v1/venues/5000"
  APIV1Venues5380 SpringVenueLink = "/api/v1/venues/5380"
)

type LeaderCategory string
const (
  BattingAverage LeaderCategory = "battingAverage"
  HomeRuns LeaderCategory = "homeRuns"
  RunsBattedIn LeaderCategory = "runsBattedIn"
)

type TicketType string
const (
  Mobile TicketType = "mobile"
  Wired TicketType = "wired"
)

type State string
const (
  California State = "California"
  DistrictOfColumbia State = "District of Columbia"
  Florida State = "Florida"
  Georgia State = "Georgia"
  Illinois State = "Illinois"
  Maryland State = "Maryland"
  Massachusetts State = "Massachusetts"
  Michigan State = "Michigan"
  Missouri State = "Missouri"
  Ohio State = "Ohio"
  Ontario State = "Ontario"
  Pennsylvania State = "Pennsylvania"
  StateArizona State = "Arizona"
  StateColorado State = "Colorado"
  StateMinnesota State = "Minnesota"
  StateNewYork State = "New York"
  StateTexas State = "Texas"
  StateWashington State = "Washington"
  Wisconsin State = "Wisconsin"
)

type XrefType string
const (
  BAMAppletvFree XrefType = "bam_appletv_free"
  BAMMlbtvFree XrefType = "bam_mlbtv_free"
  Bis XrefType = "bis"
  PostgameShowAway XrefType = "postgame_show_away"
  PostgameShowHome XrefType = "postgame_show_home"
  PregameShowAway XrefType = "pregame_show_away"
  PregameShowHome XrefType = "pregame_show_home"
)
