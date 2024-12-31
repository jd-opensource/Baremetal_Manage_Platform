import {HeightType2} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import store from 'store/index.ts';
import {CurrentInstance} from 'utils/publicClass.ts';

class TableScrollOperate {
    scrollTimer: null | number = null;
    fixedHeight: Ref<number> = ref<number>(0);
    tableMaxHeight: Ref<number> = ref<number>(0);
    headerTitle: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    operate: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    searchHeight: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});
    mitt = new CurrentInstance();
    tabsRef: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 8});
    constructor () {
        // this.activeOperate = activeOperate;
        onMounted(() => {
            // this.activeName.value = this.activeOperate.activeName.value;
            // if (['basicInfo', 'diskDetail', 'operatLog'].includes(this.activeName.value)) return;
            this.#watchNavBar();
        })
        // const type = this.router.currentRoute.value.query.type;

        // if (type === 'hardwareMonitoring') {
        //     this.#watchActive(this.router.currentRoute.value.query);
        // }
        // this.#watchActive();
        this.#resize();
        this.mitt.instanceMitt?.proxy?.$Bus?.on('device-detail-table-scroll', this.#resize);
        this.mitt.instanceMitt?.proxy?.$Bus?.on('device-detail-hardwareMonitoring-loading', this.watchLoading);
        onUnmounted(() => {
            clearTimeout(this.scrollTimer as number);
            this.mitt.instanceMitt?.proxy?.$Bus?.off('device-detail-table-scroll', this.#resize);
            this.mitt.instanceMitt?.proxy?.$Bus?.on('device-detail-hardwareMonitoring-loading', this.watchLoading);
        })
    };

    getOperateRef = (val: {value: HeightType2;})=> {
        this.operate.value = val.value;
    };

    getRef = (val: {value: HeightType2}) => {
        this.tabsRef.value = val.value;
    };

    getSearchRef = (val: {value: HeightType2}) => {
        this.searchHeight.value = val.value;
    };

    #watchNavBar = () => {
        watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
           this.scrollTimer = setTimeout(() => {
                this.#setHeight()
           }, 300)
        });
    };

    watchLoading = (newValue: {value: boolean}) => {
        if (!newValue.value) {
            this.#setHeight();
        }
    }
    // #watchActive = (query: {type: string}) => {
    //     watch(() => query, (newValue) => {
    //         if (newValue.type === 'hardwareMonitoring') {
    //             this.#setHeight();
    //         }
    //         // this.activeName.value = newValue.value;
    //     }, {deep: true});
    // };

    #resize = () => {
        window.onresize = () => {
            // if (['basicInfo', 'diskDetail', 'operatLog'].includes(this.activeName.value)) return;
            this.#setHeight();
        };
    };

    #setHeight = () => {
        nextTick(() => {
            this.fixedHeight.value = this.headerTitle.value.offsetHeight * 4.9 + this.operate.value.offsetHeight + this.tabsRef.value.offsetHeight
            const countHeight = this.searchHeight?.value?.offsetHeight + this.fixedHeight.value;
            tableScroll(countHeight, this.tableMaxHeight)
        })
    };

    getHeaderRef = (val: {value: HeightType2;}) => {
        this.headerTitle.value = val.value;
    };
}

export default TableScrollOperate;