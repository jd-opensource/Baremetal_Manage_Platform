// 登录返回的类
type LoginResponseType = {data: {result: {success: boolean;}}, loginUserName: string};

// 用户权限类
type UserPurviewType = {getUserPurview(arg0: any): Promise<void>;}
/**
 * 用户输入的类
*/
type UserInfoType = {
    required: boolean;
    message: string;
    trigger: string;
};

type key = 'loginUserName' | 'loginPassword';

/**
 * 表单提交类
*/
type RuleFormType = {
    [k in key]: string;
};

/**
 * 表单规则类
*/
interface RulesType {
    loginUserName: UserInfoType[];
    loginPassword: UserInfoType[];
};

type CheckType = {
    ruleForm: {
        loginUserName: string;
        loginPassword: string;
    };
};

export {
    CheckType,
    RulesType,
    RuleFormType,
    UserPurviewType,
    LoginResponseType
};
