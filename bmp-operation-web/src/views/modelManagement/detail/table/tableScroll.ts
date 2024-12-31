import {tableScroll} from 'utils/index.ts';
import {CurrencyType, HeightType3} from '@utils/publicType';

interface DataType {
    activeName: Ref<string>;
    headerTitle: Ref<HeightType3>;
    hasRemove: boolean;
    tableMaxHeight: Ref<number>;
    title2: Ref<HeightType3>;
};

interface ParamsType {
    active: {
        value: string;
    };
    filter: CurrencyType;
    searchTip: {
        value: boolean;
    }
}

const tableScrollOperate = (params: ParamsType) => {
    const {active, filter, searchTip}: ParamsType = params;
    const data: DataType = {
        activeName: ref<string>(''),
        hasRemove: Boolean(false),
        tableMaxHeight: ref<number>(0),
        title2: ref<HeightType3>({
            offsetHeight: 0,
            scrollHeight: 0
        }),
        headerTitle: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
    };

    data.activeName.value = active.value;

    watch(() => [filter, searchTip.value], () => { 
        getHeightTop();
    }, {deep: true})

    watch(() => active, (newValue: {value: string}) => {
        data.activeName.value = newValue.value;
        getHeightTop();
    }, {deep: true})

    onMounted(() => {
        if (data.activeName.value === 'essentialInformation') return;
        data.hasRemove = false;
        getHeightTop();
    });

    onUnmounted(() => data.hasRemove = true);

    window.onresize = () => {
        if (data.activeName.value === 'essentialInformation') return;
        getHeightTop();
    };

    const getRef = (val: DataType['title2']) => {
        data.title2 = val;
        getHeightTop();
    };

    const getHeightTop = () => {
        if (data.hasRemove) return;
        nextTick(() => {
            const addImageButton: any = document.querySelector('.add-image-btn');
            const count = searchTip.value ? 27 : 0;
            const value1 = data.headerTitle.value;
            const value2 = data.title2.value;
            const countHeight = count + addImageButton?.offsetHeight + addImageButton?.scrollHeight + value1.scrollHeight + value2.offsetHeight + value2.scrollHeight + value1.offsetHeight + 23;
            tableScroll(countHeight, data.tableMaxHeight);
        })
    };


    const getRefHeader = (val: { value: HeightType3; }) => {
        data.headerTitle.value = val.value;
    };


    return {
        ...data,
        getRefHeader,
        getRef
    }
};

export default tableScrollOperate;
