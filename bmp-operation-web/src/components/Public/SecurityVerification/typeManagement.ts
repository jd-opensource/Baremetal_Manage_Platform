import {RuleFormRefType, VerifyRulesType} from '@utils/publicType';

type RuleFormRefType2 = Omit<RuleFormRefType, 'validateField'>;

/**
 * 表单输入类型
*/
type RuleFormType = {username: string; password: string;};

/**
 * 规则类型
*/
interface RulesType {
    username: VerifyRulesType[];
    password: VerifyRulesType[];
};

export {
    RulesType,
    RuleFormType,
    RuleFormRefType2
};
