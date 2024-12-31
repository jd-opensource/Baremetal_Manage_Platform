import {HeightType2, HeightType3} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import store from 'store/index.ts';

interface TableScrollDataType {
    scrollTimer: null | number;
    searchTipHeight: number;
    fixedHeight: Ref<number>;
    tableMaxHeight: Ref<number>;
    status: boolean;
    headerTitle: Ref<HeightType3>;
    operate: Ref<HeightType3>;
    operateBtn: Ref<HeightType2>;
};

const tableScrollOperate = (filters: any['filters'], search: any['search'], searchTip: {value: boolean;}) => {
    const data: TableScrollDataType = {
        scrollTimer: null,
        searchTipHeight: 13,
        fixedHeight: ref<number>(0),
        tableMaxHeight: ref<number>(0),
        status: false,
        headerTitle: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
        operate: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
        operateBtn: ref<HeightType2>({offsetHeight: 0})
    };
    onMounted(() => {
        setScroll(setHeight(), data.tableMaxHeight);
        onResize(data.tableMaxHeight);
    });

    onUnmounted(() => {
        clearTimeout(data.scrollTimer as number)
    })

    watch(() => [
        filters.filterParams,
        search.searchParams,
        searchTip
    ], (newValue) => {
        if (newValue[2].value) {
            setScroll(data.fixedHeight.value, data.tableMaxHeight);
        }
    }, {deep: true});

    watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
        data.scrollTimer = setTimeout(() => {
            setScroll(setHeight(), data.tableMaxHeight);
        }, 200)
    });

    const onResize = (tableMaxHeight: {value: number}) => {
        window.onresize = () => {
            setScroll(setHeight(), tableMaxHeight);
        };
    };

    const setHeight = () => {
        const header = data.headerTitle.value;
        const operate = data.operate.value;
        data.fixedHeight.value = header.offsetHeight * 2.9 +  operate.offsetHeight + data.operateBtn.value.offsetHeight;
        const countHeight: number = data.fixedHeight.value;
        return countHeight;
    };

    const setScroll = (countHeight: number, tableMaxHeight: {value: number}) => {
        const params = {...filters.filterParams, ...search.searchParams};
        if (Object.values(params).some((item) => item !== '')) {
            tableScroll(countHeight + data.searchTipHeight, tableMaxHeight);
            return;
        }
        tableScroll(countHeight, tableMaxHeight);
    };

    const getBtnRef = (val: {value: HeightType2}) => {
        data.operateBtn.value = val.value;
    };

    const getHeaderRef = (val: {value: HeightType3;}) => {
        data.headerTitle.value = val.value;
    };

    const getOperateRef = (val: {value: HeightType3;})=> {
        data.operate.value = val.value;
    };

    return {
        ...data,
        getBtnRef,
        getHeaderRef,
        getOperateRef
    }
};

export default tableScrollOperate;
