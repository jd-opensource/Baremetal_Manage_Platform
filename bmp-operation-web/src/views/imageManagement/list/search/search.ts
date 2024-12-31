/**
 * @file
 * @author
*/

import {language, paginationOperate} from 'utils/publicClass.ts';

class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: any = reactive<any>({
        searchParams: {},
        selectOption: [
            {
                value: 0,
                label: language.t('imageList.search.condition.imageName')
            },
            {
                value: 1,
                label: language.t('imageList.search.condition.imageId')
            }
        ]
    });
    fn: any;

    constructor (fn: any) {
        this.fn = fn;
    };

    iptValue = (enterValue: string, selectValue: number) => {
        selectValue ? this.reactiveArr.searchParams = {imageId: enterValue} : this.reactiveArr.searchParams = {imageName: enterValue};
    };

    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} enterValue 输入框输入的值
     * @param {number} selectValue 筛选框筛选的值
    */
    enterValueIpt = (enterValue: string, selectValue: number): void => {
        selectValue ? this.reactiveArr.searchParams = {imageId: enterValue} : this.reactiveArr.searchParams = {imageName: enterValue};
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
        const {imageId, imageName} :{imageId?: string; imageName?: string;} = this.reactiveArr.searchParams;
        if (imageId || imageName) {
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
