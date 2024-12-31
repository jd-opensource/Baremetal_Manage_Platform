import {msgTip} from 'utils/index.ts';
import {addImageAPI} from 'api/model/request.ts'; 
import {language} from 'utils/publicClass.ts';

const addImageOperate = (imageList: any, props: any, emitValue: any) => {
    // loading加载态
    const isLoading: Ref<boolean> = ref<boolean>(false);

    // 确定按钮弹窗
    const determineClick = (): void => {
        isLoading.value = true;
        requestAddImage();
    };

    /**
     * 请求镜像验证接口，成功后把事件回传，关闭弹窗
    */
    const requestAddImage = (): void => {
        const paramsData = imageList.reactiveArr.arr.filter((item: {isBind: boolean;}) => !item.isBind);
        addImageAPI(
            {
                deviceTypeId: props.deviceTypeId,
                imageIds: paramsData.map((item: {imageId: string;}) => item.imageId).join(',')
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
        cancelClick,
        determineClick
    };
};

export default addImageOperate;
