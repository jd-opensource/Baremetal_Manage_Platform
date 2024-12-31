type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    describeAlertsUrl: string; // 报警历史列表
    describeAlertsExportUrl: string; // 报警历史导出
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

// idc-机房列表类
type ListType = {
    pageNumber: number;
    pageSize: number;
    isAll: boolean | number;
    exportType: string;
};


// 接口参数类型
interface ParamDataType {
    IdcListType: Partial<ListType> & Required<{name: string; level: string}>;
    idcListExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
};

interface RequestType {
    CurrencyType: {
        isBusinessProcessing: boolean;
        url: string;
        methods: string;
    }
};

export {
    TotalUrlPathType,
    ParamDataType,
    RequestType
};
