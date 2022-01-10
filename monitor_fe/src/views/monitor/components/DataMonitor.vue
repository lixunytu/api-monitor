<template>
  <div>
    <div v-for="(monitorRule, index) in ruleList" :key="index" class="m-rules">
      <el-row>
        <el-col :span="5">
          <el-select :value="monitorRule.rule" @input="updateRule(index, 'rule', $event)">
            <el-option
              v-for="item in ruleConditionList"
              :key="item.value"
              :label="item.text"
              :value="item.value"
            />
          </el-select>
        </el-col>
        <el-col :span="8" style="margin-left: 4px">
          <el-input :value="monitorRule.arg1" placeholder="键" @input="updateRule(index, 'arg1', $event)"/>
        </el-col>
        <el-col :span="8" style="margin-left: 4px">
          <el-input
            :value="monitorRule.arg2"
            :disabled="monitorRule.rule=='key-exist' || monitorRule.rule=='not-null'"
            placeholder="值"
            @input="updateRule(index, 'arg2', $event)"
          />
        </el-col>
        <el-col :span="2" style="margin-left: 4px">
          <i class="el-icon-circle-plus-outline" @click="addRule"/>
          <i class="el-icon-circle-close" style="color: #F56C6C" @click="deleteRule(index)"/>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import { ruleConditionList } from '../const'

export default {
  name: 'DataMonitor',
  model: {
    prop: 'value', // 绑定的值，通过父组件传递
    event: 'update' // 自定义时间名
  },
  props: {
    value: {
      type: Array,
      default() {
        return []
      }
    }
  },
  data() {
    return {
      ruleConditionList,
      ruleForm: { rule: 'count', arg1: '', arg2: '' }
    }
  },
  computed: {
    ruleList() {
      if (this.value && this.value.length > 0) return this.value
      return [this.$lodash.cloneDeep(this.ruleForm)]
    }
  },
  methods: {
    addRule() {
      this.ruleList.push(this.$lodash.cloneDeep(this.ruleForm));
      this.$emit('update', this.ruleList);
    },
    deleteRule(index) {
      this.ruleList.splice(index, 1);
      this.$emit('update', this.ruleList);
    },
    updateRule(index, key, val) {
      this.ruleList[index][key] = val;
      this.$emit('update', this.ruleList);
    },
  }
}
</script>

<style scoped>

</style>
