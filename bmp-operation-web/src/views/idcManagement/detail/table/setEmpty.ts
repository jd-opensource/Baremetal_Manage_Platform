/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p = document.createElement('p');
const span = document.createElement('span');

class SetEmpty {
    emitValue: any;

    constructor(props: any, emitValue: any) {
        this.emitValue = emitValue;
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: any) => {
        watch(() => props?.idcDetail?.reactiveArr?.tableData, (newValue) => {
            if (!newValue?.length) {
                nextTick(() => {
                    this.textEmpty();
                })
                return;
            }
            p.innerHTML= '';
        })
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        const customText = document.querySelector('.idc-detail .el-scrollbar__view');
        if (!customText) return;
        p.className = 'custom-tip-count';
        customText.appendChild(p)
        ai!.innerHTML = '';
        p.appendChild(span)
        span.innerHTML = `${language.t('table.empty')}`;
    };
};

export default SetEmpty;
