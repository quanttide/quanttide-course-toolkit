// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'lecture.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$LectureImpl _$$LectureImplFromJson(
  Map<String, dynamic> json,
) => _$LectureImpl(
  id: json['id'] as String,
  title: json['title'] as String,
  description: json['description'] as String,
  targets: (json['targets'] as List<dynamic>).map((e) => e as String).toList(),
  objectives: (json['objectives'] as List<dynamic>)
      .map((e) => e as String)
      .toList(),
  points: (json['points'] as List<dynamic>).map((e) => e as String).toList(),
  duration: const DurationMinutesConverter().fromJson(
    (json['duration'] as num).toInt(),
  ),
  level: $enumDecode(_$LevelEnumMap, json['level']),
);

Map<String, dynamic> _$$LectureImplToJson(_$LectureImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'title': instance.title,
      'description': instance.description,
      'targets': instance.targets,
      'objectives': instance.objectives,
      'points': instance.points,
      'duration': const DurationMinutesConverter().toJson(instance.duration),
      'level': _$LevelEnumMap[instance.level]!,
    };

const _$LevelEnumMap = {
  Level.introductory: '初级',
  Level.intermediate: '中级',
  Level.advanced: '高级',
};
