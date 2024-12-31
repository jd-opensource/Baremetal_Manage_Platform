const en: {[x: string]: {[x: string]: unknown}} = {
    size: {
        num: 'Optimize all remaining',
        all: 'All'
    },
    operate: {
        success: 'Operate Success',
        fail: 'Fail',
        doing: 'Doing',
        copy: 'Copy Success',
        refresh: 'Refresh',
        setUp: 'Settings',
        name: 'Actions',
        export: 'Export',
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        },
        tip: {
            default: 'Please select resource',
        }
    },
    searchTip: {
        title1: 'Find {0} resource matching the search and filter criteria, ',
        title2: 'Find {0} resource matching the search and filter criteria, please change the search keywords or ',
        title3: 'please go back to the list'
    },
    monitorEcharts: {
        tabs: {
            h1: '1h',
            h6: '6h',
            h12: '12h',
            d1: '1d',
            d3: '3d',
            d7: '7d',
            d14: '14d'
        },
        datePicker: {
            startDate: 'Start Date',
            endDate: 'End Date'
        },
        radio: {
            instance: 'Instance',
            disk: 'Disk',
            diskPartition: 'Disk Partition',
            netWorkInterface: 'Network Interface'
        },
        diskTip: 'Linux OS reports by DeviceName, and Windows OS reports by DriveLetter',
        diskPartitionTip: 'Linux OS reports by MountPoint, and Windows OS reports by DriveLetter',
        echartsCount: {
            period: 'Period',
            aggregationMethod: 'Aggregation Method',
            cpuUnit: 'CPU Usage（%）',
            cpu: 'CPU Usage',
            memUnit: 'Memory Usage（%）',
            mem: 'Memory Usage',
            memUsageUnit: 'Memory Used（GB）',
            memUsage: 'Memory Used',
            diskUsageUnit: 'Disk Used（GB）',
            diskUsage: 'Disk Used',
            diskUsageRateUnit: 'Disk Usage（%）',
            diskUsageRate: 'Disk Usage',
            diskReadWriteTrafficUnit: 'Disk Read/Write Traffic (Bps)',
            diskReadWriteTraffic: 'Disk Read/Write Traffic',
            diskReadTraffic: 'Disk Read Traffic ',
            diskWriteTraffic: 'Disk Write Traffic ',
            diskReadWriteIOPSUnit: 'Disk Read/Write IOPS（次/秒）',
            diskReadWriteIOPS: 'Disk Read/Write IOPS',
            diskReadIOPS: 'Disk Read IOPS',
            diskWriteIOPS: 'Disk Write IOPS',
            netWorkBpsUnit: 'Network Ingress/Egress Traffic (bps)',
            netWorkBps: 'Network Ingress/Egress Traffic ',
            netWorkIn: 'Network Ingress Traffic',
            netWorkOut: 'Network Egress Traffic',
            netWorkPackagesNumUnit: 'Network Ingress/Egress Packets（个/秒）',
            netWorkPackagesNumber: 'Network Ingress/Egress Packets',
            netWorkInPackageNum: 'Network Ingress Packets',
            netWorkOutPackageNum: 'Network Egress  Packets',
            averageLoad1MinUnit: 'Load Average 1min',
            averageLoad1Min: 'Load Average 1min',
            averageLoad5MinUnit: 'Load Average 5min',
            averageLoad5Min: 'Load Average 5min',
            averageLoad15MinUnit: 'Load Average 15min',
            averageLoad15Min: 'Load Average 15min',
            totalTCPConnectionsUnit: 'Total TCP Connections（个）',
            totalTCPConnections: 'Total TCP Connections',
            normalTCPConnectionsUnit: 'Established TCP Connections（个）',
            normalTCPConnections: 'Established TCP Connections',
            totalNumberOfProcessesUnit: 'Total Process Count（个）',
            totalNumberOfProcesses: 'Total Process Count'
        }
    },
    messageDelete: {
        header: {
            messageDelete: 'Delete Message'
        },
        tips: 'Once deleted, the message cannot be recovered. Are you sure you want to delete it?'
    },
    myMessage: {
        day90: 'Reminder 90 days before',
        day30: 'Reminder 30 days before',
        day15: 'Reminder 15 days before',
        expired: 'License expired reminder',
        header: {
            myMessage: 'Message List',
            title: '(Total {0}，',
            title2: 'of which are unread)',
        },
        radioGroup: {
            all: 'All',
            read: 'Read',
            noRead: 'Unread'
        },
        search: {
            title: 'Message Content',
            placeholder: 'Please enter the message content'
        },
        table: {
            expired: 'Expired',
            tip15: '',
            tip30: '',
            tip90: '',
            systemMessages: 'system',
            faultMessages: 'oob-monitor',
            operationMessages: 'operation',
            inbondAlert: 'inbond-alert',
            remind: 'Remind',
            failureNotification: 'Failure Notification',
            alarmTips: 'Alarm Reminder'
        },
        label: {
            messageContent: 'Message Content',
            finishTime: 'Receiving Time',
            messageType: 'Message Type',
            messageSubType: 'Message Sub Type'
        },
        footer: {
            tips: 'Please select message',
            delete: 'Delete',
            read: 'Mark as read'
        }
    },
    customFilter: {
        placeholder: 'Please Enter',
        btn: {
            confirm: 'Confirm',
            reset: 'Reset'
        }
    },
    messageDetail: {
        header: {
            title2: 'Message Detail',
            renewImmediately: 'Renew now'
        },
        label: {
            optContent: 'Operation Content',
            optPeople: 'Operator',
            optTime: 'Operation Time',
            versionName: 'Version Name',
            version: 'Version',
            startTime: 'Start Time',
            licenseEndTime: 'Expiration Time',
            licenseStatus: 'License Status',
            faultType: 'Fault Type',
            faultDesc: 'Fault Description',
            faultAlarmTime: 'Fault alarm time',
            alarmNum: 'Number of alarms',
        },
        optTips: {
            title: 'The following are the details of the operation content:',
            label: {
                instanceName: 'Instance Name',
                instanceId: 'Instance Id'
            },
            instance: 'Examples of',
            current: 'Current'
        },
        footer: {
            pre: 'Previous',
            next: 'Next'
        },
        failureNotification: 'Failure Notification',
        alarmTips: 'Alarm Reminder',
        remind: 'Remind',
        tip0: 'Hello {0} user! ',
        tip: 'Hello {0} user！Operation failed. Please try again or contact the administrator for assistance.',
        tip1: '',
        tip11: 'has expired， ',
        tip2: ' You can continue to use the cloud platform functions normally after renewal.',
        tip3: '',
        tip4: 'Here is an overview of the operation:',
        tip5: '',
        tip6: 'The following are the fault details:',
        tip7: '【Monitoring Alarm】',
        day90: '',
        day30: '',
        day15: '',
        expired: '',
        systemMessages: 'system',
        faultMessages: 'oob-monitor',
        operationMessages: 'operation',
        inbondAlert: 'inbond-alert',
        status: {
            effective: 'Effective',
            expired: 'Expired',
        }
    },
    messageSettings: {
        header: {
            messageSettings: 'Message Settings'
        },
        label: {
            emailPush: 'Associated email push',
            ipAddress: 'Mail Server',
            port: 'Mail Server Port',
            emailAddress: 'Email Address',
            emailPassword: 'Email Password',
            testEmail: 'Test email receiving address'
        },
        placeholder: {
            ipAddress: 'Please enter the SMTP server address for your email',
            port: 'Please enter the email server port',
            emailAddress: 'Please enter the email address that needs to be configured',
            emailPassword: 'Please enter your email password',
            testEmail: 'Please enter the test email receiving email address'
        },
        empty: {
            ipAddress: 'The mail server cannot be empty',
            port: 'The email server port cannot be empty',
            emailAddress: 'Email address cannot be empty',
            emailPassword: 'Email password cannot be empty',
            testEmail: 'Test email receiving email cannot be empty'
        },
        errorTip: {
            port: 'Only numbers can be entered, range 1-65535',
            emailAddress: 'Please enter the correct email address',
            emailPassword: 'Please enter the correct email password',
            testEmail: 'Please enter the correct email address'
        },
        tips: {
            pushTip: 'Enable email push, allowing users to receive email messages while receiving messages within the site; </br>' + 'Turn off email push, users can only receive internal messages and cannot receive email messages.',
            checkTip: 'Used to verify whether the email configuration information is normal.  <br/>  Click the [Save and Verify] button below to send a test email to check <br/> whether the specified mailbox can receive the email notification. <br/> There will be a certain delay in the email receiving process, please wait patiently.'
        },
        footer: {
            success: 'Verified',
            save: 'Save and verify',
            checkInfo: 'Verify email information'
        }
    },
    emailPush: {
        start: {
            title: 'Enable email push',
            content: 'Please confirm to enable email push. Once enabled, users will receive email reminders synchronously when receiving internal messages.'
        },
        close: {
            title: 'Turn off email push',
            content: 'Please confirm to turn off email push. After turning it off, users can accept internal messages and will not be able to receive email reminders.'
        }
    },
    userInfo: {
        emptyTip: {
            userName: 'Username input cannot be empty',
            password: 'Password input cannot be empty',
            cellPhone: 'Mobile phone number cannot be empty',
            email: 'Email input cannot be empty'
        },
        errorTip: {
            number: 'Only numbers can be entered',
            userName: '1-64 characters, only numbers, letters, underscore “_” and hyphen “-” are supported',
            password: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
            repeatUserName: 'Duplicated Username',
            cellPhone: 'Please enter the correct mobile phone number',
            email: 'Please enter the correct email address',
            HongKong: 'Please enter the correct mobile phone number in HongKong, China',
            MacaMacao: 'Please enter the correct mobile phone number in Macau, China',
            Taiwan: 'Please enter the correct mobile phone number in Taiwan, China'
        }
    },
    login: {
        header: {
            title: 'BMP-Operating Support Playtform STD',
            title2: 'BMP-Operating Support Playtform Pro'
        },
        formSubmit: {
            title: 'Welcome Back',
            placeholder: {
                userName: 'Username',
                passWord: 'Password'
            },
            toolTip: 'UserName does not exist password error',
            login: 'Login'
        }
    },
    loginOut: {
        title: 'Log out',
        tip: 'Confirm to login out',
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    navigationHeader: {
        title: 'BMP-Operating Support Playtform STD',
        title2: 'BMP-Operating Support Playtform Pro',
        myProfile: 'My Profile',
        securitySettings: 'Security Settings',
        apiKey: 'Personal API Keys',
        loginOut: 'Login Out'
    },
    navigationBar: {
        list: {
            computerRoomManagement: 'Data Center (DC)',
            modelManagement: 'Instance Type (IT)',
            imageManagement: 'Image',
            deviceManagement: 'Device',
            roleManagement: 'Role',
            userManagement: 'User',
            inBandMonitoringManagement: 'In-band monitoring management',
            surveillanceManagement: 'Out-of-band monitoring and management',
            messageCenter: 'Messages Center',
            userCenter: 'Personal Center'
        },
        children: {
            hardwareAlarmStatus: 'Hardware Device Status',
            resourceEcharts: 'Resource monitoring graph',
            historyAlarmInfo: 'Historical alarm information',
            allAlarmRules: 'All alarm rules',
            faultAlarmLog: 'Alarm Log',
            faultAlarmRules: 'Alarm Rules',
            myMessages: 'My Messages',
            messageSet: 'Message Settings'
        }
    },
    export: {
        title: 'Custom Export',
        content: {
            scope: 'Export Scope:',
            all: 'Export All',
            Selected: 'Export Selected instance',
            search: 'Export Search result'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    idcList: {
        header: {
            idcList: 'Idc List',
            operate: {
                refresh: 'Refresh',
                setUp: 'Settings',
                export: 'Export'
            }
        },
        content: {
            name: 'Name',
            englishCode: 'English Name',
            grade: 'Level',
            address: 'Address',
            createTime: 'Created Time',
            createPeople: 'Creator',
            updateTime: 'Update Time',
            updatePeople: 'Updater',
            operate: {
                name: 'Actions',
                edit: 'Edit'
            }
        }
    },
    idcDetail: {
        header: {
            title: 'Idc List',
            title2: 'Idc Detail',
            operate: {
                name: 'Actions',
                edit: 'Edit'
            }
        },
        basicInfo: {
            title: 'Account Basics',
            label: {
                name: 'Name',
                englishCode: 'English Name',
                grade: 'Level',
                address: 'Address',
                total: 'Cabinet Total',
                createTime: 'Created Time',
                createPeople: 'Creator',
                updateTime: 'Update Time',
                updatePeople: 'Updater'
            }
        },
        allAdministration: {
            title: 'General Settings',
            name: 'Name',
            loginAccount: 'OOB IPMI',
            loginUserName: 'Login Name',
            loginPassWord: 'Login Password',
            switchboardOneNum: 'Ethernet switch #1',
            switchboardTwoNum: 'Ethernet switch #2',
        }
    },
    securityVerification: {
        header: {
            verification: 'Security Check'
        },
        label: {
            userName: 'Username',
            passWord: 'Password'
        },
        placeholder: {
            userName: 'Username',
            passWord: 'Password'
        },
        empty: 'Password input cannot be empty',
        errTip: {
            userName: 'Username does not exist or password is wrong',
            passWord: 'Username does not exist or password is wrong',
            currencyPassword: 'Please enter 6-64 characters'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    editIdc: {
        header: {
            name: 'Edit Data center (DC)'
        },
        iptInfo: {
            other: 'Other'
        },
        label: {
            computerRoomName: 'Name',
            computerRoomCode: 'English Name',
            computerRoomGrade: 'Level',
            computerRoomAddress: 'Address',
            outsideLoginAccountNum: 'OOB IPMI Username',
            outsideLoginPsd: 'OOB IPMI Password',
            switchboardOneNum: 'Ethernet switch #1 login Username',
            switchboardOnePsd: 'Ethernet switch #1 login Password',
            switchboardTwoNum: 'Ethernet switch #2 login Username',
            switchboardTwoPsd: 'Ethernet switch #2 login Password'
        },
        placeholder: {
            computerRoomName: 'please enter the Data center (DC) name',
            computerRoomCode: 'please enter the  Data center (DC) English name',
            computerRoomGrade: 'please choose the Data center (DC) level',
            customComputerRoomGrade: 'please enter the customized Data center (DC) level',
            computerRoomAddress: 'please enter the Data center (DC) address',
            outsideLoginAccountNum: 'please enter the OOB IPMI Username',
            outsideLoginPsd: 'please enter the OOB IPMI Password',
            switchboardOneNum: 'please enter the Ethernet switch #1 login Username',
            switchboardOnePsd: 'please enter the Ethernet switch #1 login Password',
            switchboardTwoNum: 'please enter the Ethernet switch #2 login Username',
            switchboardTwoPsd: 'please enter the Ethernet switch #2 login Password'
        },
        errTip: {
            currency: 'Please enter 1-64 characters',
            currencyPassword: 'Please enter 1-128 characters',
            computerRoomNameEmpty: 'Data center (DC) name cannot be empty',
            computerRoomCodeEmpty: 'Data center (DC) English name cannot be empty',
            customComputerRoomGrade: 'Customized Data center (DC) level cannot be empty',
            computerRoomAddressEmpty: 'Data center (DC) address cannot be empty',
            outsideLoginAccountNumEmpty: 'OOB IPMI Username cannot be empty',
            outsideLoginPsdEmpty: 'OOB IPMI Password cannot be empty',
            switchboardOneNumEmpty: 'Ethernet Switch #1 login Username cannot be empty',
            switchboardOnePsdEmpty: 'Ethernet Switch #1 login Password cannot be empty',
            switchboardTwoNumEmpty: 'Ethernet Switch #2 login Username cannot be empty',
            switchboardTwoPsdEmpty: 'Ethernet switch #2 login Password cannot be empty'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    customList: {
        header: {
            name: 'Edit Columns'
        },
        text: {
            tip: 'Please select the  Colume items to display'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    modelList: {
        header: {
            modelList: 'Instance Type',
            label: {
                modelName: 'Type Name',
                modelSpecifications: 'Instance Spec (IS)'
            },
            placeholder: {
                modelName: 'Please enter the type name to search',
                modelSpecifications: 'Please enter the type Specifications to search'
            },
            operate: {
                addModel: 'Add New Type',
                refresh: 'Refresh',
                setUp: 'Settings',
                export: 'Export'
            }
        },
        content: {
            modelName: 'Type Name',
            modelType: 'Instance Type',
            machineRoomName: 'Name',
            machineRoomCode: 'English Name',
            architecture: 'Architecture',
            bootMode: 'Boot Mode',
            modelSpecification: 'Instance Spec (IS)',
            image: 'Image',
            device: 'Device',
            CPU: 'CPU',
            storage: 'Memory',
            sysDisc: 'System Volume',
            sysRAID: 'System Volume RAID Type',
            raidConfig: 'RAID Config',
            dataDisc: 'Data Volume',
            dataRAID: 'Data Volume RAID Type',
            networkCard: 'NIC',
            networkSettings: 'Network settings',
            GPU: 'GPU',
            desc: 'Description',
            operate: {
                name: 'Actions',
                edit: 'Edit',
                delete: 'Delete',
                more: 'More',
                addModel: 'Add the same model'
            }
        },
        tooltip: 'The instance type cannot be deleted, please delete the instances associated with the instance type first',
        edit: 'Models associated with devices do not support editing',
    },
    modelDetail: {
        header: {
            title: 'Instance Type ',
            title2: 'Model Details',
            specifications: 'Specifications',
            configParams: 'Config Params',
            operate: {
                name: 'Actions',
                edit: 'Edit',
                addModel: 'Add the same model',
                delete: 'Delete'
            },
            tip: 'The instance type cannot be deleted, please delete the instances associated with the instance type first',
            volumeManagement: {
                title: 'Volume Management',
                label: {
                    name: 'Volume Name',
                    type: 'Volume Type',
                    raidCan: 'RAID Config',
                    raid: 'RAID Mode',
                    disk: 'Disk Type',
                    interfaceType: 'Interface Type',
                    minNum: 'Minimum capacity of a single disk',
                    amount: 'Minimum quantity (blocks)',
                }
            }
        },
        tabs: {
            basicInfo: {
                name: 'Account Basics',
                label: {
                    modelName: 'Type Name',
                    computerRoomName: 'Name',
                    modelType: 'Instance Type',
                    architecture: 'Architecture',
                    bootMode: 'Boot Mode',
                    modelSpecifications: 'Instance Spec (IS)',
                    CPU: 'CPU',
                    storage: 'Memory',
                    sysDisc: 'System Volume',
                    raidConfig: 'RAID Config',
                    sysRAID: 'System Volume RAID Type',
                    dataDisc: 'Data Volume',
                    dataRAID: 'Data Volume RAID Type',
                    networkCard: 'NIC',
                    networkSettings: 'Network settings',
                    GPU: 'GPU',
                    height: 'Count（U）',
                    createTime: 'Created Time',
                    desc: 'Description'
                }
            },
            relationImage: {
                name: 'Associated images',
                header: {
                    addImage: 'Associate an image',
                    tooltip: 'Only 20 mirrors can be associated'
                },
                content: {
                    label: {
                        imageName: 'Image Name',
                        imageType: 'Image Type ',
                        architecture: 'Architecture',
                        operateSysPlatform: 'OS Platform',
                        operateSysVersion: 'OS Version',
                        partitionModule: 'Partition Module',
                        addTime: 'Created Time',
                        operate: {
                            name: 'Actions',
                            delete: 'Delete'
                        }
                    }
                }
            }
        },
        other: 'Other Classes'
    },
    uiPagination: {
        total: 'Total {0}',
        page: 'Total {0} page',
        piecesAndPage: '{0} / page'
    },
    addModel: {
        header: {
            title2: 'Add New Type',
            addModel: 'Add New Type'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    editModel: {
        header: {
            title2: 'Edit the Instance Type',
            editModel: 'Edit the Instance Type'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    deleteModel: {
        header: {
            deleteModel: 'Delete the Instance Type',
            batchDeleteModel: 'Batch Delete the Instance Type'
        },
        content: {
            deleteTip: 'After an instance type is deleted, its association with images is also deleted synchronously. If the same instance type is added again, you need to reassociate it with images. Please operate with caution.',
            deleteSure: 'Please confirm whether to delete the instance type：{0}？'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    addImage: {
        header: {
            addImage: 'Associate an image',
            label: {
                modelName: 'Type Name',
                choose: 'Number of selected mirrors'
            }
        },
        content: {
            label: {
                imageName: 'Image Name',
                modelName: 'Type Name',
                imageType: 'Image Type ',
                architecture: 'Architecture',
                operateSysPlatform: 'OS Platform',
                operateSysVersion: 'OS Version',
                sysPartitionModule: 'Sys Vol Partition Template'
            }
        },
        btn: {
            cancel: 'Cancel',
            save: 'Save',
            sure: 'Confirm'
        }
    },
    deleteImage: {
        header: {
            deleteImage: 'Delete the image'
        },
        tip: {
            text: 'Confirm to delete the associated image',
            delete: 'Confirm to delete the image'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    imageList: {
        header: {
            imageList: 'Image List',
            exportImage: 'Import Image',
            operate: {
                refresh: 'Refresh',
                setUp: 'Settings',
                export: 'Export'
            }
        },
        search: {
            placeholder: {
                imageName: 'Please enter the image name search',
                imageId: 'Please enter the image id search'
            },
            condition: {
                imageName: 'Image Name',
                imageId: 'Image Id'
            }
        },
        content: {
            label: {
                imageName: 'Image Name',
                imageId: 'Image ID',
                imageType: 'Image Type ',
                architecture: 'Architecture',
                operateSysPlatform: 'OS Platform',
                operateSysVersion: 'OS Version',
                format: 'Image Format',
                bootMode: 'Boot Mode',
                sysPartitionModule: 'Sys Vol Partition Template',
                imageDesc: 'Image Desp',
                numberOfAssociatedModels: 'Associated ITs',
                createTime: 'Created Time',
                operate: {
                    name: 'Actions',
                    delete: 'Delete'
                }
            }
        },
        filter: {
            presetImage: 'Preset image',
            customImage: 'Custom image'
        },
        tooltip: {
            title: 'The image cannot be deleted, please delete the image associated with the instance type first',
            image: 'Preset image does not support deletion'
        }
    },
    imageDetail: {
        header: {
            title2: 'Image Detail',
            operate: {
                title: 'Actions',
                delete: 'Delete'
            }
        },
        tooltip: {
            title: 'The image cannot be deleted, please delete the image associated with the instance type first',
            image: 'Preset image does not support deletion',
            title2: 'An instance of the current model is being created and cannot be removed'
        },
        add: {
            model: 'Add New Type'
        },
        tabs: {
            basicInfo: 'Account Basics',
            model: 'Associated Instance'
        },
        label: {
            imageName: 'Image Name',
            imageId: 'Image Id',
            imageType: 'Image Type',
            architecture: 'Architecture',
            operateSysPlatform: 'OS Platform',
            operateSysVersion: 'OS Version',
            format: 'Image Format',
            bootMode: 'Boot Mode',
            sysPartitionModule: 'Sys Vol Partition Template',
            imageDesc: 'Image Desp',
            numberOfAssociatedModels: 'Associated ITs',
            createTime: 'Created Time',
        },
        table: {
            label: {
                modelName: 'Type Name',
                modelType: 'Instance Type',
                idcName: 'Name',
                nameEn: 'English Name',
                architecture: 'Architecture',
                modelSpecifications: 'Instance Spec (IS)',
                desc: 'Desc',
                operate: {
                    title: 'Actions',
                    remove: 'Remove'
                }
            }
        },
        empty: {
            title: 'Add New Type'
        }
    },
    imageAddModel: {
        header: {
            title: 'Add New Type',
            imageName: 'Image Name',
            choose: 'Number of instance selected'
        },
        label: {
            modelName: 'Type Name',
            modelType: 'Instance Type',
            idcName: 'Name',
            nameEn: 'English Name',
            architecture: 'Architecture',
            modelSpecifications: 'Instance Spec (IS)',
            desc: 'Desc'
        },
        btn: {
            save: 'Save',
            cancel: 'Cancel'
        }
    },
    removeOperate: {
        header: {
            remove: 'Remove Instance'
        },
        content: {
            deleteTip: 'Notice : The image version cannot be used when the user creates or reinstalls the type after removal, Please operate with caution.',
            tip1: 'Are you sure to remove the [{0}] associated with [{1}{2}]?'
        },
        btn: {
            sure: 'Save',
            cancel: 'Cancel'
        }
    },
    importImage: {
        header: {
            title: 'Import Image' 
        },
        content: {
            label: {
                imageName: 'Image Name',
                architecture: 'Architecture',
                operateSysPlatform: 'OS Platform',
                operateSysVersion: 'OS Version',
                imageFormat: 'Image Format',
                bootMode: 'Boot Mode',
                systemPartitionTemplate: 'Sys Vol Partition Template',
                imageDesc: 'Image Desp',
                chooseImage: 'Choose Image'
            },
            placeholder: {
                imageName: 'please enter the image name',
                architecture: 'please select Architecture',
                operateSysPlatform: 'please select OS platform',
                customVersion: 'please enter the customized OS version',
                customOperatePlatform: 'please enter the customized OS platform',
                operateSysVersion: 'please select OS version',
                imageFormat: 'please select Image Format',
                bootMode: 'please select Boot Mode',
                imageDesc: 'please enter the image desp',
                chooseImage: 'Choose Image'
            }
        },
        emptyTip: {
            imageName: 'Image cannot be empty',
            architecture: 'please select Architecture',
            operateSysPlatform: 'OS platform cannot be empty',
            customVersion: 'please enter the customized OS version',
            customOperatePlatform: 'Customized OS platform cannot be empty',
            operateSysVersion: 'OS version cannot be empty',
            imageFormat: 'please select Image Format',
            bootMode: 'please select Boot Mode',
            noData: 'No Data'
        },
        errTip: {
            upload: 'Upload, please do not close the pop-up window',
            imageName: 'Please enter 1-64 characters',
            imageRepeat: 'Duplicate image name',
            numberLimit: 'Image cannot be empty',
            error: 'Upload Failed',
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm',
            upLoad: 'Upload'
        }
    },
    collectInfoConfirm: {
        header: {
            title: 'Re-collect device information and confirm'
        },
        tips: 'Please confirm whether to re-acquire device information?',
        checkbox: 'Clean RAID',
        tipsCount: {
            tip0: 'Note：',
            tip1: '1. The clean RAID operation is to delete the entire disk RAID. This operation is not recoverable, so please operate with caution!</br>' +
                  '2. Clearing the RAID operation will clear all data on the server disk. Please back up the data before performing this operation to avoid data loss and cause you losses.</br>'
        }
    },
    clearRaidConfirm: {
        header: {
            title: 'Clear RAID operation confirmation',
        },
        tipsCount: {
            tip0: 'Note：',
            tip1: '1. The clean RAID operation is to delete the entire disk RAID. This operation is not recoverable, so please operate with caution!</br>' +
                  '2. Clearing the RAID operation will clear all data on the server disk. Please back up the data before performing this operation to avoid data loss and cause you losses.</br>',
            tip2: 'The device you selected [{0}] will perform the RAID clear operation. Are you sure about the operation?'
        },
    },
    deviceList: {
        header: {
            deviceList: 'Device List'
        },
        radioBtn: {
            computed: 'Computer Class',
            storage: 'Storage Class',
            GPU: 'GPU',
            other: 'Other Classes'
        },
        batchOperate: {
            btn: {
                open: 'Start Up',
                close: 'Shut Down',
                restart: 'Reboot',
                recovery: 'Recycling Instance',
                edit: 'Edit instance name',
                resetPassword: 'Reset Password',
                more: 'More'
            },
            tip: {
                default: 'Please select resource',
                turnOn: 'Only support in shutdown status can be powered on',
                resetPassword: 'Only support in shutdown status  can reset passwords',
                restart: 'Only support in running status can be reboot',
                turnOff: 'Only support in running status can be shutdown',
                instanceName: 'Only created instances support editing',
                lock: 'Instance Locked',
                recovery: 'Servers that are in the "Shut down"&“Creation failed" statecan be removal'
            }
        },
        tooltip: {
            up: 'Servers that are in the "Stocked" state can be mounted',
            model: 'Please associate the model first',
            down: 'Servers that are in the "Mounted" state can be unmounted',
            delete: 'Servers that are in the "Stocked" state can be deleted',
            turnOn: 'Only support in shutdown status can be powered on',
            restart: 'Only support in running status can be reboot',
            remove: 'Servers that are in the "running"&“Shut down" statecan be removal',
            recovery: 'Servers that are in the "Shut down"&“Creation failed" statecan be removal',
            lock: 'Instance Locked',
            lock1: 'Unable to lock the instance',
            lock2: 'Unable to unlock the instance',
            reset: 'Only the device that failed the instance operation can perform the "reset instance status" operation',
            turnOff: 'Only support in running status can be shutdown',
            resetSystem: 'Only support in shutdown status  can reinstall the system',
            resetPassword: 'Only support in shutdown status  can reset passwords'
        },
        instanceStatus: {
            creating: 'Creating',
            starting: 'Starting up',
            running: 'Running',
            stopping: 'Shutting Down',
            stopped: 'Stopped',
            restarting: 'Rebooting',
            resetting_password: 'Resetting password',
            destroying: 'Destroying',
            destroyed: 'Destroyed',
            error: 'Error',
            upgrading: 'Upgrading',
            reinstalling: 'Reinstalling System'
        },
        operate: {
            instanceOperate: 'Instance Operate',
            error: {
                tip: 'Operation failed, please contact the administrator'
            },
            exportDevice: 'Import Device',
            placeholder: {
                sn: 'Please enter the SN to search',
                instanceOwner: 'Please enter the instance owner to search'
            },
            label: {
                sn: 'SN',
                instanceOwner: 'Instance Owner'
            },
            refresh: 'Refresh',
            setUp: 'Settings',
            export: 'Export'
        },
        content: {
            label: {
                sn: 'SN',
                name: 'Type Name',
                modelName: 'Instance Type',
                managementStatus: 'Management Status',
                collectionStatus: 'Collection Status',
                computerRoomName: 'Name',
                cabinetCode: 'Cabinet ID',
                uBit: 'Unit Location',
                brand: 'Brand',
                model: 'Model',
                architecture: 'Architecture',
                instanceImage: 'Image',
                instanceName: 'Instance Name',
                instanceID: 'Instance ID',
                instanceStatus: 'Instance Status',
                lockStatus: 'Locked Status',
                instanceOwner: 'Instance Owner',
                CPU: 'CPU',
                storage: 'Memory',
                sysDisc: 'System Volume',
                dataDisc: 'Data Volume',
                GPU: 'GPU',
                switchIPNetworkPortOne: 'Uplink SW IP（Port #1）',
                switchIPNetworkPortTwo: 'Uplink SW IP（Port #2）',
                networkPortOneUplinkPort: 'Uplink SW Port #1',
                networkPortTwoUplinkPort: 'Uplink SW Port #2',
                networkSettings: 'Network settings',
                outOfBandIP: 'OOB IP',
                intranceIPv4: 'Private IPv4(eth0)',
                intranceIPv4First: 'Private IPv4(eth1)',
                intranetIPv6: 'Private IPv6(eth0)',
                intranetIPv6First: 'Private IPv6(eth1)',
                createTime: 'Created Time',
                desc: 'Description',
                remark: 'Remark',
                operate: {
                    name: 'Actions',
                    up: 'Mount',
                    down: 'Unmount',
                    more: 'More',
                    lock:'Lock',
                    unlock:'Unlock',
                    delete: 'Device Delete',
                    reset: 'Reset instance state',
                    remove: 'Device Removal',
                    recovery: 'Recycling Instance',
                    open: 'Start Up',
                    close: 'Shut Down',
                    restart: 'Reboot',
                    resetSystem: 'Reinstall the system',
                    resetPassword: 'Reset Password'
                }
            }
        },
        managementStatus: {
            warehousing: 'Stocked',
            onTheShelf: 'Mounted',
            created: 'Created',
            onShelf: 'Mounting',
            unputaway: 'unputaway',
            remove: 'Removed',
            lowerShelf: 'Unmounting',
            error: 'Mount failed'
        },
        status: {
            collected: 'Collect',
            noCollected: 'No Collect',
            collecting: 'Collecting',
            collectionFailed: 'Collection Failed'
        }
    },
    resetSystemConfirm: {
        header: {
            title: 'Instructions for reinstalling the system',
            tip0: 'Note:',
            tip: '' +
                '1. Reinstalling the system is an operation to initialize the system disk and replace the operating system. This operation is not recoverable, so please operate with caution!</br>' +
                '2. This operation will clear all data on the instance system disk. Please back up the data before reinstalling the system to avoid data loss and cause you losses.</br>' +
                '3. After reinstalling the system, the system will not configure an IPv6 address for the instance by default, and you need to configure it manually.',
            tip2: 'The instance [{0}] you selected will perform a system reinstallation operation. Are you sure?'
        }
    },
    resetSystem: {
        header: {
            title: 'Reinstall system configuration',
            tip0: 'Note:',
            tip: '' +
                '1. Reinstalling the system is an operation to initialize the system disk and replace the operating system. This operation is not recoverable, so please operate with caution!</br>' +
                '2. This operation will clear all data on the instance system disk. Please back up the data before reinstalling the system to avoid data loss and cause you losses.</br>' +
                '3. After reinstalling the system, the system will not configure an IPv6 address for the instance by default, and you need to configure it manually.',
            instanceInfo: 'Instance Info',
            imageType: 'Image Type',
            system: 'System Volume',
            systemPartition: 'System Volume Partition',
            systemVolume: 'System Volume',
            systemVolumePartition: 'System Volume Partition',
            bootMode: 'Boot Mode',
            hostName: 'HostName',
            userName: 'UserName',
            setPassword: 'Set Password',
            password: 'Password',
            confirmPassword: 'Confirm Password',
            performanceMonitoring: 'Performance Monitoring',
            agent: 'Install monitoring agent'
        },
        table: {
            label: {
                instanceName: 'Instance Name',
                instanceID: 'Instance ID',
                operatingSystem: 'Operating System',
                createdTime: 'Created Time',
                sshName: 'Key Name'
            }
        },
        placeholder: {
            select: 'Please Select',
            hostName: 'Optional, computer name within the operating system'
        },
        regCheckTip: {
            hostName: 'It contains 2-64 characters in length and can be segmented into several sections with the separator (.). Uppercase and lowercase letters, numbers or hyphens (-) can be used in each section, but the dot (.) or the hyphen (-) cannot be continuously used in each section. It cannot start or end with a dot (.) or a hyphen (-).',
            password:  "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
            password2: 'The passwords are inconsistent, please re-enter'
        },
        emptyCheck: {
            sshKey: 'Key cannot be empty',
            password: 'Password cannot be empty',
            confirmPassword: 'Confirm password cannot be empty',
            repeat: 'The passwords are inconsistent, please re-enter'
        },
        radio: {
            customPassword: 'Customized Password',
            sshKeyLogin: 'Key pair login'
        },
        tip1: 'Provide real-time fault alarm'
    },
    resetPassword: {
        header: {
            title: 'Reset Password',
            tip: 'Note: After resetting the password, you need to power on the instance in the console/operation platform to take effect.',
            resetPassword1: 'The ',
            resetPassword2: ' instances you selected will be Edit reset password, Are you sure about the action?',
        },
        table: {
            label: {
                instanceName: 'Instance Name',
                instanceId: 'Instance Id',
                outOfBandIP: 'OOB IP',
                intranceIPv4: 'Private IPv4(eth0)',
            }
        },
        ipt: {
            label: {
                password: 'Password',
                newPassword: 'New Password',
                confirmPassword: 'Confirm Password'
            },
            errorTip: {
                password: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
                newPassword: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
                confirmPassword: 'The passwords are inconsistent, please re-enter'
            },
            emptyTip: {
                password: 'Password cannot be empty',
                newPassword: 'New password cannot be empty',
                confirmPassword: 'Confirm password cannot be empty'
            }
        }
    },
    batchOperate: {
        instanceName: {
            title: 'Edit instance name in batches',
            tip: 'Notice: The names of multiple instances edited in this way are named consecutively',
            tip1: 'The ',
            tip2: ' instances you selected will be Edit instance name, Are you sure about the action?',
            label: {
                instanceName: 'Instance Name',
                instanceId: 'Instance Id',
                outOfBandIP: 'OOB IP',
                intranceIPv4: 'Private IPv4(eth0)',
                newInstanceName: 'New instance name'
            },
            iptTip: {
                text: '2-128 characters, starting with uppercase and lowercase letters or Chinese, can contain numbers, ".", "_", ":" or "-"',
                empty: 'Instance name cannot be empty',
                repeat: 'Duplicate instance name'
            }
        },
        recoveryInstance: {
            title: 'Batch Recycling instance',
            tip: 'Note: Recycling will clear all data on the instance, please operate with caution.',
            recovery1: 'The ',
            recovery2: ' instances you selected will be recycled， Are you sure about the action?',
            label: {
                instanceName: 'Instance Name',
                instanceId: 'Instance Id',
                outOfBandIP: 'OOB IP',
                intranceIPv4: 'Private IPv4(eth0)'
            },
        },
        openCloseRestart: {
            open: 'Batch Startup',
            close: 'Batch Shutdown',
            restart: 'Batch Reboot',
            open1: 'The ',
            open2: ' instances you selected will be startup, Are you sure about the action?',
            close1: 'The ',
            close2: ' instances you selected will be shutdown, Are you sure about the action?',
            restart1: 'The ',
            restart2: ' instances you selected will be reboot, Are you sure about the action?',
            label: {
                instanceName: 'Instance Name',
                instanceId: 'Instance Id',
                outOfBandIP: 'OOB IP',
                intranceIPv4: 'Private IPv4(eth0)'
            }
        },
        resetPassword: {
            header: {
                title: 'Batch password reset',
                tip: 'Note:</br>' +
                    '1. Multiple instances whose passwords are reset in this way have the same password.</br>' +
                    '2. After resetting the password, you need to power on the instance in the console/operation platform to take effect.',
                resetPassword1: 'The ',
                resetPassword2: ' instances you selected will be Edit reset password, Are you sure about the action?',
            },
            table: {
                label: {
                    instanceName: 'Instance Name',
                    instanceId: 'Instance Id',
                    outOfBandIP: 'OOB IP',
                    intranceIPv4: 'Private IPv4(eth0)',
                }
            },
            ipt: {
                label: {
                    password: 'Password',
                    newPassword: 'New Password',
                    confirmPassword: 'Confirm Password'
                },
                errorTip: {
                    password: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
                    newPassword: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
                    confirmPassword: 'The passwords are inconsistent, please re-enter'
                },
                emptyTip: {
                    password: 'Password cannot be empty',
                    newPassword: 'New password cannot be empty',
                    confirmPassword: 'Confirm password cannot be empty'
                }
            }
        }
    },
    lockInstance: {
        header: {
            lock:'Lock',
            unlock:'Unlock',
        },
        lockTip: 'Confirm to lock the instance【{0}】?',
        unlockTip: 'Confirm to unlock the instance【{0}】?',
        tip: 'The "Delete" action is forbidden after the instance is locked.If you want to use it, you need to unlock it in advance.',
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    associatedModel: {
        header: {
            title: 'Associated Models'
        },
        content: {
            select: 'Select Model',
            placeholder: 'Please select the model',
            addNewModel: 'add a model',
            table: {
                label: {
                    volumeName: 'volume Name',
                    volumeType: 'volume Type',
                    raidCan: 'RAID Config',
                    raid: 'RAID Mode',
                    diskType: 'Disk Type',
                    interfaceType: 'InterfaceType',
                    minNum: 'Minimum capacity of a single disk',
                    minimumQuantity: 'Minimum quantity (blocks)',
                    associatedDisk: 'Associated Disk'
                },
                diskPlaceholder: 'Please select a disk'
            }
        }
    },
    deviceDetail: {
        header: {
            deviceList: 'Device List',
            title2: 'Device Detail'
        },
        operate: {
            name: 'Actions',
            up: 'Mount',
            down: 'Unmount',
            open: 'Start Up',
            close: 'Shut Down',
            restart: 'Reboot',
            reset: 'Reset instance state',
            remove: 'Device Removal',
            recovery: 'Recycling Instance',
            lock:'Lock',
            unlock:'Unlock',
            delete: 'Device Delete',
            resetSystem: 'Reinstall the system',
            resetPassword: 'Reset Password',
            error: {
                tip: 'Operation failed, please contact the administrator'
            }
        },
        tooltip: {
            up: 'Servers that are in the "Stocked" state can be mounted',
            model: 'Please associate the model first',
            associatedModels: 'Servers that are in the "stocked" state can be associated with machine models',
            down: 'Servers that are in the "Mounted" state can be unmounted',
            delete: 'Servers that are in the "Stocked" state can be deleted',
            remove: 'Servers that are in the "running"&“Shut down" statecan be removal',
            recovery: 'Servers that are in the "Shut down"&“Creation failed" statecan be removal',
            lock: 'Instance Locked',
            lock1: 'Unable to lock the instance',
            lock2: 'Unable to unlock the instance',
            reset: 'Only the device that failed the instance operation can perform the "reset instance status" operation',
            turnOn: 'Only support in shutdown status can be powered on',
            restart: 'Only support in running status can be reboot',
            turnOff: 'Only support in running status can be shutdown',
            resetSystem: 'Only support in shutdown status  can reinstall the system',
            resetPassword: 'Only support in shutdown status  can reset passwords'
        },
        deviceInfo: {
            deviceInfo: 'Device Details',
            sn: 'SN',
            modelName: 'Type Name',
            recapture: 'Recapture',
            managementStatus: 'Management status',
            collectionStatus: 'Collection Status',
            computerRoomName: 'Name',
            cabinetCode: 'Cabinet ID',
            uBit: 'Unit Location',
            remark: 'Remark'
        },
        instanceInfo: {
            instanceInfo: 'Instance Info',
            instanceName: 'Instance Name',
            instanceID: 'Instance ID',
            instanceImage: 'Image',
            instanceStatus: 'Instance Status',
            monitorAgentStatus: 'Monitor Agent status',
            lockStatus: 'Locked Status',
            instanceOwner: 'Instance Owner',
            createTime: 'Created Time',
            desc: 'Description'
        },
        hardwareConfiguration: {
            modelInfo: 'Model Info',
            hardwareConfiguration: 'Hardware Details',
            brand: 'Brand',
            model: 'Model',
            architecture: 'Architecture',
            CPU: 'CPU',
            storage: 'Memory',
            sysDisc: 'System Volume',
            adapeterId: 'Adapter_id (RAID card)',
            enclosure1: 'System Volume 1 enclosure #',
            slot1: 'System Volume 1 slot #',
            enclosure2: 'System Volume 2 enclosure #',
            slot2: 'System Volume 2 slot #',
            dataDisc: 'Data Volume',
            GPU: 'GPU',
            networkCard: 'NIC',
            networkSettings: 'Network settings'
        },
        networkInfo: {
            networkInfo: 'Network Details',
            intranceIPv4: 'Private IPv4（eth0）',
            intranceIPv4First: 'Private IPv4（eth1）',
            intranetIPv6: 'Private IPv6（eth0）',
            intranetIPv6First: 'Private IPv6（eth1）',
            mac1: 'MAC1（eth0）',
            mac2: 'MAC2（eth1）',
            switchIPOne: 'NIC port 1 switch IP（eth0）',
            switchIPTwo: 'NIC port 2 switch IP（eth1）',
            switchUplinkPortOne: 'NIC port 1 (eth0) switch port',
            switchUplinkPortTwo: 'NIC port 2 (eth1) switch port',
            subnetMask: 'subnet mask（eth0）',
            subnetMaskFirst: 'subnet mask（eth1）',
            networkGateway: 'gateway'
        },
        switchOutsideInfo: {
            switchOutsideInfo: 'Switches & Out-of-Band Information',
            userName1: 'NIC port 1 uplink switch login username',
            password1: 'NIC port 1 uplink switch login password',
            userName2: 'NIC port 2 uplink switch login username',
            password2: 'NIC port 2 uplink switch login password',
            outOfBandIP: 'OOB IP',
            outOfBandLoginUserName: 'OOB IPMI Username',
            outOfBandLoginPassWord: 'OOB IPMI Password'
        },
        editDesc: {
            header: {
                title: 'Description',
                title2: 'Notes'
            },
            count: {
                title: 'Modification Description:',
                title2: 'Modification Notes:'
            },
            placeholder: 'Please enter description',
            placeholder2: 'Please enter notes',
            btn: {
                cancel: 'Cancel',
                sure: 'Confirm'
            }
        },
        tabs: {
            baseInfo: 'Account Basics',
            diskDetail: 'Disk Details',
            hardwareMonitoring: 'Hardware Monitoring',
            performanceMonitoring: 'Performance Monitoring',
            operatLog: 'Operation Log'
        },
        diskDetail: {
            physicalDiskView: 'Physical disk view',
            table: {
                label: {
                    disk: 'Disk',
                    singleDiskCapacity: 'Single disk capacity',
                    diskType: 'Disk Type',
                    interfaceType: 'Interface Type',
                    enclosure: 'enclosure',
                    slot: 'slot'
                }
            },
            logicalDriveView: 'Logical disk view',
            table2: {
                label: {
                    driveLetterName: 'Driver Name',
                    singleDiskCapacity: 'Single disk capacity',
                    diskType: 'Disk Type',
                    interfaceType: 'Disk Interface',
                }
            },
            modelInfo: 'Model Info',
            modelRolInfo: 'Model roll information',
            modelName: 'Type Name',
            btn: 'Associated Instance',
            volumeManager: {
                table: {
                    label: {
                        volumeName: 'volume Name',
                        volumeType: 'volume Type',
                        raid: 'RAID Mode',
                        diskType: 'Disk Type',
                        interfaceType: 'Interface Type',
                        minNum: 'Minimum capacity of a single disk',
                        minimumQuantity: 'Minimum quantity (blocks)',
                        associatedDisk: 'Associated Disk'
                    }
                }
            }
        },
        hardwareMonitoring: {
            hardwareDeviceStatus: 'Hardware Device Status',
            alarmLogList: 'Alarm log list',
            status: {
                title: {
                    cpu: 'CPU',
                    storage: 'Memory',
                    hardDisk: 'Hard Disk',
                    networkCard: 'NIC',
                    powerSupply: 'Power Supply',
                    other: 'Other'
                },
                normal: 'Normal',
                error: 'Abnormal'
            }
        },
        table: {
            label: {
                serialNumber: 'Number',
                faultType: 'Fault Type',
                faultDesc: 'Fault Description',
                faultAlarmTime: 'Fault alarm time',
                updateTime: 'Update Time',
                alarmNum: 'Number of alarms',
                status: 'Status'
            }
        },
        status: {
            collected: 'Collect',
            noCollected: 'No Collect',
            collecting: 'Collecting',
            collectionFailed: 'Collection Failed'
        },
        operatLog: {
            header: {
                title: 'Operation Log',
                tips: '(The following list includes operation records from the past 90 days)'
            },
            search: {
                label: {
                    operatePeople: 'Operator',
                    operateTime: 'Operation Time',
                },
                placeholder: {
                    operatePeople: 'Please enter the operator',
                    startDate: 'Start Date',
                    endDate: 'End Date'
                },
                btn: {
                    search: 'Search',
                    clear: 'Clear'
                }
            },
            table: {
                label: {
                    logId: 'LogID',
                    operateName: 'Operation Name',
                    operatePeople: 'Operator',
                    operateTime: 'Operation Time',
                    operateFeedback: 'Operation Feedback',
                    failCode: 'Fault Info'
                }
            }
        }
    },
    editDesc: {
        header: {
            title: 'Description',
            title2: 'Notes'
        },
        count: {
            title: 'Modification Description:',
            title2: 'Modification Notes:'
        },
        placeholder: 'Please enter description',
        placeholder2: 'Please enter notes',
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    removeRecovery: {
        header: {
            remove: 'Device Removal',
            recovery: 'Recycling Instance'
        },
        tooltip: {
            remove: 'Please confirm  whether to remove the following device?',
            recovery: 'Please confirm whether to recycle the following instances?'
        },
        checkbox: {
            remove: 'I have understood the above precautions and agree to device removal',
            recovery: 'I have understood the above precautions and agree to recycle the instance'
        },
        operateTip: {
            remove: 'Notice:</br>' +
                '1. After the device is removed, the console will not display this instance, and this instance will not enter the inventory and cannot be purchased again</br>' +
                '2. Automatically disassociate the device from the instance owner',
            recovery: 'Notice:</br>' +
                '1. After the recovery instance operation is completed, the device is re-entered into the warehouse, and the association between the device and the instance owner is automatically released</br>' +
                '2. To avoid data loss, please confirm that the data has been backed up. The instance recovery operation will clear all data on the instance (including the system disk and data disk), and the instance will be completely deleted and cannot be restored'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    resetInstance: {
        header: {
            title: 'Reset instance state'
        },
        tooltip: {
            title: 'Notice:</br>' + 'reset instance state action reset instance state to last instance stable state',
            count: {
                tip1: 'Confirm',
                tip2: 'reset the instance status?'
            }
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    importDevice: {
        header: {
            title: 'Import Device'
        },
        label: {
            computerRoomName: 'Name',
            modelName: 'Type Name',
            deviceInfo: 'Device Details'
        },
        placeholder: {
            computerRoomName: 'please select data center name',
            modelName: 'please select instance type',
        },
        errTip: {
            upload: 'Upload, please do not close the pop-up window',
            computerRoomName: 'Data center (DC) name cannot be empty',
            modelName: 'Type name cannot be empty',
            deviceInfo: 'Device Details cannot be empty',
            error: 'Upload Failed',
        },
        noData: {
            tip: 'No Data'
        },
        operate: {
            tip1: 'If there is no suitable model, you can',
            addNewModel: ' add a model',
            templateDownLoad: 'Template Download'
        },
        btn: {
            selectionFile: 'Select a file',
            cancel: 'Cancel',
            sure: 'Confirm'
        },
        tooltip: {
            first: '1. Only 1 attachment can be selected, and the size is not more than 20M;',
            two: '2. To make a successful upload, please do not modify the colume cells in the template;',
            third: '3. Only one file can be uploaded, multiple file uploads will replace the original file;'
        }
    },
    upDownFrame: {
        operate: {
            addNewModel: 'add a model'
        },
        placeholder: {
            tip1: 'Please Select',
            tip2: 'The associated disk cannot be empty'
        },
        table: {
            label: {
                volumeName: 'volume Name',
                volumeType: 'volume Type',
                raidCan: 'RAID Config',
                raid: 'RAID Mode',
                diskType: 'Disk Type',
                interfaceType: 'Interface Type',
                minNum: 'Minimum capacity of a single disk',
                amount: 'Minimum quantity (blocks)',
                associatedDisk: 'Associated Disk'
            }
        },
        noData: {
            tip: 'No Data'
        },
        steps: {
            title1: 'Associated Instance',
            title2: 'Confirm Device'
        },
        select: {
            model: 'Select Model',
            placeholder: 'Please select the model'
        },
        titleTip: {
            up: 'Mount',
            down: 'Unmount',
            delete: 'Device Delete'
        },
        countTitleTip: {
            down: 'Note: After the server is successfully umounted, its status will be changed to "Stocked"',
            delete: 'Note: After the server is deleted, it cannot be undone. The server needs to be re-imported if it needs to be used again'
        },
        headerTitleTip: {
            up: 'Servers can only be launched as instances after they are mounted. Please confirm whether to mount the following servers?',
            down: 'Please confirm whether to unmount the following servers?',
            delete: 'Please confirm whether to delete the following servers?'
        },
        btn: {
            pre: 'Previous',
            next: 'Next step',
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    publicTable: {
        label: {
            sn: 'SN',
            instanceName: 'Instance Name',
            instanceId: 'Instance Id',
            instanceOwner: 'Instance Owner',
            cabinetCode: 'Cabinet ID',
            uBit: 'Unit Location',
            outOfBandIP: 'OOB IP',
            intranceIPv4: 'Private IPv4(eth0)'
        }
    },
    openShutDownRestart: {
        titleTip: {
            open: 'Start Up',
            close: 'Shut Down',
            restart: 'Reboot'
        },
        tooltipInfo: {
            open: 'Confirm to start up the instances：',
            close: 'Confirm to shutdown the instances：',
            restart: 'Confirm to reboot the instances：'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    modelForm: {
        header: {
            params: {
                text1: 'General parameters',
                text2: 'Configuration parameters'
            }
        },
        config: {
            configured: 'Configured RAID',
            notConfigured: 'Not Configured RAID'
        },
        radioBtn: {
            computed: 'Computer Class',
            storage: 'Storage Class',
            GPU: 'GPU',
            other: 'Other Classes'
        },
        raid: 'RAID0-stripping',
        none: 'none',
        unrestricted: 'not limited',
        netWork: {
            single: 'Single Network Interface',
            dual: 'Double Network Interfaces',
            bond: 'Single Network port bond'
        },
        specificationsType: {
            presetSpecifications: 'Preset Specifications',
            otherSpecifications: 'Other specifications'
        },
        tooltip: {
            tip: ' Add a new model using model {0} as a template',
            title: 'Instance Spec (IS) corresponds to the Type name one by one. </br>It is recommended to use [Specification Attribute]+[Instance Type]+[Architecture]. </br>Example: c1.medium.x86'
        },
        noData: {
            tip: 'No Data'
        },
        list: {
            manufacturer: 'Godson No.3'
        },
        label: {
            computerRoomName: 'Name',
            modelType: 'Instance Type',
            modelName: 'Type Name',
            architecture: 'Architecture',
            system: 'system',
            data: 'data',
            bootMode: 'Boot Mode',
            modelSpecifications: 'Instance Spec (IS)',
            desc: 'Description',
            cpu: 'CPU',
            cpuParams: 'CPU Parameters',
            processorManufacturer: 'CPU manufacturer',
            modelChoose: 'Model',
            numberPhysicalCores: 'physical cores',
            dominantFrequency: 'Clock Speed (GHz)',
            numberOfRoutes: 'Quantity',
            storage: 'Memory',
            card: 'RAID controller',
            storageParams: 'Memory Parameters',
            interface: 'interface',
            dominantFrequencyMHz: 'Clock Speed (MHz)',
            capacity: 'Capacity (GB)',
            modelNumber: 'Quantity',
            sysType: 'System volume disk type',
            sysInterfaceType: 'System volume disk interface',
            sysSingleCapacityVal: 'System volume single disk capacity',
            sysNumber: 'Number of system volume disks',
            raidConfig: 'RAID Config',
            sysRAID: 'System Volume RAID Type',
            dataType: 'Data volume disk type',
            dataInterfaceType: 'Data volume disk interface',
            dataSingleCapacityVal: 'Data volume single disk capacity',
            dataNumber: 'Number of data volume disks',
            networkSpeed: 'NIC speed (GE)',
            networkNumber: 'Number of NIC ports',
            networkSettings: 'Network settings',
            gpuBrand: 'GPU(Brand)',
            gpuModel: 'GPU(Model)',
            gpuNumber: 'Number of GPUs (number)',
            heightU: 'Count（U）',
            volumeManage: {
                title: 'Volume Management',
                label: {
                    name: 'Volume Name',
                    type: 'Volume Type',
                    raidCan: 'RAID Config',
                    raid: 'RAID Mode',
                    disk: 'Disk Type',
                    interfaceType: 'Interface Type',
                    minNum: 'Minimum capacity of a single disk',
                    amount: 'Minimum quantity (blocks)',
                    operate: 'Operate'
                },
                placeholder: {
                    tip1: 'Please Enter'
                },
                sys: {
                    tip1: 'Add at least one system volume',
                    tip2: 'Only one system volume can be added'
                },
                save: {
                    volumeManagement: 'Please save volume management',
                    empty: 'Volume management cannot be empty'
                },
                empty: {
                    name: 'Volume name cannot be empty',
                    type: 'Please select volume type',
                    raidCan: 'Please select a RAID configuration',
                    raid: 'Please select RAID mode',
                    disk: 'Please select the hard disk type',
                    interfaceType: 'Please select the interface type',
                    minNum: 'The minimum capacity of a single disk cannot be empty',
                    amount: 'Amount cannot be empty'
                },
                btn: {
                    cancel: 'Cancel',
                    add: 'Add Volume',
                    sure: 'Sure',
                    save: 'Save',
                    edit: 'Edit',
                    delete: 'Delete'
                }
            }
        },
        placeholder: {
            modelName: 'Please enter the instance type name (1-64 characters)',
            modelSpecifications: 'Please enter the instance specification',
            desc: 'please enter a description',
            select: 'please select an option',
            processorManufacturer: 'please select a CPU manufacturer',
            modelChoose: 'please select a CPU Model',
            dominantFrequency: 'please enter the Clock Speed (GHz)',
            interface: 'please select interface',
            dominantFrequencyMHz: 'please select Clock Speed (MHz)',
            capacity: 'please select Capacity (GB)',
            sysType: 'please select system volume disk type',
            sysInterfaceType: 'please select  system volume disk interface',
            sysSingleCapacityVal:'please enter the system volume single disk capacity',
            raidConfig: 'Please select a RAID configuration',
            sysRAID: 'please select system volume RAID mode',
            dataType: 'please select data volume disk type',
            dataInterfaceType: 'please select data volume disk interface',
            dataSingleCapacityVal: 'please enter the data volume single disk capacity',
            networkSpeed: 'please select  NIC speed (GE)',
            networkSettings: 'please select  bond mode',
            gpuBrand: 'please select  GPU(Brand)',
            gpuModel: 'please select  GPU(Model)',
            heightU: 'please select  Count（U）'
        },
        emptyTip: {
            modelName: 'Type name cannot be empty',
            modelSpecifications: 'Instance specification cannot be empty',
            cpu: 'CPU cannot be empty',
            processorManufacturer: 'Manufacturer cannot be empty',
            modelChoose: 'Model input cannot be empty',
            dominantFrequency: 'Clock Speed (GHz)cannot be empty',
            storage: 'memory cannot be empty',
            interface: 'interface cannot be empty',
            dominantFrequencyMHz: 'Clock Speed (MHz)cannot be empty',
            capacity: 'Capacity (GB) cannot be empty',
            sysType: 'System volume disk type cannot be empty',
            sysInterfaceType: 'System volume disk interface type cannot be empty',
            sysSingleCapacityVal: 'System volume Single Disk Capacity cannot be empty',
            raidConfig: 'RAID configuration cannot be empty',
            sysRAID: 'The system volume RAID mode cannot be empty',
            dataType: 'Data volume disk type cannot be empty',
            dataInterfaceType: 'Data  volume disk interface cannot be empty',
            dataSingleCapacityVal: 'Data volume Single Disk Capacity cannot be empty',
            networkSpeed: 'NIC speed (GE) cannot be empty',
            networkSettings: 'Bond mode cannot be empty',
            gpuBrand: 'GPU(brand) cannot be empty',
            gpuModel: 'GPU(model) cannot be empty',
            heightU: 'Count（U）cannot be empty'
        },
        errorTip: {
            number: '1-64 characters, the input content is a number, and the decimal point is allowed',
            number1: 'Only numbers can be entered',
            title: 'Please enter a valid mobile phone number',
            number2: 'Only numbers can be entered, with 2 decimal places reserved',
            number3: 'Exceeding maximum limit',
            currency: '1-64 characters, only chinese numbers, upper and lower case letters and English underscore "" And middle Dash "-"',
            currency2: 'Please enter 1-64 characters',
            modelSpecifications: 'The instance specification and Type name cannot be duplicated',
            name: 'Duplicate Type Name',
            specifications: 'Duplicated model specifications'
        }
    },
    resourceMonitor: {
        header: {
            resourceMonitor: 'Resource monitoring graph'
        },
        search: {
            userName: 'User Name',
            resourceType: 'Resource Type',
            idcName: 'Idc Name',
            instanceId: 'Instance ID',
            dimensions: 'Dimensions',
            monitoringMetrics: 'Monitoring Metrics',
            btn: {
                search: 'Search',
                clear: 'Clear'
            },
            count: {
                resourceType: 'Instance',
                dimensions: 'Instance'
            }
        },
        placeholder: {
            userName: 'Please enter Username',
            resourceType: 'Please select Resource Type',
            idcName: 'Please select Idc Name',
            instanceId: 'Please enter Instance ID',
            dimensions: 'Please select Dimensions',
            monitoringMetrics: 'Please select Monitoring Metrics'
        },
        errorTip: {
            userName: 'Please enter Username',
            resourceType: 'Please select Resource Type',
            idcName: 'Please select Idc Name',
            instanceId: 'Please enter Instance ID',
            dimensions: 'Please select Dimensions',
            monitoringMetrics: 'Please select Monitoring Metrics'
        }
    },
    historyAlarmInfo: {
        header: {
            historyAlarmInfo: 'Historical alarm information'
        },
        search: {
            label: {
                userName: 'User Name',
                alarmTime: 'Alarm Time'
            },
            placeholder: {
                userName: 'Please enter Username',
                startDate: 'Start Date',
                endDate: 'End Date',
            },
            errorTip: {
                userName: 'Please enter Username',
                alarmTime: 'Please select Alarm Time'
            },
            btn: {
                search: 'Search',
                clear: 'Clear'
            }
        },
        table: {
            label: {
                name: 'User Name',
                alarmTime: 'Alarm Time',
                resourceName: 'Resource Type',
                resourceId: 'Alarm Resources ID',
                triggerDescription: 'Triggering Conditions',
                alertValue: 'Alarm Value',
                alertLevelDescription: 'Alarm Level',
                alertPeriod: 'Duration',
                userName: 'Notification Object'
            },
            count: {
                minute: ' minute(s)'
            }
        }
    },
    allAlarmRulesList: {
        header: {
            allAlarmRulesList: 'All alarm rules'
        },
        select: {
            alarm: 'Alarm',
            normal: 'Normal',
            forbid: 'Disabled',
        },
        search: {
            label: {
                userName: 'User Name',
                ruleName: 'Rule Name',
                ruleId: 'Rule ID'
            },
            placeholder: {
                userName: 'Please enter Username',
                ruleName: 'Please enter a rule name',
                ruleId: 'Please enter a rule ID'
            },
            errorTip: {
                userName: 'Please enter Username',
                ruleName: 'Please enter a rule name',
                ruleId: 'Please enter a rule ID'
            },
            btn: {
                search: 'Search',
                clear: 'Clear'
            }
        },
        table: {
            label: {
                ruleName: 'Rule Name',
                userName: 'User Name',
                ruleId: 'Rule ID',
                resourceName: 'Resource Type',
                instanceCount: 'Instance associated number',
                triggerDescription: 'Triggering Condition',
                statusName: 'Status',
                operate: {
                    title: 'Operate',
                    lookDetail: 'Look Detail',
                    alarmHistory: 'Alarm History'
                }
            }
        }
    },
    allAlarmRulesDetail: {
        header: {
            title2: 'Alarm rule details',
            operate: {
                alarmHistory: 'Alarm History'
            }
        },
        count: {
            minute: ' minute(s)',
            hours: ' hour(s)'
        },
        baseInfo: {
            title: 'Account Basics',
            label: {
                ruleId: 'Rule Id',
                ruleName: 'RuleName',
                resourceType: 'Resource Type'
            }
        },
        alarmResources: {
            title: 'Alarm Resources',
            table: {
                label: {
                    instanceName: 'Instance Name',
                    instanceId: 'Instance ID',
                    ipv4: 'Private IPv4(eth0)',
                    ipv6: 'Private IPv6(eth0)'
                }
            }
        },
        triggeringConditions: {
            title: 'Triggering Conditions'
        },
        notificationStrategy: {
            title: 'Notification Strategy',
            label: {
                notificationCycle: 'Notification Cycle',
                effectivePeriod: 'Effective Period',
                notificationConditions: 'Notification Conditions',
                receivingChannel: 'Receiving Channel',
                notificationObject: 'Notification Object'
            }
        },
        notice: {
            email: 'Email',
            internalMessage: 'Internal message'
        }
    },
    hardwareStatusList: {
        header: {
            hardwareStatusList: 'Fault device status'
        },
        search: {
            label: {
                idcName: 'Idc Name',
                sn: 'SN'
            },
            placeholder: {
                idcName: 'Please Select',
                sn: 'Please enter the SN'
            },
            btn: {
                search: 'Search',
                clear: 'Clear'
            }
        },
        table: {
            label: {
                sn: 'SN',
                idc: 'Idc Name',
                brand: 'Brand',
                model: 'Model',
                managementStatus: 'Management Status',
                cpu: 'CPU',
                storage: 'Memory',
                hardDisk: 'Hard Disk',
                networkCard: 'NIC',
                powerSupply: 'Power Supply',
                other: 'Other',
                operate: 'Operate',
                opt: 'Historical Faults'
            }
        }
    },
    surveillance: {
        status: {
            all: 'All',
            no: 'Not Processed',
            yes: 'Processed'
        },
        accessory: {
            cpu: 'CPU',
            mem: 'Memory',
            hardDisk: 'Hard Disk',
            networkCard: 'NIC',
            powerSupply: 'Power Supply',
            voltage: 'Voltage',
            fan: 'Fan',
            temperature: 'Temperature',
            pcie: 'PCIE Bus'
        },
    },
    faultLogList: {
        header: {
            faultLogList: 'Alarm Log'
        },
        search: {
            label: {
                idcName: 'Idc Name',
                sn: 'SN',
                level: 'Failure Level',
                time: 'First fault alarm time',
                accessory: 'Faulty Accessories',
                status: 'Status'
            },
            placeholder: {
                idcName: 'Please Select',
                sn: 'Please enter the SN',
                level: 'Please Select',
                startDate: 'Start Date',
                endDate: 'End Date',
                accessory: 'Please Select',
                status: 'Please Select'
            },
            btn: {
                search: 'Search',
                clear: 'Clear'
            }
        },
        table: {
            label: {
                logId: 'LogID',
                sn: 'SN',
                idc: 'Idc Name',
                level: 'Failure Level',
                type: 'Fault Type',
                accessory: 'Faulty Accessories',
                iloip: 'iloIP',
                desc: 'Fault Desc',
                originalLog: 'Fault Original log',
                status: 'Status',
                numberOfAlarms: 'Number of alarms',
                faultDetectionTime: 'Out-of-band fault detection time',
                faultAlarmTime: 'First fault alarm time',
                updateTime: 'Update Time',
                operate: 'Operate'
            },
            operate: {
                title: 'Deal',
                tip: 'Only unprocessed fault logs can be processed'
            }
        }
    },
    faultRulesList: {
        header: {
            faultRulesList: 'Alarm Rules'
        },
        search: {
            label: {
                faultName: 'Fault Name',
                faultAccessories: 'Faulty Accessories',
                faultLevel: 'Failure Level'
            },
            placeholder: {
                faultName: 'Please enter the fault name',
                faultAccessories: 'Please Select',
                faultLevel: 'Please Select'
            },
            btn: {
                search: 'Search',
                clear: 'Clear',
                restoreDefaultSet: 'Restore default settings'
            }
        },
        table: {
            label: {
                faultName: 'Fault Name',
                faultAccessories: 'Faulty Accessories',
                faultType: 'Fault Type',
                judgmentConditions: 'Judgment Conditions',
                decisionThreshold: 'Decision Threshold',
                faultLevel: 'Failure Level',
                faultDesc: 'Fault Description',
                faultPush: 'Failure Push',
                enableStatus: 'Enabled Status',
            }
        }
    },
    faultHandingConfirm: {
        header: {
            title: 'Troubleshooting Confirmation',
            tip1: 'Note: After the fault is handled, the status changes to "Handed" and cannot be canceled. Please operate with caution.',
            tip2: 'Please confirm whether the following faults have been resolved and restored?'
        },
        label: {
            sn: 'SN',
            faultName: 'Fault Name',
            faultType: 'Fault Type',
            faultLevel: 'Failure Level',
            faultDesc: 'Fault Description',
            time: 'First fault alarm time'
        }
    },
    defaultSet: {
        header: {
            title: 'Restore default settings',
            tip1: 'Note: After restoring the default settings, all fault status will be restored to the initial state, please operate with caution',
            tip2: 'Please confirm whether you want to restore the following faults to the initial state?'
        },
        label: {
            faultName: 'Fault Name',
            faultType: 'Fault Type',
            faultLevel: 'Failure Level',
            faultDesc: 'Fault Description'
        }
    },
    faultOperation: {
        header: {
            title: 'Failure operation confirmation',
            tip1: 'Please make sure to enable the [{0}] fault?',
            tip2: 'Please make sure to disable the [{0}] fault?'
        }
    },
    faultPush: {
        header: {
            title: 'Failure push operation confirmation',
            tip1: 'Please confirm to enable fault push of [{0}]?',
            tip2: 'Are you sure to disable [{0}] fault push?'
        }
    },
    roleList: {
        header: {
            roleList: 'Role List'
        },
        operate: {
            refresh: 'Refresh',
            setUp: 'Settings',
            export: 'Export'
        },
        label: {
            role: 'Role',
            relationUser: 'Associated Users',
            jurisdiction: 'Permissions'
        }
    },
    errorTip: {
        title: 'Please enter a valid mobile phone number',
    },
    userList: {
        header: {
            userList: 'User List'
        },
        operate: {
            addUser: 'Add User',
            refresh: 'Refresh',
            setUp: 'Settings',
            export: 'Export'
        },
        search: {
            placeholder: {
                userName: 'Please enter Username'
            },
            condition: {
                userName: 'User Name'
            }
        },
        label: {
            userName: 'User Name',
            userId: 'User ID',
            role: 'Role',
            itemNum: 'Item Num',
            cellPhone: 'Mobile Phone',
            email: 'Email',
            desc: 'Description',
            createTime: 'Created Time',
            operate: {
                name: 'Actions',
                edit: 'Edit',
                delete: 'Delete',
                tooltip: {
                    delete: 'Owner Unable to delete',
                    delete2: 'Please remove the items associated with this user first'
                }
            }
        }
    },
    createKey: {
        header: {
            title: 'Create an API key'
        },
        form: {
            label: {
                keyName: 'Key Name',
                permissions: 'Permissions'
            },
            placeholder: {
                keyName: 'Please enter a key name',
                permissions: 'Please select permissions'
            }
        },
        select: {
            title: 'Read/Write'
        },
        empty: {
            keyName: 'Key name input cannot be empty',
            permissions: 'Permission cannot be empty'
        },
        errorTip: {
            keyName: '1-64 characters, only numbers, letters, underscore “_” and hyphen “-” are supported',
            repeat: 'Key Name Repeat'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    deleteKey: {
        header: {
            title: 'Delete the API key'
        },
        content: {
            title: 'Confirm to delete the key: {0}?'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    addUser: {
        header: {
            addUser: 'Add User'
        },
        label: {
            role: 'Role',
            userName: 'Name',
            password: 'Password',
            desc: 'Desc',
            cellPhone: 'Phone',
            email: 'Email'
        },
        role: {
            owner: 'Owner',
            user: 'User',
            operator: 'Operator'
        },
        default: {
            userName: '1-32 characters, only numbers, letters, underscore “_” and hyphen “-” are supported'
        },
        placeholder: {
            userName: 'Please enter Username',
            password: 'Please enter Password',
            desc: 'please enter a description',
            cellPhone: 'Please enter the Mobile Phone',
            email: 'Please enter Email '
        },
        emptyTip: {
            userName: 'Username input cannot be empty',
            password: 'Password input cannot be empty',
            cellPhone: 'Mobile phone number cannot be empty',
            email: 'Email input cannot be empty'
        },
        errorTip: {
            number: 'Only numbers can be entered',
            userName: 'Please enter a correct username',
            password: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
            repeatUserName: 'Duplicated Username',
            cellPhone: 'Please enter the correct mobile phone number',
            email: 'Please enter the correct email address',
            HongKong: 'Please enter the correct mobile phone number in HongKong, China',
            MacaMacao: 'Please enter the correct mobile phone number in Macau, China',
            Taiwan: 'Please enter the correct mobile phone number in Taiwan, China'
        },
        tooltip: {
            count: '1：Owner: Possess access to the console & operation platform, and all operation rights.</br>' +
            '2：User：Have access to the console, all operations, no access to the operating platform.</br>' +
            '3：Operator：Possess access rights to the operation platform, all operation rights (except </br>' + 'user management and role management).</br>'
        },
        phoneData: {
            China: 'China +86',
            HongKong: 'Hongkong +852',
            MacaMacao: 'Macao +853',
            Taiwan: 'Taiwan +886',
            Argentina: 'Argentina +54',
            Australia: 'Australia +61',
            Brazil: 'Brazil +55',
            Germany: 'Germany +49',
            France: 'France +33',
            SouthKorea: 'South Korea +82',
            Canada: 'Canada +1',
            USA: 'USA +1',
            Mexico: 'Mexico +52',
            SouthAfrica: 'South Africa +27',
            Japan: 'Japan +81',
            SaudiArabia: 'Saudi Arabia +966',
            Turkey: 'Turkey +90',
            Italy: 'Italy +39',
            India: 'India +91',
            Indonesia: 'Indonesia +62',
            UnitedKiongdom: 'United Kiongdom +44'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    editUser: {
        header: {
            editUser: 'Edit User'
        },
        label: {
            role: 'Role',
            userName: 'Name',
            userId: 'ID',
            desc: 'Desc',
            cellPhone: 'Phone',
            email: 'Email'
        },
        placeholder: {
            desc: 'please enter a description',
            cellPhone: 'Please enter the Mobile Phone',
            email: 'Please enter Email '
        },
        emptyTip: {
            desc: 'Description input cannot be empty',
            cellPhone: 'Mobile phone number cannot be empty',
            email: 'Email input cannot be empty'
        },
        errorTip: {
            number: 'Only numbers can be entered',
            cellPhone: 'Please enter the correct mobile phone number',
            email: 'Please enter the correct email address',
            HongKong: 'Please enter the correct mobile phone number in HongKong, China',
            MacaMacao: 'Please enter the correct mobile phone number in Macau, China',
            Taiwan: 'Please enter the correct mobile phone number in Taiwan, China'
        },
        phoneData: {
            China: 'China +86',
            HongKong: 'Hongkong +852',
            MacaMacao: 'Macao +853',
            Taiwan: 'Taiwan +886',
            Argentina: 'Argentina +54',
            Australia: 'Australia +61',
            Brazil: 'Brazil +55',
            Germany: 'Germany +49',
            France: 'France +33',
            SouthKorea: 'South Korea +82',
            Canada: 'Canada +1',
            USA: 'USA +1',
            Mexico: 'Mexico +52',
            SouthAfrica: 'South Africa +27',
            Japan: 'Japan +81',
            SaudiArabia: 'Saudi Arabia +966',
            Turkey: 'Turkey +90',
            Italy: 'Italy +39',
            India: 'India +91',
            Indonesia: 'Indonesia +62',
            UnitedKiongdom: 'United Kiongdom +44'
        },
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    myProfile: {
        header: {
            myProfile: 'My Profile'
        }
    },
    securitySettings: {
        header: {
            securitySettings: 'Security Settings'
        }
    },
    apiKey: {
        header: {
            apiKey: 'Personal API Keys'
        }
    },
    ower: {
        tips1: 'You have not yet activated this module, Please send your ',
        tips2: 'request code',
        tips3: ' and ',
        tips4: 'requirements',
        tips5: ' to ',
        tips6: ', and we will contact you as soon as possible.',
        tips7: 'You have not yet activated this module'
    },
    userCenter: {
        header: {
            userCenter: 'Personal Center',
            title1: 'Account Basics',
            title2: 'Console Preferences'
        },
        tabs: {
            myProfile: 'My Profile',
            securitySettings: 'Security Settings',
            apiKey: 'Personal API Keys',
        },
        label: {
            role: 'Role',
            userName: 'User Name',
            email: 'Email',
            phoneNumber: 'Mobile Phone',
            language: 'Language',
            time: 'Timezone'
        },
        currectPassword: 'Currect Password',
        newPassword: 'New Password',
        confirmPassword: 'Confirm Password',
        emptyTip: {
            currectPassword: 'Current password cannot be empty',
            newPassword: 'New password cannot be empty',
            confirmPassword: 'Confirm password cannot be empty'
        },
        errorTip: {
            password: "Password: 8-30 characters, and simultaneously including three types (Capital letters, lowercase letters, numbers or special characters ()`~!{'@'}#{'$'}%&*_-+={'|'}{'{'}{'}'}[]:<>,.?/)",
            password2: 'The passwords are inconsistent, please re-enter'
        },
        placeholder: {
            phoneNumber: 'Please enter phone number',
            email: 'Please enter email address',
            language: 'Please select language',
            time: 'Please select time zone'
        },
        defaultTime: '(GMT+08:00) China Time - Beijing',
        language: {
            zh: 'Chinese',
            en: 'English'
        },
        tooltip: {
            tip: 'Only supports creating 2 API keys'
        },
        table: {
            label: {
                keyName: 'Key Name',
                token: 'Token',
                permissions: 'Permissions',
                createTime: 'Created Time',
                operate: {
                    name: 'Actions',
                    copy: 'Copy',
                    delete: 'Delete'
                }
            },
            select: {
                title: 'Read/Write',
                read: 'Read Only'
            }
        },
        btn: {
            createKey: 'Create API Key',
            saveChange: 'Save Changes',
            save: 'Save'
        }
    },
    deleteUser: {
        header: {
            deleteUser: 'Delete user',
        },
        label: {
            userName: 'User Name',
            role: 'Role',
            cellPhone: 'Mobile Phone',
            email: 'Email'
        },
        deleteTip: 'Please confirm whether to delete the following users.',
        btn: {
            cancel: 'Cancel',
            sure: 'Confirm'
        }
    },
    table: {
        empty: 'No data is available',
        tip: 'You can '
    }
};

export default en;
