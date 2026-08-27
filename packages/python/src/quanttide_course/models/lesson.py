"""课程领域模型：课程树三级 Course → Lesson → Scene，加验收标准 Criterion.

字段以 qtcloud-course 服务端 API 为准；Python 属性使用 snake_case，
序列化边界经 alias 转为 API 的 camelCase。
"""

from typing import Literal, Optional

from pydantic import BaseModel, ConfigDict, Field


class _Contract(BaseModel):
    """契约基类：不可变实例 + 别名双向解析。"""

    model_config = ConfigDict(frozen=True, populate_by_name=True)


class Course(_Contract):
    id: str
    name: str
    slug: str
    description: Optional[str] = None
    status: Optional[Literal["draft", "published"]] = None
    sort_order: Optional[int] = Field(default=None, alias="sortOrder")


class Lesson(_Contract):
    id: str
    course_id: str = Field(alias="courseId")
    title: str
    slug: str
    description: Optional[str] = None
    duration: Optional[int] = None  # 课时时长（分钟），默认45
    sort_order: Optional[int] = Field(default=None, alias="sortOrder")
    status: Optional[Literal["draft", "published"]] = None
    start_scene_id: Optional[str] = Field(default=None, alias="startSceneId")
    criteria: list[str] = []


class Step(_Contract):
    order: int
    content: str


class Choice(_Contract):
    label: str
    target_scene_id: str = Field(alias="targetSceneId")


class Scene(_Contract):
    id: str
    lesson_id: str = Field(alias="lessonId")
    title: Optional[str] = None
    slug: str
    video_url: str = Field(alias="videoUrl")
    steps: list[Step] = []
    verify_tip: Optional[str] = Field(default=None, alias="verifyTip")
    choices: list[Choice]
    criteria: list[str] = []


class Criterion(_Contract):
    """验收标准：学习云 completion.criterion_id 直连本实体 ID。"""

    id: str
    lesson_id: str = Field(alias="lessonId")
    scene_id: Optional[str] = Field(default=None, alias="sceneId")  # 空表示课时级
    title: str
    description: str
