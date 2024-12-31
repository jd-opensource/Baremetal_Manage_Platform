/* eslint-disable */

// shims-vue.d.ts是为了 typescript 做的适配定义文件
// 因为.vue 文件不是一个常规的文件类型，ts 是不能理解 vue 文件的作用
// 加这一段是是告诉 ts，vue 文件是这种类型的，如果不写，import引入.vue文件会报错，其他同理
declare module '*.vue' { // declare声明宣告
    import type { DefineComponent } from 'vue'
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
    const component: DefineComponent<{}, {}, any>
    export default component
}

declare module '*.js';
declare module '*.ts';
// 解决ts文件引入图片报错
declare module '*.png';
declare module '*.jpg';
declare module '*.jpeg';
declare module 'mockjs';
