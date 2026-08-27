package course_test

import (
	"encoding/json"
	"strings"
	"testing"

	course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"
)

// 课程树标本：与 qtcloud-course seed 数据同构。
const courseTree = `{
  "id": "prod",
  "name": "生产实习",
  "slug": "prod",
  "description": "走进真实业务",
  "status": "published",
  "sortOrder": 1
}`

const lessonTree = `[
  {
    "id": "less-1",
    "courseId": "prod",
    "title": "创立故事",
    "slug": "story",
    "duration": 45,
    "sortOrder": 1,
    "status": "published",
    "startSceneId": "scen-1",
    "criteria": ["cri-1", "cri-2"]
  },
  {
    "id": "less-2",
    "courseId": "prod",
    "title": "立项",
    "slug": "kickoff",
    "sortOrder": 2
  }
]`

const sceneTree = `[
  {
    "id": "scen-1",
    "lessonId": "less-1",
    "title": "开场",
    "slug": "open",
    "videoUrl": "https://example.com/open.mp4",
    "steps": [{"order": 1, "content": "启动 Zed"}],
    "verifyTip": "确认 Zed 已启动",
    "choices": [],
    "criteria": ["cri-1"]
  }
]`

const criterionTree = `[
  {"id": "cri-1", "lessonId": "less-1", "sceneId": "scen-1", "title": "会连接 Zed", "description": "Zed 启动且可编辑文件"},
  {"id": "cri-2", "lessonId": "less-1", "title": "完成全部场景", "description": "课时内所有场景通过"}
]`

// TestCourse_UnmarshalJSON 字段与服务端 API 对齐（camelCase、omitempty 不回填零值）。
func TestCourse_UnmarshalJSON(t *testing.T) {
	var c course.Course
	if err := json.Unmarshal([]byte(courseTree), &c); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if c.ID != "prod" || c.Slug != "prod" || c.SortOrder != 1 || c.Status != "published" {
		t.Fatalf("unexpected course: %+v", c)
	}
}

func TestLesson_ParseAndBack(t *testing.T) {
	lessons, err := course.ParseLessons([]byte(lessonTree))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(lessons) != 2 || lessons[0].ID != "less-1" || lessons[0].StartSceneID != "scen-1" {
		t.Fatalf("unexpected lessons: %+v", lessons)
	}
	// 零值字段不输出（omitempty 与服务端一致）
	raw, err := json.Marshal(lessons[1])
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	for _, banned := range []string{"startSceneId", "criteria", "duration"} {
		if strings.Contains(string(raw), banned) {
			t.Fatalf("字段 %s 不应出现在零值序列化结果：%s", banned, raw)
		}
	}
}

func TestScene_UnmarshalJSON(t *testing.T) {
	scenes := []*course.Scene{}
	if err := json.Unmarshal([]byte(sceneTree), &scenes); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	s := scenes[0]
	if s.VideoURL == "" || len(s.Steps) != 1 || s.VerifyTip == "" {
		t.Fatalf("unexpected scene: %+v", s)
	}
	if s.Criteria[0] != "cri-1" {
		t.Fatalf("criteria = %v", s.Criteria)
	}
}

// TestCriterion_SceneID_Optional 场景级标准带 sceneId，课时级为空——单一事实源直连语义。
func TestCriterion_SceneID_Optional(t *testing.T) {
	cs := []*course.Criterion{}
	if err := json.Unmarshal([]byte(criterionTree), &cs); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if cs[0].SceneID != "scen-1" {
		t.Fatalf("scene-level criterion sceneId = %q", cs[0].SceneID)
	}
	if cs[1].SceneID != "" {
		t.Fatalf("lesson-level criterion sceneId should be empty, got %q", cs[1].SceneID)
	}
}
