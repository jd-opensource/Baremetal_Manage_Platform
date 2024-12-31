// 重置操作
const resetOperate = (filter: any, tableDetail: any) => {
    const reset = () => {
        filter.tableRef.value?.clearFilter();
        filter.reactiveArr.filterParams = {};
        tableDetail.init();
    };

    return {
        reset
    }
};

export default resetOperate;
