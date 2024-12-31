import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType, ListShowTooltipType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class FaultRulesStaticeData {
 
    static faultRulesInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            alertName: status2,
            alertSpec: status1,
            alertType: status1,
            alertCondition: status1,
            alertThreshold: status1,
            alertLevel: status1,
            alertDesc: status1,
            pushStatus: status2,
            useStatus: status2,
        }
    };

    static faultRulesList = [
        {
            name: global.t('faultRulesList.table.label.faultName'),
            selected: true,
            disabled: true,
            label: 'alertName',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.faultAccessories'),
            selected: true,
            disabled: false,
            label: 'alertSpec',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.faultType'),
            selected: true,
            disabled: false,
            label: 'alertType',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.judgmentConditions'),
            selected: true,
            disabled: false,
            label: 'alertCondition',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.decisionThreshold'),
            selected: true,
            disabled: false,
            label: 'alertThreshold',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.faultLevel'),
            selected: true,
            disabled: false,
            label: 'alertLevel',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.faultDesc'),
            selected: true,
            disabled: false,
            label: 'alertDesc',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.faultPush'),
            selected: true,
            disabled: true,
            label: 'pushStatus',
            showTooltip: false
        },
        {
            name: global.t('faultRulesList.table.label.enableStatus'),
            selected: true,
            disabled: true,
            label: 'useStatus',
            showTooltip: false
        }
    ];

    static faultRulesTipData: string[] = [
        'alertNameTip',
        'alertSpecTip',
        'alertTypeTip',
        'alertConditionTip',
        'alertThresholdTip',
        'alertLevelTip',
        'alertDescTip'
    ];

    static faultRulesTooltip: ListShowTooltipType = {
        alertName: {},
        alertSpec: {},
        alertType: {},
        alertCondition: {},
        alertThreshold: {},
        alertLevel: {},
        alertDesc: {}
    };

    // constructor() {
    //     for (const index in FaultRulesStaticeData.faultRulesTooltip) {
    //         FaultRulesStaticeData.faultRulesTooltip[index] = {
    //             showTooltip: false
    //         };
    //     }
    // }
};

export default FaultRulesStaticeData;
