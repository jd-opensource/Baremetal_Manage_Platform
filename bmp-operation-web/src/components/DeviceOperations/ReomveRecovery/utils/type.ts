// 表格类
interface TableDataType {
    sn: string;
    instanceName: string;
    instanceId: string;
    userName: string;
    privateIpv4: string;
    deviceId: string;
};

type ParamsType = {
    deviceId?: string;
    instanceId?: string;
};

export {
    TableDataType,
    ParamsType
}
