import {CurrencyType, RuleFormRefType, VerifyRulesType, SystemPartitionType} from '@utils/publicType';




type RuleFormRefType2 = Omit<RuleFormRefType, 'validateField'>;
interface ReactiveArrType {
    defaultBootMode: string[];
    fileParams: {result?: CurrencyType[]};
    architectureData: CurrencyType[];
    totalData: CurrencyType;
    architecture: CurrencyType[];
    osType: CurrencyType[];
    osVersion: CurrencyType[];
    bootModeData: string[];
    imageFormatData: string[];
};

/**
 * 表单类
*/
interface RuleFormType {
    imageName: string;
    architecture: string | number;
    osType: string;
    customVersion: string;
    imageFormat: string;
    bootMode: CurrencyType[] | null;
    customOperatePlatform?: string;
    version: string;
    systemPartition: SystemPartitionType[];
    description: string;
    imageFile?: CurrencyType[];
};

/**
 * 规则类
*/
interface RulesType {
    imageName: VerifyRulesType[];
    architecture: VerifyRulesType[];
    osType: VerifyRulesType[];
    customOperatePlatform: VerifyRulesType[];
    customVersion:  VerifyRulesType[];
    version: VerifyRulesType[];
    imageFormat: VerifyRulesType[];
    bootMode: VerifyRulesType[];
    systemPartition: VerifyRulesType[];
    imageFile?: VerifyRulesType[];
};


export {
    RulesType,
    RuleFormType,
    ReactiveArrType,
    RuleFormRefType2
};
