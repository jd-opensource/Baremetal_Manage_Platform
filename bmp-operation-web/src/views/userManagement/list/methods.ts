/**
 * @file
 * @author
*/

// 创建一个 require 上下文，它会匹配
const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const methods = [
    ...requireAll(require.context('./', true, /^((?!module|methods|page|typeManagement|pagination|filterStyle|setEmpty).)*\.ts$/))
];

export default methods;
