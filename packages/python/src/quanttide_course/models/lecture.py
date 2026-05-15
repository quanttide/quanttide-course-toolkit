from datetime import timedelta
from enum import Enum
from typing import Any

from pydantic import BaseModel, ConfigDict, field_serializer, field_validator


class Level(str, Enum):
    introductory = "初级"
    intermediate = "中级"
    advanced = "高级"


class Lecture(BaseModel):
    model_config = ConfigDict(frozen=True, populate_by_name=True)

    id: str
    title: str
    description: str
    targets: list[str]
    objectives: list[str]
    points: list[str]
    duration: timedelta
    level: Level

    @field_validator("targets", "objectives", "points", mode="before")
    @classmethod
    def coerce_list(cls, v: Any) -> list[str]:
        if v is None:
            return []
        if isinstance(v, list):
            return [str(item) for item in v]
        return []

    @field_validator("duration", mode="before")
    @classmethod
    def coerce_duration(cls, v: Any) -> timedelta:
        if isinstance(v, timedelta):
            return v
        if isinstance(v, int):
            return timedelta(minutes=v)
        raise ValueError(f"Expected int (minutes) or timedelta, got {type(v)}")

    @field_serializer("duration")
    def serialize_duration(self, v: timedelta) -> int:
        return int(v.total_seconds() // 60)

    def to_dict(self) -> dict[str, Any]:
        return self.model_dump(mode="json")
