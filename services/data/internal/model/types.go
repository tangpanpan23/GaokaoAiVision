package model

// AdmissionScore 录取分数表
type AdmissionScore struct {
	Id          int64  `db:"id"`
	Year        int    `db:"year"`
	Province    string `db:"province"`
	CollegeName string `db:"college_name"`
	MajorName   string `db:"major_name"`
	ScoreType   int    `db:"score_type"` // 1:文科, 2:理科, 3:综合改革
	Batch       string `db:"batch"`
	MinScore    int    `db:"min_score"`
	MinRank     int    `db:"min_rank"`
	DataSource  string `db:"data_source"`
	CreatedAt   string `db:"created_at"`
}

// QueryScoreReq 查询录取分数请求
type QueryScoreReq struct {
	Year      int    `json:"year"`
	Province  string `json:"province"`
	College   string `json:"college"`
	Major     string `json:"major"`
	ScoreType int    `json:"score_type"`
	Batch     string `json:"batch"`
	MinScore  int    `json:"min_score"`
	MaxScore  int    `json:"max_score"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

// QueryScoreTrendReq 查询分数趋势请求
type QueryScoreTrendReq struct {
	College   string `json:"college"`
	Major     string `json:"major"`
	Province  string `json:"province"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
}

// GetMatchingCollegesReq 获取匹配院校请求
type GetMatchingCollegesReq struct {
	Score     int    `json:"score"`
	Rank      int    `json:"rank"`
	Province  string `json:"province"`
	Subjects  string `json:"subjects"`
	ScoreType int    `json:"score_type"`
	Batch     string `json:"batch"`
	Limit     int    `json:"limit"`
}

// ScoreTrend 分数趋势
type ScoreTrend struct {
	Year       int     `db:"year"`
	College    string  `db:"college_name"`
	Major      string  `db:"major_name"`
	MinScore   int     `db:"min_score"`
	MinRank    int     `db:"min_rank"`
	AvgScore   float64 `db:"avg_score"`
	Batch      string  `db:"batch"`
}

// MatchingCollege 匹配的院校信息
type MatchingCollege struct {
	CollegeName string  `db:"college_name"`
	MajorName   string  `db:"major_name"`
	MinScore    int     `db:"min_score"`
	MinRank     int     `db:"min_rank"`
	Batch       string  `db:"batch"`
	Year        int     `db:"year"`
	Probability float64 `db:"probability"`
}
