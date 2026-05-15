import 'package:freezed_annotation/freezed_annotation.dart';

part 'lecture.freezed.dart';
part 'lecture.g.dart';

class DurationIso8601Converter implements JsonConverter<Duration, String> {
  const DurationIso8601Converter();

  @override
  Duration fromJson(String json) {
    final regex = RegExp(r'^PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?$');
    final match = regex.firstMatch(json);
    if (match == null) {
      throw FormatException('Invalid ISO 8601 duration: $json');
    }
    final hours = int.tryParse(match.group(1) ?? '0') ?? 0;
    final minutes = int.tryParse(match.group(2) ?? '0') ?? 0;
    final seconds = int.tryParse(match.group(3) ?? '0') ?? 0;
    return Duration(hours: hours, minutes: minutes, seconds: seconds);
  }

  @override
  String toJson(Duration object) {
    final minutes = object.inMinutes;
    final seconds = object.inSeconds.remainder(60);
    if (seconds == 0) return 'PT${minutes}M';
    final hours = object.inHours;
    final remainingMinutes = minutes.remainder(60);
    return 'PT${hours}H${remainingMinutes}M${seconds}S';
  }
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
    @DurationIso8601Converter() required Duration duration,
    required Level level,
  }) = _Lecture;

  factory Lecture.fromJson(Map<String, dynamic> json) => _$LectureFromJson(json);
}
