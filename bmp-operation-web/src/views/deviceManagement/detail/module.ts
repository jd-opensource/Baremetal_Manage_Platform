/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const pluginComp = [
    ...requireAll(require.context('./drop', true, /^((?!deviceDetail|pagination).)*\.vue$/)),
    ...requireAll(require.context('./tabs', false, /.vue$/)),
    ...requireAll(require.context('./diskDetail', false, /.vue$/)),
    ...requireAll(require.context('./performanceMonitoring', false, /^((?!echartsPage|echartsBig).)*\.vue$/)),
    ...requireAll(require.context('./deviceDetail', true, /.vue$/)),
    ...requireAll(require.context('./modelsAssociated', true, /.vue$/)),
    ...requireAll(require.context('./operatLog', false, /^((?!infoTitle).)*\.vue$/)),
    ...requireAll(require.context('./hardwareMonitoring', false, /^((?!deviceStatus|titleInfo).)*\.vue$/))
];

export default pluginComp.map(item => item.default);
