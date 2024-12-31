/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p: HTMLElement = document.createElement('p');

class SetEmpty {
    operateLog;

    constructor(operateLog: any) {
        this.operateLog = operateLog;
        // operateLog: any
        this.watchTableData();
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = () => {
        watch(() => this.operateLog?.reactiveArr.tableData, (newValue) => {
            if (!newValue?.length) {
                nextTick(() => {
                    this.textEmpty();
                })
                return;
            }
            p.remove();
        })
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        ai!.innerHTML = '';
        const customText = document.querySelector('.operate-log-list .el-scrollbar__view');
        if (!customText) return;
        const str = `
            <span class="empty-style">
                ${language.t('table.empty')}
            </span>
        `
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.className = 'custom-tip-count';
        customText.appendChild(p)
        p.innerHTML = doc.body.innerHTML;
    };
};

export default SetEmpty;
