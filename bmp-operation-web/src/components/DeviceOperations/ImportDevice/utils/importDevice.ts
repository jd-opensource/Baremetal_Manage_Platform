import {msgTip} from 'utils/index.ts';
import {importDevicesAPI} from 'api/device/request.ts';
import {language} from 'utils/publicClass.ts';
import router from 'router/index.ts';
// import {removeLong} from 'request/index.ts';
import {removeLong} from 'request/RemoveLongReq/remove.ts';

const importDevice = (formRules: any, upLoadOperate: any, heightStatus: any, emitValue: any) => {
    // loading 加载态
    const isLoading: Ref<boolean> = ref<boolean>(false);

    onMounted(() => {
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && determineClick();
        };
    });

    /**
     * 确定按钮弹窗
    */
    const determineClick = async (): Promise<void> => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            formRules.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    if (formRules.ruleForm.fileParams?.result?.length) {
                        isLoading.value = true;
                        upLoadOperate.hasError.value = false;
                        requestImportDevice();
                    }
                }
                else if (!formRules.ruleForm.deviceInfo.length) {
                    upLoadOperate.hasError.value = true;
                    heightStatus.hasClick.value = true;
                }
                else {
                    if (!valid) return;
                    msgTip.warning(language.t('importDevice.errTip.upload'));
                    upLoadOperate.hasError.value = false;
                    heightStatus.hasClick.value = true;
                }
            });
        });
    };

    /**
     * 请求导入设备接口，成功后把事件回传，关闭弹窗
    */
    const requestImportDevice = (): void => {
        importDevicesAPI(
            {
                idcId: formRules.ruleForm.idcName,
                deviceTypeId: formRules.ruleForm.modelName,
                devices: formRules.ruleForm.fileParams.result
            }
        )
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                emitValue('determine-click');
                cancelClick();
                return;
            }
        })
        .finally(() => {
            isLoading.value = false;
        });
    };

    /**
     * 取消按钮点击操作
    */
    const cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
        // 清空输入内容、异常提示状态
        formRules.ruleFormRef.value!.resetFields();
        upLoadOperate.uploads.value!.abort();
        upLoadOperate.uploads.value!.clearFiles();
        if (formRules.ruleForm.fileParams?.result?.length) {
            removeLong(router.currentRoute.value.name);
            // removeLong(router.currentRoute.value.name);
        }
    };

    return {
        isLoading,
        determineClick,
        cancelClick
    };
};

export default importDevice;
