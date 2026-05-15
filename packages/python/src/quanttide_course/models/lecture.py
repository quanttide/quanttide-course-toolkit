from enum import Enum
from typing import Any

from pydantic import BaseModel, ConfigDict, field_validator


class Level(str, Enum):
    introductory = "初级"
    intermediate = "中级"
    advanced = "高级"


class Lecture(BaseModel):
    model_config = ConfigDict(frozen=True, populate_by_name=True)

    id: str
    title: str
    description: str
    level: Level
    targets: list[str]
    objectives: list[str]
    points: list[str]
    level: Level

    @field_validator("targets", "objectives", "points", mode="before")
    @classmethod
    def coerce_list(cls, v: Any) -> list[str]:
        if v is None:
            return []
        if isinstance(v, list):
            return [str(item) for item in v]
        return []

    def to_dict(self) -> dict[str, Any]:
        return self.model_dump(mode="json")
