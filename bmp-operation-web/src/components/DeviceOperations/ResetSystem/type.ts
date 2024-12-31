import {CurrencyType, SystemPartitionType, VerifyRulesType} from '@utils/publicType';
interface ItemType {
    instanceName: string;
    instanceId: string;
    showTooltip: boolean;
};


interface DropDownItemType {
    systemPartition: SystemPartitionType[] | string;
    imageName: string;
    imageId: string;
    bootMode: {
        split: (arg0: string) => CurrencyType[];
    };
};

interface QueryType {
    deviceTypeId: string;
    idcId?: string;
    volumeType?: string;
    imageId?: string;
};

interface ReactiveArrType {
    tableData: ItemType[] | {push(arg0: ItemType): void};
    centOSData: CurrencyType[];
    imageTypekeyData: {
        name: string;
        initName: string;
        img?: string;
        childrenData?: { imageName: string; }[]
    }[];
    systemDiskData: CurrencyType[];
    keyPariData: CurrencyType[];
    systemPartitionData: SystemPartitionType[];
    bootModeData: string | string[] | DropDownItemType['bootMode'] | CurrencyType[];
};

type key = 'instanceName' | 'instanceId' | 'imageType' | 'bootMode' | 'systemVolumeRaidId' | 'hostName' | 'setType' | 'password' | 'confirmPassword' | 'sshKeyId' | 'imageId' | 'userName';


type RuleFormType = {
    [k in key]: string;
} & {
    systemPartition: SystemPartitionType[];
} & {installBmpAgent: boolean;};

interface RuleType {
    hostName: VerifyRulesType[];
    password: VerifyRulesType[];
    confirmPassword: VerifyRulesType[];
};

interface RulesCheckType {
    confirmPasswordFlag: {value: boolean};
    colorStatus: {value: boolean};
    hasKeyFlag: {value: boolean};
    errorTip: {value: string};
    hostNameCheck: unknown;
    passwordCheck: unknown;
    confirmPasswordCheck: unknown;
};

interface ChildrenDataType {
    [x: string]: string;
    imageId:string;
    imageName: string;
};

type DropDwomType = {
    initName: string;
    childrenData: ChildrenDataType[];
    name: string;
    imageName: string;
    bootMode: {
        split: (arg0: string) => CurrencyType[];
    };
};

type FormRulesOptType = {
    ruleFormRef: {
        value: {
            validate(arg0: (valid: boolean) => void): void
        };
    };
    ruleForm: {
        setType: string;
        sshKeyId: string;
    };
};

type HeightType = {style: {maxHeight: string;}; scrollHeight: number;};

interface ChooseItemType {keypairId: string;};

export {
    RuleType,
    ItemType,
    QueryType,
    HeightType,
    RuleFormType,
    DropDwomType,
    ChooseItemType,
    RulesCheckType,
    ReactiveArrType,
    ChildrenDataType,
    FormRulesOptType,
    DropDownItemType
}
