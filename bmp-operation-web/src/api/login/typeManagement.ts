type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    loginUrl: string; // 登录
    userPurviewUrl: string; // 用户权限
    loginOutUrl: string; // 退出登录
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

type UserInfo = {
    username: string;
    password: string;
};

// 接口参数类型
interface ParamDataType {
    LoginType: Required<UserInfo>;
    UserPurviewType: {};
    LoginOutType: {};
};

interface RequestType {
    CurrencyType: {
        isBusinessProcessing: boolean;
        url: string;
        methods: string;
    };
};

export {
    TotalUrlPathType,
    ParamDataType,
    RequestType
};
