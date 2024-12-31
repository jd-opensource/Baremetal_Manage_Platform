

type FilterType = {
    tableRef: {
        value: {
            clearFilter(): any;   
        };
    };
    reactiveArr: {
        filterParams: {};
    };
};

type SearchType = {
    hasClear: {
        value: boolean;
    };
    reactiveArr: {
        searchParams: {};
    };
    request(): Function;
};

interface ParamsType {
    filter: FilterType;
    search: SearchType;
};

const resetOperate = (filter: ParamsType['filter'], search: ParamsType['search']) => {

    const reset = () => {
        filter?.tableRef?.value?.clearFilter();
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
