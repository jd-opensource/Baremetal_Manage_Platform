import {msgTip} from 'utils/index.ts'; // 深拷贝数据
import {imageAddModelAPI} from 'api/image/request.ts'; // 添加机型接口
import {language} from 'utils/publicClass.ts';

const addModelOperate = (imageDeviceTypes: any, props: any, emitValue: any) => {
    // loading加载态
    const isLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 请求机型验证接口，成功后把事件回传，关闭弹窗
    */
    const determineClick = (): void => {
        const paramsData = imageDeviceTypes.reactiveArr.arr.filter((item: {isBind: boolean;}) => !item.isBind);
        isLoading.value = true;
        imageAddModelAPI(
            {
                imageId: props.imageId,
                deviceTypeIds: paramsData.map((item: {deviceTypeId: string;}) => item.deviceTypeId).join(',')
            }
        )
        .then(({data} : {data: {result: {success: boolean;}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                emitValue('determine-click');
                cancelClick();
            }
        })
        .finally(() => {
            isLoading.value = false;
        })
        .catch(() => {
            emitValue('determine-click');
            cancelClick();
        })
    };

    /**
     * 取消按钮点击操作
    */
    const cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

    return {
        isLoading,
        determineClick,
        cancelClick
    };
};

export default addModelOperate;
