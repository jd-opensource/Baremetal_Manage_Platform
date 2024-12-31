/**
 * @file
 * @author
*/

import {language} from "utils/publicClass.ts";

const monitoringNavData = (pathUrl: (arg0: string) => void, imgUrl: (arg0: string) => void) => {
    return {
        title: language.t('navigationBar.list.surveillanceManagement'),
        defaultImg: imgUrl('surveillanceDef'),
        changeImg: imgUrl('surveillance'),
        childrenCentF: true,
        path: [
            pathUrl('hardwareStatus'),
            pathUrl('faultLog'),
            pathUrl('faultRules')
        ],
        otherPath: [''],
        firstIndex: 'MonitorManagement',
        children: [
            {
                title: language.t('navigationBar.children.hardwareAlarmStatus'),
                defaultImg: imgUrl('surveillanceDef'),
                changeImg: imgUrl('surveillance'),
                path: pathUrl('hardwareStatus'),
                otherPath: [pathUrl('hardwareStatus')]
            },
            {
                title: language.t('navigationBar.children.faultAlarmLog'),
                defaultImg: imgUrl('surveillanceDef'),
                changeImg: imgUrl('surveillance'),
                path: pathUrl('faultLog'),
                otherPath: [pathUrl('faultLog')]
            },
            {
                title: language.t('navigationBar.children.faultAlarmRules'),
                defaultImg: imgUrl('surveillanceDef'),
                changeImg: imgUrl('surveillance'),
                path: pathUrl('faultRules'),
                otherPath: [pathUrl('faultRules')]
            }
        ]
    }
};

export default monitoringNavData;
