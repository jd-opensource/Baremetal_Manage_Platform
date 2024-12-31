import {VerifyRulesType} from '@utils/publicType';

/**
 * 规则类-信息
*/
interface RuleInfoType {
    description: VerifyRulesType[];
    phoneNumber: VerifyRulesType[];
    email: VerifyRulesType[];
};

/**
 * 表单类
*/
interface RuleFormType {
    type: string;
    userName: string;
    roleName: string;
    defaultProjectId: string;
    description: string;
    phoneNumber: string;
    phonePrefix: string;
    email: string;
    userId: string;
    roleId?: string;
};

export {
    RuleInfoType,
    RuleFormType
};
