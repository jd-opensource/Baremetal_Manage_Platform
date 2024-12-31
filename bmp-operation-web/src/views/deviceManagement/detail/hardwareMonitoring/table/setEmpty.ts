/**
 * @file
 * @author
*/

import {language} from 'utils/publicClass.ts';

const p: HTMLElement = document.createElement('p');
interface PropsType {
    faultLogTable: {
        reactiveArr: {
            tableData: []
        }
    }
}

class SetEmpty {

    constructor(props: PropsType) {
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: PropsType) => {
        watch(() => props?.faultLogTable?.reactiveArr.tableData, (newValue) => {
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
        const customText = document.querySelector('.hardware-monitoring .el-scrollbar__view');
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
