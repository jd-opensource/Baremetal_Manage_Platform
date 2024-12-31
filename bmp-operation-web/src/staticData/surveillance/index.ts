import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyType9} from '@utils/publicType';

// 国际化
const {global} = i18n;

class SurveillanceStaticeData {
    static faultLevel: string[] = [
        'Critical',
        'Warning'
    ];

    static faultyAccessories: string[] = [
        global.t('surveillance.accessory.cpu'),
        global.t('surveillance.accessory.mem'),
        global.t('surveillance.accessory.hardDisk'),
        global.t('surveillance.accessory.networkCard'),
        global.t('surveillance.accessory.powerSupply'),
        global.t('surveillance.accessory.voltage'),
        global.t('surveillance.accessory.fan'),
        global.t('surveillance.accessory.temperature'),
        global.t('surveillance.accessory.pcie')
    ];

    static status: CurrencyType9[] = [
        {
            text: global.t('surveillance.status.all'),
            value: ''
        },
        {
            text: global.t('surveillance.status.no'),
            value: '0'
        },
        {
            text: global.t('surveillance.status.yes'),
            value: '1'
        }
    ];
};

export default SurveillanceStaticeData;
