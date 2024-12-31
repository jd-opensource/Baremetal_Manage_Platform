type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    faultRulesUrl: string; // 故障报警规则
    faultRulesExportUrl: string; // 故障报警规则导出
    faultPushUrl: string; // 故障推送
    faultUseUrl: string; // 故障启用
    faultResetUrl: string; // 恢复默认设置
    faultLevelUrl: string; // 故障等级
    faultSpecUrl: string; // 规则配件
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
    faultRulesListType: ListType;
    faultPushType: {role_id: string; push_status: number;}
    faultUseType: {role_idd: string; use_status: number;}
    faultRulesLisExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
    faultResetType: {role_ids: string;};
    faultLevelType: {};
    faultSpecType: {};
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
