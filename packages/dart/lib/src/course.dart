/// 课程领域模型：课程树三级 Course → Lesson → Scene，加验收标准 Criterion。
///
/// 字段以 qtcloud-course 服务端 API 为准，JSON key 与 API camelCase 标签一一对应。
library;

/// 解析 Map 字段的窄化辅助。
T _require<T>(Map<String, dynamic> json, String key) => json[key] as T;

List<T> _listOf<T>(Object? raw, T Function(Map<String, dynamic>) fromJson) =>
    (raw as List?)?.cast<Map<String, dynamic>>().map(fromJson).toList() ?? [];

Map<String, dynamic> _compact(Map<String, dynamic> json) =>
    json..removeWhere((_, v) => v == null);

class Course {
  final String id;
  final String name;
  final String slug;
  final String? description;
  /// draft / published
  final String? status;
  final int? sortOrder;

  const Course({
    required this.id,
    required this.name,
    required this.slug,
    this.description,
    this.status,
    this.sortOrder,
  });

  factory Course.fromJson(Map<String, dynamic> json) => Course(
        id: _require(json, 'id'),
        name: _require(json, 'name'),
        slug: _require(json, 'slug'),
        description: json['description'],
        status: json['status'],
        sortOrder: json['sortOrder'],
      );

  Map<String, dynamic> toJson() => _compact({
        'id': id,
        'name': name,
        'slug': slug,
        'description': description,
        'status': status,
        'sortOrder': sortOrder,
      });
}

class Lesson {
  final String id;
  final String courseId;
  final String title;
  final String slug;
  final String? description;
  /// 课时时长（分钟），默认45
  final int? duration;
  final int? sortOrder;
  /// draft / published
  final String? status;
  /// 入口场景 ID
  final String? startSceneId;
  /// 引用的验收标准 ID 列表（课时总验收）
  final List<String> criteria;

  const Lesson({
    required this.id,
    required this.courseId,
    required this.title,
    required this.slug,
    this.description,
    this.duration,
    this.sortOrder,
    this.status,
    this.startSceneId,
    this.criteria = const [],
  });

  factory Lesson.fromJson(Map<String, dynamic> json) => Lesson(
        id: _require(json, 'id'),
        courseId: _require(json, 'courseId'),
        title: _require(json, 'title'),
        slug: _require(json, 'slug'),
        description: json['description'],
        duration: json['duration'],
        sortOrder: json['sortOrder'],
        status: json['status'],
        startSceneId: json['startSceneId'],
        criteria:
            (json['criteria'] as List?)?.cast<String>() ?? const [],
      );

  Map<String, dynamic> toJson() => _compact({
        'id': id,
        'courseId': courseId,
        'title': title,
        'slug': slug,
        'description': description,
        'duration': duration,
        'sortOrder': sortOrder,
        'status': status,
        'startSceneId': startSceneId,
        'criteria': criteria,
      });
}

class Step {
  final int order;
  final String content;

  const Step({required this.order, required this.content});

  factory Step.fromJson(Map<String, dynamic> json) =>
      Step(order: _require(json, 'order'), content: _require(json, 'content'));

  Map<String, dynamic> toJson() => {'order': order, 'content': content};
}

class Choice {
  final String label;
  final String targetSceneId;

  const Choice({required this.label, required this.targetSceneId});

  factory Choice.fromJson(Map<String, dynamic> json) => Choice(
        label: _require(json, 'label'),
        targetSceneId: _require(json, 'targetSceneId'),
      );

  Map<String, dynamic> toJson() => {'label': label, 'targetSceneId': targetSceneId};
}

class Scene {
  final String id;
  final String lessonId;
  final String? title;
  final String slug;
  final String videoUrl;
  final List<Step> steps;
  /// 验证方式
  final String? verifyTip;
  /// 分支选项（空数组表示终结场景；服务端契约字段始终存在）
  final List<Choice> choices;
  /// 引用的验收标准 ID 列表（每步判定）
  final List<String> criteria;

  const Scene({
    required this.id,
    required this.lessonId,
    required this.slug,
    required this.videoUrl,
    this.title,
    this.steps = const [],
    this.verifyTip,
    this.choices = const [],
    this.criteria = const [],
  });

  factory Scene.fromJson(Map<String, dynamic> json) => Scene(
        id: _require(json, 'id'),
        lessonId: _require(json, 'lessonId'),
        title: json['title'],
        slug: _require(json, 'slug'),
        videoUrl: _require(json, 'videoUrl'),
        steps: _listOf(json['steps'], Step.fromJson),
        verifyTip: json['verifyTip'],
        choices: _listOf(json['choices'], Choice.fromJson),
        criteria: (json['criteria'] as List?)?.cast<String>() ?? const [],
      );

  Map<String, dynamic> toJson() => _compact({
        'id': id,
        'lessonId': lessonId,
        'title': title,
        'slug': slug,
        'videoUrl': videoUrl,
        'steps': steps.map((s) => s.toJson()).toList(),
        'verifyTip': verifyTip,
        'choices': choices.map((c) => c.toJson()).toList(),
        'criteria': criteria,
      });
}

class Criterion {
  final String id;
  final String lessonId;
  /// 所属场景；空表示课时级标准
  final String? sceneId;
  final String title;
  /// 判定规则（什么算做对）
  final String description;

  const Criterion({
    required this.id,
    required this.lessonId,
    this.sceneId,
    required this.title,
    required this.description,
  });

  factory Criterion.fromJson(Map<String, dynamic> json) => Criterion(
        id: _require(json, 'id'),
        lessonId: _require(json, 'lessonId'),
        sceneId: json['sceneId'],
        title: _require(json, 'title'),
        description: _require(json, 'description'),
      );

  Map<String, dynamic> toJson() => _compact({
        'id': id,
        'lessonId': lessonId,
        'sceneId': sceneId,
        'title': title,
        'description': description,
      });
}
