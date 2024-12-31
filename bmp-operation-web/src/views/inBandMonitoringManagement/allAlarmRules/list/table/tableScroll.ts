import {HeightType, HeightType2} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import {TimerOperate} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class TableScrollOperate extends TimerOperate {
    fixedHeight: Ref<number> = ref<number>(0);
    tableMaxHeight: Ref<number> = ref<number>(0);
    headerTitle: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    operate: Ref<HeightType> = ref<HeightType>({offsetHeight: 0, offsetTop: 0});
    searchHeight: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});;
    constructor () {
        super();
        onMounted(async () => {
            // this.fixedHeight.value =  this.headerTitle.value.offsetHeight + this.headerTitle.value.scrollHeight +  this.operate.value.offsetHeight
            // const countHeight = this.searchHeight?.value?.offsetHeight +  this.searchHeight?.value?.scrollHeight + this.fixedHeight.value;
            tableScroll(this.#setHeight(), this.tableMaxHeight);
            this.#resize();
           await this.#watchNavBar();
        });

        onUnmounted(() => clearTimeout(this.timer));
    };

    searchRef = (val: {value: HeightType2}) => {
        this.searchHeight.value = val.value;
    };

    #watchNavBar = () => {
        watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
            this.timer = setTimeout(() => {
                // const countHeight = this.searchHeight?.value?.offsetHeight + this.searchHeight?.value?.scrollHeight + this.fixedHeight.value;
                tableScroll(this.#setHeight(), this.tableMaxHeight);
            }, 200)
        });
    };

    #resize = () => {
        window.onresize = () => {
            // const countHeight = this.searchHeight?.value?.offsetHeight + this.searchHeight?.value?.scrollHeight + fixedHeight;
            tableScroll(this.#setHeight(), this.tableMaxHeight);
        };
    };

    #setHeight = () => {
        // this.headerTitle.value.scrollHeight 
        this.fixedHeight.value = this.headerTitle.value.offsetHeight * 3 +  this.operate.value.offsetHeight + 5;
        // this.searchHeight?.value?.scrollHeight
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
