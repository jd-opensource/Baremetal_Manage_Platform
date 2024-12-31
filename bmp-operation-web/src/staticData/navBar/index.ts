/**
 * @file
 * @author
*/

// 创建一个 require 上下文，它会匹配
const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const svgFiles = [
    ...requireAll(require.context('../../views/', true, /\/navRouter.ts$/)),
]

const [deviceNavData, idcNavData, imageNavData, inBandMonitoringNavData, _, modelNavData, monitoringNavData, personalCenterNavData, roleNavData, userNavData] = svgFiles.map((item) => item?.default?? item);

// 所有数据列表
const eyeryData = (proxy: {
    $defInfo: {
        routerPath: (arg0: string) => void,
        imagePath: (arg0: string) => void
    }
}) =>  {
    const routerPath = proxy.$defInfo.routerPath;
    const imagePath = proxy.$defInfo.imagePath;
    return [
        idcNavData(routerPath, imagePath),
        modelNavData(routerPath, imagePath),
        imageNavData(routerPath, imagePath),
        deviceNavData(routerPath, imagePath),
        inBandMonitoringNavData(routerPath, imagePath),
        monitoringNavData(routerPath, imagePath),
        roleNavData(routerPath, imagePath),
        userNavData(routerPath, imagePath),
        svgFiles[4].userNavData(routerPath, imagePath),
        svgFiles[4].adminUserNavData(routerPath, imagePath),
        personalCenterNavData(routerPath, imagePath)
    ];
};

export default eyeryData;
