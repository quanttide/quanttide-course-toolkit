import 'package:freezed_annotation/freezed_annotation.dart';

part 'lecture.freezed.dart';
part 'lecture.g.dart';

enum Level {
  @JsonValue('初级') introductory('初级'),
  @JsonValue('中级') intermediate('中级'),
  @JsonValue('高级') advanced('高级');

  final String label;
  const Level(this.label);
}

@freezed
class Lecture with _$Lecture {
  const factory Lecture({
    required String id,
    required String title,
    required String description,
    required Level level,
    required List<String> targets,
    required List<String> objectives,
    required List<String> points,
  }) = _Lecture;

  factory Lecture.fromJson(Map<String, dynamic> json) => _$LectureFromJson(json);
}
