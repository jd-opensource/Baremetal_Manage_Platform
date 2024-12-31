/**
 * @file
 * @author
*/

class RegularContent {
    // 中国大陆手机号正则
    static phoneReg: RegExp = /^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[189]))\d{8}$/;
    // 中国香港手机号正则
    static hongKongPhoneReg: RegExp = /^([5|6|9])\d{7}$/;
    // 中国澳门手机号正则
    static aOmenPhoneReg: RegExp = /^6\d{7}$/;
    // 中国台湾手机号正则
    static taiWanPhoneReg: RegExp = /^[0][9]\d{8}$/;
    // 邮箱正则
    static emailReg: RegExp = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    // 多邮箱正则
    static moreEmailReg: RegExp = /^((([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6}\,))*(([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})))$/;
    // 名称正则
    static nameReg: RegExp = /^[-0-9a-zA-Z_]{1,64}$/;
    // 名称正则
    static name1Reg: RegExp = /^[-0-9a-zA-Z_]{1,32}$/;
    // 机型名称正则
    static modelNameReg: RegExp = /^(?:[-0-9a-zA-Z_]|[\u3400-\u4DB5\u4E00-\u9FEA\uFA0E\uFA0F\uFA11\uFA13\uFA14\uFA1F\uFA21\uFA23\uFA24\uFA27-\uFA29]|[\uD840-\uD868\uD86A-\uD86C\uD86F-\uD872\uD874-\uD879][\uDC00-\uDFFF]|\uD869[\uDC00-\uDED6\uDF00-\uDFFF]|\uD86D[\uDC00-\uDF34\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1D\uDC20-\uDFFF]|\uD873[\uDC00-\uDEA1\uDEB0-\uDFFF]|\uD87A[\uDC00-\uDFE0])+$/;
    // 密码正则
    static passwordReg: RegExp = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[\d\*\(\)`~!@#\$%&_\-+=\|\{\}\[\]:\";\'<>,.\?\/\）])[a-zA-Z\d\*\(\)`~!@#\$%&_\-+=\|\{\}\[\]:\";\'<>,.\?\/\）]{8,30}$/;
    // 只能输入正整数（包括0）、小数点保留两位小数
    static capacityReg: RegExp = /^(([0-9][0-9]*)|(([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2})))$/;
    // 只能输入正整数（包括0）
    static numberReg: RegExp = /^([0]|[1-9][0-9]*)$/;
    // 实例名称
    static instanceNameReg: RegExp = new RegExp('^[a-zA-Z\\\u4e00-\\\u9fa5]([0-9a-zA-Z\\\u4e00-\\\u9fa5-_.:]{1,127})$');
    // 特殊字符
    static specialCharacters: RegExp = /[\@\#\$\%\^\&\*\+\'{'\'}\=\[\]\【\】\)\）\（\(\!\！\`\~\·\;\；\"\,\.\，\。\?\？\、\——\\\>\<\》\《\|]/g;
    // 主机名正则
    static hostNameReg: RegExp = /(^((?!(^[-.])|(-$)|(--+)|(\.$)|(\.\.+))[0-9A-Za-z-.{}]){2,64}$)|(^\s*$)/;
    // 验证是否含有某一个字符串
    static hasEqualReg: (type: string) => RegExp = (type: string) => new RegExp(type);
    // 数字-不包括0
    static NumReg: RegExp = /^([1-9][0-9]*)$/;
};

export default RegularContent;
