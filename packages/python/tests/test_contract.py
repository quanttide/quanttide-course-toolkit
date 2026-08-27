"""契约测试：以根 tests/ 的 Schema + Fixture 验证 Python 模型."""

import json
from pathlib import Path

from quanttide_course.models.lesson import (
    Choice,
    Course,
    Criterion,
    Lesson,
    Scene,
    Step,
)

ROOT = Path(__file__).parents[3] / "tests"
FIXTURES = ROOT / "fixtures"
SCHEMAS = ROOT / "schemas"


def fixture(name: str) -> str:
    return (FIXTURES / name).read_text(encoding="utf-8")


def fixture_dict(name: str) -> dict:
    return json.loads(fixture(name))


def has_required(data: dict, schema_name: str) -> None:
    """校验 fixture 覆盖 Schema 的必填字段（轻量实现对齐 Go 侧）。"""
    schema = json.loads((SCHEMAS / f"{schema_name}.json").read_text(encoding="utf-8"))
    missing = [k for k in schema.get("required", []) if k not in data]
    assert not missing, f"{schema_name} 缺少必填字段: {missing}"


def _assert_choice(c: Choice) -> None:
    assert c.label == "继续" and c.target_scene_id == "scene-2"


def _choice_dict() -> dict:
    return {"label": "继续", "targetSceneId": "scene-2"}


def test_course_contract():
    data = fixture_dict("course.json")
    has_required(data, "course")
    c = Course.model_validate_json(fixture("course.json"))
    assert c.id == "prod" and c.slug == "prod" and c.status == "published"


def test_lesson_contract():
    data = fixture_dict("lesson.json")
    has_required(data, "lesson")
    l = Lesson.model_validate_json(fixture("lesson.json"))
    assert l.course_id == "prod" and l.start_scene_id == "scen-1"
    assert len(l.criteria) == 2

    # Round-trip：dump 回 API 形态（camelCase）后可再次解析，值不变
    again = Lesson.model_validate(l.model_dump(by_alias=True))
    assert again == l


def test_scene_contract():
    data = fixture_dict("scene.json")
    has_required(data, "scene")
    s = Scene.model_validate_json(fixture("scene.json"))
    assert s.video_url.endswith("open.mp4") and len(s.steps) == 1 and isinstance(s.steps[0], Step)
    # choices 空数组 = 终结场景（服务端契约：字段始终存在）
    assert s.choices == []
    assert Scene.model_validate(s.model_dump(by_alias=True)) == s


def test_criterion_contract():
    data = fixture_dict("criterion.json")
    has_required(data, "criterion")
    cr = Criterion.model_validate_json(fixture("criterion.json"))
    assert cr.scene_id == "scen-1" and cr.title == "会连接 Zed"

    # 无 sceneId 表示课时级标准（schema 中非必填、模型默认 None）
    lesson_level = Criterion(**{
        "id": "cri-2",
        "lessonId": "less-1",
        "title": "完成全部场景",
        "description": "课时内所有场景通过",
    })
    assert lesson_level.scene_id is None
