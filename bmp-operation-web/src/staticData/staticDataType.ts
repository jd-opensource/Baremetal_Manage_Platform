import {CurrencyStatusType} from '../utils/publicType';

/**
 * 通用类-列表
*/
type CurrencyType = {
    text: string;
    label: string | number;
    info: any;
    filterParams: string;
};

/**
 * 导航类
*/
type NavType = {
    title: string;
    defaultImg: string;
    changeImg: string;
    path: string | string[];
    detailPath: string;
};

/**
 * 自定义数据信息类
*/
type InfoType = Record<string, Required<CurrencyStatusType>>;

/**
 * 角色数据类
*/
type RoleNameDataType = {
    filterParams: string;
    showInfo: string;
    text: string;
    value: number;
};

/**
 * 类集合
*/
interface StaticDataType {
    NavigationBarListaType: Required<NavType> & Partial<{children?: NavType[];}>;
    CurrencyListType: Required<{value: number | string;}> & Partial<CurrencyType>;
    CustomCurrencyType: Omit<CurrencyStatusType, 'required'> & Required<{name: string; disabled: boolean; showTooltip: boolean;}> & Partial<{label: string;}>;
    CustomInfoType: InfoType;
    ModelListCustomInfoType: InfoType;
    ImageListCustomInfoType: InfoType;
    DeviceListCustomType: InfoType;
    MessageListCustomType: InfoType;
    HardwareType: InfoType;
    FaultLogType: InfoType;
    FaultRulesType: InfoType;
    RoleListCustomInfoType: InfoType;
    UserListCustomInfoType: InfoType;
    AuditLogsInfoType: InfoType;
    AllAlarmRulesInfoInfoType: InfoType;
    FilterDataType: Required<{text: string; value: number;}>;
    roleId: RoleNameDataType[];
};

type AllCustomInfoType = {
    idcInfo: StaticDataType['CustomInfoType'];
    modelInfo: StaticDataType['ModelListCustomInfoType'];
    imageInfo: StaticDataType['ImageListCustomInfoType'];
    deviceInfo: StaticDataType['DeviceListCustomType'];
    messageInfo: StaticDataType['MessageListCustomType'];
    deviceFaultLogInfo: StaticDataType['DeviceListCustomType'];
    faultLogInfo: StaticDataType['HardwareType'];
    hardwareInfo: StaticDataType['FaultLogType'];
    faultRulesInfo: StaticDataType['FaultRulesType'];
    roleInfo: StaticDataType['RoleListCustomInfoType'];
    userInfo: StaticDataType['UserListCustomInfoType'];
    auditLogsInfo: StaticDataType['AuditLogsInfoType'];
    allAlarmRulesInfo: StaticDataType['AllAlarmRulesInfoInfoType'];
};

type key = 'idcList' | 'modelList' | 'imageList' | 'deviceList' | 'messageList' | 'deviceFaultLogList' | 'hardwareList' | 'faultLogList' | 'faultRulesList' | 'roleList' | 'userList' | 'auditLogsList' | 'allAlarmRulesList';


type AllCustomListType = {
    [k in key]: StaticDataType['CustomCurrencyType'][];
};


export {
    AllCustomInfoType,
    AllCustomListType,
};
