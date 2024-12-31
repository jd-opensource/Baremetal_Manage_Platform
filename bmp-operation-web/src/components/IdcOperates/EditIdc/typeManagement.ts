import {RuleFormRefType, VerifyRulesType} from '@utils/publicType';

/**
 * 表单-输入项ts类型
*/
interface FormType {
    value: string;
    idcId: string;
    name: string;
    shortname: string;
    nameEn: string;
    level?: string;
    customGrade?: string;
    address: string;
    iloUser: string;
    iloPassword: string;
    switchUser1: string;
    switchPassword1: string;
    switchUser2: string;
    switchPassword2: string;
};

/**
 * 表单-规则ts类型
*/
interface RulesType {
    name: VerifyRulesType[];
    nameEn: VerifyRulesType[];
    level: VerifyRulesType[];
    customGrade: VerifyRulesType[];
    address: VerifyRulesType[];
    iloUser: VerifyRulesType[];
    iloPassword: VerifyRulesType[];
    switchUser1: VerifyRulesType[];
    switchPassword1: VerifyRulesType[];
    switchUser2: VerifyRulesType[];
    switchPassword2: VerifyRulesType[];
};

export {
    FormType,
    RulesType,
    RuleFormRefType
};
