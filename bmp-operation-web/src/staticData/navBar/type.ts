
/**
 * 导航类
*/
type NavType = {
    title: string;
    defaultImg: string;
    changeImg: string;
    path: string | string[];
    // detailPath: string;
    otherPath: [any] | [any, string];
};

type  NavigationBarListaType = Required<NavType> & Partial<{children?: NavType[];}>;

export default NavigationBarListaType;