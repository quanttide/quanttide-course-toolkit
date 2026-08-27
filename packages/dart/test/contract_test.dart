// 契约测试：以根 tests/ 的 Schema + Fixture 验证 Dart 模型。
import 'dart:convert';
import 'dart:io';

import 'package:test/test.dart';
import 'package:quanttide_course/quanttide_course.dart';

final fixturesDir = Directory('${Directory.current.path}/../../tests/fixtures');
final schemasDir = Directory('${Directory.current.path}/../../tests/schemas');

Map<String, dynamic> fixture(String name) =>
    jsonDecode(File('${fixturesDir.path}/$name').readAsStringSync())
        as Map<String, dynamic>;

void hasRequired(Map<String, dynamic> data, String schemaName) {
  final schema = jsonDecode(
    File('${schemasDir.path}/$schemaName.json').readAsStringSync(),
  ) as Map<String, dynamic>;
  final missing =
      ((schema['required'] as List?) ?? []).cast<String>().where((k) => !data.containsKey(k));
  expect(missing, isEmpty, reason: '$schemaName 缺少必填字段');
}

void main() {
  test('course 契约', () {
    final data = fixture('course.json');
    hasRequired(data, 'course');
    final c = Course.fromJson(data);
    expect(c.id, 'prod');
    expect(c.slug, 'prod');
    expect(c.status, 'published');
  });

  test('lesson 契约 + round-trip', () {
    final data = fixture('lesson.json');
    hasRequired(data, 'lesson');
    final l = Lesson.fromJson(data);
    expect(l.courseId, 'prod');
    expect(l.startSceneId, 'scen-1');
    expect(l.criteria.length, 2);

    // Round-trip：序列化回 API 形态后可再次解析，值不变
    final again = Lesson.fromJson(l.toJson());
    expect(again.id, l.id);
    expect(again.courseId, l.courseId);
    expect(again.startSceneId, l.startSceneId);
    expect(again.criteria, l.criteria);
  });

  test('scene 契约', () {
    final data = fixture('scene.json');
    hasRequired(data, 'scene');
    final s = Scene.fromJson(data);
    expect(s.videoUrl.endsWith('open.mp4'), isTrue);
    expect(s.steps.length, 1);
    expect(s.choices, isEmpty); // 空数组 = 终结场景（字段始终存在）
  });

  test('criterion 契约', () {
    final data = fixture('criterion.json');
    hasRequired(data, 'criterion');
    final cr = Criterion.fromJson(data);
    expect(cr.sceneId, 'scen-1');
    expect(cr.title, '会连接 Zed');

    // 无 sceneId 表示课时级标准
    final lessonLevel = const Criterion(
      id: 'cri-2',
      lessonId: 'less-1',
      title: '完成全部场景',
      description: '课时内所有场景通过',
    );
    expect(lessonLevel.sceneId, isNull);
  });
}
