// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"
)

type EnumMood string

const (
	EnumMoodHappy   EnumMood = "happy"
	EnumMoodSad     EnumMood = "sad"
	EnumMoodNeutral EnumMood = "neutral"
)

func (e *EnumMood) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumMood(s)
	case string:
		*e = EnumMood(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumMood: %T", src)
	}
	return nil
}

type NullEnumMood struct {
	EnumMood EnumMood `json:"enum_mood"`
	Valid    bool     `json:"valid"` // Valid is true if EnumMood is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEnumMood) Scan(value interface{}) error {
	if value == nil {
		ns.EnumMood, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EnumMood.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEnumMood) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EnumMood), nil
}

type User struct {
	UserID   int32        `json:"user_id"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Email    string       `json:"email"`
	Mood     NullEnumMood `json:"mood"`
}
