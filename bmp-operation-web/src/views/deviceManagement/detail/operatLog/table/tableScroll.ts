/**
 * @file
 * @author
*/

import {HeightType, HeightType2} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import {TimerOperate, CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';


class TableScrollOperate extends TimerOperate {
    tabsHeight: number = 116;
    fixedHeight: Ref<number> = ref<number>(0);
    tableMaxHeight: Ref<number> = ref<number>(0);
    headerTitle: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 60});
    operate: Ref<HeightType> = ref<HeightType>({offsetHeight: 0, offsetTop: 0});
    searchHeight: any;
    infoHeaderHeight: any;
    mitt = new CurrentInstance();

    constructor () {
        super();
        onMounted(async () => {
            tableScroll(this.#setHeight(), this.tableMaxHeight);
            await this.#watchNavBar();
        });
        this.#resize();
        onUnmounted(() => {
            clearTimeout(this.timer);
        });
    };

    searchRef = (val: {value: {offsetHeight: number}}) => {
        this.searchHeight = val;
    };

    infoHeaderRef = (val: {value: {offsetHeight: number}}) => {
        this.infoHeaderHeight = val;
    };

    #watchNavBar = () => {
        watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
            this.timer = setTimeout(() => {
                tableScroll(this.#setHeight(), this.tableMaxHeight);
            }, 180)
        });
    };

    #resize = () => {
        window.onresize = () => {
            if (this.searchHeight?.value?.offsetHeight !== void 0) {
                tableScroll(this.#setHeight(), this.tableMaxHeight);
                return;
            }
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('device-detail-table-scroll', 'hardwareMonitoring');
        };
    };

    #setHeight = () => {
        this.fixedHeight.value =  (this.headerTitle?.value?.offsetHeight * 3) + this.operate?.value?.offsetHeight + this.tabsHeight;
        const countHeight = this.searchHeight?.value?.offsetHeight + this.fixedHeight.value + this.infoHeaderHeight.value.offsetHeight;
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