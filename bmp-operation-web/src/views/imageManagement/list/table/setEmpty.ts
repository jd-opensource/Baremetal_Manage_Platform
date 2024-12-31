/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p = document.createElement('p');
class SetEmpty {
    emitValue: any;

    constructor(props: any, emitValue: any) {
        this.emitValue = emitValue;
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: any) => {
        watch(() => props?.imageList?.reactiveArr.tableData, (newValue) => {
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
        ai!.innerHTML = '';
        const customText = document.querySelector('.image-list .el-scrollbar__view');
        if (!customText) return;
        p.className = 'custom-tip-count';
        const str = `
            <span class="empty-style">
                ${language.t('table.empty')}，${language.t('table.tip')}
            </span>
            <span class="empty-text">
                ${language.t('imageList.header.exportImage')}
            </span>
        `
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.className = 'custom-tip-count';
        p.innerHTML = doc.body.innerHTML;
        customText.appendChild(p)
        const clickOpt = document.querySelector('.image-list .custom-tip-count .empty-text');
        clickOpt!.addEventListener('click', () => {
           this.emitValue('empty-click')
        })
    };
};

export default SetEmpty;
