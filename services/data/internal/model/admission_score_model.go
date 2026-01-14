package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdmissionScoreModel = (*customAdmissionScoreModel)(nil)

type (
	// AdmissionScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdmissionScoreModel.
	AdmissionScoreModel interface {
		admissionScoreModel
		FindByCondition(req *QueryScoreReq) ([]AdmissionScore, error)
		FindTrends(req *QueryScoreTrendReq) ([]ScoreTrend, error)
		FindMatchingColleges(req *GetMatchingCollegesReq) ([]MatchingCollege, error)
	}

	customAdmissionScoreModel struct {
		*defaultAdmissionScoreModel
	}
)

// NewAdmissionScoreModel returns a model for the database table.
func NewAdmissionScoreModel(host string, port int, user, password, database, charset string) AdmissionScoreModel {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		user, password, host, port, database, charset)

	conn := sqlx.NewMysql(dsn)
	return &customAdmissionScoreModel{
		defaultAdmissionScoreModel: newAdmissionScoreModel(conn),
	}
}

// FindByCondition 根据条件查询录取分数
func (m *customAdmissionScoreModel) FindByCondition(req *QueryScoreReq) ([]AdmissionScore, error) {
	var conditions []string
	var args []interface{}

	if req.Year > 0 {
		conditions = append(conditions, "year = ?")
		args = append(args, req.Year)
	}
	if req.Province != "" {
		conditions = append(conditions, "province = ?")
		args = append(args, req.Province)
	}
	if req.College != "" {
		conditions = append(conditions, "college_name LIKE ?")
		args = append(args, "%"+req.College+"%")
	}
	if req.Major != "" {
		conditions = append(conditions, "major_name LIKE ?")
		args = append(args, "%"+req.Major+"%")
	}
	if req.ScoreType > 0 {
		conditions = append(conditions, "score_type = ?")
		args = append(args, req.ScoreType)
	}
	if req.Batch != "" {
		conditions = append(conditions, "batch = ?")
		args = append(args, req.Batch)
	}
	if req.MinScore > 0 {
		conditions = append(conditions, "min_score >= ?")
		args = append(args, req.MinScore)
	}
	if req.MaxScore > 0 {
		conditions = append(conditions, "min_score <= ?")
		args = append(args, req.MaxScore)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`SELECT %s FROM %s %s ORDER BY min_score DESC LIMIT ? OFFSET ?`,
		admissionScoreRows, m.table, whereClause)

	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (req.Page - 1) * pageSize

	args = append(args, pageSize, offset)

	var resp []AdmissionScore
	err := m.QueryRowsNoCache(&resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return []AdmissionScore{}, nil
	default:
		return nil, err
	}
}

// FindTrends 查询分数趋势
func (m *customAdmissionScoreModel) FindTrends(req *QueryScoreTrendReq) ([]ScoreTrend, error) {
	query := `SELECT year, college_name, major_name, MIN(min_score) as min_score,
		MIN(min_rank) as min_rank, AVG(min_score) as avg_score, batch
		FROM admission_score
		WHERE college_name = ? AND province = ?`

	args := []interface{}{req.College, req.Province}

	if req.Major != "" {
		query += " AND major_name LIKE ?"
		args = append(args, "%"+req.Major+"%")
	}

	if req.StartYear > 0 {
		query += " AND year >= ?"
		args = append(args, req.StartYear)
	}

	if req.EndYear > 0 {
		query += " AND year <= ?"
		args = append(args, req.EndYear)
	}

	query += " GROUP BY year, college_name, major_name, batch ORDER BY year DESC"

	var resp []ScoreTrend
	err := m.QueryRowsNoCache(&resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return []ScoreTrend{}, nil
	default:
		return nil, err
	}
}

// FindMatchingColleges 根据分数和位次获取匹配的院校
func (m *customAdmissionScoreModel) FindMatchingColleges(req *GetMatchingCollegesReq) ([]MatchingCollege, error) {
	query := `SELECT college_name, major_name, min_score, min_rank, batch, year,
		CASE
			WHEN min_rank <= ? THEN 0.9
			WHEN min_rank <= ? THEN 0.7
			WHEN min_rank <= ? THEN 0.5
			ELSE 0.3
		END as probability
		FROM admission_score
		WHERE province = ? AND score_type = ? AND min_rank >= ?
		ORDER BY probability DESC, min_rank ASC LIMIT ?`

	// 计算不同的位次区间
	rank1 := req.Rank - 1000 // 高概率区间
	rank2 := req.Rank + 2000 // 中等概率区间
	rank3 := req.Rank + 5000 // 低概率区间

	args := []interface{}{rank1, rank2, rank3, req.Province, req.ScoreType, req.Rank - 10000}

	limit := req.Limit
	if limit <= 0 {
		limit = 50
	}
	args = append(args, limit)

	var resp []MatchingCollege
	err := m.QueryRowsNoCache(&resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return []MatchingCollege{}, nil
	default:
		return nil, err
	}
}
