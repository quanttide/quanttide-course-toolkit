// Package course 提供课程领域模型。
package course

import (
	"encoding/json"
	"fmt"
)

// Level 课程难度层级。
type Level string

const (
	LevelIntroductory Level = "初级"
	LevelIntermediate Level = "中级"
	LevelAdvanced     Level = "高级"
)

// Lecture 讲次：课程内容的基本组织单元。
//
// Targets 为适合人群；Objectives 为学习目标；Points 为知识点清单。
// 字段语义与 Dart / Python SDK 保持一致，标签取自课程数据 JSON 形态。
type Lecture struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Level       Level    `json:"level"`
	Targets     []string `json:"targets"`
	Objectives  []string `json:"objectives"`
	Points      []string `json:"points"`
}

// UnmarshalJSON 反序列化时容忍缺失或非列表的列表字段（如 null）。
func (l *Lecture) UnmarshalJSON(data []byte) error {
	type alias struct {
		ID          string          `json:"id"`
		Title       string          `json:"title"`
		Description string          `json:"description"`
		Level       Level           `json:"level"`
		Targets     json.RawMessage `json:"targets"`
		Objectives  json.RawMessage `json:"objectives"`
		Points      json.RawMessage `json:"points"`
	}
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	l.ID, l.Title, l.Description, l.Level = a.ID, a.Title, a.Description, a.Level

	targets, err := coerceList(a.Targets)
	if err != nil {
		return fmt.Errorf("targets: %w", err)
	}
	objectives, err := coerceList(a.Objectives)
	if err != nil {
		return fmt.Errorf("objectives: %w", err)
	}
	points, err := coerceList(a.Points)
	if err != nil {
		return fmt.Errorf("points: %w", err)
	}
	l.Targets, l.Objectives, l.Points = targets, objectives, points
	return nil
}

// coerceList 将 raw 归一为字符串列表：缺失、null 或非列表一律空切片，
// 与 Dart / Python SDK 的宽松语义保持一致。
func coerceList(raw json.RawMessage) ([]string, error) {
	var out []string
	if len(raw) == 0 || string(raw) == "null" {
		return []string{}, nil
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return []string{}, nil
	}
	return out, nil
}

// Validate 校验必需字段与合法难度层级。
func (l *Lecture) Validate() error {
	if l.ID == "" {
		return fmt.Errorf("id 必填")
	}
	if l.Title == "" {
		return fmt.Errorf("title 必填")
	}
	switch l.Level {
	case LevelIntroductory, LevelIntermediate, LevelAdvanced:
	default:
		return fmt.Errorf("level 非法：%q", l.Level)
	}
	return nil
}
