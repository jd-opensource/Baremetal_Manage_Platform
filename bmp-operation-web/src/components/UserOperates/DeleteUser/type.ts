
type key = 'userName' | 'roleName' | 'phoneNumber' | 'userId' | 'email';

/**
 * 表格数据类
*/
type ListDataType = {
    [k in key]: string;
};


export {
    ListDataType
};
