/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p = document.createElement('p');

interface PropsType {
    tableData: []
}

class SetEmpty {
    emitValue: Function;

    constructor(props: PropsType, emitValue: Function) {
        this.emitValue = emitValue;
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: PropsType) => {
        watch(() => props?.tableData, (newValue) => {
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
        ai!.innerHTML = '';
        const customText = document.querySelector('.image-detail .el-scrollbar__view');
        if (!customText) return;
        const str = `
            <span class="empty-style">
                ${language.t('table.empty')}，${language.t('table.tip')}
            </span>
            <span class="empty-text">
                ${language.t('imageDetail.empty.title')}
            </span>
        `
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.className = 'custom-tip-count';
        p.innerHTML = doc.body.innerHTML;
        customText.appendChild(p)
        const clickOpt = document.querySelector('.image-detail-model .custom-tip-count .empty-text');
        clickOpt!.addEventListener('click', () => {
            this.emitValue('empty-click')
        });
    };
};

export default SetEmpty;
