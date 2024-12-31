import {HeightType3} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';
import {CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';

interface TableScrollDataType {
    scrollTimer: null | number;
    searchTipHeight: number;
    fixedHeight: Ref<number>;
    tableMaxHeight: Ref<number>;
    status: boolean;
    headerTitle: Ref<HeightType3>;
    operate: Ref<HeightType3>;
};

const tableScrollOperate = (filters: any, search: any, searchTip: {value: boolean}, path: string) => {
    const data: TableScrollDataType = {
        scrollTimer: null,
        searchTipHeight: 13,
        fixedHeight: ref<number>(0),
        tableMaxHeight: ref<number>(0),
        status: false,
        headerTitle: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
        operate: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0})
    };

    const mitt = new CurrentInstance();
    const pathUrl: string = mitt.proxy.$defInfo.routerPath('myMessage');
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
        search.radioGroup,
        searchTip,
    ], (newValue: any) => {
        if (newValue[3].value) {
            if (path !== pathUrl) return;
            setScroll(setHeight(), data.tableMaxHeight);
        }
    }, {deep: true})

    watch(() => store.navigationBarStatus.hasNavigationBar, (_: boolean) => {
        if (path !== pathUrl) return;
        data.scrollTimer = setTimeout(() => {
            setScroll(setHeight(), data.tableMaxHeight);
        }, 200)
    });

    const onResize = (tableMaxHeight: {value: number}) => {
        if (path !== pathUrl) return;
        window.onresize = async () => {
            await mitt.instanceMitt?.proxy?.$Bus?.emit('custom-filter-onresize', 1);
            setScroll(setHeight(), tableMaxHeight);
        };
    };

    const setHeight = () => {
        const header = data.headerTitle.value;
        const operate = data.operate.value;
        data.fixedHeight.value = header.offsetHeight * 2.9 +  operate.offsetHeight + 35;
        const countHeight: number = data.fixedHeight.value;
        return countHeight;
    };

    const setScroll = (countHeight: number, tableMaxHeight: {value: number}) => {
        const messageType: string = filters.filterParams?.messageType??'';
        const messageSubType: string = filters.filterParams?.messageSubType??'';
        const content: string = search.searchParams?.content??'';
        const hasRead: string = search.radioGroup?.hasRead??'';
        if ([messageType, messageSubType, content, hasRead].some((item: string) => item !== '')) {
            tableScroll(countHeight + data.searchTipHeight, tableMaxHeight);
            return;
        }
        tableScroll(countHeight, tableMaxHeight);
    };

    const getHeaderRef = (val: {value:HeightType3}) => {
        data.headerTitle.value = val.value;
    };

    const getOperateRef = (val: {value: HeightType3})=> {
        data.operate.value = val.value;
    };

    return {
        ...data,
        // getBtnRef,
        getHeaderRef,
        getOperateRef
    }
};

export default tableScrollOperate;
