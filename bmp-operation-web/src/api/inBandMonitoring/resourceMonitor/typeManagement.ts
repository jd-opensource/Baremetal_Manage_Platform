type AddReadonly<T> = {readonly [P in keyof T]: T[P]};

type UrlPathType = {
    monitorDataUrl: string; // 监控图表
};

// url地址类型
type TotalUrlPathType = AddReadonly<UrlPathType>;

export {
    TotalUrlPathType
};
