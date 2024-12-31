import {PublicType6, CurrencyStatusType, CurrencyStatusType2} from '@utils/publicType';

// 需要显示的列表信息类
type CheckListArrType = Partial<PublicType6> & Required<CurrencyStatusType2> & Required<{showTooltip: boolean;}>;

// 自定义信息类
type HasCustomInfoType = {[x: number]: CurrencyStatusType};

type ParamsType = {pageKey?: string; pageValue: HasCustomInfoType};

export {
    ParamsType,
    CheckListArrType,
    HasCustomInfoType
};
