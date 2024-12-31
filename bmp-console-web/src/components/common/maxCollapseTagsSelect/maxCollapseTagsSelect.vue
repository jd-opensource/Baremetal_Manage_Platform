<template>
  <div>
    <el-select v-model="internalValue" multiple @change="handleChange">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
      </el-option>
      <template v-if="hiddenTagsCount > 0" #default="{ option, selected }">
        <span>+{{ hiddenTagsCount }} more</span>
      </template>
    </el-select>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, toRefs, watch, ref, computed } from 'vue';

export default defineComponent({
  name: 'MaxCollapseTagsSelect',
  props: {
    modelValue: {
      type: Array as PropType<Array<any>>,
      default: () => [],
    },
    options: {
      type: Array as PropType<Array<{ label: string; value: any }>>,
      default: () => [],
    },
    maxCollapseTags: {
      type: Number,
      default: 3,
    },
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const { modelValue, maxCollapseTags } = toRefs(props);
    const internalValue = ref([...modelValue.value]);

    watch(modelValue, (newValue) => {
      internalValue.value = [...newValue];
    });

    const handleChange = (value: Array<any>) => {
      emit('update:modelValue', value);
    };

    const hiddenTagsCount = computed(() => {
      return internalValue.value.length > maxCollapseTags.value
        ? internalValue.value.length - maxCollapseTags.value
        : 0;
    });

    return {
      internalValue,
      handleChange,
      hiddenTagsCount,
    };
  },
});
</script>