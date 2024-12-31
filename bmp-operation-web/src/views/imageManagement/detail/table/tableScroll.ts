import {HeightType3} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';

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
}

const tableScrollOperate = (active: ParamsType['active']) => {
    const data: DataType = {
        activeName: ref<string>(''),
        hasRemove: Boolean(false),
        headerTitle: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
        tableMaxHeight: ref<number>(0),
        title2: ref<HeightType3>({
            offsetHeight: 0,
            scrollHeight: 0
        })
    };

    data.activeName.value = active.value;

    onMounted(() => {
        if (data.activeName.value === 'basicInfo') return;
        getHeightTop();
    });

    onUnmounted(() => data.activeName.value = '');

    window.onresize = () => {
        if (data.activeName.value === 'basicInfo') return;
        getHeightTop();
    };

    watch(() => active, (newValue: {value: string}) => {
        data.activeName.value = newValue.value;
        getHeightTop();
    }, {deep: true})

    const getRef = (val: DataType['title2']) => {
        data.title2 = val;
        getHeightTop();
    };


    const getRefHeader = (val: { value: HeightType3 }) => {
        data.headerTitle.value = val.value;
    };

 
    const getHeightTop = () => {
        if (data.activeName.value === 'basicInfo') return;
        nextTick(() => {
            const addImageButton: any = document.querySelector('.image-add-model');
            const value1 = data.headerTitle.value;
            const value2 = data.title2.value;
            const countHeight = addImageButton?.offsetHeight + addImageButton?.scrollHeight + value1.scrollHeight + value2.offsetHeight + value2.scrollHeight + value1.offsetHeight + 23;
            tableScroll(countHeight, data.tableMaxHeight);
        })
    };
    return {
        ...data,
        getHeightTop,
        getRefHeader,
        getRef
    }
};

export default tableScrollOperate;
