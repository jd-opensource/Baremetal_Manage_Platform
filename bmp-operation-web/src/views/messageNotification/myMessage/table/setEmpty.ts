/**
 * @file
 * @author
*/

import {language, CurrentInstance} from 'utils/publicClass.ts';

const p = document.createElement('p');
const span = document.createElement('span');

class SetEmpty {
    mitt = new CurrentInstance();

    constructor() {
        this.mitt.instanceMitt?.proxy?.$Bus?.on('message-list', this.watchTableData);
        onUnmounted(() => {
            p.innerHTML = '';
            this.mitt.instanceMitt?.proxy?.$Bus?.off('message-list', this.watchTableData);
        });
    };

    watchTableData = (val: any) => {
        if (!val?.length) {
            nextTick(() => {
                this.textEmpty();
            })
            return;
        }
        p.remove();
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        const customText = document.querySelector('.my-message-list .el-scrollbar__view');
        if (!customText) return;
        ai!.innerHTML = '';
        p.className = 'custom-tip-count';
        customText.appendChild(p)
        p.appendChild(span)
        span.innerHTML = `${language.t('table.empty')}`;
    };
};

export default SetEmpty;

