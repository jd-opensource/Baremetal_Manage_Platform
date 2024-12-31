type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    timeOneUrl: string; // 时区
    userInfoUrl: string; // 用户信息
    setUserInfoUrl: string; // 修改用户信息
    deleteApiKeyUrl: string; // 删除apikey
    createApiKeyUrl: string; // 创建apikey
    getApiKeyUrl: string; // 获取apikey
    revisePasswordUrl: string; // 修改密码
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;


/**
 * 接口参数类型
*/
interface ParamDataType {
    TimeOneType: {};
    UserInfoType: {};
    GetApiKeyType: {
        isAll: string;
    };
    CreateApiKeyType: {
        name: string;
        readOnly: number;
        type: string;
    };
    DeleteApiKeyType: {
        apiKey: string;
    };
    SetUserInfoType: {
        email: string;
        phoneNumber: string;
        language: string;
        timezone: string;
        userName: string;
        phonePrefix: string;
    };
    RevisePasswordType: {
        oldPassword: string;
        password: string;
    };
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
