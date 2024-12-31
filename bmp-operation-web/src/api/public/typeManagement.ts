type SetType<T> = T extends {} ? any : T;

type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    resourcesUrl: string; // 精确查询
    securityVerificationUrl: string; // 安全验证
    customListUrl: string; // 自定义列表
    setCustomListUrl: string; // 设置自定义列表
    surveillanceCustomListUrl: string; // 监控自定义列表
    surveillanceSetCustomListUrl: string; // 监控设置自定义列表
    getOssUrl: string; // oss数据
    queryImagesUrl: string; // 镜像类型
    systemRaidsUrl: string; // raids
    querySystemPartitionUrl: string; // 系统盘分区
    keyInfoUrl: string; // 密钥信息
    localUrl: string; // 权限列表
};

type UserInfo = {
    username: string;
    password: string;
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

// oss数据类
type OssDataType = {
    osName: string;
    osType: string;
    platform: string;
    osVersion: string;
    isAll: string | number;
};

// 接口参数类型
interface ParamDataType {
    ResourcesType: Partial<{
        name: string;
        userName: string;
        imageName: string;
        deviceType: string;
    }>;
    SecurityVerificationType: Required<UserInfo>;
    SetCustomListType: Required<{pageKey: string; pageValue: SetType<{}>;}>;
    CustomListType: Omit<ParamDataType['SetCustomListType'], 'pageValue'>;
    OssTyoe: Partial<OssDataType>;
    QueryImagesType: {idcId: string; deviceTypeId: string;};
    SystemRaidsType: {deviceTypeId: string; volumeType: string;};
    SystemPartitionType: {deviceTypeId: string; imageId: string;};
    KeyInfoType: {isAll: string;};
};

// 请求类型
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
