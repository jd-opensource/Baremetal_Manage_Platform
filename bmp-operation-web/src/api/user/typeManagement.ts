type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    userListUrl: string; // 用户列表
    addUserUrl: string; // 添加用户
    editUserUrl: string; // 编辑用户
    deleteUserUrl: string; // 删除用户
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
 * 操作用户类
*/
type OperateUserType = {
    description: string;
    email: string;
    language: string;
    phoneNumber: string;
    phonePrefix: string;
    roleId: string;
    userName: string;
    userId: string;
    defaultProjectId: string;
};

/**
 * 接口参数类型
*/
interface ParamDataType {
    UserListType: Partial<ListType> & Partial<{roleId: string; defaultProjectId: string; userName: string;}>;
    UserListExportType: Omit<ListType, 'pageNumber' | 'pageSize'>; 
    AddUserType: Omit<ParamDataType['EditUserType'], 'defaultProjectId' & 'userId'>;
    EditUserType: Required<OperateUserType>;
    DeleteUserType: Required<{userId: string;}>;
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
