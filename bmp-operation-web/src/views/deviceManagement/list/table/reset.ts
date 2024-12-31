type FilterType = {
    tableRef?: {
        value?: {
            clearFilter(): void;   
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
    request(): void;
};

interface ParamsType {
    filter: FilterType;
    search: SearchType;
};

const resetOperate = (filter: ParamsType['filter'], search: ParamsType['search']) => {

    const reset = () => {
        sessionStorage?.removeItem('deviceTypeId');
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
