type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    idcListUrl: string; // 机房列表
    idcDetailUrl: string; // 机房详情
    editIdcUrl: string; // 编辑机房
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

// 编辑机房类
type EditIdcType = {
    idcId: string;
    address: string;
    iloPassword: string;
    iloUser: string;
    level: string;
    name: string;
    switchPassword1: string
    switchPassword2: string
    switchUser1: string
    switchUser2: string
    nameEn: string;
};

// 接口参数类型
interface ParamDataType {
    IdcListType: Partial<ListType> & Required<{name: string; level: string}>;
    idcListExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
    IdcDetailType: Partial<{idcid?: string; show: string;}>
    EditIdcType: Required<EditIdcType>;
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
