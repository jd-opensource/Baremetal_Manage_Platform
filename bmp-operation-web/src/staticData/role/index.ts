import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class RoleStaticData {
    static roleList = [
        {
            name: global.t('roleList.label.role'),
            selected: true,
            disabled: true,
            label: 'roleName',
            showTooltip: false
        },
        {
            name: global.t('roleList.label.relationUser'),
            selected: true,
            disabled: false,
            label: 'userName',
            showTooltip: false
        },
        {
            name: global.t('roleList.label.jurisdiction'),
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];

    static roleInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            operate: status2,
            roleName: status2,
            userName: status1
        }
    };
};
export default RoleStaticData;
