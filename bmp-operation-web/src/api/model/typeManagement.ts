type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    modelListUrl: string; // 机型列表
    modelDetailUrl: string; // 机型详情
    relationImageListUrl: string; // 机型详情-关联镜像列表
    addImageUrl: string; // 机型详情-添加镜像
    addModelUrl: string; // 添加机型
    editModelUrl: string; // 编辑机型
    deleteImageUrl: string; // 删除关联镜像
    deleteModelUrl: string; // 删除机型
    getRaidsUrl: string; // raid
    filterListUrl: string; // filter List
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

type ModelRequiredType = {
    deviceTypeId: string;
    name: string;
    deviceType: string;
    deviceSeries: string;
    architecture: string;
    height: string;
    nicAmount: string | number;
    nicRate: string;
    interfaceMode: string;
    systemVolumeType: string;
    systemVolumeInterfaceType: string;
    systemVolumeSize: string | number;
    systemVolumeUnit: string;
    systemVolumeAmount: string | number;
    gpuAmount: string | number;
    gpuManufacturer: string;
    gpuModel: string;
    dataVolumeType: string;
    dataVolumeInterfaceType: string;
    dataVolumeSize: string | number;
    dataVolumeUnit: string;
    dataVolumeAmount: string | number;
    raidId: string;
};

type ModelPartialType = {
    id: string;
    idcId: string;
    description: string;
    cpuManufacturer: string;
    cpuModel: string;
    cpuAmount: string | number;
    cpuCores: string | number;
    cpuFrequency: string | number;
    memType: string;
    memSize: string;
    memAmount: string | number;
    memFrequency: string;
};

// 接口参数类型
interface ParamDataType {
    ModeListType: Partial<ListType> & Partial<{idcId: string; deviceSeries: string;}>;
    ModelListExportType: Partial<ListType> & Partial<{idcId: string; deviceSeries: string;}>;
    ModelDetailType: Required<{id: string;}>;
    RelationImageListType: Required<{deviceTypeId: string}> | Partial<{imageId: string;}>;
    AddImageType: Required<ParamDataType['RelationImageListType']>;
    ModelCurrencyType: Required<ModelRequiredType> & Partial<ModelPartialType>;
    DeleteImageType: Required<ParamDataType['RelationImageListType']>;
    DeleteModelType: Required<ParamDataType['ModelDetailType']>;
    RaidsType: {};
    FilterListType: {};
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
