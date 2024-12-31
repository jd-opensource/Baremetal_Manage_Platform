type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    devicesListUrl: string; // 设备列表
    deviceTemplateUrl: string; // 设备模板
    devicesUpUrl: string; // 设备上架
    devicesDownUrl: string; // 设备下架
    devicesRemoveUrl: string; // 设备移除
    instanceRecoveryUrl: string; // 回收实例
    resetInstanceUrl: string; // 重置实例
    devicesEditUrl: string; // 设备编辑
    devicesDeleteUrl: string; // 设备删除
    devicesStartUrl: string; // 设备开机
    devicesResetPasswordUrl: string; // 设备重置密码
    devicesBatchStartUrl: string; // 设备批量开机
    devicesStopUrl: string; // 设备关机
    devicesBatchStopUrl: string; // 设备批量关机
    devicesRestartUrl: string; // 设备重启
    devicesBatchRestartUrl: string; // 设备批量重启
    deviceBatchResetPasswordUrl: string; // 设备批量重置密码
    lockUrl: string; // 锁定
    unlockUrl: string; // 解锁
    resetSystemUrl: string; // 重装系统
    devicesDetailUrl: string; // 设备详情
    devicesCollectUrl: string; // 设备采集
    instanceDetailUrl: string; // 实例详情
    importDevicesUrl: string; // 导入设备
    batchEditInstanceNameUrl: string; // 批量编辑实例名称
    batchRecoveryInstanceUrl: string; // 批量回收实例
    deviceStatusUrl: string; // 获取设备带外总体状态
    consolelogsUrl: string; // 设备详情-硬件监控-报警日志列表接口
    disksDetailUrl: string; // 磁盘详情
    describeAssociateDisksUrl: string; // 获取设备关联磁盘信息
    associateDisksUrl: string; // 关联设备磁盘
    describeVolumesUrl: string; // 根据机型获取卷列表信息
    associateUrl: string; // 设备关联机型
    auditLogsUrl: string; // 设备详情-操作日志
    auditLogsTypesUrl: string; // 获取操作日志类型的接口
    monitorDataUrl: string; // 监控图表
    desrcibeTagsUrl: string; // 监控tags
    desrcibeAgentStatusUrl: string; // 监控agent状态
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

type DeviceOperateType = {
    deviceId: string;
    sns: string;
};

// 接口参数类型
interface ParamDataType {
    DeleteImageType: Required<{deviceTypeId: string}> | Partial<{imageId: string;}>;
    DeleteModelType: Required<{id: string;}>;
    DevicesListType: Partial<ListType> & Partial<{sn: string; deviceTypeId: string; manageStatus: string; iloIp: string; instanceOwer: string;}>;
    DevicesListExportType:  Partial<ListType> & Partial<{sn: string; deviceTypeId: string; manageStatus: string; iloIp: string; instanceOwer: string;}>;
    ConsolelogsType: ListType & {sn: string};
    DeviceUpType: Required<DeviceOperateType>;
    DeviceDownType: Required<ParamDataType['DeviceUpType']>;
    DevicesRemoveType: Required<{deviceId: string}>;
    InstanceRecoveryType: Required<{instanceId: string;}>;
    ResetInstanceType: Required<{instanceId: string;}>;
    DevicesEditType: {id?: string; description: string;};
    DeviceDeleteType: Required<{deviceId: string;}>;
    DevicesStartType: Required<{instanceId: string;}>;
    DevicesResetPasswordType: Required<{instanceId: string; idcId: string; password: string;}>;
    DevicesStopType: Required<{instanceId: string;}>;
    DevicesRestartType: Required<ParamDataType['DevicesStartType']>;
    BatchResetPasswordType: {instanceIds: string[]; password: string;};
    LockType: {instanceId: string;};
    UnlockType: {instanceId: string;};
    DevicesDetailType: {deviceId?: string; show: string};
    InstanceDetailType: {instanceId: string};
    ResetSystemType: any;
    DeviceTemplateType: {};
    ImportDevicesType: {idcId: string; deviceTypeId: string};
    BatchEditInstanceNameType: any;
    BatchRecoveryInstanceType: any;
    deviceStatusType: {pageNumber: number; pageSize: number; idcId: string; sn: string;};
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
