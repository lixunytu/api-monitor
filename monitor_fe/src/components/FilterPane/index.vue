<template>
  <el-row :gutter="10" class="filter-container">
    <el-col
      v-for="item in filterList"
      :key="item.id"
      class="filter-item"
      :xs="24"
      :sm="12"
      :md="item.colSpan || item.md || 6"
      :lg="item.colSpan || item.lg || 6"
      :xl="item.colSpan || item.xl || 3"
    >
      <!--下拉选择-->
      <template v-if="item.type === 'select'">
        <el-select
          v-model="searchData[item.id]"
          class="block"
          clearable
          :placeholder="item.placeholder"
          :filterable="item.filterable"
          @change="onChangeSelect(item.id)"
        >
          <el-option v-for="i in item.child" :key="i.id" :label="i.text || i.name" :value="i.id" />
        </el-select>
      </template>

      <!--输入框-->
      <template v-if="item.type === 'input'">
        <el-input
          v-model="searchData[item.id]"
          :placeholder="item.placeholder"
          clearable
        />
      </template>

      <!--时间范围选择-->
      <template v-if="item.type === 'rangeDate'">
        <el-date-picker
          v-model="searchData[item.id]"
          :clearable="false"
          style="width: 100%"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          value-format="timestamp"
        />
      </template>

      <!--日期选择-->
      <template v-if="item.type === 'date'">
        <el-date-picker
          v-model="searchData[item.id]"
          value-format="timestamp"
          :clearable="false"
          type="date"
          placeholder="选择日期"
        />
      </template>

    </el-col>

    <el-col :xs="24" :sm="12" :md="6" :lg="6" :xl="3">
      <el-button
        type="primary"
        icon="el-icon-search"
        @click="onClickSearch"
      >筛选</el-button>
      <el-button
        icon="el-icon-brush"
        @click="onClickReset"
      >重置</el-button>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'FilterPane',
  props: {
    filterList: {
      type: Array,
      default: () => []
    },
    searchData: { // 用于初始化数据
      type: Object,
      default: () => {}
    }
  },
  methods: {
    onClickSearch() {
      const searchData = this.$lodash.cloneDeep(this.searchData);
      Object.keys(searchData).map(key => {
        if (searchData[key] === null || searchData[key] === undefined || searchData[key] === '') {
          delete searchData[key];
        }
      });
      this.$emit('search', searchData);
    },
    onClickReset() {
      const searchData = {};
      this.$emit('reset', searchData);
    },
    onChangeSelect(id) {
      this.$emit('changeSelect', id);
    }
  }
}
</script>

<style lang="scss" scope>
.filter-container {
  padding-bottom: 0;

  .filter-item {
    display: inline-block;
    vertical-align: middle;
    margin-bottom: 10px;
  }
}
</style>
