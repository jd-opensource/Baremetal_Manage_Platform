// 系统盘分区模板类
type SystemPartitionType = {format: string; point: string; size: number;};

/**
 * params 参数类
*/
interface ParamsType {
    pageSize: number;
    pageNumber: number;
};

/**
 * 筛选params参数
*/
type FilterParamsType = {
    source?: string;
    architecture: string;
    osType: string;
};

/**
 * 复杂类
*/
interface ReactiveArrType {
    filterParams?: FilterParamsType;
};

/**
 * text类
*/
type TextType = {
    [x: string]: {
        text: string;
    };
};

interface OSSType {
    source: TextType;
    osType: TextType;
    architecture: TextType;
    getOssData(): Function;
};

export {
    OSSType,
    ParamsType,
    ReactiveArrType,
    FilterParamsType,
    SystemPartitionType
};
