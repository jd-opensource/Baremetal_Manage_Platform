import {HeightType3} from '@utils/publicType';
import {tableScroll} from 'utils/index.ts';

const tableScrollOperate = () => {
    const data = {
        tableMaxHeight: ref<number>(0),
        headerTitle: ref<HeightType3>({
            offsetHeight: 0,
            scrollHeight: 0
        })
    };

    onMounted(() => {
        getHeightTop();
    });

    onUnmounted(() => {
        window.onresize = null;
        document.onkeyup = () => {return;}
    });

    window.onresize = () => {
        getHeightTop();
    };

    const getHeaderRef = (val: Ref<HeightType3>) => {
        data.headerTitle = val;
    };

    const getHeightTop = () => {
        const createKeyBtn: HTMLElement | null = document!.querySelector('.create-key');
        nextTick(() => {
            const countHeight = createKeyBtn!.offsetHeight + data.headerTitle.value.scrollHeight + data.headerTitle.value.offsetHeight + 120;
            tableScroll(countHeight, data.tableMaxHeight);
        })
    };

    return {
        ...data,
        getHeaderRef
    }
};

export default tableScrollOperate;
