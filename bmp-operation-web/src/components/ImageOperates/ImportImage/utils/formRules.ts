import {RuleFormType, RuleFormRefType2} from '../typeManagement';

interface DataType {
    ruleFormRef: Ref<RuleFormRefType2 | undefined>;
    ruleForm: RuleFormType;
}

const formRules = () => {
    const data: DataType = {
        // 表单ref
        ruleFormRef: ref<RuleFormRefType2>(),
        // 表单提交项
        ruleForm: reactive<RuleFormType>({
            imageFile: [], // 选择镜像
            imageName: '', // 镜像名称
            architecture: '', // 体系架构
            osType: '', // 操作系统平台
            customOperatePlatform: '', // 自定义操作系统平台
            customVersion: '', // 自定义操作版本
            imageFormat: '', // 镜像格式
            bootMode: null, // 引导模式
            version: '', // 操作系统版本
            systemPartition: [
                {
                    format: "xfs",
                    point: "/",
                    size: 51200
                },
                {
                    format: "swap",
                    point: "swap",
                    size: 10240
                }
            ], // 系统盘分区模块
            description: '' // 镜像描述
        })
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType2}): void => {
        data.ruleFormRef.value = formEl.value;
    };

    return {
        ...data,
        getFormRef
    };
};

export default formRules;
