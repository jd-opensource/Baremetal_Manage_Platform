/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p = document.createElement('p');
const span = document.createElement('span');

class SetEmpty {
    constructor(props: any) {
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: any) => {
        watch(() => props?.roleList?.reactiveArr.tableData, (newValue) => {
            if (!newValue?.length) {
                nextTick(() => {
                    this.textEmpty();
                })
                return;
            }
            p.innerHTML= '';
        }, {deep: true})
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        const customText = document.querySelector('.role-list .el-scrollbar__view');
        if (!customText) return;
        ai!.innerHTML = '';
        p.className = 'custom-tip-count';
        customText.appendChild(p)
        p.appendChild(span)
        span.innerHTML = `${language.t('table.empty')}`;
    };
};

export default SetEmpty;
