import {language, CurrentInstance} from 'utils/publicClass';

const p = document.createElement('p');

class SetEmpty {
    mitt: any = new CurrentInstance();
    props;

    constructor(props: any) {
        this.props = props;
        this.watchTableData(props);
        this.mitt.instanceMitt?.proxy?.$Bus?.on('all-alarm-rule-list', this.watchTableData);
        onUnmounted(() =>{
            p.innerHTML = '';
            this.mitt.instanceMitt?.proxy?.$Bus?.off('all-alarm-rule-list', this.watchTableData);
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
        if (ai === null) return;
        ai!.innerHTML = '';
        const customText: HTMLElement | null = document.querySelector('.all-alarm-rule-list .el-scrollbar__view');
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
