import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType, OsStoreType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class ImageStaticData {
    static source: OsStoreType[] = [ // 镜像类型
        {
            value: 1,
            text: global.t('imageList.filter.presetImage'),
            filterParams: 'common'
        },
        {
            value: 2,
            text: global.t('imageList.filter.customImage'),
            filterParams: 'user_defined'
        }
    ];
    static imageList = [
        {
            name: global.t('imageList.content.label.imageName'), // 镜像名称
            selected: true,
            disabled: true,
            label: 'imageName',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.imageId'), // 镜像ID
            selected: true,
            disabled: false,
            label: 'imageId',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.imageType'), // 镜像类型
            selected: true,
            disabled: false,
            label: 'source',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.architecture'), // 体系架构
            selected: true,
            disabled: false,
            label: 'architecture',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.operateSysPlatform'), // 操作系统平台
            selected: true,
            disabled: false,
            label: 'osType',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.operateSysVersion'), // 操作系统版本
            selected: true,
            disabled: false,
            label: 'osVersion',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.format'), // 镜像格式
            selected: true,
            disabled: false,
            label: 'format',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.bootMode'), // 引导模式
            selected: true,
            disabled: false,
            label: 'bootMode',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.sysPartitionModule'), // 系统盘分区模块
            selected: true,
            disabled: false,
            label: 'systemPartition',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.imageDesc'), // 镜像描述
            selected: true,
            disabled: false,
            label: 'description',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.numberOfAssociatedModels'), // 已关联机型数
            selected: true,
            disabled: false,
            label: 'deviceTypeNum',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.createTime'), // 创建时间
            selected: true,
            disabled: false,
            label: 'createdTime',
            showTooltip: false
        },
        {
            name: global.t('imageList.content.label.operate.name'), // 操作
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];

    static imageDetailTipData: string[] = [
        'nameTip',
        'deviceSeriesNameTip',
        'idcNameTip',
        'idcNameEnTip',
        'architectureTip',
        'deviceTypeTip',
        'descriptionTip'
    ];

    static imageListTipData: string[] = [
        'imageNameTip',
        'imageIdTip',
        'osTypeTip',
        'osVersionTip',
        'descriptionTip'
    ];

    static imageModelAddTipData: string[] = [
        'nameTip',
        'idcNameTip',
        'idcNameEnTip',
        'deviceTypeTip',
        'descriptionTip'
    ];

    static imageInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            imageName: status2,
            imageId: status1,
            source: status1,
            architecture: status1,
            osType: status1,
            osVersion: status1,
            format: status1,
            bootMode: status1,
            systemPartition: status1,
            description: status1,
            deviceTypeNum: status1,
            createdTime: status1,
            operate: status2
        }
    };

    static otherData: string[] = ['其它', 'other', 'Other'];

    static otherSystemData :string[] = ['Windows', '其它', 'other', 'Other'];
};

export default ImageStaticData;
