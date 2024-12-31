<template>
    <ui-form
        class="edit-idc-rule-form"
        label-position="left"
        :label-width="setDiffClass('160px', '115px')"
        :model="formRuleOperate.ruleForm"
        :rules="formRuleOperate.rules"
        @rule-form-ref="formRuleOperate.getFormRef"
    >
        <!-- 机房名称 -->
        <ui-form-item
            prop="name"
            :class="[
                setDiffClass('set-en', ''),
                formRuleOperate.idcNameFlag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.computerRoomName')"
        >
            <ui-input
                type="text"
                v-model.trim="formRuleOperate.ruleForm.name"
                :placeholder="$t('editIdc.placeholder.computerRoomName')"
                @blur="formRuleOperate.validate('name', formRuleOperate.idcNameFlag)"
            />
        </ui-form-item>
        <!-- 机房英文名称 -->
        <ui-form-item
            prop="nameEn"
            :class="[
                setDiffClass('set-en', ''),
                formRuleOperate.nameEnFlag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.computerRoomCode')"
        >
            <ui-input
                type="text"
                v-model.trim="formRuleOperate.ruleForm.nameEn"
                :placeholder="$t('editIdc.placeholder.computerRoomCode')"
                @blur="formRuleOperate.validate('nameEn', formRuleOperate.nameEnFlag)"
            />
        </ui-form-item>
        <!-- 机房等级 -->
        <ui-form-item
            prop="level"
            :class="setDiffClass('set-en', '')"
            :label="$t('editIdc.label.computerRoomGrade')"
        >
            <ui-select
                v-model="formRuleOperate.ruleForm.level"
                :placeholder="$t('editIdc.placeholder.computerRoomGrade')"
            >
                <ui-option
                    style="fontSize: 12px"
                    v-for="item in editIdcOperate.machineRoomGrade"
                    :key="item.value"
                    :label="item.label"
                    :value="item.label"
                />
            </ui-select>
        </ui-form-item>
        <!-- 自定义机房等级 -->
        <ui-form-item
            prop="customGrade"
            v-if="IdcStaticData.otherData.includes((formRuleOperate.ruleForm.level as string))"
            :class="[formRuleOperate.customGradeFlag.value ? 'set-empty' : '']"
        >
            <ui-input
                v-model="formRuleOperate.ruleForm.customGrade"
                :placeholder="$t('editIdc.placeholder.customComputerRoomGrade')"
                @blur="formRuleOperate.validate('customGrade', formRuleOperate.customGradeFlag)"
            />
        </ui-form-item>
        <!-- 机房地址 -->
        <ui-form-item
            prop="address"
            :class="setDiffClass('set-address', 'edit-idc-rule-form-address')" 
            :label="$t('editIdc.label.computerRoomAddress')"
        >
            <ui-input
                type="textarea"
                maxlength="256"
                v-model="formRuleOperate.ruleForm.address"
                show-word-limit
                :placeholder="$t('editIdc.placeholder.computerRoomAddress')"
            />
        </ui-form-item>
        <!-- 带外登录账号 -->
        <ui-form-item
            prop="iloUser"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.iloUserFlag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.outsideLoginAccountNum')"
        >
            <ui-input
                type="text"
                v-model="formRuleOperate.ruleForm.iloUser"
                :placeholder="$t('editIdc.placeholder.outsideLoginAccountNum')"
                @blur="formRuleOperate.validate('iloUser', formRuleOperate.iloUserFlag)"
            />
        </ui-form-item>
        <!-- 带外登录密码 -->
        <ui-form-item
            prop="iloPassword"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.iloPasswordFlag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.outsideLoginPsd')"
        >
            <ui-input
                type="password"
                v-model="formRuleOperate.ruleForm.iloPassword"
                :placeholder="$t('editIdc.placeholder.outsideLoginPsd')"
                @blur="formRuleOperate.validate('iloPassword', formRuleOperate.iloPasswordFlag)"
            />
        </ui-form-item>
        <!-- 交换机1登录账号 -->
        <ui-form-item
            prop="switchUser1"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.switchUser1Flag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.switchboardOneNum')"
        >
            <ui-input
                type="text"
                v-model="formRuleOperate.ruleForm.switchUser1"
                :placeholder="$t('editIdc.placeholder.switchboardOneNum')"
                @blur="formRuleOperate.validate('switchUser1', formRuleOperate.switchUser1Flag)"
            />
        </ui-form-item>
        <!-- 交换机1登录密码 -->
        <ui-form-item
            prop="switchPassword1"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.switchPassword1Flag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.switchboardOnePsd')"
        >
            <ui-input
                type="password"
                v-model="formRuleOperate.ruleForm.switchPassword1"
                :placeholder="$t('editIdc.placeholder.switchboardOnePsd')"
                @blur="formRuleOperate.validate('switchPassword1', formRuleOperate.switchPassword1Flag)"
            />
        </ui-form-item>
        <!-- 交换机2登录账号 -->
        <ui-form-item
            prop="switchUser2"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.switchUser2Flag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.switchboardTwoNum')"
        >
            <ui-input
                type="text"
                v-model="formRuleOperate.ruleForm.switchUser2"
                :placeholder="$t('editIdc.placeholder.switchboardTwoNum')"
                @blur="formRuleOperate.validate('switchUser2', formRuleOperate.switchUser2Flag)"
            />
        </ui-form-item>
        <!-- 交换机2登录密码 -->
        <ui-form-item
            prop="switchPassword2"
            :class="[
                setDiffClass('set-public-style', 'set-width'),
                formRuleOperate.switchPassword2Flag.value ? 'set-empty' : ''
            ]"
            :label="$t('editIdc.label.switchboardTwoPsd')"
        >
            <ui-input
                type="password"
                v-model="formRuleOperate.ruleForm.switchPassword2"
                :placeholder="$t('editIdc.placeholder.switchboardTwoPsd')"
                @blur="formRuleOperate.validate('switchPassword2', formRuleOperate.switchPassword2Flag)"
            />
        </ui-form-item>
    </ui-form>
</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts'; // 设置不同样式
import IdcStaticData from 'staticData/idc/index.ts';


// props类型
interface PropsType {
    formRuleOperate: any;
    editIdcOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>