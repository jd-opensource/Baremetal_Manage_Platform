/**
 * @file
 * @author
*/

// 按需引入引入echarts核心模块
import * as echarts from 'echarts/core';
// 引入折线图
import {LineChart} from 'echarts/charts';
// 引入必需的组件
import {TitleComponent, TooltipComponent, LegendComponent, GridComponent} from 'echarts/components';
// 引入Canvas渲染器
import {CanvasRenderer} from 'echarts/renderers';
 // 引入缩放组件
import {DataZoomComponent} from 'echarts/components';

// 注册必需组件
echarts.use([
    TitleComponent,
    TooltipComponent,
    LegendComponent,
    GridComponent,
    LineChart,
    CanvasRenderer,
    DataZoomComponent
]);

export default echarts;
