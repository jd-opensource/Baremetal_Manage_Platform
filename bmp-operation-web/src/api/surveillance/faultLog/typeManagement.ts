type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    faultLogUrl: string; // 故障报警日志
    faultLogExportUrl: string; // 故障报警日志导出
    faultLogDealUrl: string; // 故障报警日志处理
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

/**
 * idc-机房列表类
*/
type ListType = {
    pageNumber: number;
    pageSize: number;
    isAll: boolean | number;
    exportType: string;
};

/**
 * 接口参数类型
*/
interface ParamDataType {
    faultLogListType: ListType;
    faultLogExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
    faultLogDealType: {logid: number};
};

/**
 * 请求类型
*/
interface ReqType {
    CurrencyType: {
        isBusinessProcessing: boolean;
        url: string;
        methods: string;
    }
};

// 请求类型
type RequestType = Required<ReqType>;

export {
    RequestType,
    ParamDataType,
    TotalUrlPathType
};
