type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    hardwareStatusUrl: string; // 硬件报警状态
    hardwareStatusExportUrl: string; // 硬件报警状态导出
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

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
    hardwareStatusType: ListType;
    hardwareStatusExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
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
