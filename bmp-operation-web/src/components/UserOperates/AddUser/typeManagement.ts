import {VerifyRulesType} from '@utils/publicType';

/**
 * 规则类-信息
*/
interface RuleInfoType {
    userName: VerifyRulesType[];
    password: VerifyRulesType[];
    phoneNumber: VerifyRulesType[];
    email: VerifyRulesType[];
};

/**
 * 表单类信息
*/
interface RuleFormType {
    userName: string; // 用户名
    password: string; // 密码
    roleId: string;
    description: string; // 描述
    phonePrefix: string; // 手机号类型
    phoneNumber: string; // 手机号
    email: string; // 邮箱
};

interface AddUserFormRulesType {
    ruleForm: RuleFormType;
    phoneFlag: boolean;
    validate(arg0: string, arg1: boolean): Function;
};

export {
    RuleInfoType,
    RuleFormType,
    AddUserFormRulesType
};
