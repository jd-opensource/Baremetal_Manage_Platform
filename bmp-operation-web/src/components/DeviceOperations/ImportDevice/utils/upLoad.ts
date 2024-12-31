import {CurrencyType, UpLoadsType} from '@utils/publicType';
import {msgTip} from 'utils/index.ts';
import {language, CurrentInstance} from 'utils/publicClass.ts';

class UpLoadOperate {
    // 上传
    uploads: Ref<UpLoadsType | undefined> = ref<UpLoadsType>();
    // 是否显示异常提示-默认false
    hasError: Ref<boolean> = ref<boolean>(false);
    formRulesOperate: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor (formRulesOperate: any) {
        this.formRulesOperate = formRulesOperate;
    };

    /**
     * 上传的event事件
    */
    uploadEvent = (refs: {value: UpLoadsType}): void => {
        this.uploads.value = refs.value;
        
    };

    handleChange = (_: unknown, data: Array<CurrencyType>) => {
        this.formRulesOperate.ruleForm.deviceInfo = data;
        this.uploads.value!.submit();
    };

    /**
     * 上传成功的返回
    */
    uploadSuccess = (data: CurrencyType): void => {
        this.formRulesOperate.ruleForm.fileParams = data;
        this.hasError.value = false;
    };

    /**
     * 点击删除上传信息时触发
    */
    removeChange = (): boolean => {
        this.formRulesOperate.ruleForm.deviceInfo = [];
        return this.hasError.value = false;
    };

    progressChange = (info: {percent: number | {toFixed(arg0: number): number;}}) => {
        this.hasError.value = false;
        info.percent = Number(info.percent.toFixed(2));
    };
    /**
     * 仅限上传一次，第二次上次替换第一次的值时-触发
     * @param {Object} files 上传的文件
    */
    handleExceed = (files: CurrencyType[]): void => {
        this.uploads?.value!.clearFiles();
        const file: CurrencyType = files[0];
        this.uploads.value!.handleStart(file);
        this.hasError.value = false;
    };

    /**
     * 上传失败的返回
     * @param {Object} error 失败信息
    */
    errorChange = (error: Error & string): void => {
        const status = JSON.parse(JSON.stringify(error)).status;
        if (status === 402) {
            const path: string = this.instanceProxy!.$defInfo.routerPath('login') as unknown as string;
            window.open(path, '_self');
            return;
        }
        const errorInfo = JSON.parse(error.message);
        msgTip.error(errorInfo?.error?.message??language.t('importDevice.errTip.error'));
    };
};

export default UpLoadOperate;
