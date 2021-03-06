// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gof/app/dao/internal"
)

// userScoresDao is the manager for logic model data accessing and custom defined data operations functions management. 
// You can define custom methods on it to extend its functionality as you wish.
type userScoresDao struct {
	*internal.UserScoresDao
}

var (
	// UserScores is globally public accessible object for table user_scores operations.
	UserScores = userScoresDao{
		internal.NewUserScoresDao(),
	}
)

// Fill with you ideas below.