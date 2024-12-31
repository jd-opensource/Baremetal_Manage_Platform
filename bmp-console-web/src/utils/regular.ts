/**
 * 中国大陆手机号正则
*/
export const phoneReg: RegExp = /^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8}$/;

/**
 * 中国香港手机号正则
*/
export const hongKongPhoneReg: RegExp = /^([5|6|9])\d{7}$/;

/**
 * 中国澳门手机号正则
*/
export const aOmenPhoneReg: RegExp = /^6\d{7}$/;

/**
 * 中国台湾手机号正则
*/
export const taiWanPhoneReg: RegExp = /^[0][9]\d{8}$/;

/**
 * 其他国家手机号正则
*/
export const otherPhoneReg: RegExp = /^\d*$/;

/**
 * 邮箱正则
*/
export const emailReg: RegExp = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

/**
 * 密码正则
 */
export const passwordReg: RegExp = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[\d\*\(\)`~!@#\$%&_\-+=\|\{\}\[\]:\";\'<>,.\?\/\）])[a-zA-Z\d\*\(\)`~!@#\$%&_\-+=\|\{\}\[\]:\";\'<>,.\?\/\）]{8,30}$/;

/**
 * 用户名正则
*/
export const userNameReg: RegExp = /^[-0-9a-zA-Z_]{1,64}$/;

/**
 * 用户名正则
*/
export const projectNameReg: RegExp = /^[-0-9a-zA-Z\u4e00-\u9fa5_]{1,64}$/;

/**
 * 实例名称正则
*/
export const instanceNameReg: RegExp = /^[a-zA-Z\u4e00-\u9fa5]([0-9a-zA-Z\u4e00-\u9fa5-_.:]{1,127})$/;

/**
 * host名称正则
*/
export const hostNameReg: RegExp = /(^((?!(^-)|(-$)|(--+)|(^\\.)|(\\.$)|(\\.\\.+))[0-9A-Za-z-.]){2,64}$)|(^\s*$)/;

/**
 * 规则名阈值正则
*/
export const ruleThresholdReg: RegExp = /^(0|[1-9]\d*)(\.\d{1,2})?$/;

/**
 * 规则名阈值%正则
*/
export const ruleThresholdPrecentReg: RegExp = /^(100(\.0+)?|[0-9]{1,2}(\.\d+)?)$/;
