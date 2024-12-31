import UiRow from './Row/Row.vue'; // Layout布局 - row组件
import UiCol from './Col/Col.vue'; // Layout布局 - col组件
import UiTag from './Tag/Tag.vue'; // 标签
import UiForm from './Form/Form.vue'; // form表单组件
import UiIcon from './Icon/Icon.vue'; // icon组件
import UiMenu from './Menu/Menu.vue'; // 菜单栏组件
import UiTabs from './Tabs/Tabs.vue'; // 标签页组件
import UiStep from './Step/Step.vue'; // 步骤条
import UiSteps from './Steps/Steps.vue'; // 步骤条
import EmptyText from './Text/Text.vue'; // 空组件-文案
import UiRadio from './Radio/Radio.vue'; // 单选框组件
import UiInput from './Input/Input.vue'; // input输入框组件
import UiTable from './Table/Table.vue'; // table表格组件
import UiSwitch from './Switch/Switch.vue'; // switch开关组件
import UiUpload from './UpLoad/UpLoad.vue'; // 上传组件
import UiDialog from './DiaLog/DiaLog.vue'; // 对话框组件
import UiCustomMask from './Mask/Mask.vue'; // 自定义蒙层组件
import UiButton from './Button/Button.vue'; // button按钮组件
import UiSelect from './Select/Select.vue'; // select选择器组件
import UiOption from './Option/Option.vue'; // 选择器组件
import UiTooltip from './ToolTip/ToolTip.vue'; // 文字提示组件
import UiFormItem from './FormItem/FormItem.vue'; // 表单-form-item组件
import UiCheckbox from './CheckBox/CheckBox.vue'; // 多选框组件
import UiMenuItem from './MenuItem/MenuItem.vue'; // 菜单栏-menu-item组件
import UiProgress from './Progress/Progress.vue'; // 进度条
import UiScrollBar from './ScrollBar/ScrollBar.vue'; // scroll高度
import UiMenuItemGroup from './MenuItemGroup/MenuItemGroup.vue'; // 菜单栏-menu-item-group组件
import UiSubMenu from './SubMenu/SubMenu.vue'; // 菜单栏-sub-menu组件
import UiDropdown from './DropDown/DropDown.vue'; // 下拉菜单组件
import UiPagination from './Pagination/Pagination.vue'; // 分页组件
import UiRadioGroup from './RadioGroup/RadioGroup.vue'; // 单项选择器组件
import UiInputNumber from './InputNumber/InputNumber.vue'; // 数字输入框组件
import UiTableColumn from './TableColumn/TableColumn.vue'; // 表格-table-column组件
import UiDatePicker from './DatePicker/DatePicker.vue'; // 日期
import UiRadioButton from './RadioButton/RadioButton.vue'; // 单选框-button类组件

import UiDropdownItem from './DropDownItem/DropDownItem.vue'; // 下拉菜单-drop-down-item组件
import UiConfigProvider from './ConfigProvider/ConfigProvider.vue'; // 国际化组件，用于element-plus内部国际化切换
/**
 * 所有组件
*/
const components: {[x: string]: Function;}[] = [
    UiRow,
    UiCol,
    UiTag,
    UiMenu,
    UiForm,
    UiTabs,
    UiStep,
    UiSteps,
    UiIcon,
    UiRadio,
    UiInput,
    UiTable,
    UiSwitch,
    UiButton,
    UiSelect,
    UiOption,
    UiUpload,
    UiDialog,
    EmptyText,
    UiTooltip,
    UiFormItem,
    UiCheckbox,
    UiMenuItem,
    UiProgress,
    UiScrollBar,
    UiMenuItemGroup,
    UiSubMenu,
    UiDropdown,
    UiCustomMask,
    UiRadioGroup,
    UiPagination,
    UiDatePicker,
    UiRadioButton,
    UiInputNumber,
    UiTableColumn,
    UiDropdownItem,
    UiConfigProvider
];

/**
 * 所有组件名
*/
const componentsName: string[] = [
    'UiRow',
    'UiCol',
    'UiTag',
    'UiMenu',
    'UiForm',
    'UiTabs',
    'UiStep',
    'UiSteps',
    'UiIcon',
    'UiRadio',
    'UiInput',
    'UiTable',
    'UiSwitch',
    'UiButton',
    'UiSelect',
    'UiOption',
    'UiUpload',
    'UiDialog',
    'EmptyText',
    'UiTooltip',
    'UiFormItem',
    'UiCheckbox',
    'UiMenuItem',
    'UiProgress',
    'UiScrollBar',
    'UiMenuItemGroup',
    'UiSubMenu',
    'UiDropdown',
    'UiCustomMask',
    'UiRadioGroup',
    'UiPagination',
    'UiDatePicker',
    'UiRadioButton',
    'UiInputNumber',
    'UiTableColumn',
    'UiDropdownItem',
    'UiConfigProvider'
];

/**
 * 遍历，依次挂载
*/
export default {
    // install 方法 接收参数app
    install(app: {component: Function;}): void {
        components.forEach((comp, index: number) => {
            app.component(componentsName[index], comp);
        });
    }
};
