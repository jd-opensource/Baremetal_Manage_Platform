/**
 * @file
 * @author
*/
import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType, OsStoreType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class AllAlarmRulesStatic {
    static allAlarmRulesList = [
        {
            name: global.t('allAlarmRulesList.table.label.userName'), // 用户名称
            selected: true,
            disabled: true,
            label: 'userName',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.ruleName'), // 规则名称
            selected: true,
            disabled: false,
            label: 'ruleName',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.ruleId'), // 规则ID
            selected: true,
            disabled: false,
            label: 'ruleId',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.triggerDescription'), // 资源类型
            selected: true,
            disabled: false,
            label: 'resourceName',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.instanceCount'), // 实例关联数量
            selected: true,
            disabled: false,
            label: 'instanceCount',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.triggerDescription'), // 触发条件
            selected: true,
            disabled: false,
            label: 'triggerDescription',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.statusName'), // 状态
            selected: true,
            disabled: false,
            label: 'statusName',
            showTooltip: false
        },
        {
            name: global.t('allAlarmRulesList.table.label.operate.title'), // 操作
            selected: true,
            disabled: true,
            label: 'operate',
            showTooltip: false
        }
    ];
    static allAlarmRulesInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) => {
        return {
            userName: status2,
            ruleName: status1,
            ruleId: status1,
            resourceName: status1,
            instanceCount: status1,
            triggerDescription: status1,
            statusName: status1,
            operate: status2
        }
    };
    static allAlarmRulesListTipData: string[] = [
        'ruleNameTip',
        'ruleIdTip',
        'triggerDescriptionTip'
    ];

    static allAlarmRulesStatusData: OsStoreType[] = [
        {
            filterParams: '1',
            text: global.t('allAlarmRulesList.select.normal'),
            value: 1
        },
        {
            filterParams: '2',
            text: global.t('allAlarmRulesList.select.forbid'),
            value: 2
        },
        {
            filterParams: '3',
            // Alert
            text: global.t('allAlarmRulesList.select.alarm'),
            value: 3
        }
    ];
}

export default AllAlarmRulesStatic;
