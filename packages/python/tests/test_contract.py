import json
from pathlib import Path

from quanttide_course.models.lecture import Lecture

ROOT = Path(__file__).parents[3] / "tests"
SCHEMAS = ROOT / "schemas"
FIXTURES = ROOT / "fixtures"


def _validate_schema(data: dict, schema: dict) -> list[str]:
    errors = []
    props = schema.get("properties", {})
    required = schema.get("required", [])

    for field in required:
        if field not in data or data[field] is None:
            errors.append(f"missing required field: {field}")

    for field, value in data.items():
        if field not in props:
            continue
        prop = props[field]
        ptype = prop.get("type")
        if ptype == "string" and not isinstance(value, str):
            errors.append(f"{field}: expected string, got {type(value).__name__}")
        if ptype == "array" and not isinstance(value, list):
            errors.append(f"{field}: expected array, got {type(value).__name__}")

    return errors


class TestLectureContract:
    def test_fixture_conforms_to_schema(self):
        schema = json.loads((SCHEMAS / "lecture.json").read_text())
        fixture = json.loads((FIXTURES / "lecture.json").read_text())
        errors = _validate_schema(fixture, schema)
        assert errors == [], f"Schema validation errors: {errors}"

    def test_fixture_round_trip(self):
        fixture = json.loads((FIXTURES / "lecture.json").read_text())
        lecture = Lecture.model_validate(fixture)
        data = lecture.to_dict()
        restored = Lecture.model_validate(data)
        assert restored == lecture

    def test_model_serialization_conforms_to_schema(self):
        schema = json.loads((SCHEMAS / "lecture.json").read_text())
        lecture = Lecture(
            id="lec_test",
            title="Test",
            description="Test description",
            targets=["target1"],
            objectives=["obj1"],
            points=["pt1"],
            level="初级",
        )
        data = lecture.to_dict()
        errors = _validate_schema(data, schema)
        assert errors == [], f"Schema validation errors: {errors}"

    def test_minimal_lecture_conforms(self):
        schema = json.loads((SCHEMAS / "lecture.json").read_text())
        lecture = Lecture(
            id="lec_min",
            title="Minimal",
            description="",
            targets=[],
            objectives=[],
            points=[],
            level="初级",
        )
        data = lecture.to_dict()
        errors = _validate_schema(data, schema)
        assert errors == []
