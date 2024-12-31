import {SearchArrType} from '../typeManagement';
import {language, paginationOperate} from 'utils/publicClass.ts';

class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: SearchArrType = reactive<SearchArrType>({
        searchParams: {},
        selectOption: [
            {
                value: 0,
                label: language.t('modelList.header.label.modelName')
            },
            {
                value: 1,
                label: language.t('modelList.header.label.modelSpecifications')
            }
        ]
    });
    fn: any;

    constructor (fn: any) {
        this.fn = fn;
    };

    iptValue = (enterValue: string, selectValue: number) => {
        selectValue ? this.reactiveArr.searchParams = {deviceType: enterValue} : this.reactiveArr.searchParams = {name: enterValue};
    };

    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} enterValueIpt 输入框输入的值
     * @param {number} selectValue 筛选框筛选的值
    */
    enterValueIpt = (enterValueIpt: string, selectValue: number): void => {
        selectValue ? this.reactiveArr.searchParams = {deviceType: enterValueIpt} : this.reactiveArr.searchParams = {name: enterValueIpt};
        this.request();
    };

    /**
     * 搜索框筛选
     * @param {number} val 搜索框切换的搜索key
     * @return {number} selectVal 对应的key
    */
    changeSelect = (val: number): number => {
        return this.selectVal.value = val;
    };

    clearClick = (val: string) => {
        if (!val) {
            this.selectChange();
        }
    };

    selectChange = () => {
        const {deviceType, name}: {deviceType?: string; name?: string;} = this.reactiveArr.searchParams;
        if (deviceType || name) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    request = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn();
    };
};

export default SearchOperate;
