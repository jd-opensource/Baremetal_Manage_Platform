/**
 * 标签页类
*/
type TabsPangeDataType = {
    label: string;
    name: string;
};

/**
 * store库存储的状态值-类
*/
export interface StoreType {
    hasNavigationBar: boolean;
};

/**
 * store库-oss数据类
*/
export interface StoreOssType {
    ossData: any[];
};

/**
 * 复杂数据类型type
*/
export interface ReactiveType {
    tableRef: any;
    filterParams: any;
    detail: any;
    tabsPangeData: TabsPangeDataType[];
    tableData: any[];
};

export interface FilterParamsType {
    deviceTypeId?: string;
    architecture: string;
    osType: string;
};
