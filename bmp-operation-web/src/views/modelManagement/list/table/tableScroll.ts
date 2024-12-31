import {HeightType3, HeightType2} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import {TableScrollType} from '../typeManagement';
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

const tableScrollOperate = (filters: TableScrollType['filters'], search: TableScrollType['search'], searchTip: {value: boolean}) => {
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
        searchTip,
    ], (newValue: any) => {
        if (newValue[2].value) {
            setScroll(setHeight(), data.tableMaxHeight);
        }
    }, {deep: true})

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
        const deviceSeries: string = filters.filterParams?.deviceSeries??'';
        const name: string = search.searchParams?.name??'';
        const deviceType: string = search.searchParams?.deviceType??'';
        if ([deviceSeries, name, deviceType].some((item: string) => item !== '')) {
            tableScroll(countHeight + data.searchTipHeight, tableMaxHeight);
            return;
        }
        tableScroll(countHeight, tableMaxHeight);
    };

    const getBtnRef = (val: {value: HeightType2}) => {
        data.operateBtn.value = val.value;
    };

    const getHeaderRef = (val: {value:HeightType3}) => {
        data.headerTitle.value = val.value;
    };

    const getOperateRef = (val: {value: HeightType3})=> {
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
