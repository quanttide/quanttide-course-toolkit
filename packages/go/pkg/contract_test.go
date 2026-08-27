// 契约测试：以根 tests/ 的 Schema + Fixture 为单一事实源验证 Go 模型。
//
// 当前为轻量实现（零依赖）：校验必填字段存在性（与 Schema required 一致）
// 并反序列化标本。完整 JSON Schema 校验待引入校验依赖后升级。
package course_test

import (
	"encoding/json"
	"os"
	"testing"

	course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"
)

func fixture(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile("../../../tests/fixtures/" + name)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return data
}

// hasFields 校验解码后的对象包含 schema 要求的必填字段。
func hasFields(t *testing.T, data []byte, required ...string) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	for _, k := range required {
		if _, ok := m[k]; !ok {
			t.Fatalf("%s 缺少必填字段 %q", "fixture", k)
		}
	}
	return m
}

func TestContract_Fixtures_Unmarshal(t *testing.T) {
	var c course.Course
	if err := json.Unmarshal(fixture(t, "course.json"), &c); err != nil {
		t.Fatal(err)
	}
	hasFields(t, fixture(t, "course.json"), "id", "name", "slug")
	if c.ID != "prod" || c.Status != "published" || c.SortOrder != 1 {
		t.Fatalf("course = %+v", c)
	}

	var l course.Lesson
	if err := json.Unmarshal(fixture(t, "lesson.json"), &l); err != nil {
		t.Fatal(err)
	}
	hasFields(t, fixture(t, "lesson.json"), "id", "courseId", "title", "slug")
	if l.CourseID != "prod" || l.StartSceneID != "scen-1" || len(l.Criteria) != 2 {
		t.Fatalf("lesson = %+v", l)
	}

	var s course.Scene
	if err := json.Unmarshal(fixture(t, "scene.json"), &s); err != nil {
		t.Fatal(err)
	}
	hasFields(t, fixture(t, "scene.json"), "id", "lessonId", "slug", "videoUrl", "choices")
	if s.VideoURL == "" || len(s.Steps) != 1 || s.Choices == nil {
		t.Fatalf("scene = %+v", s)
	}

	var cr course.Criterion
	if err := json.Unmarshal(fixture(t, "criterion.json"), &cr); err != nil {
		t.Fatal(err)
	}
	hasFields(t, fixture(t, "criterion.json"), "id", "lessonId", "title", "description")
	if cr.SceneID != "scen-1" || cr.Title != "会连接 Zed" {
		t.Fatalf("criterion = %+v", cr)
	}
}

// TestContract_RoundTrip 反序列化 → 再序列化，契约字段保持一致。
func TestContract_RoundTrip(t *testing.T) {
	var l course.Lesson
	if err := json.Unmarshal(fixture(t, "lesson.json"), &l); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}
	var back course.Lesson
	if err := json.Unmarshal(out, &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Criteria) != len(l.Criteria) || back.ID != l.ID || back.CourseID != l.CourseID ||
		back.StartSceneID != l.StartSceneID || back.Duration != l.Duration {
		t.Fatalf("round-trip mismatch: got %#v want %#v", back, l)
	}
	for i := range l.Criteria {
		if back.Criteria[i] != l.Criteria[i] {
			t.Fatalf("criteria[%d] mismatch", i)
		}
	}
}
