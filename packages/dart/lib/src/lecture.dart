import 'package:freezed_annotation/freezed_annotation.dart';

part 'lecture.freezed.dart';
part 'lecture.g.dart';

class DurationMinutesConverter implements JsonConverter<Duration, int> {
  const DurationMinutesConverter();

  @override
  Duration fromJson(int json) => Duration(minutes: json);

  @override
  int toJson(Duration object) => object.inMinutes;
}

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
    @DurationMinutesConverter() required Duration duration,
    required Level level,
  }) = _Lecture;

  factory Lecture.fromJson(Map<String, dynamic> json) => _$LectureFromJson(json);
}
