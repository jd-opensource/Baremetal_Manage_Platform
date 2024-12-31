import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {TableDataType, ParamsType} from './type';
import {SuccessType} from '@utils/publicType';
import {devicesRemoveAPI, instanceRecoveryAPI} from 'api/device/request.ts';

class RemoveRecovery {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 使用mitt传值
    instanceMitt: any = getCurrentInstance();
    // 表格数据
    tableData: TableDataType[] = reactive<TableDataType[]>([]);
    checkboxStatus: Ref<boolean> = ref<boolean>(false);
    props: {title: string} = {title: ''};
    emitValue: any;

    constructor(props: {title: string}, emitValue: any) {
        this.props = props;
        this.emitValue = emitValue;
        // 监听mitt-remove-recovery-table，获取表格数据
        this.instanceMitt?.proxy?.$Bus?.on('device-table', this.getTableFn);
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.#getRequestApi();
    };

    #getRequestApi = () => {
        const requestData = new Map([
            ['remove', devicesRemoveAPI],
            ['recovery', instanceRecoveryAPI]
        ]);
        return this.#requestRemoveRecovery(requestData.get(this.props.title), this.props.title)
    };

    /**
     * @param {Function} requestUrl url请求地址
    */
    #requestRemoveRecovery = (requestUrl: (arg0: ParamsType) => Promise<{data: SuccessType}>, type: string) => {
        requestUrl(
            {
                ...this.#setParams(type)
            }
        )
        .then(({data} : {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        }).finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            this.emitValue('determine-click', type);
        });
    };

    #setParams = (type: string) => {
        const params = new Map([
            ['remove', {deviceId: this.tableData[0].deviceId}],
            ['recovery', {instanceId: this.tableData[0].instanceId}]
        ]);
        return params.get(type);
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
        this.instanceMitt.proxy.$Bus.off('device-table', this.getTableFn);
    };

    /**
     * 获取表格数据函数
     * @param {Object} item 表格数据
    */
    getTableFn: unknown = ({list}: {list: TableDataType;}) => {
        this.tableData.push(list);
    };
};

export default RemoveRecovery;
