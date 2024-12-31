import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class HardwareStaticData {

    static hardwareTipData: string[] = [
        'snTip',
        'idcNameTip',
        'brandTip',
        'modelTip'
    ];

    static hardwareInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            sn: status2,
            idcName: status1,
            brand: status1,
            model: status1,
            manageStatus: status1,
            cpuStatus: status1,
            memStatus: status1,
            diskStatus: status1,
            nicStatus: status1,
            powerStatus: status1,
            otherStatus: status1,
            operate: status2
        }
    };

    static hardwareList = [
        {
            name: global.t('hardwareStatusList.table.label.sn'),
            selected: true,
            disabled: true,
            label: 'sn',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.idc'),
            selected: true,
            disabled: false,
            label: 'idcName',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.brand'),
            selected: true,
            disabled: false,
            label: 'brand',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.model'),
            selected: true,
            disabled: false,
            label: 'model',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.managementStatus'),
            selected: true,
            disabled: false,
            label: 'manageStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.cpu'),
            selected: true,
            disabled: false,
            label: 'cpuStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.storage'),
            selected: true,
            disabled: false,
            label: 'memStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.hardDisk'),
            selected: true,
            disabled: false,
            label: 'diskStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.networkCard'),
            selected: true,
            disabled: false,
            label: 'nicStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.powerSupply'),
            selected: true,
            disabled: false,
            label: 'powerStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.other'),
            selected: true,
            disabled: false,
            label: 'otherStatus',
            showTooltip: false
        },
        {
            name: global.t('hardwareStatusList.table.label.operate'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];
};

export default HardwareStaticData;
