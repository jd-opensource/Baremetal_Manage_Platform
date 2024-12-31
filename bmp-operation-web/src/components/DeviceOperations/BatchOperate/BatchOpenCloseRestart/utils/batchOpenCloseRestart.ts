import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {SuccessType, CurrencyType} from '@utils/publicType';
import {devicesBatchStopAPI, devicesBatchStartAPI, devicesBatchRestartAPI} from 'api/device/request.ts';

class BatchOpenCloseRestart {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    props: {
        types: string;
        selectArr: CurrencyType[];
    };
    emitValue: any;

    constructor (props: {types: string; selectArr: CurrencyType[];}, emitValue: any) {
        this.props = props;
        this.emitValue = emitValue;
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestOpenDownRestart();
    };

    requestOpenDownRestart = () => {
        const requestApi = new Map([
            ['open', devicesBatchStartAPI],
            ['close', devicesBatchStopAPI],
            ['restart', devicesBatchRestartAPI]
        ]);
        return this.getDevicesOperate(requestApi.get(this.props.types));
    };

    getDevicesOperate = (requestUrl: (arg0: {instanceIds: string[]}) => Promise<{data: SuccessType}>) => {
        requestUrl(
            {
                instanceIds: this.props.selectArr.map((item: CurrencyType) => item.instanceId)
            }
        )
        .then(({data} : {data: SuccessType}) => {
            data?.result?.success && msgTip.success(language.t('operate.success'));
        })
        .finally(() => {
            this.isLoading.value = false;
            this.cancelClick();
            this.emitValue('determine-click');
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        this.emitValue('dia-log-close', false);
    };
};

export default BatchOpenCloseRestart;
