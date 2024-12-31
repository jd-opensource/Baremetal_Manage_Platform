/**
 * @file
 * @author
*/

import {language, CurrentInstance} from 'utils/publicClass.ts';

const p = document.createElement('p');

class SetEmpty {
    createKey1: {createKeyClick(): void};
    mitt = new CurrentInstance();

    constructor(createKey1: {createKeyClick(): void}) {
        this.createKey1 = createKey1;
        // this.watchTableData();
        this.mitt.instanceMitt?.proxy?.$Bus?.on('api-key-data', this.watchTableData);
        onUnmounted(() =>{
            p.innerHTML = '';
            this.mitt.instanceMitt?.proxy?.$Bus?.off('api-key-data', this.watchTableData);
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
        ai!.innerHTML = '';
        const customText = document.querySelector('.user-centet-table .el-scrollbar__view');
        if (!customText) return;
        const str = `
            <span class="empty-style">
                ${language.t('table.empty')}，${language.t('table.tip')}
            </span>
            <span class="empty-text">
                ${language.t('userCenter.btn.createKey')}
            </span>
        `
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.className = 'custom-tip-count';
        p.innerHTML = doc.body.innerHTML;
        customText.appendChild(p)
        const clickOpt = document.querySelector('.user-centet-table .custom-tip-count .empty-text');
        clickOpt!.addEventListener('click', () => {
            this.createKey1.createKeyClick();
        });
    };
};

export default SetEmpty;
