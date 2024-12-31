import {language} from 'utils/publicClass';

const p = document.createElement('p');

class SetEmpty {
    constructor(props: any) {
        this.watchTableData(props);
        onUnmounted(() => p.innerHTML = '');
    };

    watchTableData = (props: any) => {
        watch(() => props?.faultLog?.reactiveArr.tableData, (newValue) => {
            if (!newValue?.length) {
                nextTick(() => {
                    this.textEmpty();
                })
                return;
            }
            p.innerHTML = '';
        }, {deep: true})
    };

    textEmpty = () => {
        const ai = document.querySelector('.el-table__empty-text');
        ai!.innerHTML = '';
        const customText: HTMLElement | null = document.querySelector('.fault-log-list .el-scrollbar__view');
        p.className = 'custom-tip-count';
        const str2 = `
            ${language.t('table.empty')}
        `;
        this.otherEmpty(str2, customText!);
    };

    otherEmpty = (str: string, customText: HTMLElement) => {
        const parser = new DOMParser();
        const doc = parser.parseFromString(str, 'text/html');
        p.innerHTML = doc.body.innerHTML;
        customText?.appendChild(p);
    }
};

export default SetEmpty;
