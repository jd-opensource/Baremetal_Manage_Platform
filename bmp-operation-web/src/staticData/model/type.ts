/**
 * 通用类-列表
*/
type CurrencyType = {
    text: string;
    label: string | number;
    info: any;
    filterParams: string;
};

type ModelType = Required<{value: number | string;}> & Partial<CurrencyType>;

export default ModelType;
