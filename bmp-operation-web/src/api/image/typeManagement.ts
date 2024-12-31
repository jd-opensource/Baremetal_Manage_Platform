type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    imagesUrl: string; // 镜像列表
    imagesDetailUrl: string; // 镜像详情
    imageDeviceTypesUrl: string; // 镜像详情-机型列表-添加机型列表
    imageAddModelUrl: string; // 镜像详情-机型列表-添加机型
    deleteImageModelUrl: string; // 镜像详情-机型列表-移除关联机型
    imageEditUrl: string; // 镜像详情-编辑描述
    imagesDeleteUrl: string; // 删除镜像
    importImageUrl: string; // 导入镜像
    imageOssUrl: string; // 镜像oss数据
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

type SysPartitionType = {
    format: string;
    point: string;
    size: number;
};

interface ImportImage {
    imageFile: unknown;
    imageName: string;
    architecture: string;
    description: string;
    systemPartition: SysPartitionType[];
    osType: string;
    osVersion: string;
};

// 接口参数类型
interface ParamDataType {
    ImagesType: Partial<ListType> &  Partial<{imageId: string; imageName: string; deviceTypeId: string; osId: string;}>;
    ImagesModelType: {
        imageId: string;
        architecture: string;
        isBind: string;
    };
    ImageAddModelType: {
        imageId: string;
        deviceTypeIds: string;
    };
    ImageEditType: {imageId: string; description: string;};
    DeleteImageModelType: {
        imageId: string;
        deviceTypeId: string;
    };
    ImagesDeleteType: Required<{imageId: string;}>;
    ImportImageType: ImportImage;
    ImageOssType: {};
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
