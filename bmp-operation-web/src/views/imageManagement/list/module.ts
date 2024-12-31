/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const pluginComp = [
    ...requireAll(require.context('./', true, /^((?!imageList).)*\.vue$/))
].map(item => item.default);

export default pluginComp;