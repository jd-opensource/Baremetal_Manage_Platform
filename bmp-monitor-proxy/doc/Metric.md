# 云物理服务器


metric中文展示名称 |metric英文展示名称 | metric名字 | 值类型 | downsample-aggregator | grouping-aggregator | 变化率展示 | tags(必须/不必须) | 上报时间间隔 | 上报 | monitor展示 | 单位转换公式(convertFrom/convertTo) | 前端是否展示 | 是否可用于创建告警规则 | 创建报警规则的特殊需求 | 漏报检查(每小时)预期/阈值
---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ----| ----| ---- | ---- 
CPU使用率 | CPU Usage | cps.cpu.util | 时刻 | avg |  | n | - | 60s | % | % |  | y | y | 无 |
内存使用率 | Memory Usage | cps.memory.util | 时刻 | avg |  | n | - | 60s | % | % |   | y | y | 无 |
内存使用量 | Memory Used | cps.memory.used | 时刻 | avg |  | n | - | 60s |Byte| GB |1/1000000000   | y | y | 无 |
硬盘使用量 | Disk Used | cps.disk.used | 时刻 | avg |  | n | device | 60s | Byte | GB | 1/1000000000 | y | y | 无 |
硬盘使用率 | Disk Usage | cps.disk.util | 时刻 | avg |  | n | device | 60s | % | % |  | y | y | 无 |
硬盘读流量 | Disk Read Traffic | cps.disk.bytes.read | 累计 | last |  | y | device | 60s | Byte | Bps |  | y | y | 无 |
硬盘写流量 | Disk Write Traffic | cps.disk.bytes.write | 累计 | last |  | y | device | 60s | Byte | Bps |  | y | y | 无 |
硬盘读IOPS | Disk Read IOPS | cps.disk.counts.read | 累计 | last |  | y | device | 60s | 次 | 次/秒 |  | y | y | 无 |
硬盘写IOPS | Disk Write IOPS | cps.disk.counts.write | 累计 | last |  | y | device | 60s | 次 | 次/秒 |  | y | y | 无 |
网卡进流量 | Network Ingress Traffic | cps.network.bytes.ingress | 累计 | last |  | y | - | 60s | Byte | bps |8/1  | y | y | 无 |
网卡出流量 | Network Egress Traffic | cps.network.bytes.egress | 累计 | last |  | y | - | 60s | Byte | bps | 8/1 | y | y | 无 | 
网络进包量 | Network Ingress Packets | cps.network.packets.ingress | 累计 | last |  | y | - | 60s | 个 | 个/秒 |  | y | y | 无 | 
网络出包量 | Network Egress Packets | cps.network.packets.egress | 累计 | last |  | y | - | 60s | 个 | 个/秒 |  | y | y | 无 | 
平均负载1min | Load Average 1min | cps.avg.load1 | 时刻 | avg |  | n | - | 60s | - | 无 |  | y | y | 无 |
平均负载5min | Load Average 5min | cps.avg.load5 | 时刻 | avg |  | n | - | 60s | - | 无 |  | y | y | 无 |
平均负载15min | Load Average 15min | cps.avg.load15 | 时刻 | avg |  | n | - | 60s | - | 无 |  | y | y | 无 | 
TCP总连接数 | Total TCP Connections | cps.tcp.connect.total | 时刻 | avg |  | n | - | 60s | 个 | 个 |  | y | y | 无 |
TCP正常连接数 | Established TCP Connections | cps.tcp.connect.established | 时刻 | avg |  | n | - | 60s | 个 | 个 |  | y | y | 无 |
总进程数 | Total Process Count | cps.process.total | 时刻 | avg |  | n | - | 60s | 个 | 个 |  | y | y | 无 |
