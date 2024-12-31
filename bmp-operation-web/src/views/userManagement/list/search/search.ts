import {SearchArrType} from '../typeManagement';
import {language, paginationOperate} from 'utils/publicClass.ts';

class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    userList: {
        requestOperate(arg0: {}, arg1: string): void;
    };
    reactiveArr: SearchArrType = reactive<SearchArrType>({
        searchParams: {},
        selectOption: [
            {
                value: 0,
                label: language.t('userList.search.condition.userName')
            }
        ]
    });

    constructor(userList: { requestOperate(arg0: {}, arg1: string): void; }) {
        this.userList = userList;
    };

    iptValue = (enterValue: string) => {
        this.reactiveArr.searchParams = {userName: enterValue};
    };


    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} enterValueIpt 输入框输入的值
    */
    enterValueIpt = (enterValueIpt: string): void => {
        this.reactiveArr.searchParams = {userName: enterValueIpt};
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
        const {userName} :{userName?: string;} = this.reactiveArr.searchParams;
        if (userName) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    request = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.userList.requestOperate(this.reactiveArr.searchParams, 'search');
    };

};

export default SearchOperate;
