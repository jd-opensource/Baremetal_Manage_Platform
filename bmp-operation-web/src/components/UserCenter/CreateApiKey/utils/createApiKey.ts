import {createApiKeyAPI} from 'api/userCenter/request.ts';
import {language} from 'utils/publicClass.ts';
import {msgTip} from 'utils/index.ts';

const createApiKey = (formRules: any, emitValue: any) => {
    const isLoading: Ref<boolean> = ref<boolean>(false);

    onUnmounted(() => {
        document.onkeyup = () => {return;}
    });

    document.onkeyup = (event: {keyCode: number;}) => {
        event.keyCode === 13 && determineClick();
    };

    /**
     * 请求创建密钥接口，成功后把事件回传，关闭弹窗
    */
    const determineClick = async () => {
        isLoading.value = true;
        // 判断输入项是否符合条件
        await nextTick(() => {
            formRules.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    createApiKey();
                }
                else {
                    isLoading.value = false;
                }
            });
        });
    };

    const createApiKey = () => {
        createApiKeyAPI(
            {
                ...formRules.ruleForm
            }
        )
        .then(({data} : {data: {result: {apikeyId: string}}}) => {
            if (data?.result?.apikeyId?.length) {
                msgTip.success(language.t('operate.success'));
                cancelClick();
                emitValue('determine-click');
                return;
            }
        })
        .finally(() => {
            isLoading.value = false;
        })
        .catch(() => {
            cancelClick();
            emitValue('determine-click');
        });
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

export default createApiKey;

/**
 * 创建密钥
*/
// class CreateApiKey {
//     isLoading: Ref<boolean> = ref<boolean>(false)

//     constructor () {
//         onUnmounted(() => {
//             document.onkeyup = () => {return;}
//         })
//         document.onkeyup = (event: {keyCode: number;}) => {
//             event.keyCode === 13 && this.determineClick();
//         };
//     };

//     /**
//      * 请求创建密钥接口，成功后把事件回传，关闭弹窗
//     */
//     determineClick = async () => {
//         this.isLoading.value = true;
//         // 判断输入项是否符合条件
//         await nextTick(() => {
//             formRulesOperate.ruleFormRef.value!.validate((valid: boolean) => {
//                 // 输入符合条件后才可以继续后续操作
//                 if (valid) {
//                     this.createApiKey();
//                 }
//                 else {
//                     this.isLoading.value = false;
//                 }
//             });
//         });
//     };

//     createApiKey = () => {
//         createApiKeyAPI(
//             {
//                 ...formRulesOperate.ruleForm
//             }
//         )
//         .then(({data} : {data: {result: {apikeyId: string}}}) => {
//             if (data?.result?.apikeyId?.length) {
//                 msgTip.success(language.t('operate.success'));
//                 this.cancelClick();
//                 emitValue('determine-click');
//                 return;
//             }
//         })
//         .finally(() => {
//             this.isLoading.value = false;
//         })
//         .catch(() => {
//             this.cancelClick();
//             emitValue('determine-click');
//         });
//     };

//     /**
//      * 取消按钮点击操作
//     */
//     cancelClick = (): void => {
//         // 回传父组件当前页码，可以处理相关操作
//         emitValue('dia-log-close', false);
//     };
// };