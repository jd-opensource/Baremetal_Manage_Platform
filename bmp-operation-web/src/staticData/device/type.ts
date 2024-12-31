/**
 * 通用类-列表
*/
type CurrencyType = {
    text: string;
    label: string | number;
    info: any;
    filterParams: string;
};

type DeviceType = Required<{value: number | string;}> & Partial<CurrencyType>;

export default DeviceType;
