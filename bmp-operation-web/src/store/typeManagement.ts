import {CpuDataType, CurrencyType, MemDataType, CurrencyStatusType, VerifyRulesType, CurrencyType2, PublicType6, OsStoreType} from '@utils/publicType';

// checkboxArr ts
type CheckListArrType = PublicType6[] & Omit<CurrencyStatusType, 'required'>[] & {disabled: boolean}[];

/**
 * 自定义列表-数据信息-类
*/
type CustomDataInfoType = {
    pageKey: string;
    data: Record<string, CurrencyStatusType>;
};

interface IndexStateType {
    customDataInfo: CustomDataInfoType[];
};

interface FormSubmitType extends VerifyRulesType {};

interface StateType {
    data: {
        radioBtn: string[];
        modelCPUBtn: string[];
        storageData: string[];
    };
    raidsData: [];
    ruleForm: CurrencyType2 | {[x: string]: [] | {}};
    rules: {};
};

/* modelForm.ts */

interface SetCpuMemFormSubmitType {
    cpuData: {label: string; info: CpuDataType['info']}[];
    cpuInfo: string;
    memInfo: string;
    memData: {label: string; info: MemDataType['info']}[];
};

interface SetFromType {
    [x: string]: string | number;
    modelType: string;
    raidId: string;
    sysRaid: never;
    deviceSeries: string;
    systemVolumeSize: number;
    dataVolumeSize: number;
};

type SysRaidType = {join(arg0: string): never;};

type ValType = {value: boolean};

/**
 * state ts
*/
interface OssStateType {
    osType: OsStoreType[];
    messageType: OsStoreType[];
    architecture: OsStoreType[];
    customFilterList: {label: string; value: string; checkboxStatus: boolean}[];
    source: OsStoreType[];
    manageStatus: OsStoreType[];
    collectStatus: OsStoreType[];
    operation: OsStoreType[];
    deviceSeries: OsStoreType[];
    deviceTypeId: OsStoreType[];
    result: OsStoreType[];
    status: OsStoreType[];
};

/**
 * idcStore
*/
interface IdcStoreStateType {
    idcData: CurrencyType[] | { newIdcName: string | undefined;}[];
    loading: boolean;
    idcDataNoOpt: CurrencyType[];
};

/**
 * faultLevelStore
*/
interface FaultLevelStateType {
    loading: boolean;
    faultLevelData: string[];
};

/**
 * faultSpecStore
*/
interface FaultSpecStateType {
    loading: boolean;
    faultSpecData: string[];
};


export {
    ValType,
    StateType,
    SetFromType,
    SysRaidType,
    OsStoreType,
    OssStateType,
    IndexStateType,
    FormSubmitType,
    CheckListArrType,
    IdcStoreStateType,
    CurrencyStatusType,
    FaultSpecStateType,
    FaultLevelStateType,
    SetCpuMemFormSubmitType
};
