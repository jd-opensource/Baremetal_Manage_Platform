# Changelog

## [2.0.7] - 2024-08-09
### Added
- 新增日志同步到rsyslog
- 修复format分区顺序和partition不一致时导致分区显示乱序问题
- 新增Ubuntu22.04版本ramos
- 新增Ubuntu20.04版本ramos
- 修复cannot import name 'MutableMapping' from 'collections'

## [2.0.6] - 2024-08-05
### Fixed
- 修改刷新数据到硬盘的方式

## [2.0.5] - 2024-08-05
### Fixed
- 修复写入文件的数据未刷到硬盘

## [2.0.4] - 2024-08-02
### Changed
- 采集信息返回网卡bus_info

## [2.0.3] - 2024-07-22
### Added
- 修复硬盘发现较慢导致makeraid的sn值获取错误

## [2.0.2] - 2024-07-19
### Added
- 采集信息指令新增返回系统架构字段

## [2.0.1] - 2024-05-28
### Added
- RamOS构建新增版本号
- 新增Ubuntu22.04RamOS

### Changed
- 修改采集指令
- 优化raid指令

## [2.0.0] - 2023-03-03
### Added
- 升级到Python3，接口有重大调整，不兼容老系统