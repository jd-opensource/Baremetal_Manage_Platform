import {HeightType, HeightType2} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import {TimerOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class TableScrollOperate extends TimerOperate {
    fixedHeight: Ref<number> = ref<number>(0);
    tableMaxHeight: Ref<number> = ref<number>(0);
    headerTitle: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    operate: Ref<HeightType> = ref<HeightType>({offsetHeight: 0, offsetTop: 0});
    searchHeight: any;
    constructor () {
        super();
        onMounted(async () => {
            tableScroll(this.#setHeight(), this.tableMaxHeight);
            this.#resize();
           await this.#watchNavBar();
        });

        onUnmounted(() => clearTimeout(this.timer));
    };

    searchRef = (val: {value: {offsetHeight: number}}) => {
        this.searchHeight = val;
    };

    #watchNavBar = () => {
        watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
            this.timer = setTimeout(() => {
                tableScroll(this.#setHeight(), this.tableMaxHeight);
            }, 200)
        });
    };

    #resize = () => {
        window.onresize = () => {
            tableScroll(this.#setHeight(), this.tableMaxHeight);
        };
    };

    #setHeight = () => {
        this.fixedHeight.value =  (this.headerTitle.value.offsetHeight * 3) + this.operate.value.offsetHeight + 6;
        const countHeight = this.searchHeight?.value?.offsetHeight + this.fixedHeight.value;
        return countHeight;
    };

    getHeaderRef = (val: {value: HeightType2;}) => {
        this.headerTitle.value = val.value;
    };

    getOperateRef = (val: {value: HeightType;})=> {
        this.operate.value = val.value;
    };
}

export default TableScrollOperate;