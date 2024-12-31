const resetOperate = (imageList: any, filterOperate: any) => {

    const reset = () => {
        imageList.tableRef?.value?.clearFilter();
        filterOperate.reactiveArr.filterParams = {source: '', osType: ''};
        imageList.init();
    };

    return {
        reset
    }
};

export default resetOperate;
