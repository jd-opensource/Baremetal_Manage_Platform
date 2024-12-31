type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    describeRulesUrl: string; // 全部报警规则列表
    describeRulesExportUrl: string; // 全部报警规则导出
    describeRulesDetailUrl: string; // 全部报警规则详情
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
