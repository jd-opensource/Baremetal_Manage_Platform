const zh: {[x: string]: {[x: string]: unknown}} =  {
    size: {
        num: '剩余全部',
        all: '全部'
    },
    operate: {
        success: '操作成功',
        fail: '操作失败',
        doing: '操作中',
        copy: '复制成功',
        refresh: '刷新',
        setUp: '设置',
        export: '导出',
        name: '操作',
        btn: {
            cancel: '取消',
            sure: '确定'
        },
        tip: {
            default: '请选择资源',
        }
    },
    searchTip: {
        title1: '查询到符合搜索及筛选条件的资源共 {0} 条，',
        title2: '查询到符合搜索及筛选条件的资源共 {0} 条，请更换搜索关键字或',
        title3: '返回列表'
    },
    monitorEcharts: {
        tabs: {
            h1: '1小时',
            h6: '6小时',
            h12: '12小时',
            d1: '1天',
            d3: '3天',
            d7: '7天',
            d14: '14天'
        },
        datePicker: {
            startDate: '开始日期',
            endDate: '结束日期'
        },
        radio: {
            instance: '实例',
            disk: '磁盘',
            diskPartition: '磁盘用量',
            netWorkInterface: '网口'
        },
        diskTip: 'Linux系统按设备名上报， Windows系统按盘符上报',
        diskPartitionTip: 'Linux系统按分区挂载点上报，Windows系统按盘符上报',
        echartsCount: {
            period: '周期',
            aggregationMethod: '聚合方式',
            cpuUnit: 'CPU使用率（%）',
            cpu: 'CPU使用率',
            memUnit: '内存使用率（%）',
            mem: '内存使用率',
            memUsageUnit: '内存使用量（GB）',
            memUsage: '内存使用量',
            diskUsageUnit: '磁盘使用量（GB）',
            diskUsage: '磁盘使用量',
            diskUsageRateUnit: '磁盘使用率（%）',
            diskUsageRate: '磁盘使用率',
            diskReadWriteTrafficUnit: '磁盘读/写流量（Bps）',
            diskReadWriteTraffic: '磁盘读/写流量',
            diskReadTraffic: '磁盘读流量',
            diskWriteTraffic: '磁盘写流量',
            diskReadWriteIOPSUnit: '磁盘读/写IOPS（次/秒）',
            diskReadWriteIOPS: '磁盘读/写IOPS',
            diskReadIOPS: '磁盘读IOPS',
            diskWriteIOPS: '磁盘写IOPS',
            netWorkBpsUnit: '网口进/出流量（bps）',
            netWorkBps: '网口进/出流量',
            netWorkIn: '网口进流量',
            netWorkOut: '网口出流量',
            netWorkPackagesNumUnit: '网络进/出包数量（个/秒）',
            netWorkPackagesNumber: '网络进/出包数量',
            netWorkInPackageNum: '网络进包量',
            netWorkOutPackageNum: '网络出包量',
            averageLoad1MinUnit: '平均负载1min',
            averageLoad1Min: '平均负载1min',
            averageLoad5MinUnit: '平均负载5min',
            averageLoad5Min: '平均负载5min',
            averageLoad15MinUnit: '平均负载15min',
            averageLoad15Min: '平均负载15min',
            totalTCPConnectionsUnit: 'TCP总连接数（个）',
            totalTCPConnections: 'TCP总连接数',
            normalTCPConnectionsUnit: 'TCP正常连接数（个）',
            normalTCPConnections: 'TCP正常连接数',
            totalNumberOfProcessesUnit: '总进程数（个）',
            totalNumberOfProcesses: '总进程数'
        }
    },
    userInfo: {
        emptyTip: {
            userName: '用户名输入不能为空',
            password: '密码输入不能为空',
            cellPhone: '手机号输入不能为空',
            email: '邮箱地址输入不能为空'
        },
        errorTip: {
            number: '只能输入数字',
            userName: '1~64 字符，只支持数字、大小写字母、英文下划线 “_”及中划线“-”',
            password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
            repeatUserName: '用户名重复',
            cellPhone: '请输入正确的手机号',
            email: '请输入正确的邮箱地址',
            HongKong: '请输入正确的中国香港手机号',
            MacaMacao: '请输入正确的中国澳门手机号',
            Taiwan: '请输入正确的中国台湾手机号'
        }
    },
    login: {
        header: {
            title: '裸金属运营支撑平台（标准版）',
            title2: '裸金属运营支撑平台（专业版）'
        },
        formSubmit: {
            title: '欢迎回来',
            placeholder: {
                userName: '请输入用户名',
                passWord: '请输入密码'
            },
            toolTip: '用户名不存在或者密码错误',
            login: '登录'
        }
    },
    loginOut: {
        title: '退出登录',
        tip: '确定退出登录',
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    navigationHeader: {
        title: '裸金属运营支撑平台（标准版）',
        title2: '裸金属运营支撑平台（专业版）',
        myProfile: '我的账户',
        securitySettings: '安全设置',
        apiKey: '个人API密钥',
        loginOut: '退出登录'
    },
    navigationBar: {
        list: {
            computerRoomManagement: '机房管理',
            modelManagement: '机型管理',
            imageManagement: '镜像管理',
            deviceManagement: '设备管理',
            roleManagement: '角色管理',
            userManagement: '用户管理',
            inBandMonitoringManagement: '带内监控管理',
            surveillanceManagement: '带外监控管理',
            messageCenter: '消息中心',
            userCenter: '个人中心'
        },
        children: {
            hardwareAlarmStatus: '硬件设备状态',
            resourceEcharts: '资源监控图',
            historyAlarmInfo: '历史报警信息',
            allAlarmRules: '全部报警规则',
            faultAlarmLog: '故障报警日志',
            faultAlarmRules: '故障报警规则',
            myMessages: '我的消息',
            messageSet: '消息设置'
        }
    },
    export: {
        title: '自定义导出',
        content: {
            scope: '导出范围：',
            all: '导出全部',
            Selected: '导出选中实例',
            search: '仅导出搜索结果'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    idcList: {
        header: {
            idcList: '机房列表',
            operate: {
                refresh: '刷新',
                setUp: '设置',
                export: '导出'
            }
        },
        content: {
            name: '机房名称',
            englishCode: '机房英文名称',
            grade: '机房等级',
            address: '机房地址',
            createTime: '创建时间',
            createPeople: '创建人',
            updateTime: '更新时间',
            updatePeople: '更新人',
            operate: {
                name: '操作',
                edit: '编辑'
            }
        }
    },
    idcDetail: {
        header: {
            title: '机房列表',
            title2: '机房详情',
            operate: {
                name: '操作',
                edit: '编辑'
            }
        },
        basicInfo: {
            title: '基本信息',
            label: {
                name: '机房名称',
                englishCode: '机房英文名称',
                grade: '机房等级',
                address: '机房地址',
                total: '机柜总数',
                createTime: '创建时间',
                createPeople: '创建人',
                updateTime: '更新时间',
                updatePeople: '更新人'
            }
        },
        allAdministration: {
            title: '全局管理',
            name: '名称',
            loginAccount: '带外',
            loginUserName: '登录账号',
            loginPassWord: '登录密码',
            switchboardOneNum: '交换机1',
            switchboardTwoNum: '交换机2',
        }
    },
    messageDelete: {
        header: {
            messageDelete: '删除消息'
        },
        tips: '删除后消息将无法恢复，您确定要删除吗？'
    },
    myMessage: {
        day90: '前90天提醒',
        day30: '前30天提醒',
        day15: '前15天提醒',
        expired: '已到期提醒',
        header: {
            myMessage: '消息列表',
            title: '(共 {0} 条，其中',
            title2: '条未读)'
        },
        radioGroup: {
            all: '全部',
            read: '已读',
            noRead: '未读'
        },
        search: {
            title: '消息内容',
            placeholder: '请输入消息内容搜索'
        },
        table: {
            expired: '已到期',
            tip15: '',
            tip30: '',
            tip90: '',
            systemMessages: '系统消息',
            faultMessages: '故障消息',
            operationMessages: '操作消息',
            inbondAlert: '报警消息',
            remind: '提醒',
            failureNotification: '失败通知',
            alarmTips: '报警提醒'
        },
        label: {
            messageContent: '消息内容',
            finishTime: '接收时间',
            messageType: '消息类型',
            messageSubType: '消息子类型'
        },
        footer: {
            tips: '请选择消息',
            delete: '删除',
            read: '标为已读'
        }
    },
    customFilter: {
        placeholder: '请输入',
        btn: {
            confirm: '筛选',
            reset: '重置'
        }
    },
    messageSettings: {
        header: {
            messageSettings: '消息设置'
        },
        label: {
            emailPush: '是否关联邮箱推送',
            ipAddress: '邮件服务器',
            port: '邮件服务器端口',
            emailAddress: '邮箱地址',
            emailPassword: '邮箱密码',
            testEmail: '测试邮件接收邮箱'
        },
        placeholder: {
            ipAddress: '请输入您邮箱的smtp服务器地址',
            port: '请输入邮件服务器端口',
            emailAddress: '请输入需要配置的邮箱地址',
            emailPassword: '请输入邮箱密码',
            testEmail: '请输入测试邮件接收邮箱'
        },
        empty: {
            ipAddress: '邮件服务器不能为空',
            port: '邮件服务器端口不能为空',
            emailAddress: '邮箱地址不能为空',
            emailPassword: '邮箱密码不能为空',
            testEmail: '测试邮件接收邮箱不能为空'
        },
        errorTip: {
            port: '只能输入数字，范围1-65535',
            emailAddress: '请输入正确的邮箱地址',
            emailPassword: '请输入正确的邮箱密码',
            testEmail: '请输入正确的邮箱地址'
        },
        tips: {
            pushTip: '开启邮箱推送，用户接收站内信的同时可接收邮箱消息;关闭邮箱推送，用户仅可接收站内信，无法接收邮箱消息',
            checkTip: '用于验证邮箱配置信息是否正常。 单击下方【保存并验证】按钮发送测试邮件，查<br/>看指定邮箱是否可以接收到邮件通知，邮件接收过程存在一定延时，请耐心等待。'
        },
        footer: {
            success: '已验证',
            save: '保存并验证',
            checkInfo: '校验邮箱信息'
        }
    },
    messageDetail: {
        header: {
            title2: '消息详情',
            renewImmediately: '立即续费'
        },
        label: {
            optContent: '操作内容',
            optPeople: '操作人',
            optTime: '操作时间',
            versionName: '版本名称',
            version: '版本号',
            startTime: '开始时间',
            licenseEndTime: '到期时间',
            licenseStatus: '许可状态',
            faultType: '故障类型',
            faultDesc: '故障描述',
            faultAlarmTime: '故障报警时间',
            alarmNum: '报警次数',
        },
        optTips: {
            title: '以下为操作内容详情：',
            label: {
                instanceName: '实例名称',
                instanceId: '实例ID'
            },
            instance: '的实例',
            current: '当前'
        },
        footer: {
            pre: '上一条',
            next: '下一条'
        },
        failureNotification: '失败通知',
        alarmTips: '报警提醒',
        remind: '提醒',
        tip0: '用户{0}您好！',
        tip: '用户{0}您好！操作失败，请重新操作或联系管理员处理',
        tip1: '',
        tip11: '已到期，',
        tip2: '续费后可继续正常使用云平台功能。',
        tip3: '',
        tip4: '以下为操作概览：',
        tip5: '',
        tip6: '以下为故障详情：',
        tip7: '【监控报警】',
        day90: '',
        day30: '',
        day15: '',
        expired: '',
        systemMessages: '系统消息',
        faultMessages: '故障消息',
        operationMessages: '操作消息',
        inbondAlert: '报警消息',
        status: {
            effective: '有效',
            expired: '已过期'
        }
    },
    emailPush: {
        start: {
            title: '开启邮箱推送',
            content: '请确认开启邮箱推送，开启后，用户在接收站内信时将同步接收邮箱提醒。'
        },
        close: {
            title: '关闭邮箱推送',
            content: '请确认关闭邮箱推送，关闭后，用户可接收站内信，将无法接收到邮箱提醒。'
        }
    },
    securityVerification: {
        header: {
            verification: '安全验证'
        },
        label: {
            userName: '用户名',
            passWord: '密码'
        },
        placeholder: {
            userName: '请输入用户名',
            passWord: '请输入密码'
        },
        empty: '密码输入不能为空',
        errTip: {
            userName: '用户名不存在或者密码错误',
            passWord: '用户名不存在或者密码错误',
            currencyPassword: '请输入6-64位字符'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    editIdc: {
        header: {
            name: '编辑机房'
        },
        iptInfo: {
            other: '其它'
        },
        label: {
            computerRoomName: '机房名称',
            computerRoomCode: '机房英文名称',
            computerRoomGrade: '机房等级',
            computerRoomAddress: '机房地址',
            outsideLoginAccountNum: '带外登录账号',
            outsideLoginPsd: '带外登录密码',
            switchboardOneNum: '交换机1登录账号',
            switchboardOnePsd: '交换机1登录密码',
            switchboardTwoNum: '交换机2登录账号',
            switchboardTwoPsd: '交换机2登录密码'
        },
        placeholder: {
            computerRoomName: '请输入机房名称',
            computerRoomCode: '请输入机房英文名称',
            computerRoomGrade: '请选择机房等级',
            customComputerRoomGrade: '请输入自定义机房等级',
            computerRoomAddress: '请输入机房地址',
            outsideLoginAccountNum: '请输入带外登录账号',
            outsideLoginPsd: '请输入带外登录密码',
            switchboardOneNum: '请输入交换机1登录账号',
            switchboardOnePsd: '请输入交换机1登录密码',
            switchboardTwoNum: '请输入交换机2登录账号',
            switchboardTwoPsd: '请输入交换机2登录密码'
        },
        errTip: {
            currency: '请输入1-64位字符',
            currencyPassword: '请输入1-128位字符',
            computerRoomNameEmpty: '机房名称不能为空',
            computerRoomCodeEmpty: '机房英文名称不能为空',
            customComputerRoomGrade: '自定义机房等级输入不能为空',
            computerRoomAddressEmpty: '机房地址输入不能为空',
            outsideLoginAccountNumEmpty: '带外登录账号不能为空',
            outsideLoginPsdEmpty: '带外登录密码不能为空',
            switchboardOneNumEmpty: '交换机1登录账号不能为空',
            switchboardOnePsdEmpty: '交换机1登录密码不能为空',
            switchboardTwoNumEmpty: '交换机2登录账号不能为空',
            switchboardTwoPsdEmpty: '交换机2登录密码不能为空'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    customList: {
        header: {
            name: '自定义列表'
        },
        text: {
            tip: '请选择要显示在列表的信息'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    modelList: {
        header: {
            modelList: '机型列表',
            label: {
                modelName: '机型名称',
                modelSpecifications: '机型规格'
            },
            placeholder: {
                modelName: '请输入机型名称搜索',
                modelSpecifications: '请输入机型规格搜索'
            },
            operate: {
                addModel: '添加机型',
                refresh: '刷新',
                setUp: '设置',
                export: '导出'
            }
        },
        content: {
            modelName: '机型名称',
            modelType: '机型类型',
            machineRoomName: '机房名称',
            machineRoomCode: '机房英文名称',
            architecture: '体系架构',
            bootMode: '引导模式',
            modelSpecification: '机型规格',
            image: '镜像',
            device: '设备',
            CPU: 'CPU',
            storage: '内存',
            sysDisc: '系统盘',
            sysRAID: '系统盘RAID',
            raidConfig: 'RAID配置',
            dataDisc: '数据盘',
            dataRAID: '数据盘RAID',
            networkCard: '网卡',
            networkSettings: '网络设置',
            GPU: 'GPU',
            desc: '描述',
            operate: {
                name: '操作',
                edit: '编辑',
                delete: '删除',
                more: '更多',
                addModel: '添加相同机型'
            }
        },
        edit: '已关联设备的机型不支持编辑',
        tooltip: '无法删除机型，请先删除该机型关联的实例资源'
    },
    uiPagination: {
        total: '共 {0} 条',
        page: '共 {0} 页',
        piecesAndPage: '{0} 条/页'
    },
    addModel: {
        header: {
            title2: '添加机型',
            addModel: '添加机型'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    editModel: {
        header: {
            title2: '编辑机型',
            editModel: '编辑机型'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    deleteModel: {
        header: {
            deleteModel: '删除机型',
            batchDeleteModel: '批量删除机型'
        },
        content: {
            deleteTip: '删除机型后，该机型与镜像关联关系也同步删除，再次添加机型需要重新添加关联镜像，请谨慎操作。',
            deleteSure: '请确定是否删除机型：{0}？'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    modelDetail: {
        header: {
            title: '机型管理',
            title2: '机型详情',
            specifications: '规格参数',
            configParams: '配置参数',
            operate: {
                name: '操作',
                edit: '编辑',
                addModel: '添加相同机型',
                delete: '删除'
            },
            tip: '无法删除机型，请先删除该机型关联的实例资源',
            volumeManagement: {
                title: '卷管理',
                label: {
                    name: '卷名称',
                    type: '卷类型',
                    raidCan: 'RAID配置',
                    raid: 'RAID模式',
                    disk: '磁盘类型',
                    interfaceType: '接口类型',
                    minNum: '单盘最小容量',
                    amount: '最低数量(块)'
                }
            }
        },
        tabs: {
            basicInfo: {
                name: '基本信息',
                label: {
                    modelName: '机型名称',
                    computerRoomName: '机房名称',
                    modelType: '机型类型',
                    architecture: '体系架构',
                    bootMode: '引导模式',
                    modelSpecifications: '机型规格',
                    CPU: 'CPU',
                    storage: '内存',
                    sysDisc: '系统盘',
                    raidConfig: 'RAID配置',
                    sysRAID: '系统盘RAID',
                    dataDisc: '数据盘',
                    dataRAID: '数据盘RAID',
                    networkCard: '网卡',
                    networkSettings: '网络设置',
                    GPU: 'GPU',
                    height: '高度(U)',
                    createTime: '创建时间',
                    desc: '描述'
                }
            },
            relationImage: {
                name: '关联镜像',
                header: {
                    addImage: '添加镜像',
                    tooltip: '仅支持关联20个镜像'
                },
                content: {
                    label: {
                        imageName: '镜像名称',
                        imageType: '镜像类型',
                        architecture: '体系架构',
                        operateSysPlatform: '操作系统平台',
                        operateSysVersion: '操作系统版本',
                        partitionModule: '分区模块',
                        addTime: '添加时间',
                        operate: {
                            name: '操作',
                            delete: '删除'
                        }
                    }
                }
            }
        },
        other: '其它'
    },
    addImage: {
        header: {
            addImage: '添加镜像',
            label: {
                modelName: '机型名称',
                choose: '已选镜像个数'
            }
        },
        content: {
            label: {
                imageName: '镜像名称',
                modelName: '机型名称',
                imageType: '镜像类型',
                architecture: '体系架构',
                operateSysPlatform: '操作系统平台',
                operateSysVersion: '操作系统版本',
                sysPartitionModule: '系统盘分区'
            }
        },
        btn: {
            cancel: '取消',
            save: '保存',
            sure: '确定'
        }
    },
    deleteImage: {
        header: {
            deleteImage: '删除镜像'
        },
        tip: {
            text: '确认删除关联镜像',
            delete: '确定删除镜像'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    imageList: {
        header: {
            imageList: '镜像列表',
            exportImage: '导入镜像',
            operate: {
                refresh: '刷新',
                setUp: '设置',
                export: '导出'
            }
        },
        search: {
            placeholder: {
                imageName: '请输入镜像名称搜索',
                imageId: '请输入镜像ID搜索'
            },
            condition: {
                imageName: '镜像名称',
                imageId: '镜像ID'
            }
        },
        content: {
            label: {
                imageName: '镜像名称',
                imageId: '镜像ID',
                imageType: '镜像类型',
                architecture: '体系架构',
                operateSysPlatform: '操作系统平台',
                operateSysVersion: '操作系统版本',
                format: '镜像格式',
                bootMode: '引导模式',
                sysPartitionModule: '系统盘分区',
                imageDesc: '镜像描述',
                numberOfAssociatedModels: '已关联机型数',
                createTime: '创建时间',
                operate: {
                    name: '操作',
                    delete: '删除'
                }
            }
        },
        filter: {
            presetImage: '预置镜像',
            customImage: '自定义镜像'
        },
        tooltip: {
            title: '无法删除镜像，请先删除该镜像与机型的关联关系',
            image: '预置镜像不支持删除'
        }
    },
    associatedModel: {
        header: {
            title: '关联机型'
        },
        content: {
            select: '选择机型',
            placeholder: '请选择机型',
            addNewModel: '添加新机型',
            table: {
                label: {
                    volumeName: '卷名称',
                    volumeType: '卷类型',
                    raidCan: 'RAID配置',
                    raid: 'RAID模式',
                    diskType: '磁盘类型',
                    interfaceType: '接口类型',
                    minNum: '单盘最小容量',
                    minimumQuantity: '最低数量(块)',
                    associatedDisk: '关联磁盘'
                },
                diskPlaceholder: '请选择磁盘'
            }
        }
    },
    imageDetail: {
        header: {
            title2: '镜像详情',
            operate: {
                title: '操作',
                delete: '删除'
            }
        },
        tooltip: {
            title: '无法删除镜像，请先删除该镜像与机型的关联关系',
            image: '预置镜像不支持删除',
            title2: '当前机型有实例在创建中，无法移除'
        },
        add: {
            model: '添加机型'
        },
        tabs: {
            basicInfo: '基本信息',
            model: '关联机型'
        },
        label: {
            imageName: '镜像名称',
            imageId: '镜像ID',
            imageType: '镜像类型',
            architecture: '体系架构',
            operateSysPlatform: '操作系统平台',
            operateSysVersion: '操作系统版本',
            format: '镜像格式',
            bootMode: '引导模式',
            sysPartitionModule: '系统盘分区',
            imageDesc: '镜像描述',
            numberOfAssociatedModels: '已关联机型数',
            createTime: '创建时间',
        },
        table: {
            label: {
                modelName: '机型名称',
                modelType: '机型类型',
                idcName: '机房名称',
                nameEn: '机房英文名称',
                architecture: '体系架构',
                modelSpecifications: '机型规格',
                desc: '描述',
                operate: {
                    title: '操作',
                    remove: '移除'
                }
            }
        },
        empty: {
            title: '添加机型'
        }
    },
    imageAddModel: {
        header: {
            title: '添加机型',
            imageName: '镜像名称',
            choose: '已选择机型个数'
        },
        label: {
            modelName: '机型名称',
            modelType: '机型类型',
            idcName: '机房名称',
            nameEn: '机房英文名称',
            architecture: '体系架构',
            modelSpecifications: '机型规格',
            desc: '描述'
        },
        btn: {
            save: '保存',
            cancel: '取消'
        }
    },
    removeOperate: {
        header: {
            remove: '移除机型'
        },
        content: {
            deleteTip: '注意:移除后用户创建/重装该机型时无法使用该镜像版本，请谨慎操作!',
            tip1: '确定将【{0}】所关联的【{1}{2}】移除吗？'
        },
        btn: {
            sure: '确定',
            cancel: '取消'
        }
    },
    importImage: {
        header: {
            title: '导入镜像' 
        },
        content: {
            label: {
                imageName: '镜像名称',
                architecture: '体系架构',
                operateSysPlatform: '操作系统平台',
                operateSysVersion: '操作系统版本',
                imageFormat: '镜像格式',
                bootMode: '引导模式',
                systemPartitionTemplate: '系统盘分区',
                imageDesc: '镜像描述',
                chooseImage: '选择镜像'
            },
            placeholder: {
                imageName: '请输入镜像名称',
                architecture: '请选择体系架构',
                operateSysPlatform: '请选择操作系统平台',
                customOperatePlatform: '请输入自定义操作系统平台',
                customVersion: '请输入自定义操作系统版本',
                operateSysVersion: '请选择操作系统版本',
                imageFormat: '请选择镜像格式',
                bootMode: '请选择引导模式',
                imageDesc: '请输入镜像描述',
                chooseImage: '选择镜像'
            }
        },
        emptyTip: {
            imageName: '镜像输入不能为空',
            architecture: '请选择体系架构',
            operateSysPlatform: '请选择操作系统平台',
            customOperatePlatform: '请输入自定义操作系统平台',
            customVersion: '请输入自定义操作系统版本',
            operateSysVersion: '请选择操作系统版本',
            imageFormat: '请选择镜像格式',
            bootMode: '请选择引导模式',
            noData: '暂无数据'
        },
        errTip: {
            upload: '上传中，请勿关闭弹窗。',
            imageName: '请输入1-64个字符',
            imageRepeat: '镜像名称重复',
            numberLimit: '镜像不能为空',
            error: '上传失败'
        },
        btn: {
            cancel: '取消',
            sure: '确定',
            upLoad: '上传'
        }
    },
    collectInfoConfirm: {
        header: {
            title: '重新采集设备信息确认'
        },
        tips: '请确定是否重新执行设备信息重新采集操作？',
        checkbox: '清空RAID',
        tipsCount: {
            tip0: '注意：',
            tip1: '1. 清除RAID操作是针对整个磁盘RAID进行删除操作，该操作不可恢复，请谨慎操作！</br>' +
                  '2. 清除RAID操作会清空服务器磁盘的所有数据，请备份好数据后再执行该操作，以免数据丢失给您造成损失。</br>'
        }
    },
    clearRaidConfirm: {
        header: {
            title: '清空RAID操作确认',
        },
        tipsCount: {
            tip0: '注意：',
            tip1: '1. 清除RAID操作是针对整个磁盘RAID进行删除操作，该操作不可恢复，请谨慎操作！</br>' +
                  '2. 清除RAID操作会清空服务器磁盘的所有数据，请备份好数据后再执行该操作，以免数据丢失给您造成损失。</br>',
            tip2: '您选择的设备【{0}】将执行清空RAID操作，您是否确定操作？'
        },
    },
    deviceList: {
        header: {
            deviceList: '设备列表'
        },
        radioBtn: {
            computed: '计算型',
            storage: '存储型',
            GPU: 'GPU',
            other: '其它'
        },
        batchOperate: {
            btn: {
                open: '开机',
                close: '关机',
                restart: '重启',
                recovery: '回收实例',
                edit: '编辑实例名称',
                resetPassword: '重置密码',
                more: '更多'
            },
            tip: {
                default: '请选择资源',
                turnOn: '仅关机状态实例可以开机',
                resetPassword: '仅关机状态实例可以重置密码',
                restart: '仅运行状态实例可以重启',
                turnOff: '仅运行状态实例可以关机',
                instanceName: '仅已创建实例支持编辑',
                lock: '实例已锁定',
                recovery: '只有处于“已关机”和“创建失败”的设备能进行“回收实例”操作'
            }
        },
        tooltip: {
            up: '只有处于“已入库”状态下的设备才能上架',
            model: '请先关联机型',
            down: '只有处于“已上架”状态下的设备才能下架',
            delete: '只有处于“已入库”状态下的服务器才能删除',
            remove: '只有处于“运行中”和“已关机”的实例能进行“设备移除”操作',
            lock: '实例已锁定',
            lock1: '无法锁定该实例',
            lock2: '无法解锁该实例',
            recovery: '只有处于“已关机”和“创建失败”的设备能进行“回收实例”操作',
            reset: '只有实例操作失败的设备能进行“重置实例状态”操作',
            turnOn: '仅关机状态实例可以开机',
            restart: '仅运行状态实例可以重启',
            turnOff: '仅运行状态实例可以关机',
            resetSystem: '仅关机状态实例可以重装系统',
            resetPassword: '仅关机状态实例可以重置密码'
        },
        instanceStatus: {
            creating: '创建中',
            starting: '开机中',
            running: '运行',
            stopping: '关机中',
            stopped: '关机',
            restarting: '重启中',
            resetting_password: '重置密码中',
            destroying: '销毁中',
            destroyed: '已销毁',
            error: '错误',
            upgrading: '调整配置中',
            reinstalling: '重装系统中'
        },
        operate: {
            instanceOperate: '实例操作',
            error: {
                tip: '操作失败，请联系管理员处理'
            },
            exportDevice: '导入设备',
            placeholder: {
                sn: '请输入SN搜索',
                instanceOwner: '请输入实例属主搜索'
            },
            label: {
                sn: 'SN',
                instanceOwner: '实例属主'
            },
            refresh: '刷新',
            setUp: '设置',
            export: '导出'
        },
        content: {
            label: {
                sn: 'SN',
                name: '机型名称',
                modelName: '机型类型',
                managementStatus: '管理状态',
                collectionStatus: '采集状态',
                computerRoomName: '机房名称',
                cabinetCode: '机柜编码',
                uBit: '所在U位',
                brand: '品牌',
                model: '型号',
                architecture: '体系架构',
                instanceImage: '镜像',
                instanceName: '实例名称',
                instanceID: '实例ID',
                instanceStatus: '实例状态',
                lockStatus: '锁定状态',
                instanceOwner: '实例属主',
                CPU: 'CPU',
                storage: '内存',
                sysDisc: '系统盘',
                dataDisc: '数据盘',
                GPU: 'GPU',
                switchIPNetworkPortOne: '网口1交换机IP',
                switchIPNetworkPortTwo: '网口2交换机IP',
                networkPortOneUplinkPort: '网口1（eth0）上联端口',
                networkPortTwoUplinkPort: '网口2（eth1）上联端口',
                networkSettings: '网络设置',
                outOfBandIP: '带外IP',
                intranceIPv4: '内网IPv4(eth0)',
                intranceIPv4First: '内网IPv4(eth1)',
                intranetIPv6: '内网IPv6(eth0)',
                intranetIPv6First: '内网IPv6(eth1)',
                createTime: '创建时间',
                desc: '描述',
                remark: '备注',
                operate: {
                    name: '操作',
                    up: '上架',
                    down: '下架',
                    more: '更多',
                    lock:'锁定',
                    unlock:'解锁',
                    delete: '设备删除',
                    reset: '重置实例状态',
                    remove: '设备移除',
                    recovery: '回收实例',
                    open: '开机',
                    close: '关机',
                    restart: '重启',
                    resetSystem: '重装系统',
                    resetPassword: '重置密码'
                }
            }
        },
        managementStatus: {
            warehousing: '已入库',
            onTheShelf: '已上架',
            created: '已创建',
            onShelf: '上架中',
            remove: '已移除',
            unputaway: '已下架',
            lowerShelf: '下架中',
            error: '上架失败'
        },
        status: {
            collected: '已采集',
            noCollected: '未采集',
            collecting: '采集中',
            collectionFailed: '采集失败'
        }
    },
    resetSystemConfirm: {
        header: {
            title: '重装系统须知',
            tip0: '注意：',
            tip: '' +
                '1. 重装系统是针对系统盘做初始化、更换操作系统的操作，该操作不可恢复，请谨慎操作！</br>' +
                '2. 该操作会清空实例系统盘所有数据，请备份好数据后再执行重装系统操作，以免数据丢失给您造成损失。</br>' +
                '3. 重装系统后，系统默认不会为实例配置IPv6地址，您需要手工进行配置。',
            tip2: '您选择的实例【{0}】将执行重装系统操作，您是否确定操作？'
        }
    },
    resetSystem: {
        header: {
            title: '重装系统配置',
            tip0: '注意：',
            tip: '' +
                '1. 重装系统是针对系统盘做初始化、更换操作系统的操作，该操作不可恢复，请谨慎操作！</br>' +
                '2. 该操作会清空实例系统盘所有数据，请备份好数据后再执行重装系统操作，以免数据丢失给您造成损失。</br>' +
                '3. 重装系统后，系统默认不会为实例配置IPv6地址，您需要手工进行配置。',
            instanceInfo: '实例信息',
            imageType: '镜像类型',
            system: '系统盘',
            systemPartition: '系统盘分区',
            systemVolume: '系统卷',
            systemVolumePartition: '系统卷分区',
            bootMode: '引导模式',
            hostName: 'HostName',
            userName: '用户名',
            setPassword: '设置密码',
            password: '密码',
            confirmPassword: '确认密码',
            performanceMonitoring: '性能监控',
            agent: '安装监控Agent'
        },
        table: {
            label: {
                instanceName: '实例名称',
                instanceID: '实例ID',
                operatingSystem: '操作系统',
                createdTime: '创建时间',
                sshName: '密钥名称'
            }
        },
        placeholder: {
            select: '请选择',
            hostName: '可选填，操作系统内部的计算机名'
        },
        regCheckTip: {
            hostName: '长度为 2-64 个字符，允许使用点号(.)分隔 字符成多段，每段允许使用大小写字母、数字或连字符(-)，但不能连续使用点号(.)或连字符(-)。不能以点号(.)或连字符(-)开头或结尾',
            password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
            password2: '密码不一致，请重新输入'
        },
        emptyCheck: {
            sshKey: '密钥不能为空',
            password: '密码不能为空',
            confirmPassword: '确认密码不能为空',
            repeat: '密码不一致，请重新输入'
        },
        radio: {
            customPassword: '自定义密码',
            sshKeyLogin: '密钥登录'
        },
        tip1: '提供实时故障告警'
    },
    resetPassword: {
        header: {
            title: '重置密码',
            tip: '注意：重置密码后，需要在控制台/运营平台对实例进行开机才能生效',
            resetPassword1: '您选择的 ',
            resetPassword2: ' 台实例将执行重置密码操作，您是否确定操作？',
        },
        table: {
            label: {
                instanceName: '实例名称',
                instanceId: '实例ID',
                outOfBandIP: '带外IP',
                intranceIPv4: '内网IPv4(eth0)'
            }
        },
        ipt: {
            label: {
                password: '密码',
                newPassword: '新密码',
                confirmPassword: '确认密码'
            },
            errorTip: {
                password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
                newPassword: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
                confirmPassword: '密码不一致，请重新输入'
            },
            emptyTip: {
                password: '密码不能为空',
                newPassword: '新密码不能为空',
                confirmPassword: '确认密码不能为空',
            }
        }
    },
    batchOperate: {
        instanceName: {
            title: '批量编辑实例名称',
            tip: '注意：通过此方式编辑的多台实例名称为连续命名',
            tip1: '您选择的 ',
            tip2: ' 台实例将执行批量编辑实例名称操作，您是否确定操作？',
            label: {
                instanceName: '实例名称',
                instanceId: '实例ID',
                outOfBandIP: '带外IP',
                intranceIPv4: '内网IPv4(eth0)',
                newInstanceName: '新实例名称'
            },
            iptTip: {
                text: '2-128个字符，以大小写字母或中文开头，可包含数字、“.”、“_”、“:”或“-”',
                empty: '实例名称不能为空',
                repeat: '实例名称重复'
            }
        },
        recoveryInstance: {
            title: '批量回收实例',
            tip: '注意：回收实例将会清空实例上的所有数据，请谨慎操作。',
            recovery1: '您选择的 ',
            recovery2: ' 台实例将执行回收实例操作，您是否确定操作？',
            label: {
                instanceName: '实例名称',
                instanceId: '实例ID',
                outOfBandIP: '带外IP',
                intranceIPv4: '内网IPv4(eth0)'
            },
        },
        openCloseRestart: {
            open: '批量开机',
            close: '批量关机',
            restart: '批量重启',
            open1: '您选择的 ',
            open2: ' 台实例将执行开机操作，您是否确定操作？',
            close1: '您选择的 ',
            close2: ' 台实例将执行关机操作，您是否确定操作？',
            restart1: '您选择的 ',
            restart2: ' 台实例将执行重启操作，您是否确定操作？',
            label: {
                instanceName: '实例名称',
                instanceId: '实例ID',
                outOfBandIP: '带外IP',
                intranceIPv4: '内网IPv4(eth0)'
            }
        },
        resetPassword: {
            header: {
                title: '批量重置密码',
                tip: '注意：</br>' +
                    '1. 通过此方式重置密码的多个实例密码均相同</br>' +
                    '2. 重置密码后，需要在控制台/运营平台对实例进行开机才能生效',
                resetPassword1: '您选择的 ',
                resetPassword2: ' 台实例将执行重置密码操作，您是否确定操作？',
            },
            table: {
                label: {
                    instanceName: '实例名称',
                    instanceId: '实例ID',
                    outOfBandIP: '带外IP',
                    intranceIPv4: '内网IPv4(eth0)'
                }
            },
            ipt: {
                label: {
                    password: '密码',
                    newPassword: '新密码',
                    confirmPassword: '确认密码'
                },
                errorTip: {
                    password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
                    newPassword: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
                    confirmPassword: '密码不一致，请重新输入'
                },
                emptyTip: {
                    password: '密码不能为空',
                    newPassword: '新密码不能为空',
                    confirmPassword: '确认密码不能为空',
                }
            }
        }
    },
    lockInstance: {
        header: {
            lock:'锁定',
            unlock:'解锁',
        },
        lockTip:'确定锁定实例【{0}】?',
        unlockTip:'确定解锁实例【{0}】?',
        tip: '锁定实例后“删除”操作不可用，如需使用需提前解锁',
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    deviceDetail: {
        header: {
            deviceList: '设备列表',
            title2: '设备详情'
        },
        operate: {
            name: '操作',
            up: '上架',
            down: '下架',
            open: '开机',
            close: '关机',
            restart: '重启',
            reset: '重置实例状态',
            remove: '设备移除',
            recovery: '回收实例',
            lock:'锁定',
            unlock:'解锁',
            delete: '设备删除',
            resetSystem: '重装系统',
            resetPassword: '重置密码',
            error: {
                tip: '操作失败，请联系管理员处理'
            }
        },
        tooltip: {
            up: '只有处于“已入库”状态下的设备才能上架',
            model: '请先关联机型',
            associatedModels: '只有处于“已入库”状态下的设备才能关联机型',
            down: '只有处于“已上架”状态下的设备才能下架',
            delete: '只有处于“已入库”状态下的服务器才能删除',
            remove: '只有处于“运行中”和“已关机”的实例能进行“设备移除”操作',
            recovery: '只有处于“已关机”和“创建失败”的设备能进行“回收实例”操作',
            lock: '实例已锁定',
            lock1: '无法锁定该实例',
            lock2: '无法解锁该实例',
            reset: '只有实例操作失败的设备能进行“重置实例状态”操作',
            turnOn: '仅关机状态实例可以开机',
            restart: '仅运行状态实例可以重启',
            turnOff: '仅运行状态实例可以关机',
            resetPassword: '仅关机状态实例可以重置密码',
            resetSystem: '仅关机状态实例可以重装系统'
        },
        deviceInfo: {
            deviceInfo: '设备信息',
            sn: 'SN',
            modelName: '机型名称',
            recapture: '重新采集',
            managementStatus: '管理状态',
            collectionStatus: '采集状态',
            computerRoomName: '机房名称',
            cabinetCode: '机柜编码',
            uBit: '所在U位',
            remark: '备注'
        },
        instanceInfo: {
            instanceInfo: '实例信息',
            instanceName: '实例名称',
            instanceID: '实例ID',
            instanceImage: '镜像',
            instanceStatus: '实例状态',
            monitorAgentStatus: '监控Agent状态',
            lockStatus: '锁定状态',
            instanceOwner: '实例属主',
            createTime: '创建时间',
            desc: '描述'
        },
        hardwareConfiguration: {
            modelInfo: '机型信息',
            hardwareConfiguration: '硬件配置',
            brand: '品牌',
            model: '型号',
            architecture: '体系架构',
            CPU: 'CPU',
            storage: '内存',
            sysDisc: '系统盘',
            adapeterId: 'adapeter_id（RAID卡）',
            enclosure1: '系统盘1背板号（enclosure）',
            slot1: '系统盘1槽（slot）',
            enclosure2: '系统盘2背板号（enclosure）',
            slot2: '系统盘2槽（slot）',
            dataDisc: '数据盘',
            GPU: 'GPU',
            networkCard: '网卡',
            networkSettings: '网络设置'
        },
        networkInfo: {
            networkInfo: '网络信息',
            intranceIPv4: '内网IPv4（eth0）',
            intranceIPv4First: '内网IPv4（eth1）',
            intranetIPv6: '内网IPv6（eth0）',
            intranetIPv6First: '内网IPv6（eth1）',
            mac1: 'MAC1（eth0）',
            mac2: 'MAC2（eth1）',
            switchIPOne: '交换机IP1（eth0）',
            switchIPTwo: '交换机IP2（eth1）',
            switchUplinkPortOne: '交换机上联端口1（eth0）',
            switchUplinkPortTwo: '交换机上联端口2（eth1）',
            subnetMask: '子网掩码（eth0）',
            subnetMaskFirst: '子网掩码（eth1）',
            networkGateway: '网关'
        },
        switchOutsideInfo: {
            switchOutsideInfo: '交换机&带外信息',
            userName1: '网口1交换机登录用户名',
            password1: '网口1交换机登录密码',
            userName2: '网口2交换机登录用户名',
            password2: '网口2交换机登录密码',
            outOfBandIP: '带外IP',
            outOfBandLoginUserName: '带外登录用户名',
            outOfBandLoginPassWord: '带外登录密码'
        },
        editDesc: {
            header: {
                title: '描述',
                title2: '备注'
            },
            count: {
                title: '修改描述：',
                title2: '修改备注：'
            },
            placeholder: '请输入描述',
            placeholder2: '请输入备注',
            btn: {
                cancel: '取消',
                sure: '确定'
            }
        },
        tabs: {
            baseInfo: '基本信息',
            diskDetail: '磁盘详情',
            hardwareMonitoring: '硬件监控',
            performanceMonitoring: '性能监控',
            operatLog: '操作日志'
        },
        diskDetail: {
            physicalDiskView: '物理盘视图',
            table: {
                label: {
                    disk: '磁盘',
                    singleDiskCapacity: '单盘容量',
                    diskType: '磁盘类型',
                    interfaceType: '接口类型',
                    enclosure: '背板号(enclosure)',
                    slot: '槽位(slot)'
                }
            },
            logicalDriveView: '逻辑盘视图',
            table2: {
                label: {
                    driveLetterName: '盘符名称',
                    singleDiskCapacity: '单盘容量',
                    diskType: '磁盘类型',
                    interfaceType: '接口类型',
                }
            },
            modelRolInfo: '机型卷信息',
            modelInfo: '机型信息',
            modelName: '机型名称',
            btn: '关联机型',
            volumeManager: {
                table: {
                    label: {
                        volumeName: '卷名称',
                        volumeType: '卷类型',
                        raid: 'RAID模式',
                        diskType: '磁盘类型',
                        interfaceType: '接口类型',
                        minNum: '单盘最小容量',
                        minimumQuantity: '最低数量(块)',
                        associatedDisk: '关联磁盘'
                    }
                }
            }
        },
        hardwareMonitoring: {
            hardwareDeviceStatus: '硬件设备状态',
            alarmLogList: '报警日志列表',
            status: {
                title: {
                    cpu: 'CPU',
                    storage: '内存',
                    hardDisk: '硬盘',
                    networkCard: '网卡',
                    powerSupply: '电源',
                    other: '其他'
                },
                normal: '正常',
                error: '异常'
            }
        },
        table: {
            label: {
                serialNumber: '序号',
                faultType: '故障类型',
                faultDesc: '故障描述',
                faultAlarmTime: '故障报警时间',
                updateTime: '更新时间',
                alarmNum: '报警次数',
                status: '状态'
            }
        },
        status: {
            collected: '已采集',
            noCollected: '未采集',
            collecting: '采集中',
            collectionFailed: '采集失败'
        },
        operatLog: {
            header: {
                title: '操作日志',
                tips: '(以下列表包括了近90天的操作记录)'
            },
            search: {
                label: {
                    operatePeople: '操作人',
                    operateTime: '操作时间',
                },
                placeholder: {
                    operatePeople: '请输入操作人',
                    startDate: '开始日期',
                    endDate: '结束日期'
                },
                btn: {
                    search: '搜索',
                    clear: '清空'
                }
            },
            table: {
                label: {
                    logId: 'LogID',
                    operateName: '操作名称',
                    operatePeople: '操作人',
                    operateTime: '操作时间',
                    operateFeedback: '操作反馈',
                    failCode: '错误码'
                }
            }
        }
    },
    editDesc: {
        header: {
            title: '描述',
            title2: '备注'
        },
        count: {
            title: '修改描述：',
            title2: '修改备注：'
        },
        placeholder: '请输入描述',
        placeholder2: '请输入备注',
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    removeRecovery: {
        header: {
            remove: '设备移除',
            recovery: '回收实例'
        },
        tooltip: {
            remove: '请确定是否移除以下设备？',
            recovery: '请确定是否回收以下实例？'
        },
        checkbox: {
            remove: '我已知晓以上注意事项并同意设备移除',
            recovery: '我已知晓以上注意事项并同意回收实例'
        },
        operateTip: {
            remove: '注意：</br>' +
                '1. 设备移除操作后控制台将不显示此实例，同时此实例不会进入库存，无法再次购买</br>' +
                '2. 自动解除设备和实例属主的关联关系',
            recovery: '注意：</br>' +
                '1. 回收实例操作完成后设备重新入库，自动解除设备和实例属主的关联关系</br>' +
                '2. 为避免数据丢失，请确认数据已备份，回收实例操作会清空实例上所有数据（包含系统盘和数据盘），彻底删除实例且无法恢复'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    resetInstance: {
        header: {
            title: '重置实例状态'
        },
        tooltip: {
            title: '注意：</br>' + '重置实例状态操作实例状态重置为上一个实例稳定状态',
            count: {
                tip1: '确定将',
                tip2: '重置实例状态？'
            }
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    importDevice: {
        header: {
            title: '导入设备'
        },
        label: {
            computerRoomName: '机房名称',
            modelName: '机型名称',
            deviceInfo: '设备信息'
        },
        placeholder: {
            computerRoomName: '请选择机房',
            modelName: '请选择机型',
        },
        errTip: {
            upload: '上传中，请勿关闭弹窗。',
            computerRoomName: '请选择机房名称',
            modelName: '请选择机型名称',
            deviceInfo: '请选择设备信息',
            error: '上传失败',
        },
        noData: {
            tip: '暂无数据'
        },
        operate: {
            tip1: '如无合适机型，可以',
            addNewModel: '添加新机型',
            templateDownLoad: '模板下载'
        },
        btn: {
            selectionFile: '选取文件',
            cancel: '取消',
            sure: '确定'
        },
        tooltip: {
            first: '1. 仅支持选取1个附件，且大小不超过20M；',
            two: '2. 为保证上传成功，请勿对模板中的单元格进行更改；',
            third: '3. 仅支持一个文件上传，多个文件上传会替换原文件；'
        }
    },
    upDownFrame: {
        operate: {
            addNewModel: '添加新机型',
        },
        placeholder: {
            tip1: '请选择',
            tip2: '关联磁盘不能为空'
        },
        table: {
            label: {
                volumeName: '卷名称',
                volumeType: '卷类型',
                raidCan: 'RAID配置',
                raid: 'RAID模式',
                diskType: '磁盘类型',
                interfaceType: '接口类型',
                minNum: '单盘最小容量',
                amount: '最低数量(块)',
                associatedDisk: '关联磁盘'
            }
        },
        noData: {
            tip: '暂无数据'
        },
        steps: {
            title1: '关联机型',
            title2: '确认设备'
        },
        select: {
            model: '选择机型',
            placeholder: '请选择机型'
        },
        titleTip: {
            up: '上架',
            down: '下架',
            delete: '设备删除'
        },
        countTitleTip: {
            down: '注意：设备下架成功后，服务器状态将变更为“已入库”',
            delete: '注意：设备删除执行后无法撤销，如需再次使用需要重新导入'
        },
        headerTitleTip: {
            up: '设备上架后才可创建使用，请确定是否上架以下设备？',
            down: '请确定是否下架以下设备？',
            delete: '请确定是否删除以下设备？'
        },
        btn: {
            pre: '上一步',
            next: '下一步',
            cancel: '取消',
            sure: '确定'
        }
    },
    publicTable: {
        label: {
            sn: 'SN',
            instanceName: '实例名称',
            instanceId: '实例ID',
            instanceOwner: '实例属主',
            cabinetCode: '机柜编码',
            uBit: '所在U位',
            outOfBandIP: '带外IP',
            intranceIPv4: '内网IPv4(eth0)'
        }
    },
    openShutDownRestart: {
        titleTip: {
            open: '开机',
            close: '关机',
            restart: '重启'
        },
        tooltipInfo: {
            open: '确定开启实例：',
            close: '确定关闭实例：',
            restart: '确定重启实例：'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    modelForm: {
        header: {
            params: {
                text1: '常规参数',
                text2: '配置参数'
            }
        },
        config: {
            configured: '需配置RAID',
            notConfigured: '无需配置RAID'
        },
        radioBtn: {
            computed: '计算型',
            storage: '存储型',
            GPU: 'GPU',
            other: '其它'
        },
        netWork: {
            single: '单网口',
            dual: '双网口',
            bond: '单网口bond'
        },
        specificationsType: {
            presetSpecifications: '预置规格',
            otherSpecifications: '其它规格'
        },
        tooltip: {
            tip: ' 以机型{0}为模板添加新机型',
            title: '机型规格与机型名称一一对应，建议可以使用【规格属性】+【实例类型】+【体系架构】。</br>例:c1.medium.x86'
        },
        noData: {
            tip: '暂无数据'
        },
        list: {
            manufacturer: '龙芯3号'
        },
        raid: '单盘RAID0',
        none: '无',
        unrestricted: '无限制',
        label: {
            computerRoomName: '机房名称',
            modelType: '机型类型',
            modelName: '机型名称',
            architecture: '体系架构',
            bootMode: '引导模式',
            modelSpecifications: '机型规格',
            desc: '描述',
            cpu: 'CPU',
            cpuParams: 'CPU参数',
            processorManufacturer: '处理器厂商',
            modelChoose: '型号',
            numberPhysicalCores: '物理核数',
            dominantFrequency: '主频(GHz)',
            system: '系统卷',
            card: '阵列卡',
            data: '数据卷',
            numberOfRoutes: '数量',
            storage: '内存',
            storageParams: '内存参数',
            interface: '接口',
            dominantFrequencyMHz: '主频(MHz)',
            capacity: '容量(GB)',
            modelNumber: '数量',
            sysType: '系统盘类型',
            sysInterfaceType: '系统盘接口类型',
            sysSingleCapacityVal: '系统盘单块容量',
            sysNumber: '系统盘数量(块) ',
            raidConfig: 'RAID配置',
            sysRAID: '系统盘RAID',
            dataType: '数据盘类型',
            dataInterfaceType: '数据盘接口类型',
            dataSingleCapacityVal: '数据盘单块容量',
            dataNumber: '数据盘数量(块) ',
            networkSpeed: '网卡速度(GE)',
            networkNumber: '网口数量(个)',
            networkSettings: '网口设置',
            gpuBrand: 'GPU(品牌)',
            gpuModel: 'GPU(型号)',
            gpuNumber: 'GPU数量(个)',
            heightU: '高度(U)',
            volumeManage: {
                title: '卷管理',
                label: {
                    name: '卷名称',
                    type: '卷类型',
                    raidCan: 'RAID配置',
                    raid: 'RAID模式',
                    disk: '磁盘类型',
                    interfaceType: '接口类型',
                    minNum: '单盘最小容量',
                    amount: '最低数量(块)',
                    operate: '操作'
                },
                placeholder: {
                    tip1: '请输入'
                },
                empty: {
                    name: '卷名称不能为空',
                    type: '请选择卷类型',
                    raidCan: '请选择RAID配置',
                    raid: '请选择RAID模式',
                    disk: '请选择磁盘类型',
                    interfaceType: '请选择接口类型',
                    minNum: '单盘最小容量不能为空',
                    amount: '数量不能为空',
                },
                save: {
                    volumeManagement: '请保存卷管理',
                    empty: '卷管理不能为空'
                },
                sys: {
                    tip1: '至少添加一个系统卷',
                    tip2: '只能添加一个系统卷'
                },
                btn: {
                    cancel: '取消',
                    add: '添加卷',
                    sure: '确定',
                    save: '保存',
                    edit: '编辑',
                    delete: '删除'
                }
            }
        },
        placeholder: {
            modelName: '请输入机型名称（1-64个字符）',
            modelSpecifications: '请输入机型规格',
            desc: '请输入描述',
            select: '请选择',
            processorManufacturer: '请选择处理器厂商',
            modelChoose: '请输入型号',
            dominantFrequency: '请选择主频(GHz)',
            interface: '请选择接口',
            dominantFrequencyMHz: '请选择主频(MHz)',
            capacity: '请选择容量(GB)',
            sysType: '请选择系统盘类型',
            sysInterfaceType: '请选择系统盘接口类型',
            sysSingleCapacityVal: '请输入系统盘单块容量',
            raidConfig: '请选择RAID配置',
            sysRAID: '请选择系统盘RAID',
            dataType: '请选择数据盘类型',
            dataInterfaceType: '请选择数据盘接口类型',
            dataSingleCapacityVal: '请输入数据盘单块容量',
            networkSpeed: '请选择网卡速度(GE)',
            networkSettings: '请选择网口设置',
            gpuBrand: '请选择GPU(品牌)',
            gpuModel: '请选择GPU(型号)',
            heightU: '请选择高度(U)'
        },
        emptyTip: {
            modelName: '机型名称输入不能为空',
            modelSpecifications: '机型规格输入不能为空',
            cpu: 'CPU不能为空',
            processorManufacturer: '处理器厂商不能为空',
            modelChoose: '型号输入不能为空',
            dominantFrequency: '主频(GHz)不能为空',
            storage: '内存不能为空',
            interface: '接口不能为空',
            dominantFrequencyMHz: '主频(MHz)不能为空',
            capacity: '容量(GB)不能为空',
            sysType: '系统盘类型不能为空',
            sysInterfaceType: '系统盘接口类型不能为空',
            sysSingleCapacityVal: '系统盘单块容量输入不能为空',
            raidConfig: 'RAID配置不能为空',
            sysRAID: '系统盘RAID不能为空',
            dataType: '数据盘类型不能为空',
            dataInterfaceType: '数据盘接口类型不能为空',
            dataSingleCapacityVal: '数据盘单块容量输入不能为空',
            networkSpeed: '网卡速度(GE)不能为空',
            networkSettings: '网络设置不能为空',
            gpuBrand: 'GPU(品牌)不能为空',
            gpuModel: 'GPU(型号)不能为空',
            heightU: '高度(U)不能为空'
        },
        errorTip: {
            number: '1-64字符，输入内容为数字，允许输入小数点',
            number1: '只能输入数字',
            title: '请输入合法的手机号码',
            number2: '只能输入数字，小数保留2位',
            number3: '超出最大限制',
            currency: '1-64字符，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”"',
            currency2: '请输入1-64个字符',
            modelSpecifications: '机型规格与机型名称不能重复',
            name: '机型名称重复',
            specifications: '机型规格重复'
        }
    },
    resourceMonitor: {
        header: {
            resourceMonitor: '资源监控图'
        },
        search: {
            userName: '用户名',
            resourceType: '资源类型',
            idcName: '机房',
            instanceId: '实例ID',
            dimensions: '维度',
            monitoringMetrics: '监控指标',
            btn: {
                search: '搜索',
                clear: '清空'
            },
            count: {
                resourceType: '实例',
                dimensions: '实例'
            }
        },
        placeholder: {
            userName: '请输入用户名',
            resourceType: '请选择资源类型',
            idcName: '请选择机房',
            instanceId: '请输入实例ID',
            dimensions: '请选择维度',
            monitoringMetrics: '请选择监控指标'
        },
        errorTip: {
            userName: '请输入用户名',
            resourceType: '请选择资源类型',
            idcName: '请选择机房',
            instanceId: '请输入实例ID',
            dimensions: '请选择维度',
            monitoringMetrics: '请选择监控指标'
        }
    },
    historyAlarmInfo: {
        header: {
            historyAlarmInfo: '历史报警信息'
        },
        search: {
            label: {
                userName: '用户名',
                alarmTime: '报警时间'
            },
            placeholder: {
                userName: '请输入用户名',
                startDate: '开始日期',
                endDate: '结束日期'
            },
            errorTip: {
                userName: '请输入用户名',
                alarmTime: '请选择报警时间'
            },
            btn: {
                search: '搜索',
                clear: '清空'
            }
        },
        table: {
            label: {
                name: '用户名',
                alarmTime: '报警时间',
                resourceName: '资源类型',
                resourceId: '报警资源ID',
                triggerDescription: '触发条件',
                alertValue: '报警值',
                alertLevelDescription: '报警级别',
                alertPeriod: '持续时间',
                userName: '通知对象'
            },
            count: {
                minute: '分钟'
            }
        }
    },
    allAlarmRulesList: {
        header: {
            allAlarmRulesList: '全部报警规则'
        },
        select: {
            alarm: '报警',
            normal: '正常',
            forbid: '已禁用',
        },
        search: {
            label: {
                userName: '用户名',
                ruleName: '规则名称',
                ruleId: '规则ID'
            },
            placeholder: {
                userName: '请输入用户名',
                ruleName: '请输入规则名称',
                ruleId: '请输入规则ID'
            },
            errorTip: {
                userName: '请输入用户名',
                ruleName: '请输入规则名称',
                ruleId: '请输入规则ID'
            },
            btn: {
                search: '搜索',
                clear: '清空'
            }
        },
        table: {
            label: {
                ruleName: '规则名称',
                userName: '用户名',
                ruleId: '规则ID',
                resourceName: '资源类型',
                instanceCount: '实例关联数量',
                triggerDescription: '触发条件',
                statusName: '状态',
                operate: {
                    title: '操作',
                    lookDetail: '查看详情',
                    alarmHistory: '报警历史'
                }
            }
        }
    },
    allAlarmRulesDetail: {
        header: {
            title2: '报警规则详情',
            operate: {
                alarmHistory: '报警历史'
            }
        },
        count: {
            minute: '分钟',
            hours: '小时'
        },
        baseInfo: {
            title: '基本信息',
            label: {
                ruleId: '规则ID',
                ruleName: '规则名称',
                resourceType: '资源类型'
            }
        },
        alarmResources: {
            title: '报警资源',
            table: {
                label: {
                    instanceName: '实例名称',
                    instanceId: '实例ID',
                    ipv4: '内网IPV4(eth0)',
                    ipv6: '内网IPV6(eth0)'
                }
            }
        },
        triggeringConditions: {
            title: '触发条件'
        },
        notificationStrategy: {
            title: '通知策略',
            label: {
                notificationCycle: '通知周期',
                effectivePeriod: '有效时段',
                notificationConditions: '通知条件',
                receivingChannel: '接收渠道',
                notificationObject: '通知对象'
            }
        },
        notice: {
            email: '邮件',
            internalMessage: '站内信'
        }
    },
    hardwareStatusList: {
        header: {
            hardwareStatusList: '硬件设备状态'
        },
        search: {
            label: {
                idcName: '机房',
                sn: 'SN'
            },
            placeholder: {
                idcName: '请选择',
                sn: '请输入SN号'
            },
            btn: {
                search: '搜索',
                clear: '清空'
            }
        },
        table: {
            label: {
                sn: 'SN',
                idc: '机房',
                brand: '品牌',
                model: '型号',
                managementStatus: '管理状态',
                cpu: 'CPU',
                storage: '内存',
                hardDisk: '硬盘',
                networkCard: '网卡',
                powerSupply: '电源',
                other: '其他',
                operate: '操作',
                opt: '历史故障'
            }
        }
    },
    surveillance: {
        status: {
            all: '全部',
            no: '未处理',
            yes: '已处理'
        },
        accessory: {
            cpu: 'CPU',
            mem: '内存',
            hardDisk: '硬盘',
            networkCard: '网卡',
            powerSupply: '电源',
            voltage: '电压',
            fan: '风扇',
            temperature: '温度',
            pcie: 'PCIE 总线'
        },
    },
    faultLogList: {
        header: {
            faultLogList: '故障报警日志'
        },
        search: {
            label: {
                idcName: '机房',
                sn: 'SN',
                level: '故障等级',
                time: '首次故障报警时间',
                accessory: '故障配件',
                status: '状态'
            },
            placeholder: {
                idcName: '请选择',
                sn: '请输入SN号',
                level: '请选择',
                startDate: '开始日期',
                endDate: '结束日期',
                accessory: '请选择',
                status: '请选择'
            },
            btn: {
                search: '搜索',
                clear: '清空'
            }
        },
        table: {
            label: {
                logId: 'LogID',
                sn: 'SN',
                idc: '机房',
                level: '故障等级',
                type: '故障类型',
                accessory: '故障配件',
                iloip: 'iloIP',
                desc: '故障描述',
                originalLog: '故障原始日志',
                status: '状态',
                numberOfAlarms: '报警次数',
                faultDetectionTime: '带外故障发现时间',
                faultAlarmTime: '首次故障报警时间',
                updateTime: '更新时间',
                operate: '操作'
            },
            operate: {
                title: '处理',
                tip: '仅未处理的故障日志可处理'
            }
        }
    },
    faultRulesList: {
        header: {
            faultRulesList: '故障报警规则'
        },
        search: {
            label: {
                faultName: '故障名称',
                faultAccessories: '故障配件',
                faultLevel: '故障等级'
            },
            placeholder: {
                faultName: '请输入故障名称',
                faultAccessories: '请选择',
                faultLevel: '请选择'
            },
            btn: {
                search: '搜索',
                clear: '清空',
                restoreDefaultSet: '恢复默认设置'
            }
        },
        table: {
            label: {
                faultName: '故障名称',
                faultAccessories: '故障配件',
                faultType: '故障类型',
                judgmentConditions: '判定条件',
                decisionThreshold: '判定阈值',
                faultLevel: '故障等级',
                faultDesc: '故障描述',
                faultPush: '故障推送',
                enableStatus: '启用状态',
            }
        }
    },
    faultHandingConfirm: {
        header: {
            title: '故障处理确认',
            tip1: '注意:故障处理后状态变更为“已处理”，不可取消，请谨慎操作。',
            tip2: '请确认是否处理恢复以下故障?'
        },
        label: {
            sn: 'SN',
            faultName: '故障名称',
            faultType: '故障类型',
            faultLevel: '故障等级',
            faultDesc: '故障描述',
            time: '首次故障报警时间'
        }
    },
    defaultSet: {
        header: {
            title: '恢复默认设置',
            tip1: '注意：恢复默认设置后所有的故障状态将恢复到初始状态，请谨慎操作',
            tip2: '请确认是否要恢复以下故障至初始状态？'
        },
        label: {
            faultName: '故障名称',
            faultType: '故障类型',
            faultLevel: '故障等级',
            faultDesc: '故障描述'
        }
    },
    faultOperation: {
        header: {
            title: '故障操作确认',
            tip1: '请确定启用【{0}】故障？',
            tip2: '请确定禁用【{0}】故障？'
        }
    },
    faultPush: {
        header: {
            title: '故障推送操作确认',
            tip1: '请确定启用【{0}】的故障推送？',
            tip2: '请确定禁用【{0}】的故障推送？'
        }
    },
    roleList: {
        header: {
            roleList: '角色列表'
        },
        operate: {
            refresh: '刷新',
            setUp: '设置',
            export: '导出'
        },
        label: {
            role: '角色',
            relationUser: '关联用户',
            jurisdiction: '权限'
        }
    },
    userList: {
        header: {
            userList: '用户列表'
        },
        operate: {
            addUser: '添加用户',
            refresh: '刷新',
            setUp: '设置',
            export: '导出'
        },
        search: {
            placeholder: {
                userName: '请输入用户名'
            },
            condition: {
                userName: '用户名'
            }
        },
        label: {
            userName: '用户名',
            userId: '用户ID',
            role: '角色',
            itemNum: '项目数',
            cellPhone: '手机',
            email: '邮箱',
            desc: '描述',
            createTime: '创建时间',
            operate: {
                name: '操作',
                edit: '编辑',
                delete: '删除',
                tooltip: {
                    delete: '平台拥有者无法删除',
                    delete2: '请先移除此用户关联的项目'
                }
            }
        }
    },
    createKey: {
        header: {
            title: '创建API密钥'
        },
        form: {
            label: {
                keyName: '密钥名称',
                permissions: '权限'
            },
            placeholder: {
                keyName: '请输入密钥名称',
                permissions: '请选择权限'
            }
        },
        select: {
            title: '读/写'
        },
        empty: {
            keyName: '密钥名称输入不能为空',
            permissions: '权限不能为空'
        },
        errorTip: {
            keyName: '1-64字符，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”"',
            repeat: '密钥名称重复'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    ower: {
        tips1: '您尚未开通此项模块，请您将',
        tips2: '请求码',
        tips3: '和',
        tips4: '需求',
        tips5: '发送电子邮件至 ',
        tips6: '，我们将尽快与您联系。',
        tips7: '您尚未开通此项模块'
    },
    deleteKey: {
        header: {
            title: '删除API密钥'
        },
        content: {
            title: '请确认是否删除密钥：{0}？'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    addUser: {
        header: {
            addUser: '添加用户'
        },
        label: {
            role: '角色',
            userName: '用户名',
            password: '密码',
            desc: '描述',
            cellPhone: '手机',
            email: '邮箱'
        },
        role: {
            owner: '平台拥有者',
            user: '普通用户',
            operator: '运营人员'
        },
        default: {
            userName: '1~32 字符，只支持数字、大小写字母、英文下划线 “_”及中划线“-”'
        },
        placeholder: {
            userName: '请输入用户名',
            password: '请输入密码',
            desc: '请输入描述',
            cellPhone: '请输入手机号',
            email: '请输入邮箱地址'
        },
        emptyTip: {
            userName: '用户名输入不能为空',
            password: '密码输入不能为空',
            cellPhone: '手机号输入不能为空',
            email: '邮箱地址输入不能为空'
        },
        errorTip: {
            number: '只能输入数字',
            userName: '请输入正确的用户名',
            password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
            repeatUserName: '用户名重复',
            cellPhone: '请输入正确的手机号',
            email: '请输入正确的邮箱地址',
            HongKong: '请输入正确的中国香港手机号',
            MacaMacao: '请输入正确的中国澳门手机号',
            Taiwan: '请输入正确的中国台湾手机号'
        },
        tooltip: {
            count: '1：平台拥有者，拥有访问控制台&运营平台访问权限，全部操作权限。</br>' +
            '2：普通用户，拥有控制台访问权限，所有操作权限，无运营平台访问权限。</br>' +
            '3：运营人员，拥有运营平台访问权限，所有操作权限（除用户管理和角色管理）。'
        },
        phoneData: {
            China: '中国大陆 +86',
            HongKong: '中国香港 +852',
            MacaMacao: '中国澳门 +853',
            Taiwan: '中国台湾 +886',
            Argentina: '阿根廷 +54',
            Australia: '澳大利亚 +61',
            Brazil: '巴西 +55',
            Germany: '德国 +49',
            France: '法国 +33',
            SouthKorea: '韩国 +82',
            Canada: '加拿大 +1',
            USA: '美国 +1',
            Mexico: '墨西哥 +52',
            SouthAfrica: '南非 +27',
            Japan: '日本 +81',
            SaudiArabia: '沙特阿拉伯 +966',
            Turkey: '土耳其 +90',
            Italy: '意大利 +39',
            India: '印度 +91',
            Indonesia: '印度尼西亚 +62',
            UnitedKiongdom: '英国 +44'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    errorTip: {
        title: '请输入合法的手机号码'
    },
    editUser: {
        header: {
            editUser: '编辑用户'
        },
        label: {
            role: '角色',
            userName: '用户名',
            userId: '用户ID',
            desc: '描述',
            cellPhone: '手机',
            email: '邮箱'
        },
        placeholder: {
            desc: '请输入描述',
            cellPhone: '请输入手机号',
            email: '请输入邮箱地址'
        },
        emptyTip: {
            desc: '描述输入不能为空',
            cellPhone: '手机号输入不能为空',
            email: '邮箱地址输入不能为空'
        },
        errorTip: {
            number: '只能输入数字',
            cellPhone: '请输入正确的手机号',
            email: '请输入正确的邮箱地址',
            HongKong: '请输入正确的中国香港手机号',
            MacaMacao: '请输入正确的中国澳门手机号',
            Taiwan: '请输入正确的中国台湾手机号'
        },
        phoneData: {
            China: '中国大陆 +86',
            HongKong: '中国香港 +852',
            MacaMacao: '中国澳门 +853',
            Taiwan: '中国台湾 +886',
            Argentina: '阿根廷 +54',
            Australia: '澳大利亚 +61',
            Brazil: '巴西 +55',
            Germany: '德国 +49',
            France: '法国 +33',
            SouthKorea: '韩国 +82',
            Canada: '加拿大 +1',
            USA: '美国 +1',
            Mexico: '墨西哥 +52',
            SouthAfrica: '南非 +27',
            Japan: '日本 +81',
            SaudiArabia: '沙特阿拉伯 +966',
            Turkey: '土耳其 +90',
            Italy: '意大利 +39',
            India: '印度 +91',
            Indonesia: '印度尼西亚 +62',
            UnitedKiongdom: '英国 +44'
        },
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    myProfile: {
        header: {
            myProfile: '我的账户'
        }
    },
    securitySettings: {
        header: {
            securitySettings: '安全设置'
        }
    },
    apiKey: {
        header: {
            apiKey: '个人API密钥'
        }
    },
    userCenter: {
        header: {
            userCenter: '个人中心',
            title1: '基本信息',
            title2: '偏好设置'
        },
        tabs: {
            myProfile: '我的账户',
            securitySettings: '安全设置',
            apiKey: '个人API密钥',
        },
        label: {
            role: '角色',
            userName: '用户名',
            email: '邮箱',
            phoneNumber: '手机',
            language: '语言',
            time: '时区'
        },
        currectPassword: '当前密码',
        newPassword: '新密码',
        confirmPassword: '确认密码',
        emptyTip: {
            currectPassword: '当前密码输入不能为空',
            newPassword: '新密码输入不能为空',
            confirmPassword: '确认密码输入不能为空'
        },
        errorTip: {
            password: "8-30个字符，同时包含三项(大写字母，小写字母，数字或特殊符号()`~!{'@'}#{'{'}{'}'}{'$'}%&*_-+={'|'}[]:\";\'<>,.?/)",
            password2: '密码不一致，请重新输入'
        },
        placeholder: {
            phoneNumber: '请输入手机号',
            email: '请输入邮箱地址',
            language: '请选择语言',
            time: '请选择时区'
        },
        defaultTime: '(GMT+08:00)中国时间-北京',
        language: {
            zh: '中文',
            en: 'English'
        },
        tooltip: {
            tip: '仅支持创建 2 个密钥'
        },
        table: {
            label: {
                keyName: '密钥名称',
                token: 'Token',
                permissions: '权限',
                createTime: '创建时间',
                operate: {
                    name: '操作',
                    copy: '复制',
                    delete: '删除'
                }
            },
            select: {
                title: '读/写',
                read: '只读'
            }
        },
        btn: {
            createKey: '创建密钥',
            saveChange: '更改',
            save: '保存'
        }
    },
    deleteUser: {
        header: {
            deleteUser: '删除用户',
        },
        label: {
            userName: '用户名',
            role: '角色',
            cellPhone: '手机',
            email: '邮箱'
        },
        deleteTip: '请确认是否删除以下用户?',
        btn: {
            cancel: '取消',
            sure: '确定'
        }
    },
    table: {
        empty: '暂无数据',
        tip: '您可以'
    }
};

export default zh;

