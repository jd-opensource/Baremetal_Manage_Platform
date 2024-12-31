import {ResetType} from '../typeManagement';

const resetOperate = (filter: ResetType['filter'], search: ResetType['search']) => {

    // 重置
    const reset = () => {
        filter.tableRef.value?.clearFilter();
        filter.reactiveArr.filterParams = {};
        search.hasClear.value = true;
        search.reactiveArr.searchParams = {};
        search.request();
    };

    return {
        reset
    };
};

export default resetOperate;
