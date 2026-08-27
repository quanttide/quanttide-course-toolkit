package course_test

import (
	"strings"
	"testing"

	course "github.com/quanttide/quanttide-course-toolkit/packages/go/pkg"
)

// 标本独立于模板书写：若有人改动模板导致契约漂移，此处会先红。
var routeSpec = map[string]string{
	"RouteCourses":        "/courses",
	"RouteCourse":         "/courses/cour-1",
	"RouteLessons":        "/lessons",
	"RouteLesson":         "/lessons/less-1",
	"RouteScenes":         "/scenes/scen-1",
	"RouteCriteria":       "/criteria",
	"RouteCriterion":      "/criteria/cri-1",
	"RouteCourseLessons":  "/courses/prod/lessons",
	"RouteLessonScenes":   "/lessons/less-1/scenes",
	"RouteLessonCriteria": "/lessons/less-1/criteria",
	"RouteSceneCriteria":  "/scenes/scen-1/criteria",
}

func TestRoutes_MatchSpec(t *testing.T) {
	cases := []struct {
		name string
		got  string
	}{
		{"RouteCourses", course.RouteCourses},
		{"RouteCourse", course.CoursePath("cour-1")},
		{"RouteLessons", course.RouteLessons},
		{"RouteLesson", course.LessonPath("less-1")},
		{"RouteScenes", course.ScenePath("scen-1")},
		{"RouteCriteria", course.RouteCriteria},
		{"RouteCriterion", course.CriterionPath("cri-1")},
		{"RouteCourseLessons", course.CourseLessonsPath("prod")},
		{"RouteLessonScenes", course.LessonScenesPath("less-1")},
		{"RouteLessonCriteria", course.LessonCriteriaPath("less-1")},
		{"RouteSceneCriteria", course.SceneCriteriaPath("scen-1")},
	}
	for _, c := range cases {
		want := routeSpec[c.name]
		if c.got != want {
			t.Errorf("%s = %q, want %q", c.name, c.got, want)
		}
	}
}

// TestRoutes_TemplatesContainWildcards 通配段名与服务端路径变量一致。
func TestRoutes_TemplatesContainWildcards(t *testing.T) {
	if !strings.Contains(course.RouteCourseLessons, "{courseId}") ||
		!strings.Contains(course.RouteLessonScenes, "{lessonId}") ||
		!strings.Contains(course.RouteLessonCriteria, "{lessonId}") ||
		!strings.Contains(course.RouteSceneCriteria, "{sceneId}") {
		t.Fatal("子资源模板通配段名漂移")
	}
}
