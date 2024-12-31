import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyType5, CurrencyType7, CurrencyType8, CurrencyStatusType} from '@utils/publicType';
import ModelType from './type';

// 国际化
const {global} = i18n;

class ModelStaticData {
    static processorManufacturerData: CurrencyType5[] = [
        {
            value: 1,
            label: 'Intel'
        },
        {
            value: 2,
            label: 'AMD'
        },
        {
            value: 3,
            label: 'Ampere'
        },
        {
            value: 4,
            label: 'PHYTIUM'
        },
        {
            value: 5,
            label: 'Hygon'
        },
        {
            value: 6,
            label: 'Kunpeng920'
        },
        {
            value: 7,
            label: '龙芯3号'
        }
    ];

    static interfaceData: CurrencyType5[] = [
        {
            label: 'DDR',
            value: 1
        },
        {
            label: 'DDR2',
            value: 2
        },
        {
            label: 'DDR3',
            value: 3
        },
        {
            label: 'DDR4',
            value: 4
        },
        {
            label: 'DDR5',
            value: 5
        }
    ];

    static dominantFrequencyMHzData: CurrencyType7[] = [
        {
            label: 2400,
            value: 1
        },
        {
            label: 2666,
            value: 2
        },
        {
            label: 2933,
            value: 3
        },
        {
            label: 3200,
            value: 4
        }
    ];

    static capacityData: CurrencyType7[] = [
        {
            label: 16,
            value: 1
        },
        {
            label: 32,
            value: 2
        },
        {
            label: 64,
            value: 3
        },
        {
            label: 128,
            value: 4
        },
        {
            label: 256,
            value: 5
        },
        {
            label: 384,
            value: 6
        }
    ];

    static systemTypeData: CurrencyType5[] = [
        {
            label: 'SSD',
            value: 1
        },
        {
            label: 'HDD',
            value: 2
        }
    ];

    static systemInterfaceTypeData: CurrencyType5[] = [
        {
            label: 'SATA',
            value: 1
        },
        {
            label: 'SAS',
            value: 2
        },
        {
            label: 'NVME',
            value: 3
        }
    ];

    static sysSingleCapacityData: CurrencyType5[] = [
        {
            label: 'GB',
            value: 1
        },
        {
            label: 'TB',
            value: 2
        }
    ];

    static dataTypeData: CurrencyType5[] = [
        {
            label: 'SSD',
            value: 1
        },
        {
            label: 'HDD',
            value: 2
        }
    ];

    static dataInterfaceTypeData: CurrencyType5[] = [
        {
            label: 'SATA',
            value: 1
        },
        {
            label: 'SAS',
            value: 2
        },
        {
            label: 'NVME',
            value: 3
        }
    ];
    static dataSingleCapacityData: CurrencyType5[] = [
        {
            label: 'GB',
            value: 1
        },
        {
            label: 'TB',
            value: 2
        }
    ];
    static networkSpeedData: CurrencyType7[] = [
        {
            label: 1,
            value: 1
        },
        {
            label: 10,
            value: 2
        },
        {
            label: 25,
            value: 3
        },
        {
            label: 100,
            value: 4
        }
    ];
    static networkSetUpData: CurrencyType8[] = [
        {
            label: global.t('modelForm.netWork.single'),
            value: 'single'
        }
    ];
    static networkSetUpData1: CurrencyType8[] = [
        {
            label: global.t('modelForm.netWork.bond'),
            value: 'bond'
        },
        {
            label: global.t('modelForm.netWork.dual'),
            value: 'dual'
        }
    ];
    static gpuBrandData: CurrencyType5[] = [
        {
            label: 'NVIDIA',
            value: 1
        }
    ];
    static gpuModelData: CurrencyType5[] = [
        {
            label: 'T4',
            value: 1
        },
        {
            label: 'P40',
            value: 2
        },
        {
            label: 'V100',
            value: 3
        },
        {
            label: 'A100',
            value: 4
        }
    ];
    static heightUData: CurrencyType7[] = [
        {
            label: 2,
            value: 1
        },
        {
            label: 4,
            value: 2
        }
    ];
    static modelTypeData: ModelType[] = [
        {
            filterParams: 'computer',
            text: global.t('deviceList.radioBtn.computed'),
            value: 1
        },
        {
            filterParams: 'storage',
            text: global.t('deviceList.radioBtn.storage'),
            value: 2
        },
        {
            filterParams: 'gpu',
            text: global.t('deviceList.radioBtn.GPU'),
            value: 3
        },
        {
            filterParams: 'other',
            text: global.t('deviceList.radioBtn.other'),
            value: 4
        },
    ];
    static modelList = [
        {
            name: global.t('modelList.content.modelName'), // 机型名称
            selected: true,
            disabled: true,
            label: 'name',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.modelType'), // 机型类型
            selected: true,
            disabled: false,
            label: 'deviceSeries',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.machineRoomName'), // 机房名称
            selected: true,
            disabled: false,
            label: 'idcName',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.machineRoomCode'), // 机房编码
            selected: true,
            disabled: false,
            label: 'idcRegion',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.architecture'), // 体系架构
            selected: true,
            disabled: false,
            label: 'architecture',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.bootMode'), // 引导模式
            selected: true,
            disabled: false,
            label: 'bootMode',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.modelSpecification'), // 机型规格
            selected: true,
            disabled: false,
            label: 'deviceType',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.image'), // 镜像
            selected: true,
            disabled: false,
            label: 'image',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.device'), // 设备
            selected: true,
            disabled: false,
            label: 'deviceCount',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.CPU'), // CPU
            selected: true,
            disabled: false,
            label: 'cpu',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.storage'), // 内存
            selected: true,
            disabled: false,
            label: 'memory',
            showTooltip: false
        },
        // {
        //     name: global.t('modelList.content.sysDisc'), // 系统盘
        //     selected: true,
        //     disabled: false,
        //     label: 'systemVolume',
        //     showTooltip: false
        // },
        // {
        //     name: global.t('modelList.content.raidConfig'), // RAID配置
        //     selected: true,
        //     disabled: false,
        //     label: 'raidCan',
        //     showTooltip: false
        // },
        // {
        //     name: global.t('modelList.content.sysRAID'), // 系统盘RAID
        //     selected: true,
        //     disabled: false,
        //     label: 'systemVolumeRaid',
        //     showTooltip: false
        // },
        // {
        //     name: global.t('modelList.content.dataDisc'), // 数据盘
        //     selected: true,
        //     disabled: false,
        //     label: 'dataVolume',
        //     showTooltip: false
        // },
        {
            name: global.t('modelList.content.networkCard'), // 网卡
            selected: true,
            disabled: false,
            label: 'nicInfo',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.networkSettings'), // 网络设置
            selected: true,
            disabled: false,
            label: 'nic',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.GPU'), // GPU
            selected: true,
            disabled: false,
            label: 'gpu',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.desc'), // 描述
            selected: true,
            disabled: false,
            label: 'remark',
            showTooltip: false
        },
        {
            name: global.t('modelList.content.operate.name'), // 操作
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];
    static modelInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) => {
        return {
            name: status2,
            deviceSeries: status1,
            idcName: status1,
            idcRegion: status1,
            architecture: status1,
            bootMode: status1,
            deviceType: status1,
            image: status1,
            deviceCount: status1,
            cpu: status1,
            memory: status1,
            // systemVolume: status1,
            // raidCan: status1,
            // systemVolumeRaid: status1,
            // dataVolume: status1,
            nicInfo: status1,
            nic: status1,
            gpu: status1,
            remark: status1,
            operate: status2
        }
    };

    static modelTipData: string[] = [
        'nameTip',
        'deviceSeriesNameTip',
        'idcNameTip',
        'idcNameEnTip',
        'architectureTip',
        'deviceTypeTip',
        'cpuInfoTip',
        'memInfoTip',
        'svInfoTip',
        'raidCanTip',
        'raidTip',
        'dvInfoTip',
        'nicInfoTip',
        'interfaceModeNameTip',
        'gpuInfoTip',
        'descriptionTip'
    ];

    static modelDetailTipData: string[] = [
        'imageNameTip',
        'osTypeTip',
        'osVersionTip'
    ];

    static modelSpecData: string[] = ['预置规格', 'Preset Specifications'];

    static noRaidData: string[] = ['NORAID', 'NO RAID'];

    static noRaidData2: string[] = ['', 'NO RAID', 'NORAID'];

    static otherSpecData: string[] = ['Other specifications', '其它规格'];

    static singleRaid0: string[] = ['RAID0-stripping', '单盘RAID0'];
};

export default ModelStaticData;
