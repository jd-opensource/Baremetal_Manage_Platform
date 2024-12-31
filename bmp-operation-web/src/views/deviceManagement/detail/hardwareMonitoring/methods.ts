/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const methods = [
    ...requireAll(require.context('./', true, /^((?!methods|module|tableScroll|setEmpty).)*\.ts$/),)
]

export default methods;
