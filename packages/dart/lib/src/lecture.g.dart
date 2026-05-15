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
  level: $enumDecode(_$LevelEnumMap, json['level']),
  targets: (json['targets'] as List<dynamic>).map((e) => e as String).toList(),
  objectives: (json['objectives'] as List<dynamic>)
      .map((e) => e as String)
      .toList(),
  points: (json['points'] as List<dynamic>).map((e) => e as String).toList(),
);

Map<String, dynamic> _$$LectureImplToJson(_$LectureImpl instance) =>
    <String, dynamic>{
      'id': instance.id,
      'title': instance.title,
      'description': instance.description,
      'level': _$LevelEnumMap[instance.level]!,
      'targets': instance.targets,
      'objectives': instance.objectives,
      'points': instance.points,
    };

const _$LevelEnumMap = {
  Level.introductory: '初级',
  Level.intermediate: '中级',
  Level.advanced: '高级',
};
