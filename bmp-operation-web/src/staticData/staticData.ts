import IdcStaticData from './idc/index';
import ModelStaticData from './model/index';
import ImageStaticData from './image/index';
import DeviceStaticData from './device/index';
import RoleStaticData from './role/index';
import UserStaticData from './user/index';
import {CurrencyStatusType} from '@utils/publicType';
import {AllCustomInfoType, AllCustomListType} from './staticDataType'; // 对应ts状态类
import FaultRulesStaticeData from './surveillance/faultRules/index';
import HardwareStaticData from './surveillance/hardware/index';
import FalutLogStaticData from './surveillance/faultLog/index';
import MessageStaticData from './message/index';
import AllAlarmRulesStatic from './inBandMonitoring/allAlarmRules/index';

// 通用状态
const currencyStatus: CurrencyStatusType = {
    required: false,
    selected: true
};

// 禁止更改状态
const prohibitChangesStatus: CurrencyStatusType = {
    required: true,
    selected: true
};

class AllStaticData {
    static allCustomListInfo: AllCustomInfoType = {
        idcInfo: IdcStaticData.idcInfo(currencyStatus, prohibitChangesStatus),
        modelInfo: ModelStaticData.modelInfo(currencyStatus, prohibitChangesStatus),
        imageInfo: ImageStaticData.imageInfo(currencyStatus, prohibitChangesStatus),
        deviceInfo: DeviceStaticData.deviceInfo(currencyStatus, prohibitChangesStatus),
        deviceFaultLogInfo: DeviceStaticData.deviceFaultLogInfo(currencyStatus, prohibitChangesStatus),
        messageInfo: MessageStaticData.messageInfo(currencyStatus, prohibitChangesStatus),
        hardwareInfo: HardwareStaticData.hardwareInfo(currencyStatus, prohibitChangesStatus),
        faultLogInfo: FalutLogStaticData.faultLogInfo(currencyStatus, prohibitChangesStatus),
        faultRulesInfo: FaultRulesStaticeData.faultRulesInfo(currencyStatus, prohibitChangesStatus),
        roleInfo: RoleStaticData.roleInfo(currencyStatus, prohibitChangesStatus),
        userInfo: UserStaticData.userInfo(currencyStatus, prohibitChangesStatus),
        auditLogsInfo: DeviceStaticData.auditLogsInfo(currencyStatus, prohibitChangesStatus),
        allAlarmRulesInfo: AllAlarmRulesStatic.allAlarmRulesInfo(currencyStatus, prohibitChangesStatus)
    };

    static allCustomList: AllCustomListType = {
        idcList: IdcStaticData.idcList,
        modelList: ModelStaticData.modelList,
        imageList: ImageStaticData.imageList,
        deviceList: DeviceStaticData.deviceList,
        deviceFaultLogList: DeviceStaticData.deviceFaultLogList,
        messageList: MessageStaticData.messageList,
        hardwareList: HardwareStaticData.hardwareList,
        faultLogList: FalutLogStaticData.faultLogList,
        faultRulesList: FaultRulesStaticeData.faultRulesList,
        roleList: RoleStaticData.roleList,
        userList: UserStaticData.userList,
        auditLogsList: DeviceStaticData.auditLogsList,
        allAlarmRulesList: AllAlarmRulesStatic.allAlarmRulesList
    };

    static noFoundData: string[] = ['找不到对象', 'Not found'];
};


export default AllStaticData;
