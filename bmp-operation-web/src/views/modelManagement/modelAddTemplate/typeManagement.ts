import {CurrencyType, CpuDataType, MemDataType, VerifyRulesType} from '@utils/publicType';

type key = 'name' | 'value';

type key1 = 'hasClick' | 'hasClick1' | 'hasClick2' | 'hasClick3' | 'hasClick4' | 'sysRaidFlag' | 'raidConfig';

/**
 * 机型类型
*/
type RadioBtnType = {
    [k in key]: string;
};

// type RadioBtnType<T> = {  [key in keyof T]: string}

type IdcNameRadioBtnType = {
    find(arg0: (item: {title: string;}) => boolean): {idcId: string; title: string;};
};

/**
 * 机房名称类
*/
type ComputerRoomRadioBtnType = {
    title: string;
    idcId: string;
    showTooltip: boolean;
};

/**
 * 通用操作类
*/
interface ProcessingOperationsType {
    setModelType: any;
    // processingOperations(formSubmit: CurrencyType): CurrencyType;
    // cpuPreset(formSubmit: CurrencyType): CurrencyType;
    // memPreset(formSubmit: CurrencyType): CurrencyType;
    // setCpuMemData(arg0: boolean, arg1: boolean, arg2: CurrencyType): Promise<void>;
    // setFormData(formSubmit: CurrencyType): CurrencyType;
    // setErrorHeight(arg0: Ref<boolean>[], arg1: boolean, arg2: boolean, arg3: Ref<boolean>, arg4: Ref<boolean>, arg5: any, arg6: any): never;
};

/**
 * 复杂数据类
*/
interface ReactiveDataType {
    radioBtn: RadioBtnType[];
    computerRoomRadioBtn: ComputerRoomRadioBtnType[];
    modelCPUBtn: string[];
    storageData: string[];
};

/**
 * 规则验证类
*/
interface RuleFormType {
    templateName: string;
    deviceTypeName: string;
    bootMode: string[] | string;
    machineRoomName: string;
    cpuSpec: string;
    memSpec: string;
    idcId: string;
    modelType: string;
    arrayCard: string;
    // raidCan: any;
    name: string;
    architecture: string | number;
    deviceType: string;
    description: string;
    modelCPU: string;
    cpuInfo: string;
    cpuManufacturer: string;
    cpuModel: string;
    cpuCores: number | string;
    cpuFrequency: string;
    cpuAmount: string | number;
    modelStorage: string;
    memInfo: string;
    memType: string;
    memFrequency: string | number;
    memSize: string | number;
    memAmount: string | number;
    volumeManagerTableData: any;
    noConfigurationData: any;
    // systemVolumeType: string;
    // systemVolumeInterfaceType: string;
    // systemVolumeSize: string | number | null;
    // systemVolumeUnit: string | number;
    // systemVolumeAmount: string | number;
    // sysRaid: string[];
    // dataVolumeType: string;
    // dataVolumeInterfaceType: string;
    // dataVolumeSize: string | null;
    // dataVolumeUnit: string | number;
    // dataVolumeAmount: number;
    nicRate: number | string;
    nicAmount: number;
    interfaceMode: string | number;
    gpuManufacturer: string;
    gpuModel: string;
    gpuAmount: number;
    height: string | number;
    cpuData: CpuDataType[];
    memData: MemDataType[];
};

/**
 * 规则类
*/
interface RulesType {
    name: VerifyRulesType[];
    architecture: VerifyRulesType[];
    deviceType: VerifyRulesType[];
    cpuInfo: VerifyRulesType[];
    cpuManufacturer: VerifyRulesType[];
    cpuModel: VerifyRulesType[];
    cpuCores: VerifyRulesType[];
    cpuFrequency: VerifyRulesType[];
    memInfo: VerifyRulesType[];
    memType: VerifyRulesType[];
    memFrequency: VerifyRulesType[];
    memSize: VerifyRulesType[];
    systemVolumeType: VerifyRulesType[];
    systemVolumeInterfaceType: VerifyRulesType[];
    systemVolumeSize: VerifyRulesType[];
    systemVolumeUnit: VerifyRulesType[];
    sysRaid: VerifyRulesType[];
    dataVolumeType: VerifyRulesType[];
    dataVolumeInterfaceType: VerifyRulesType[];
    dataVolumeSize: VerifyRulesType[];
    dataVolumeUnit: VerifyRulesType[];
    dataVolumeAmount: VerifyRulesType[];
    nicRate: VerifyRulesType[];
    nicAmount: VerifyRulesType[];
    interfaceMode: VerifyRulesType[];
    gpuManufacturer: VerifyRulesType[];
    gpuModel: VerifyRulesType[];
    gpuAmount: VerifyRulesType[];
    height: VerifyRulesType[];
};

interface PropsReactiveType {
    architecture: string[];
    raidData: any;
    volumeTypeData: string[];
    diskTypeData: {label: string; value: string}[];
    noRaidInterfaceData: {label: string; value: string}[];
    interfaceData: {label: string; value: string}[];
    raidConfig: string[];
    noSysRaid: {label: string; value: string}[];
    sysRaid: {label: string; value: string}[];
    sysDiskRaid0: {label: string; value: string}[];
    bootModeData: {label: string; select: boolean;}[];
    computerRoomRadioBtn: {title: string; idcId: string; showTooltip: boolean;}[];
};

interface RadiosPropsType {
    hasTemplate: boolean;
    formData: {value: CurrencyType;}; cpuSpec: string; memSpec: string;
};

export {
    key1,
    CpuDataType,
    MemDataType,
    RulesType,
    RuleFormType,
    ProcessingOperationsType,
    RadiosPropsType,
    PropsReactiveType,
    ReactiveDataType,
    IdcNameRadioBtnType
};
