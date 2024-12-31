// import {SearchArrType} from '../typeManagement';
import {paginationOperate, language} from 'utils/publicClass.ts';
class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: any = reactive<any>({
        detail: '',
        radioGroup: {
            hasReadName: language.t('myMessage.radioGroup.all'),
            hasRead: ''
        },
        searchParams: {
            // hasRead: '',
            detail: ''
        },
        selectOption: [
            {
                value: 0,
                label: language.t('myMessage.search.title')
            }
        ]
    });
    fn: Function;

    constructor (fn: Function) {
        this.fn = fn;
        this.reactiveArr.radioGroup.hasRead = this.#setRadioRead();
        this.watchRadio();
    };

    #setRadioRead = () => {
        const type = new Map([
            [language.t('myMessage.radioGroup.all'), ''],
            [language.t('myMessage.radioGroup.read'), '1'],
            [language.t('myMessage.radioGroup.noRead'), '0']
        ]);
        return type.get(this.reactiveArr.radioGroup.hasReadName);
    }

    iptValue = (enterValue: string) => {
        this.reactiveArr.detail = enterValue;
        // this.reactiveArr.searchParams = {detail: enterValue}
        // selectValue ? this.reactiveArr.searchParams = {deviceType: enterValue} : this.reactiveArr.searchParams = {name: enterValue};
    };

    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} enterValueIpt 输入框输入的值
    */
    enterValueIpt = (enterValueIpt: string): void => {
        this.reactiveArr.searchParams = {detail: this.reactiveArr.detail}
        this.reactiveArr.searchParams = {detail: enterValueIpt};
        // selectValue ? this.reactiveArr.searchParams = {deviceType: enterValueIpt} : this.reactiveArr.searchParams = {name: enterValueIpt};
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
        const {detail}: {detail?: string;} = this.reactiveArr.searchParams;
        if (detail) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    noReadClick = () => {
        this.reactiveArr.radioGroup.hasReadName = language.t('myMessage.radioGroup.noRead')
    }

    watchRadio = () => {
        watch(() => this.reactiveArr.radioGroup.hasReadName, (_: string) => {
            this.reactiveArr.radioGroup.hasRead = this.#setRadioRead();
            this.fn('listTimer');
        })
    }

    request = () => {
        if (paginationOperate.pageNumber.value > 1) {
            paginationOperate.pageNumber.value = 1;
            return;
        }
        this.fn('listTimer');
    };
};

export default SearchOperate;
