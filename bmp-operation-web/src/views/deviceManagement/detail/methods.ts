/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const methods = [
    ...requireAll(require.context('./drop', true, /^((?!dropDownOpt).)*\.ts$/)),
    ...requireAll(require.context('./tabs', false, /.ts$/)),
    ...requireAll(require.context('./diskDetail', false, /.ts$/)),
    ...requireAll(require.context('./deviceDetail', true, /.ts$/)),
    ...requireAll(require.context('./modelsAssociated', false, /.ts$/)),
    ...requireAll(require.context('./hardwareMonitoring/table', false, /tableScroll.ts$/))
];


export default methods.map(item => item.default);
