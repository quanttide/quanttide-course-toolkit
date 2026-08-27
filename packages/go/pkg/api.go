package course

import "strings"

// 课程域 API 路由模板（唯一事实源：qtcloud-course 服务端路由表）。
//
// Route* 常量含 {参数} 通配段，段名即服务端注册的路径变量名，
// 服务端以 Method + 常量注册路由；消费方用下方构造函数生成具体路径。
// 构造函数与常量在同一文件内推导，保证永不漂移。
const (
	// 静态资源集合路径。
	RouteCourses  = "/courses"
	RouteLessons  = "/lessons"
	RouteScenes   = "/scenes"
	RouteCriteria = "/criteria"

	// 单资源路径（{id}）。
	RouteCourse    = "/courses/{id}"
	RouteLesson    = "/lessons/{id}"
	RouteScene     = "/scenes/{id}"
	RouteCriterion = "/criteria/{id}"

	// 子资源路径。
	RouteCourseLessons  = "/courses/{courseId}/lessons"
	RouteLessonScenes   = "/lessons/{lessonId}/scenes"
	RouteLessonCriteria = "/lessons/{lessonId}/criteria"
	RouteSceneCriteria  = "/scenes/{sceneId}/criteria"
)

// fill 将 route 模板中的 {key} 通配段替换为 value。
func fill(route, key, value string) string {
	return strings.ReplaceAll(route, "{"+key+"}", value)
}

// CoursePath 构造指定课程的路径。
func CoursePath(id string) string { return fill(RouteCourse, "id", id) }

// LessonPath 构造指定课时的路径。
func LessonPath(id string) string { return fill(RouteLesson, "id", id) }

// ScenePath 构造指定场景的路径。
func ScenePath(id string) string { return fill(RouteScene, "id", id) }

// CriterionPath 构造指定验收标准的路径。
func CriterionPath(id string) string { return fill(RouteCriterion, "id", id) }

// CourseLessonsPath 构造课程课时列表的路径。
func CourseLessonsPath(courseID string) string { return fill(RouteCourseLessons, "courseId", courseID) }

// LessonScenesPath 构造课时场景列表的路径。
func LessonScenesPath(lessonID string) string { return fill(RouteLessonScenes, "lessonId", lessonID) }

// LessonCriteriaPath 构造课时级验收标准列表的路径。
func LessonCriteriaPath(lessonID string) string {
	return fill(RouteLessonCriteria, "lessonId", lessonID)
}

// SceneCriteriaPath 构造场景级验收标准列表的路径。
func SceneCriteriaPath(sceneID string) string { return fill(RouteSceneCriteria, "sceneId", sceneID) }
