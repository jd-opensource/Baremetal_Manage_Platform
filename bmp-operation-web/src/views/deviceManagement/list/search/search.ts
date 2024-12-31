import {language, paginationOperate} from 'utils/publicClass.ts';

class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: any = reactive<any>({
        searchParams: {},
        selectOption: [ // search搜索框搜索名称
            {
                value: 0,
                label: language.t('deviceList.operate.label.sn')
            },
            {
                value: 1,
                label: language.t('deviceList.operate.label.instanceOwner')
            }
        ]
    });
    fn: any;

    constructor (fn: any) {
        this.fn = fn;
    };

    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} enterValueIpt 输入框输入的值
     * @param {number} selectValue 筛选框筛选的值
    */
    enterValueIpt = (enterValueIpt: string, selectValue: number): void => {
        selectValue ? this.reactiveArr.searchParams = {userName: enterValueIpt} : this.reactiveArr.searchParams = {sn: enterValueIpt};
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
        const {userName, sn} :{userName?: string; sn?: string;} = this.reactiveArr.searchParams;
        if (userName || sn) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    request = () => {
        // this.deviceList.hasCheckAll.value = false;
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn();
        // this.deviceList.setFilterData(this.reactiveArr, 'search', this.hasClear);
        // this.imageList.getImageList();
    };

};

export default SearchOperate;
