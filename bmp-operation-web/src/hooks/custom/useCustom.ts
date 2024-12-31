/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import {CustomOperate} from 'utils/publicClass.ts';
import AllStaticData from 'staticData/staticData.ts'; // 自定义列表信息 自定义列表数据

const {idcInfo, modelInfo, imageInfo, deviceInfo, messageInfo, roleInfo, userInfo, hardwareInfo, faultLogInfo, faultRulesInfo, auditLogsInfo, allAlarmRulesInfo} = AllStaticData.allCustomListInfo;
const {idcList, modelList, imageList, deviceList, messageList, roleList, userList, hardwareList, faultLogList, faultRulesList, auditLogsList, allAlarmRulesList} = AllStaticData.allCustomList;

const useCustom = (type: string) => {
    const infoList = new Map([
        ['idcList', [idcInfo, idcList]],
        ['deviceTypeList', [modelInfo, modelList]],
        ['imageList', [imageInfo, imageList]],
        ['deviceList', [deviceInfo, deviceList]],
        ['roleList', [roleInfo, roleList]],
        ['userList', [userInfo, userList]],
        [methodsTotal.humpConversion('deviceStatusList'), [hardwareInfo, hardwareList]],
        [methodsTotal.humpConversion('alertLogList'), [faultLogInfo, faultLogList]],
        [methodsTotal.humpConversion('alertRuleList'), [faultRulesInfo, faultRulesList]],
        [methodsTotal.lineConverting('messageList'), [messageInfo, messageList]],
        ['auditLogsList', [auditLogsInfo, auditLogsList]],
        ['monitorRulesList', [allAlarmRulesInfo, allAlarmRulesList]]
    ]);

    // @ts-ignore
    const info = new CustomOperate(type, infoList.get(type)[0], infoList.get(type)[1]);

    return {
        info
    }
}

export default useCustom;