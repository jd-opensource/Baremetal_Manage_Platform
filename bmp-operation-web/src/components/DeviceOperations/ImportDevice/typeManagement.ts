import {CurrencyType, RuleFormRefType, VerifyRulesType} from '@utils/publicType';

type RuleFormRefType2 = Omit<RuleFormRefType, 'validateField'>;
// 'modelName' | 
type key = 'idcName' | 'deviceInfo';

type RulesType = {
    [k in key]: VerifyRulesType[];
};

interface RuleFormType {
    fileParams: CurrencyType;
    idcName: string;
    // modelName: string;
    deviceInfo: CurrencyType[];
};

export {
    RulesType,
    RuleFormType,
    // RuleFormRefType,
    RuleFormRefType2
};
