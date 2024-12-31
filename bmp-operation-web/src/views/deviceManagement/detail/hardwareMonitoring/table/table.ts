import {consolelogsAPI} from 'api/device/request.ts';
import {paginationOperate, language} from 'utils/publicClass.ts';
import {msgTip, methodsTotal} from 'utils/index.ts';
import DeviceStaticData from 'staticData/device/index.ts';

class FaultLogTable {
    isTableLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr: any = reactive({
        tableData: []
    });
    errorCode: Ref<number> = ref<number>(0);
    query: {sn: string;};

    constructor (query: {sn: string;}) {
        this.query = query;
        // this.getTableList();
    };

    refresh = () => {
        this.getTableList();
    };

    getTableList = () => {
        this.isTableLoading.value = true;
        const params = methodsTotal.toLine({
            pageNum: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value
        });
        consolelogsAPI({
            ...params,
            sn: this.query.sn
        }).then(({data}: any) => {
            if (data?.result?.detail?.length) {
                const newResult = methodsTotal.lineConverting(data.result);
                newResult.detail.forEach((item: any, index: number) => {
                    DeviceStaticData.deviceFaultLogTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                })
                paginationOperate.total.value = newResult.totalCount;
                this.reactiveArr.tableData = newResult.detail;
                return;
            }
            return Promise.reject('');
        })
        .catch(({message}: {message: string}) => {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            if (['module license disabled'].includes(message)) {
                return;
            }
            message && msgTip.error(message);
        })
        .finally(() => this.isTableLoading.value = false);
    };

    setStatus = (status: number | string) => {
        const val: string = String(status);
        const newStatus = new Map([
            ['0', language.t('surveillance.status.no')]
        ]);
        return newStatus.get(val)?? language.t('surveillance.status.yes')
    };

    setColor = (status: number | string) => {
        const val: string = String(status);
        const newColor = new Map([
            ['0', '#ff4d4f']
        ]);
        return newColor.get(val)?? '#43b561';
    };
};

export default FaultLogTable;