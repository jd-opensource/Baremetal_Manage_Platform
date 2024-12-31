/**
 * 通用类-列表
*/
type CurrencyType = {
    text: string;
    label: string | number;
    info: any;
    filterParams: string;
};

/**
 * 角色数据类
*/
type RoleNameDataType = {
    filterParams: string;
    showInfo: string;
    text: string;
    value: number;
};

type UserType = Required<{value: number | string;}> & Partial<CurrencyType>;

export {
    UserType,
    RoleNameDataType
};
