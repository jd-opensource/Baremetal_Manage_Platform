/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

// 将两个上下文合并为一个数组
const pluginComp = [
    ...requireAll(require.context('./', true, /^((?!userList).)*\.vue$/))
];

export default pluginComp;
