import {ScrollOperate} from 'utils/publicClass.ts';
import {HeightType3} from '@utils/publicType';

class TableScrollOperate {
    tableMaxHeight: Ref<number> = ref<number>(0);
    headerTitle: Ref<HeightType3> = ref<HeightType3>({offsetHeight: 0, scrollHeight: 0});
    operate: Ref<HeightType3> = ref<HeightType3>({offsetHeight: 0, scrollHeight: 0});

    constructor() {
        onMounted(() => {
            const countHeight = this.headerTitle.value.offsetHeight + this.headerTitle.value.scrollHeight + this.operate.value.offsetHeight + this.operate.value.scrollHeight + 50;
            new ScrollOperate(countHeight, this.tableMaxHeight);
        });
    };

    getHeaderRef = (val: {value: HeightType3;}) => {
        this.headerTitle.value = val.value;
    };

    getOperateRef = (val: {value: HeightType3;})=> {
        this.operate.value = val.value;
    };
};

export default TableScrollOperate;
