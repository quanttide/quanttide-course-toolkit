import 'package:freezed_annotation/freezed_annotation.dart';

part 'lecture.freezed.dart';
part 'lecture.g.dart';

enum Level {
  @JsonValue('初级') introductory,
  @JsonValue('中级') intermediate,
  @JsonValue('高级') advanced;
}

@freezed
class Lecture with _$Lecture {
  const factory Lecture({
    required String id,
    required String title,
    required String description,
    required List<String> targets,
    required List<String> objectives,
    required List<String> points,
    required Level level,
  }) = _Lecture;

  factory Lecture.fromJson(Map<String, dynamic> json) => _$LectureFromJson(json);
}
