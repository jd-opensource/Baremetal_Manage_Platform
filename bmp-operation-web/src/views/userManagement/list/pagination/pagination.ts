import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 分页监听
*/
class PaginationWatch {
    // 构造器
    constructor (props: any) {
        this.watchSizeNumber(props);
    };

    /**
     * 监听分页页数、条数变化
    */
    watchSizeNumber = (props: any) => {
        watch(() => [
            paginationOperate.pageSize.value,
            paginationOperate.pageNumber.value
        ], (_: boolean[]): void => {
            if (paginationOperate.routerChange.value) return;
            if (props.search.reactiveArr.searchParams?.userName) {
                props.userList.requestOperate(props.search.reactiveArr.searchParams, 'search');
                return;
            }
            props.userList.requestOperate(props.filter.reactiveArr.filterParams, 'filter');
        });
    };
};

export default PaginationWatch;
