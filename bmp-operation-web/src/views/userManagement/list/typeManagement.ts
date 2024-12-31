import {PaginationType, LocationQueryValue, CurrencyType} from '@utils/publicType';

// 筛选参数类
type FilterParamsType = {roleId: string;};

type TableRef = {
    [x: string]: unknown;
    value?: {
        clearFilter(): unknown;
    };
};

/**
 * oss store参数类
*/
type OSStoreType = {
    filterParams: string;
    showInfo: string;
    text: string;
    value: number;
};

/**
 * 参数类
*/
type ParamsType = PaginationType
    | PaginationType & Partial<FilterParamsType>
    | PaginationType & Partial<{userName: string | undefined | LocationQueryValue[]}>
    | PaginationType & Partial<FilterParamsType> & Partial<{userName: string | undefined | LocationQueryValue[]}>;

type key = 'userName' | 'roleId' | 'roleName' | 'userId' | 'phoneNumber' | 'email';
/**
 * 用户公用类
*/
type UserType = {
    [k in key]: string;
};

// 删除用户类
type DeleteUserType = Required<UserType>;

/**
 * 编辑用户类
*/
interface EditUserType extends Required<UserType> {
    [x: string]: string | boolean;
    defaultProjectId: string;
    showTooltip: boolean;
    phonePrefix: string;
    description: string;
};

// 用户数据类
interface UsersDataType extends Required<EditUserType>{};

/**
 * 复杂数据类型类
*/
interface ReactiveArrType {
    filterParams?: FilterParamsType;
    tableData?: EditUserType[];
};

// 数据列表类
interface DataListType {tableData: EditUserType[];};

type DataOperateType = {requestOperate(arg0: CurrencyType, arg1: string): void;};

interface ClassParamsType {
    filterOperates: {
        filterEmptyInfo: {
            deleteEmtpyData: void;
        }
        reactiveArr: {
            filterParams: {roleId: string;};
        };
        tableRef: TableRef;
    };
    searchOperate: {
        hasClear: {
            value: boolean;
        };
        reactiveArr: {
            searchParams: {userName: string};
        };
    },
    userList: {
        reactiveArr: {
            tableData: {length: number};
        };
        sessionId: {value: string};
        resetData(arg0: ClassParamsType['searchOperate']): () => void;
        requestOperate(): () => void;
        getUserList(): () => void;
    };
    refreshOperate: {
        reset(): void;
    }
};

interface UserListType {
    sessionId: {
        value: string;
    };
    requestOperate(arg1: FilterParamsType, arg2: string): void;
};

interface SearchArrType {
    searchParams: {
        userName?: string;
    };
    selectOption: {
        value: number;
        label: string;
    }[];
};

export {
    TableRef,
    ParamsType,
    OSStoreType,
    DataListType,
    EditUserType,
    UserListType,
    SearchArrType,
    UsersDataType,
    DeleteUserType,
    ClassParamsType,
    DataOperateType,
    ReactiveArrType,
    FilterParamsType
};
