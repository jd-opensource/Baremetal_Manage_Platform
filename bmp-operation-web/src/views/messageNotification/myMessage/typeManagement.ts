/**
 * text类
*/
type TextType = {
    [x: string]: {
        text: string;
    };
};

interface OSSType {
    deviceSeries: TextType;
    manageStatus: TextType;
    deviceTypeId: TextType;
    getModelList(): void;
    getOssData(): void;
};
interface SearchArrType {
    searchParams: {
        sn?: string;
        userName?: string;
    };
    selectOption: {
        value: number;
        label: string;
    }[];
};


/**
 * filter参数类
*/
// messageSubType?: string; content?: string; hasRead?: string;
type FilterParamsType = {messageType?: string; messageSubType?: string;};


/**
 * 复杂数据类型类
*/
type ReactiveArrType = {filterParams?: FilterParamsType;};

type HeightType = {offsetHeight: number; offsetTop: number;};

type OperateHeightType = Pick<HeightType, 'offsetHeight'>;

// table ts
type TableType = {devices: string;} & {[x: string]: string | boolean;};
// 跳转路由 ts
type JumpRouterType = {sn: string; deviceId: string};
// 导出类型
type ExportType = {devicesListExportAPI: Function; hasExportData: {value: boolean;}; exportData: Function;};

export {
    OSSType,
    SearchArrType,
    ReactiveArrType,
    FilterParamsType,
    HeightType,
    OperateHeightType,
    TableType,
    JumpRouterType,
    ExportType
};
