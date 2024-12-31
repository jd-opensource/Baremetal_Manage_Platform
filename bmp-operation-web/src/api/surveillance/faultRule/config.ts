/**
 * @file
 * @author
*/

import {methodsTotal} from 'utils/index.ts';
import {TotalUrlPathType} from './typeManagement'; // url地址类

const defaultPath: string = methodsTotal.humpSplit('/v1/oobAlert/rules');

// url地址集合
const urlPath: TotalUrlPathType =  {
    faultRulesUrl: defaultPath, // 故障报警规则
    faultRulesExportUrl: defaultPath, // 故障报警规则导出
    faultPushUrl: `${defaultPath}/${methodsTotal.humpSplit('changePush')}`, // 故障推送push
    faultUseUrl: `${defaultPath}/${methodsTotal.humpSplit('changeUse')}`, // 故障启用禁用
    faultResetUrl: `${defaultPath}/reset`, // 恢复默认设置
    faultLevelUrl: `${defaultPath}/${methodsTotal.humpSplit('alertLevelList')}`, // 故障等级列表
    faultSpecUrl: `${defaultPath}/${methodsTotal.humpSplit('alertSpecList')}`, // 故障配件列表  
};

export default urlPath;
