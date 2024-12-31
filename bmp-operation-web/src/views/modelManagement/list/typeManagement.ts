import {CurrencyType1 } from "@utils/publicType";

type TableType = {name: string; deviceType: string;} & CurrencyType1; 

type JumpRouterType = Required<{type: string; deviceTypeId: string; architecture: string;}>;

// 导出类型
type ExportType = {modelListExportAPI: Function; hasExportData: {value: boolean;}; exportData: Function;};

/**
 * text类
*/
type TextType = {
    [x: string]: {
        text: string;
    };
};

interface OSSType {
    deviceSeries: TextType;
    getOssData(): void;
};

interface SearchArrType {
    searchParams: {
        name?: string;
        deviceType?: string;
    };
    selectOption: {
        value: number;
        label: string;
    }[];
};

/**
 * 参数类型
*/
interface ParamType {
    pageNumber: number;
    pageSize: number;
};

interface RefreshType {
    // modelList: {
    //     reactiveArr: {
    //         tableData: string | never[];
    //     };
    //     getModelList(): void;
    // };
    modelList: { reactiveArr: { tableData: string | unknown[]; }; getModelList: () => void; };
    reset: {
        reset(): void;
    };
}

/**
 * filter参数类
*/
type FilterParamsType = {deviceSeries?: string;};


/**
 * 复杂数据类型类
*/
type ReactiveArrType = {filterParams?: FilterParamsType;};
interface ResetType {
    filter: {
        tableRef: {
            value: {
                clearFilter: () => void;
            };
        };
        reactiveArr: {
            filterParams: {};
        };
    };
    search: {
        hasClear: {
            value: boolean;
        };
        reactiveArr: {
            searchParams: {};
        };
        request(): void;
    }
};


interface TableScrollType {
    filters: {
        filterParams: {
            deviceSeries?: string;
        };
    };
    search: {
        searchParams: {
            name?: string;
            deviceType?: string;
        };
    }
}

export {
    OSSType,
    TableType,
    JumpRouterType,
    ExportType,
    SearchArrType,
    RefreshType,
    TableScrollType,
    ResetType,
    ReactiveArrType,
    ParamType
};

