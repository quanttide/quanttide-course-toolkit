import re
from datetime import timedelta
from enum import Enum
from typing import Any

from pydantic import BaseModel, ConfigDict, field_serializer, field_validator

_ISO_DURATION_RE = re.compile(r"^PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?$")


def _parse_duration(value: str) -> timedelta:
    match = _ISO_DURATION_RE.match(value)
    if not match:
        raise ValueError(f"Invalid ISO 8601 duration: {value}")
    hours = int(match.group(1) or "0")
    minutes = int(match.group(2) or "0")
    seconds = int(match.group(3) or "0")
    return timedelta(hours=hours, minutes=minutes, seconds=seconds)


def _format_duration(delta: timedelta) -> str:
    total = int(delta.total_seconds())
    hours, remainder = divmod(total, 3600)
    minutes, seconds = divmod(remainder, 60)
    parts = ["PT"]
    if hours:
        parts.append(f"{hours}H")
    if minutes:
        parts.append(f"{minutes}M")
    if seconds:
        parts.append(f"{seconds}S")
    return "".join(parts)


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
        if isinstance(v, str):
            return _parse_duration(v)
        raise ValueError(f"Expected ISO 8601 string or timedelta, got {type(v)}")

    @field_serializer("duration")
    def serialize_duration(self, v: timedelta) -> str:
        return _format_duration(v)

    def to_dict(self) -> dict[str, Any]:
        return self.model_dump(mode="json")
