/**
 * @file
 * @author
*/

const dafaultUrl: string = '/messages';
// const defaultUrl2: string = '/v1';

// url地址集合
const urlPath =  {
    messageListUrl: dafaultUrl, // 消息列表
    statisticUrl: `${dafaultUrl}/statistic`, // 获取message总数和未读数
    messageTypesUrl: `${dafaultUrl}/getMessageTypes`, // 获取消息类型、子类型
    hasUnreadMessageUrl: `${dafaultUrl}/hasUnreadMessage`, // 获取有没有未读消息
    messageByIdUrl: `${dafaultUrl}/getMessageById`, // 获取消息详情
    doReadUrl: `${dafaultUrl}/doRead`, // 消息标为已读
    deleteUrl: `${dafaultUrl}/delete`, // 删除消息
    describeMailUrl: `${dafaultUrl}/describeMail`, // 读消息设置
    dialMailUrl: `${dafaultUrl}/dialMail`, // 消息设置-测试邮箱信息接口
    saveMailUrl: `${dafaultUrl}/saveMail`, // 消息设置-保存接口
    savelsPushMailUrl: `${dafaultUrl}/saveIsPushMail`, // 是否关联邮箱推送
};

export default urlPath;
