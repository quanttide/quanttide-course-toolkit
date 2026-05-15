import 'dart:convert';
import 'dart:io';

import 'package:test/test.dart';
import 'package:quanttide_course/quanttide_course.dart';

Map<String, dynamic> _snakeToCamel(Map<String, dynamic> input) {
  final result = <String, dynamic>{};
  for (final entry in input.entries) {
    final camel = entry.key.replaceAllMapped(
      RegExp(r'_(.)'),
      (m) => m.group(1)!.toUpperCase(),
    );
    result[camel] = entry.value;
  }
  return result;
}

String _fixturePath(String name) =>
    File('${Directory.current.path}/../../tests/fixtures/$name').path;

String _schemaPath(String name) =>
    File('${Directory.current.path}/../../tests/schemas/$name').path;

void main() {
  group('Lecture contract', () {
    test('schema definition is valid JSON', () {
      final path = _schemaPath('lecture.json');
      final raw = File(path).readAsStringSync();
      final schema = jsonDecode(raw) as Map<String, dynamic>;
      expect(schema['title'], 'Lecture');
      expect(schema['required'], contains('id'));
      expect(schema['required'], contains('title'));
      expect(schema['required'], isNot(contains('duration')));
    });

    test('fixture deserializes correctly', () {
      final path = _fixturePath('lecture.json');
      final raw = jsonDecode(File(path).readAsStringSync()) as Map<String, dynamic>;
      final camel = _snakeToCamel(raw);
      final lecture = Lecture.fromJson(camel);
      expect(lecture.id, 'lec_001');
      expect(lecture.title, 'Python 基础');
      expect(lecture.level, Level.introductory);
    });

    test('round-trip', () {
      final path = _fixturePath('lecture.json');
      final raw = jsonDecode(File(path).readAsStringSync()) as Map<String, dynamic>;
      final camel = _snakeToCamel(raw);
      final lecture = Lecture.fromJson(camel);
      final json = lecture.toJson();
      expect(json['id'], 'lec_001');
      expect(json['title'], 'Python 基础');
      expect(json['level'], '初级');
    });

    test('minimal instance serializes', () {
      final lecture = Lecture(
        id: 'lec_min',
        title: 'Minimal',
        description: '',
        targets: [],
        objectives: [],
        points: [],
        level: Level.introductory,
      );
      final json = lecture.toJson();
      expect(json['id'], 'lec_min');
      expect(json['title'], 'Minimal');
    });
  });
}
