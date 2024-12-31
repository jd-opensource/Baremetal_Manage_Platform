/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p: HTMLElement = document.createElement('p');
interface PropsType {
    deviceList: {
        reactiveArr: {
            tableData: []
        }
    }
}

class SetEmpty {
    emitValue: Function;

    constructor(props: PropsType, emitValue: Function) {
        this.emitValue = emitValue;
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: PropsType) => {
        watch(() => props?.deviceList?.reactiveArr.tableData, (newValue) => {
            if (!newValue?.length) {
                nextTick(() => {
                    this.textEmpty();
                })
                return;
            }
            p.remove();
        }, {deep: true})
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        ai!.innerHTML = '';
        const customText = document.querySelector('.device-list .el-scrollbar__view');
        if (!customText) return;
        const str = `
            <span class="empty-style">
                ${language.t('table.empty')}，${language.t('table.tip')}
            </span>
            <span class="empty-text">
                ${language.t('deviceList.operate.exportDevice')}
            </span>
        `
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.className = 'custom-tip-count';
        customText.appendChild(p)
        p.innerHTML = doc.body.innerHTML;
        const clickOpt = document.querySelector('.device-list .custom-tip-count .empty-text');
        clickOpt!.addEventListener('click', () => {
           this.emitValue('empty-click')
        })
    };
};

export default SetEmpty;
