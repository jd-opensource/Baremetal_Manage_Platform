/**
 * @file
 * @author
*/

import {paginationOperate} from 'utils/publicClass.ts';

/**
 * 分页监听
*/
class PaginationOperate {

    // 构造器
    constructor (faultLogTable: any) {
        this.watchSizeNumber(faultLogTable);
    };

    /**
     * 监听分页页数、条数变化
    */
    watchSizeNumber = (faultLogTable: any) => {
        
        watch(() => [
            paginationOperate.pageSize.value,
            paginationOperate.pageNumber.value
        ], (_: boolean[]): void => {
            if (paginationOperate.routerChange.value) return;
            faultLogTable.refresh();
        });
    };
};

export default PaginationOperate;