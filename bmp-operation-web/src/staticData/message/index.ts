/**
 * @file
 * @author
*/

import i18n from 'lib/i18n/index.ts'; // 国际化
import {CurrencyStatusType, OsStoreType} from '@utils/publicType';

// 国际化
const {global} = i18n;

class MessageStaticData {
    static messageType: OsStoreType[] = [ // 镜像类型
        {
            value: 1,
            text: global.t('myMessage.table.inbondAlert'),
            filterParams: global.t('myMessage.table.inbondAlert')
        },
        {
            value: 2,
            text: global.t('myMessage.table.systemMessages'),
            filterParams: global.t('myMessage.table.systemMessages')
        },
        {
            value: 3,
            text: global.t('myMessage.table.faultMessages'),
            filterParams: global.t('myMessage.table.faultMessages')
        },
        {
            value: 4,
            text: global.t('myMessage.table.operationMessages'),
            filterParams: global.t('myMessage.table.operationMessages')
        },
    ];

    static messageTipData: string[] = [
        'finishTimeTip',
        'messageTypeTip',
        'messageSubTypeTip'
    ];

    static messageList = [
        {
            name: global.t('myMessage.label.messageContent'),
            selected: true,
            disabled: true,
            label: 'detail',
            showTooltip: false
        },
        {
            name: global.t('myMessage.label.finishTime'),
            selected: true,
            disabled: false,
            label: 'finishTime',
            showTooltip: false
        },
        {
            name: global.t('myMessage.label.messageType'),
            selected: true,
            disabled: false,
            label: 'messageType',
            showTooltip: false
        },
        {
            name: global.t('myMessage.label.messageSubType'),
            selected: true,
            disabled: true,
            label: 'messageSubType',
            showTooltip: false
        }
    ];

    static messageInfo = (status1: CurrencyStatusType, status2: CurrencyStatusType) =>  {
        return {
            detail: status2,
            finishTime: status1,
            messageType: status1,
            messageSubType: status2
        }
    };

    static btnData: {id: string; text: string;}[] = [
        {
            id: 'prevId',
            text: global.t('messageDetail.footer.pre')
        },
        {
            id: 'nextId',
            text: global.t('messageDetail.footer.next')
        }
    ];

    static templateData: {prop: string; label: string; hasShow: string; hasTime: boolean; hasOther: boolean;}[] = [
        {
            prop: 'messageSubType',
            label: global.t('messageDetail.label.optContent'),
            hasShow: 'optMessage',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'userName',
            label: global.t('messageDetail.label.optPeople'),
            hasShow: 'optMessage',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'finishTime',
            label: global.t('messageDetail.label.optTime'),
            hasShow: 'optMessage',
            hasTime: true,
            hasOther: false
        },
        {
            prop: 'licenseName',
            label: global.t('messageDetail.label.versionName'),
            hasShow: 'unexpiredExpired',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'licenseNumber',
            label: global.t('messageDetail.label.version'),
            hasShow: 'unexpiredExpired',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'licenseStartTime',
            label: global.t('messageDetail.label.startTime'),
            hasShow: 'unexpiredExpired',
            hasTime: true,
            hasOther: false
        },
        {
            prop: 'licenseEndTime',
            label: global.t('messageDetail.label.licenseEndTime'),
            hasShow: 'unexpiredExpired',
            hasTime: true,
            hasOther: false
        },
        {
            prop: 'licenseEndTime',
            label: global.t('messageDetail.label.licenseStatus'),
            hasShow: 'unexpiredExpired',
            hasTime: false,
            hasOther: true
        },
        {
            prop: 'logid',
            label: 'LogID',
            hasShow: 'faultMessage',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'faultType',
            label: global.t('messageDetail.label.faultType'),
            hasShow: 'faultMessage',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'content',
            label: global.t('messageDetail.label.faultDesc'),
            hasShow: 'faultMessage',
            hasTime: false,
            hasOther: false
        },
        {
            prop: 'alertTime',
            label: global.t('messageDetail.label.faultAlarmTime'),
            hasShow: 'faultMessage',
            hasTime: true,
            hasOther: false
        },
        {
            prop: 'alertCount',
            label: global.t('messageDetail.label.alarmNum'),
            hasShow: 'faultMessage',
            hasTime: false,
            hasOther: false
        }
    ];

    static instanceTemplateData: {prop: string; label: string;}[] = [
        {
            prop: 'instanceName',
            label: global.t('messageDetail.optTips.label.instanceName'),
        },
        {
            prop: 'instanceId',
            label: global.t('messageDetail.optTips.label.instanceId'),
        }
    ];
}

export default MessageStaticData;
