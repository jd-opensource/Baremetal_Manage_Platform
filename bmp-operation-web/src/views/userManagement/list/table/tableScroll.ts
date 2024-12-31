import {tableScroll} from 'utils/index.ts';
import {HeightType2, HeightType3} from '@utils/publicType';
import {ClassParamsType} from '../typeManagement';
import {TimerOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class TableScrollOperate extends TimerOperate {
    searchTipHeight: number = 13;
    fixedHeight: Ref<number> = ref<number>(0);
    tableMaxHeight: Ref<number> = ref<number>(0);
    status: boolean = false;
    filters: ClassParamsType['filterOperates'];
    search: ClassParamsType['searchOperate'];
    headerTitle: Ref<HeightType3> = ref<HeightType3>({offsetHeight: 0, scrollHeight: 0});
    operate: Ref<HeightType3> = ref<HeightType3>({offsetHeight: 0, scrollHeight: 0});
    operateBtn: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    searchTip: {
        value: boolean;
    }
    constructor(filter: ClassParamsType['filterOperates'], search: ClassParamsType['searchOperate'], searchTip: {value: boolean;}) {
        super();
        this.filters = filter;
        this.search = search;
        this.searchTip = searchTip;
        onMounted(async() => {
            this.#setScroll(this.#setHeight(), this.tableMaxHeight);
            this.#onResize(this.tableMaxHeight);
            await this.#watchNavBar();
        });
        this.#watchFilterSeach();

        onUnmounted(() => {
            clearTimeout(this.timer);
        })
    };

    #watchFilterSeach = () => {
        watch(() => [
            this.filters.reactiveArr.filterParams,
            this.search.reactiveArr.searchParams,
            this.searchTip
        ], (newValue: any) => {
            if (newValue[2].value) {
                this.#setScroll(this.fixedHeight.value, this.tableMaxHeight);
            }
        }, {deep: true})
    };

    #watchNavBar = () => {
        watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
            this.timer = setTimeout(() => {
                this.#setScroll(this.#setHeight(), this.tableMaxHeight);
            }, 200)
        });
    };

    #onResize = (tableMaxHeight: {value: number}) => {
        window.onresize = () => {
            this.#setScroll(this.#setHeight(), tableMaxHeight);
        };
    };

    #setHeight = () => {
        const header = this.headerTitle.value;
        const operate = this.operate.value;
        this.fixedHeight.value = header.offsetHeight * 2.9 +  operate.offsetHeight + this.operateBtn.value.offsetHeight;
        const countHeight: number = this.fixedHeight.value;
        return countHeight;
    };

    #setScroll = (countHeight: number, tableMaxHeight: {value: number}) => {
        const params = {
            ...this.filters.reactiveArr.filterParams,
            ...this.search.reactiveArr.searchParams
        }
        if (Object.values(params).some((item: string) => item !== '')) {
            tableScroll(countHeight + this.searchTipHeight, tableMaxHeight);
            return;
        }
        tableScroll(countHeight, tableMaxHeight);
    };

    getBtnRef = (val: {value: HeightType2}) => {
        this.operateBtn.value = val.value;
    };


    getHeaderRef = (val: {value: HeightType3;}) => {
        this.headerTitle.value = val.value;
    };

    getOperateRef = (val: {value: HeightType3;})=> {
        this.operate.value = val.value;
    };
};

export default TableScrollOperate;
