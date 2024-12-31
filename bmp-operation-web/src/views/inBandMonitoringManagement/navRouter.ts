/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const inBandMonitoringNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.inBandMonitoringManagement'),
        defaultImg: imgUrl('defaultInBandMonitoringDef'),
        changeImg: imgUrl('defaultInBandMonitoringActive'),
        childrenCentF: true,
        path: [
            pathUrl('resourceMonitor'),
            pathUrl('historyAlarmInfo'),
            pathUrl('alarmRulesList'),
            pathUrl('alarmRulesDetail')
        ],
        otherPath: [''],
        firstIndex: 'InBandMonitoringManagement',
        children: [
            {
                title: language.t('navigationBar.children.resourceEcharts'),
                defaultImg: imgUrl('defaultInBandMonitoringDef'),
                changeImg: imgUrl('defaultInBandMonitoringActive'),
                path: pathUrl('resourceMonitor'),
                otherPath: [pathUrl('resourceMonitor')]
            },
            {
                title: language.t('navigationBar.children.historyAlarmInfo'),
                defaultImg: imgUrl('defaultInBandMonitoringDef'),
                changeImg: imgUrl('defaultInBandMonitoringActive'),
                path: pathUrl('historyAlarmInfo'),
                otherPath: [pathUrl('historyAlarmInfo')]
            },
            {
                title: language.t('navigationBar.children.allAlarmRules'),
                defaultImg: imgUrl('defaultInBandMonitoringDef'),
                changeImg: imgUrl('defaultInBandMonitoringActive'),
                path: pathUrl('alarmRulesList'),
                otherPath: [pathUrl('alarmRulesList'), pathUrl('alarmRulesDetail')]
            },
        ]
    }
};

export default inBandMonitoringNavData;
