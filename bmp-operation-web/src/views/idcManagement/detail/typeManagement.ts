import {Router, CurrencyType, LocationQueryValue} from '@utils/publicType';

type RouterType = {router: Router; jumpRouter(path: string): void;};

/**
 * 新table数据类型
*/
interface NewTableType {
    name: string;
    user: string;
    password: string;
    hasPassword: boolean;
    type: string;
};

interface ReactiveArrType {
    detailInfo: CurrencyType;
    tableData: NewTableType[]
};

interface PasswordType {
    validate: {
        securityVerificationDiaLog: {
            value: boolean;
        };
        hasEditOperate: {
            value: boolean;
        };
    }
    idcDetail: {
        password: {
            value: string;
        }
    }
};

export {
    RouterType,
    NewTableType,
    CurrencyType,
    PasswordType,
    ReactiveArrType,
    LocationQueryValue
};
