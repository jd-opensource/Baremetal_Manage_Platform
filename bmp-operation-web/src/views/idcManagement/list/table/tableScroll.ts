/**
 * @file
 * @author
*/

import {HeightType3} from '@utils/publicType';
import {ScrollOperate} from 'utils/publicClass.ts';

interface DataType {
    tableMaxHeight: Ref<number>;
    headerTitle: Ref<HeightType3>;
    operate: Ref<HeightType3>;
};

const tableScrollOperate = () => {
    const data: DataType = {
        tableMaxHeight: ref<number>(0),
        headerTitle: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0}),
        operate: ref<HeightType3>({offsetHeight: 0, scrollHeight: 0})
    };

    onMounted(() => {
        const countHeight = data.headerTitle.value.offsetHeight + data.headerTitle.value.scrollHeight + data.operate.value.offsetHeight + data.operate.value.scrollHeight + 50;
        new ScrollOperate(countHeight, data.tableMaxHeight);
    });

    onUnmounted(() => window.onresize = null);

    const getHeaderRef = (val: {value: HeightType3;}) => {
        data.headerTitle.value = val.value;
    };

    const getOperateRef = (val: {value: HeightType3;})=> {
        data.operate.value = val.value;
    };

    return {
        ...data,
        getHeaderRef,
        getOperateRef
    };
};

export default tableScrollOperate;
