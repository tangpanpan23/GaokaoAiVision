package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

// cache keys
const (
	cacheAdmissionScoreIdPrefix = "cache:admissionScore:id:"
)
