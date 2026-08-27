package course_test

import (
	"encoding/json"
	"os"
	"testing"

	course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"
)

// TestLecture_UnmarshalFixture 用共享标本验证字段对齐。
func TestLecture_UnmarshalFixture(t *testing.T) {
	data, err := os.ReadFile("../../../tests/fixtures/lecture.json")
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	var lec course.Lecture
	if err := json.Unmarshal(data, &lec); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if lec.ID != "lec_001" || lec.Title != "Python 基础" {
		t.Fatalf("unexpected lecture: %s %s", lec.ID, lec.Title)
	}
	if lec.Level != course.LevelIntroductory {
		t.Fatalf("level = %q", lec.Level)
	}
	if len(lec.Objectives) != 3 || len(lec.Points) != 7 || len(lec.Targets) != 1 {
		t.Fatalf("list fields not coerced: %d/%d/%d", len(lec.Targets), len(lec.Objectives), len(lec.Points))
	}
	if err := (&lec).Validate(); err != nil {
		t.Fatalf("validate: %v", err)
	}
}

// TestLecture_NullListsCoerceToEmpty 缺失与 null 的列表字段归一为空切片。
func TestLecture_NullListsCoerceToEmpty(t *testing.T) {
	// 与 Python SDK 语义一致：null、缺失与非列表（如数字）都归一为空切片。
	data := []byte(`{"id":"lec_002","title":"空","description":"","level":"中级","targets":null,"objectives":[],"points":3}`)
	var lec course.Lecture
	if err := json.Unmarshal(data, &lec); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(lec.Targets) != 0 || len(lec.Objectives) != 0 || len(lec.Points) != 0 {
		t.Fatalf("coercion wrong: %#v", lec)
	}
}

// TestLecture_Validate_LevelInvalid 非法层级被拒绝。
func TestLecture_Validate_LevelInvalid(t *testing.T) {
	lec := &course.Lecture{ID: "x", Title: "y", Level: "骨灰级"}
	if err := lec.Validate(); err == nil {
		t.Fatal("expected level validation error")
	}
}
