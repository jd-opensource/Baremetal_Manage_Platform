/**
 * @file
 * @author
*/

import {Router, LocationQueryValue} from 'vue-router'; // 路由 路由value类

// 将类型转为字符串有一定的限制，仅支持以下类型
export type CanStringified = string | number | bigint | boolean | null | undefined;

// 将支持的类型转化为字符串
export type Stringify<T extends CanStringified> = `${T}`

export type SuccessType = {result: {success: boolean;}};

export const identity = <T>(arg: T): T => {
    return arg;
};

/**
 * 通用状态类
*/
export type CurrencyStatusType = {
    required: boolean;
    selected: boolean;
};

export type CurrencyStatusType2 = {
    disabled: boolean;
    selected: boolean;
};

export interface OsStoreType {
    filterParams: string;
    showInfo?: string;
    text: string;
    value: number;
};

export type CurrencyType = {[x: string]: string;};
export type CurrencyType1 = {[x: string]: string | boolean;};
export type CurrencyType2 = {[x: string]: string | number;};
export type CurrencyType3 = {[x: string]: unknown;};
export type CurrencyType4 = {[x: string]: {selected: boolean;}};
export type CurrencyType5 = {value: number; label: string;};
export type CurrencyType6 = {name: string; label: string;};
export type CurrencyType7 = {value: number; label: number;};
export type CurrencyType8 = {value: string; label: string;};
export type CurrencyType9 = {value: string | number; text: string;};
export type PublicType6 = {name: number; label: string;};

export type StatusParamsType<T extends (type: boolean) => boolean> = T;

export type StatusType<T extends () => boolean> = T;

export type HeightType = {offsetHeight: number; offsetTop: number;};

export type HeightType2 = Pick<HeightType, 'offsetHeight'>;

export type HeightType3 = {offsetHeight: number; scrollHeight: number;};

export type ListShowTooltipType = {[x: string]: {showTooltip?: boolean}};

/**
 * 显示列表的数据类
*/
export type CheckBoxType = {
    name: string;
    checkbox: boolean;
    disabled: boolean;
};

/**
 * 自定义信息类
*/
export interface CustomInfoType {
    hasCustomInfo?: Record<string, {selected: boolean;}>;
    checkListArr?: CheckBoxType[];
};

// 显示列表的数据类
export type CheckListArrType = CheckBoxType[];


export interface RequestParamsType extends CurrencyType2 {
    exportType: string;
    isAll: string;
};

// 分页类
export type PaginationType = Required<{
    pageNumber: number;
    pageSize: number;
}>;

export interface TabsDataType {
    tabsPangeData: {
        label: string;
        name: string;
    }[];
};

// 添加用户 & 编辑用户
export interface RuleFormRefType {
    resetFields(): Function;
    clearValidate(arg0: string): Function;
    validateField(arg0: string, arg1: (valid: string) => void): Function;
    validate(arg0: (valid: boolean) => void): Function;
};

export interface UpLoadsType {
    submit(): Function;
    clearFiles(): Function;
    handleStart(arg0: CurrencyType): Function;
    abort(): Function;
};

export type VerifyRulesType = {
    required: boolean;
    trigger?: string;
    message?: string;
    validator?: unknown;
};

export type SystemPartitionType = {format: string; point: string; size: number;};

export type CallbackType = (arg0?: Error | string) => void;

export interface CpuDataType extends CurrencyType5 {
    info: {
        cpuAmount: number;
        cpuCores: number;
        cpuFrequency: string;
        cpuManufacturer: string;
        cpuModel: string;
    }
};

export interface MemDataType extends CurrencyType5 {
    info: {
        memAmount: number;
        memFrequency: number;
        memSize: number;
        memType: string;
    }
};

/**
 * 接收重复请求的数组类型
*/
export interface RepeatRequestArrType {
    methodUrl: string;
    params: CurrencyType3;
    f: Function;
};

export {
    Router,
    LocationQueryValue
};
