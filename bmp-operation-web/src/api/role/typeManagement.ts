type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    roleListUrl: string; // 角色列表
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
    roleListType: ListType;
};

/**
 * 请求类型
*/
interface ReqType {
    CurrencyType: {
        isBusinessProcessing: boolean;
        url: string;
        methods: string;
        params: ParamDataType['roleListType'];
    }
};

// 请求类型
type RequestType = Required<ReqType>;

export {
    RequestType,
    ParamDataType,
    TotalUrlPathType
};
