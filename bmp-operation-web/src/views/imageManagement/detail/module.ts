// import TabsData from './tabs/tabs.vue';
// import DescEdit from './descEdit/desc.vue';
// import TableData from './table/tableData.vue';
// import BaseInfo from './baseInfo/baseInfo.vue';
// import ModelAdd from './modelAdd/modelAdd.vue';
// import ModelRemove from './remove/remove.vue';
// import DropDownOperate from './drop/dropDown.vue';
// import Pagination from './pagination/pagination.vue';
// import ImageDelete from './drop/imageDelete/imageDelete.vue';

// const pluginComp = {
//     TabsData,
//     DescEdit,
//     TableData,
//     BaseInfo,
//     ModelAdd,
//     ModelRemove,
//     DropDownOperate,
//     Pagination,
//     ImageDelete
// };

// export default pluginComp;

/**
 * @file
 * @author
*/

const requireAll = <T extends {keys(): {map(arg0: T): any}}>(requireContext: T) => {
    // keys 返回字典的所有的key 
    return requireContext.keys().map(requireContext);
}

const pluginComp = [
    ...requireAll(require.context('./', true, /^((?!imageDetail).)*\.vue$/))
].map(item => item.default);

export default pluginComp;
